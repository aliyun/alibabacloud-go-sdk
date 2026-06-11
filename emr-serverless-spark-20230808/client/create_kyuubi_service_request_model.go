// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKyuubiServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeInstance(v string) *CreateKyuubiServiceRequest
	GetComputeInstance() *string
	SetKyuubiConfigs(v string) *CreateKyuubiServiceRequest
	GetKyuubiConfigs() *string
	SetKyuubiReleaseVersion(v string) *CreateKyuubiServiceRequest
	GetKyuubiReleaseVersion() *string
	SetName(v string) *CreateKyuubiServiceRequest
	GetName() *string
	SetPublicEndpointEnabled(v bool) *CreateKyuubiServiceRequest
	GetPublicEndpointEnabled() *bool
	SetQueue(v string) *CreateKyuubiServiceRequest
	GetQueue() *string
	SetReleaseVersion(v string) *CreateKyuubiServiceRequest
	GetReleaseVersion() *string
	SetReplica(v int32) *CreateKyuubiServiceRequest
	GetReplica() *int32
	SetSparkConfigs(v string) *CreateKyuubiServiceRequest
	GetSparkConfigs() *string
}

type CreateKyuubiServiceRequest struct {
	// The specifications of the Kyuubi service.
	//
	// example:
	//
	// 2c8g
	ComputeInstance *string `json:"computeInstance,omitempty" xml:"computeInstance,omitempty"`
	// The Kyuubi service configuration.
	//
	// example:
	//
	// [{\\"key\\":\\"kyuubi.engine.share.level\\",\\"value\\":\\"USER\\"}]
	KyuubiConfigs *string `json:"kyuubiConfigs,omitempty" xml:"kyuubiConfigs,omitempty"`
	// The Kyuubi service database engine version.
	//
	// example:
	//
	// 1.9.2-0.0.2
	KyuubiReleaseVersion *string `json:"kyuubiReleaseVersion,omitempty" xml:"kyuubiReleaseVersion,omitempty"`
	// The Kyuubi service name.
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
	// The queue where the Kyuubi service runs.
	//
	// example:
	//
	// dev_queue
	Queue *string `json:"queue,omitempty" xml:"queue,omitempty"`
	// The Spark database engine version.
	//
	// example:
	//
	// esr-4.6.0 (Spark 3.5.2, Scala 2.12)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The number of replicas for high availability (HA).
	//
	// example:
	//
	// 3
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// The Spark configuration.
	//
	// example:
	//
	// [{\\"key\\":\\"spark.app.name\\",\\"value\\":\\"test\\"}]
	SparkConfigs *string `json:"sparkConfigs,omitempty" xml:"sparkConfigs,omitempty"`
}

func (s CreateKyuubiServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateKyuubiServiceRequest) GetComputeInstance() *string {
	return s.ComputeInstance
}

func (s *CreateKyuubiServiceRequest) GetKyuubiConfigs() *string {
	return s.KyuubiConfigs
}

func (s *CreateKyuubiServiceRequest) GetKyuubiReleaseVersion() *string {
	return s.KyuubiReleaseVersion
}

func (s *CreateKyuubiServiceRequest) GetName() *string {
	return s.Name
}

func (s *CreateKyuubiServiceRequest) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *CreateKyuubiServiceRequest) GetQueue() *string {
	return s.Queue
}

func (s *CreateKyuubiServiceRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *CreateKyuubiServiceRequest) GetReplica() *int32 {
	return s.Replica
}

func (s *CreateKyuubiServiceRequest) GetSparkConfigs() *string {
	return s.SparkConfigs
}

func (s *CreateKyuubiServiceRequest) SetComputeInstance(v string) *CreateKyuubiServiceRequest {
	s.ComputeInstance = &v
	return s
}

func (s *CreateKyuubiServiceRequest) SetKyuubiConfigs(v string) *CreateKyuubiServiceRequest {
	s.KyuubiConfigs = &v
	return s
}

func (s *CreateKyuubiServiceRequest) SetKyuubiReleaseVersion(v string) *CreateKyuubiServiceRequest {
	s.KyuubiReleaseVersion = &v
	return s
}

func (s *CreateKyuubiServiceRequest) SetName(v string) *CreateKyuubiServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateKyuubiServiceRequest) SetPublicEndpointEnabled(v bool) *CreateKyuubiServiceRequest {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *CreateKyuubiServiceRequest) SetQueue(v string) *CreateKyuubiServiceRequest {
	s.Queue = &v
	return s
}

func (s *CreateKyuubiServiceRequest) SetReleaseVersion(v string) *CreateKyuubiServiceRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CreateKyuubiServiceRequest) SetReplica(v int32) *CreateKyuubiServiceRequest {
	s.Replica = &v
	return s
}

func (s *CreateKyuubiServiceRequest) SetSparkConfigs(v string) *CreateKyuubiServiceRequest {
	s.SparkConfigs = &v
	return s
}

func (s *CreateKyuubiServiceRequest) Validate() error {
	return dara.Validate(s)
}
