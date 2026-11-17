// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKyuubiServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeInstance(v string) *UpdateKyuubiServiceRequest
	GetComputeInstance() *string
	SetExpectedReplica(v int32) *UpdateKyuubiServiceRequest
	GetExpectedReplica() *int32
	SetKyuubiConfigs(v string) *UpdateKyuubiServiceRequest
	GetKyuubiConfigs() *string
	SetKyuubiReleaseVersion(v string) *UpdateKyuubiServiceRequest
	GetKyuubiReleaseVersion() *string
	SetName(v string) *UpdateKyuubiServiceRequest
	GetName() *string
	SetPublicEndpointEnabled(v bool) *UpdateKyuubiServiceRequest
	GetPublicEndpointEnabled() *bool
	SetQueue(v string) *UpdateKyuubiServiceRequest
	GetQueue() *string
	SetRefresh(v bool) *UpdateKyuubiServiceRequest
	GetRefresh() *bool
	SetReleaseVersion(v string) *UpdateKyuubiServiceRequest
	GetReleaseVersion() *string
	SetReplica(v int32) *UpdateKyuubiServiceRequest
	GetReplica() *int32
	SetRestart(v bool) *UpdateKyuubiServiceRequest
	GetRestart() *bool
	SetScale(v bool) *UpdateKyuubiServiceRequest
	GetScale() *bool
	SetSparkConfigs(v string) *UpdateKyuubiServiceRequest
	GetSparkConfigs() *string
}

type UpdateKyuubiServiceRequest struct {
	// The specifications of the Kyuubi service.
	//
	// example:
	//
	// 2c8g
	ComputeInstance *string `json:"computeInstance,omitempty" xml:"computeInstance,omitempty"`
	// The expected number of instances.
	//
	// example:
	//
	// 5
	ExpectedReplica *int32 `json:"expectedReplica,omitempty" xml:"expectedReplica,omitempty"`
	// The Kyuubi service configurations.
	//
	// example:
	//
	// []
	KyuubiConfigs *string `json:"kyuubiConfigs,omitempty" xml:"kyuubiConfigs,omitempty"`
	// The Kyuubi service DPI engine version.
	//
	// example:
	//
	// 1.9.2-0.0.2
	KyuubiReleaseVersion *string `json:"kyuubiReleaseVersion,omitempty" xml:"kyuubiReleaseVersion,omitempty"`
	// The name.
	//
	// example:
	//
	// dev_serverless_spark
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to enable public network access.
	//
	// example:
	//
	// true
	PublicEndpointEnabled *bool `json:"publicEndpointEnabled,omitempty" xml:"publicEndpointEnabled,omitempty"`
	// The running queue for the Kyuubi service.
	//
	// example:
	//
	// dev_queue
	Queue *string `json:"queue,omitempty" xml:"queue,omitempty"`
	// Specifies whether to update the configurations.
	//
	// example:
	//
	// true
	Refresh *bool `json:"refresh,omitempty" xml:"refresh,omitempty"`
	// The Spark DPI engine version number.
	//
	// example:
	//
	// esr-4.6.0 (Spark 3.5.2, Scala 2.12)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The number of replicas enabled for high availability.
	//
	// example:
	//
	// 0
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// Specifies whether to restart.
	//
	// example:
	//
	// false
	Restart *bool `json:"restart,omitempty" xml:"restart,omitempty"`
	// Specifies whether to update the number of replicas.
	//
	// example:
	//
	// true
	Scale *bool `json:"scale,omitempty" xml:"scale,omitempty"`
	// The Spark configurations.
	//
	// example:
	//
	// []
	SparkConfigs *string `json:"sparkConfigs,omitempty" xml:"sparkConfigs,omitempty"`
}

func (s UpdateKyuubiServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKyuubiServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateKyuubiServiceRequest) GetComputeInstance() *string {
	return s.ComputeInstance
}

func (s *UpdateKyuubiServiceRequest) GetExpectedReplica() *int32 {
	return s.ExpectedReplica
}

func (s *UpdateKyuubiServiceRequest) GetKyuubiConfigs() *string {
	return s.KyuubiConfigs
}

func (s *UpdateKyuubiServiceRequest) GetKyuubiReleaseVersion() *string {
	return s.KyuubiReleaseVersion
}

func (s *UpdateKyuubiServiceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKyuubiServiceRequest) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *UpdateKyuubiServiceRequest) GetQueue() *string {
	return s.Queue
}

func (s *UpdateKyuubiServiceRequest) GetRefresh() *bool {
	return s.Refresh
}

func (s *UpdateKyuubiServiceRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *UpdateKyuubiServiceRequest) GetReplica() *int32 {
	return s.Replica
}

func (s *UpdateKyuubiServiceRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdateKyuubiServiceRequest) GetScale() *bool {
	return s.Scale
}

func (s *UpdateKyuubiServiceRequest) GetSparkConfigs() *string {
	return s.SparkConfigs
}

func (s *UpdateKyuubiServiceRequest) SetComputeInstance(v string) *UpdateKyuubiServiceRequest {
	s.ComputeInstance = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetExpectedReplica(v int32) *UpdateKyuubiServiceRequest {
	s.ExpectedReplica = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetKyuubiConfigs(v string) *UpdateKyuubiServiceRequest {
	s.KyuubiConfigs = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetKyuubiReleaseVersion(v string) *UpdateKyuubiServiceRequest {
	s.KyuubiReleaseVersion = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetName(v string) *UpdateKyuubiServiceRequest {
	s.Name = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetPublicEndpointEnabled(v bool) *UpdateKyuubiServiceRequest {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetQueue(v string) *UpdateKyuubiServiceRequest {
	s.Queue = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetRefresh(v bool) *UpdateKyuubiServiceRequest {
	s.Refresh = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetReleaseVersion(v string) *UpdateKyuubiServiceRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetReplica(v int32) *UpdateKyuubiServiceRequest {
	s.Replica = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetRestart(v bool) *UpdateKyuubiServiceRequest {
	s.Restart = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetScale(v bool) *UpdateKyuubiServiceRequest {
	s.Scale = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) SetSparkConfigs(v string) *UpdateKyuubiServiceRequest {
	s.SparkConfigs = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) Validate() error {
	return dara.Validate(s)
}
