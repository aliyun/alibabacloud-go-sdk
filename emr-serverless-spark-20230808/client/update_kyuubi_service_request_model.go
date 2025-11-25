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
	SetReleaseVersion(v string) *UpdateKyuubiServiceRequest
	GetReleaseVersion() *string
	SetReplica(v int32) *UpdateKyuubiServiceRequest
	GetReplica() *int32
	SetRestart(v bool) *UpdateKyuubiServiceRequest
	GetRestart() *bool
	SetSparkConfigs(v string) *UpdateKyuubiServiceRequest
	GetSparkConfigs() *string
}

type UpdateKyuubiServiceRequest struct {
	// example:
	//
	// 2c8g
	ComputeInstance *string `json:"computeInstance,omitempty" xml:"computeInstance,omitempty"`
	// example:
	//
	// []
	KyuubiConfigs *string `json:"kyuubiConfigs,omitempty" xml:"kyuubiConfigs,omitempty"`
	// example:
	//
	// 1.9.2-0.0.2
	KyuubiReleaseVersion *string `json:"kyuubiReleaseVersion,omitempty" xml:"kyuubiReleaseVersion,omitempty"`
	// example:
	//
	// dev_serverless_spark
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	PublicEndpointEnabled *bool `json:"publicEndpointEnabled,omitempty" xml:"publicEndpointEnabled,omitempty"`
	// example:
	//
	// dev_queue
	Queue *string `json:"queue,omitempty" xml:"queue,omitempty"`
	// example:
	//
	// esr-4.6.0 (Spark 3.5.2, Scala 2.12)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// 0
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// example:
	//
	// false
	Restart *bool `json:"restart,omitempty" xml:"restart,omitempty"`
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

func (s *UpdateKyuubiServiceRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *UpdateKyuubiServiceRequest) GetReplica() *int32 {
	return s.Replica
}

func (s *UpdateKyuubiServiceRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdateKyuubiServiceRequest) GetSparkConfigs() *string {
	return s.SparkConfigs
}

func (s *UpdateKyuubiServiceRequest) SetComputeInstance(v string) *UpdateKyuubiServiceRequest {
	s.ComputeInstance = &v
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

func (s *UpdateKyuubiServiceRequest) SetSparkConfigs(v string) *UpdateKyuubiServiceRequest {
	s.SparkConfigs = &v
	return s
}

func (s *UpdateKyuubiServiceRequest) Validate() error {
	return dara.Validate(s)
}
