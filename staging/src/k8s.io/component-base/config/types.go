/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClientConnectionConfiguration contains details for constructing a client.
type ClientConnectionConfiguration struct {
	// kubeconfig is the path to a KubeConfig file.
	Kubeconfig string
	// acceptContentTypes defines the Accept header sent by clients when connecting to a server, overriding the
	// default value of 'application/json'. This field will control all connections to the server used by a particular
	// client.
	AcceptContentTypes string
	// contentType is the content type used when sending data to the server from this client.
	ContentType string
	// qps controls the number of queries per second allowed for this connection.
	QPS float32
	// burst allows extra queries to accumulate when a client is exceeding its rate.
	Burst int32
}

// LeaderElectionConfiguration 定义了启用了 leader election 的组件的 leader election 客户端的配置信息。
// 在 Kubernetes 中，leader election 是一种机制，用于确保在集群中只有一个副本可以执行某些关键任务，例如控制器的管理或其他需要协调的操作。Leader election 通过选举出一个 leader 来实现这一目标，而 LeaderElectionConfiguration 则定义了这个过程中客户端的行为和属性。
// LeaderElectionConfiguration 包括以下属性：
// leaderElect: 一个 bool 类型的值，指示该组件是否启用 leader election。
// leaseDuration: 一个 time.Duration 类型的值，表示 leader 任期的持续时间。
// renewDeadline: 一个 time.Duration 类型的值，表示 leader 必须在此时间之前续约，否则将丧失领导权。
// retryPeriod: 一个 time.Duration 类型的值，表示在丧失领导权后重新尝试获取领导权之前的等待时间。
// 通过配置这些属性，LeaderElectionConfiguration 可以控制 leader election 客户端的行为，从而确保集群中只有一个副本可以执行关键任务。
type LeaderElectionConfiguration struct {
	// leaderElect 是一个用于实现领导者选举的机制，在执行主循环之前，使领导选举客户端能够获得领导权。在运行复制组件以实现高可用性时，可以启用这个机制。
	// 在分布式系统中，为了确保高可用性和数据一致性，通常会采用领导者选举的方式来选择一个节点作为领导者（或称为主节点），
	// 负责协调整个系统的操作。其他节点则作为从节点，根据领导者的指示执行相应的任务。
	LeaderElect bool
	// leaseDuration是指在领导者选举启用的情况下，非领导者候选人在观察到领导权续约后等待的时间间隔，
	// 然后尝试获取一个已被领导者占据但未经续约的领导者位置。实际上，这是领导者在被替换之前可以被停止的最长时间间隔。
	LeaseDuration metav1.Duration
	// renewDeadline是指在领导者选举启用的情况下，当前领导者尝试更新其领导位置的时间间隔，
	// 以避免失去领导权。如果当前领导者无法在该时间间隔内成功更新其领导位置，则它将停止领导并启动新一轮的领导者选举。
	RenewDeadline metav1.Duration
	// retryPeriod 是指客户端在尝试获取和续约领导权之间等待的持续时间。这仅适用于启用了 leader election 的情况。
	RetryPeriod metav1.Duration
	// resourceLock 是一个用于指示在 leader election 循环中将被用于锁定的资源对象类型的属性。
	// 在 Kubernetes 中，leader election 机制通过使用一种资源对象来实现锁定，以确保只有一个副本可以成为 leader。
	// 这个资源对象可以是 Kubernetes API 中的任何一种类型，例如 ConfigMap、Lease 或 Endpoints 等。
	ResourceLock string
	// resourceName 表示在领导者选举周期中用于加锁的资源对象的名称。
	ResourceName string
	// resourceNamespace 表示在领导者选举周期中用于加锁的资源对象所在的命名空间（namespace）。
	// 在 Kubernetes 等容器编排平台中，领导者选举通常是通过创建一个特殊的 Kubernetes 对象来实现的，
	// 这个对象被称为 Lease。Lease 对象包含了一个租约（lease），用于表示当前节点是否拥有领导权。
	// 当节点获得领导权时，它会创建一个 Lease 对象，并在其中记录自己的标识和租约信息。其他节点可以定期检查这个 Lease 对象，以确定当前的领导者是谁。
	// 为了确保在多个节点同时参与领导者选举时不会出现竞争条件，需要对 Lease 对象进行加锁，
	// 以保证同一时间只有一个节点可以修改这个对象。这个锁通常是通过 Kubernetes API Server 提供的分布式锁机制来实现的。
	// 在这个机制中，每个锁都与一个资源对象相关联，而这个资源对象通常是一个 Kubernetes 对象，例如 ConfigMap、Secret、Pod 等。
	ResourceNamespace string
}

// DebuggingConfiguration holds configuration for Debugging related features.
type DebuggingConfiguration struct {
	// enableProfiling enables profiling via web interface host:port/debug/pprof/
	EnableProfiling bool
	// enableContentionProfiling enables block profiling, if
	// enableProfiling is true.
	EnableContentionProfiling bool
}
