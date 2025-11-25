// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKyuubiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetKyuubiServiceResponseBodyData) *GetKyuubiServiceResponseBody
	GetData() *GetKyuubiServiceResponseBodyData
	SetRequestId(v string) *GetKyuubiServiceResponseBody
	GetRequestId() *string
}

type GetKyuubiServiceResponseBody struct {
	Data *GetKyuubiServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 8CE06D75-E6A2-505D-9B4B-31DEE3D98A04
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetKyuubiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKyuubiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetKyuubiServiceResponseBody) GetData() *GetKyuubiServiceResponseBodyData {
	return s.Data
}

func (s *GetKyuubiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKyuubiServiceResponseBody) SetData(v *GetKyuubiServiceResponseBodyData) *GetKyuubiServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetKyuubiServiceResponseBody) SetRequestId(v string) *GetKyuubiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKyuubiServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKyuubiServiceResponseBodyData struct {
	// example:
	//
	// 2c8g
	ComputeInstance *string `json:"computeInstance,omitempty" xml:"computeInstance,omitempty"`
	// example:
	//
	// 1749456094000
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 150978934701****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// kyuubi-cn-beijing-internal.spark.emr.aliyuncs.com
	InnerEndpoint *string `json:"innerEndpoint,omitempty" xml:"innerEndpoint,omitempty"`
	// example:
	//
	// []
	KyuubiConfigs *string `json:"kyuubiConfigs,omitempty" xml:"kyuubiConfigs,omitempty"`
	// example:
	//
	// 1.9.2-0.0.2
	KyuubiReleaseVersion *string `json:"kyuubiReleaseVersion,omitempty" xml:"kyuubiReleaseVersion,omitempty"`
	// Kyuubi Service ID。
	//
	// example:
	//
	// kb-4e209b04588***95f04ad3538ae4
	KyuubiServiceId *string `json:"kyuubiServiceId,omitempty" xml:"kyuubiServiceId,omitempty"`
	// example:
	//
	// dev_serverless_spark
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// emr-spark-kyuubi-gateway-cn-beijing.aliyuncs.com
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xml:"publicEndpoint,omitempty"`
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
	// []
	SparkConfigs *string `json:"sparkConfigs,omitempty" xml:"sparkConfigs,omitempty"`
	// example:
	//
	// 1749456094000
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetKyuubiServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKyuubiServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKyuubiServiceResponseBodyData) GetComputeInstance() *string {
	return s.ComputeInstance
}

func (s *GetKyuubiServiceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetKyuubiServiceResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetKyuubiServiceResponseBodyData) GetInnerEndpoint() *string {
	return s.InnerEndpoint
}

func (s *GetKyuubiServiceResponseBodyData) GetKyuubiConfigs() *string {
	return s.KyuubiConfigs
}

func (s *GetKyuubiServiceResponseBodyData) GetKyuubiReleaseVersion() *string {
	return s.KyuubiReleaseVersion
}

func (s *GetKyuubiServiceResponseBodyData) GetKyuubiServiceId() *string {
	return s.KyuubiServiceId
}

func (s *GetKyuubiServiceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetKyuubiServiceResponseBodyData) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *GetKyuubiServiceResponseBodyData) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *GetKyuubiServiceResponseBodyData) GetQueue() *string {
	return s.Queue
}

func (s *GetKyuubiServiceResponseBodyData) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *GetKyuubiServiceResponseBodyData) GetReplica() *int32 {
	return s.Replica
}

func (s *GetKyuubiServiceResponseBodyData) GetSparkConfigs() *string {
	return s.SparkConfigs
}

func (s *GetKyuubiServiceResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetKyuubiServiceResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetKyuubiServiceResponseBodyData) SetComputeInstance(v string) *GetKyuubiServiceResponseBodyData {
	s.ComputeInstance = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetCreateTime(v string) *GetKyuubiServiceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetCreator(v string) *GetKyuubiServiceResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetInnerEndpoint(v string) *GetKyuubiServiceResponseBodyData {
	s.InnerEndpoint = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetKyuubiConfigs(v string) *GetKyuubiServiceResponseBodyData {
	s.KyuubiConfigs = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetKyuubiReleaseVersion(v string) *GetKyuubiServiceResponseBodyData {
	s.KyuubiReleaseVersion = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetKyuubiServiceId(v string) *GetKyuubiServiceResponseBodyData {
	s.KyuubiServiceId = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetName(v string) *GetKyuubiServiceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetPublicEndpoint(v string) *GetKyuubiServiceResponseBodyData {
	s.PublicEndpoint = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetPublicEndpointEnabled(v bool) *GetKyuubiServiceResponseBodyData {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetQueue(v string) *GetKyuubiServiceResponseBodyData {
	s.Queue = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetReleaseVersion(v string) *GetKyuubiServiceResponseBodyData {
	s.ReleaseVersion = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetReplica(v int32) *GetKyuubiServiceResponseBodyData {
	s.Replica = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetSparkConfigs(v string) *GetKyuubiServiceResponseBodyData {
	s.SparkConfigs = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetStartTime(v string) *GetKyuubiServiceResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) SetState(v string) *GetKyuubiServiceResponseBodyData {
	s.State = &v
	return s
}

func (s *GetKyuubiServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
