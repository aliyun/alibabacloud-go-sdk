// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListKyuubiServicesResponseBodyData) *ListKyuubiServicesResponseBody
	GetData() *ListKyuubiServicesResponseBodyData
	SetRequestId(v string) *ListKyuubiServicesResponseBody
	GetRequestId() *string
}

type ListKyuubiServicesResponseBody struct {
	Data *ListKyuubiServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListKyuubiServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesResponseBody) GetData() *ListKyuubiServicesResponseBodyData {
	return s.Data
}

func (s *ListKyuubiServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKyuubiServicesResponseBody) SetData(v *ListKyuubiServicesResponseBodyData) *ListKyuubiServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListKyuubiServicesResponseBody) SetRequestId(v string) *ListKyuubiServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKyuubiServicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListKyuubiServicesResponseBodyData struct {
	KyuubiServices []*ListKyuubiServicesResponseBodyDataKyuubiServices `json:"kyuubiServices,omitempty" xml:"kyuubiServices,omitempty" type:"Repeated"`
}

func (s ListKyuubiServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesResponseBodyData) GetKyuubiServices() []*ListKyuubiServicesResponseBodyDataKyuubiServices {
	return s.KyuubiServices
}

func (s *ListKyuubiServicesResponseBodyData) SetKyuubiServices(v []*ListKyuubiServicesResponseBodyDataKyuubiServices) *ListKyuubiServicesResponseBodyData {
	s.KyuubiServices = v
	return s
}

func (s *ListKyuubiServicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListKyuubiServicesResponseBodyDataKyuubiServices struct {
	// example:
	//
	// 4C16G
	ComputeInstance *string `json:"computeInstance,omitempty" xml:"computeInstance,omitempty"`
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 103*******
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// kyuubi-cn-hangzhou-internal.spark.emr.aliyuncs.com
	InnerEndpoint *string `json:"innerEndpoint,omitempty" xml:"innerEndpoint,omitempty"`
	// example:
	//
	// kyuubi.conf.key=value1
	//
	// kyuubi.conf.key1=value2
	KyuubiConfigs        *string `json:"kyuubiConfigs,omitempty" xml:"kyuubiConfigs,omitempty"`
	KyuubiReleaseVersion *string `json:"kyuubiReleaseVersion,omitempty" xml:"kyuubiReleaseVersion,omitempty"`
	// KyuubiServer ID。
	//
	// example:
	//
	// kb-070104e7631242448d12a1377c309f30
	KyuubiServiceId *string `json:"kyuubiServiceId,omitempty" xml:"kyuubiServiceId,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// emr-spark-kyuubi-gateway-cn-hangzhou.aliyuncs.com
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xml:"publicEndpoint,omitempty"`
	// example:
	//
	// dev_queue
	Queue *string `json:"queue,omitempty" xml:"queue,omitempty"`
	// example:
	//
	// esr-4.2.0 (Spark 3.5.2, Scala 2.12)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// 3
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// example:
	//
	// spark.conf.key=value1
	//
	// spark.conf.key1=value2
	SparkConfigs *string `json:"sparkConfigs,omitempty" xml:"sparkConfigs,omitempty"`
	// example:
	//
	// 2024-11-23 09:22:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListKyuubiServicesResponseBodyDataKyuubiServices) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiServicesResponseBodyDataKyuubiServices) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetComputeInstance() *string {
	return s.ComputeInstance
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetCreator() *string {
	return s.Creator
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetInnerEndpoint() *string {
	return s.InnerEndpoint
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetKyuubiConfigs() *string {
	return s.KyuubiConfigs
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetKyuubiReleaseVersion() *string {
	return s.KyuubiReleaseVersion
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetKyuubiServiceId() *string {
	return s.KyuubiServiceId
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetName() *string {
	return s.Name
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetQueue() *string {
	return s.Queue
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetReplica() *int32 {
	return s.Replica
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetSparkConfigs() *string {
	return s.SparkConfigs
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetStartTime() *string {
	return s.StartTime
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) GetState() *string {
	return s.State
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetComputeInstance(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.ComputeInstance = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetCreateTime(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.CreateTime = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetCreator(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.Creator = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetInnerEndpoint(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.InnerEndpoint = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetKyuubiConfigs(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.KyuubiConfigs = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetKyuubiReleaseVersion(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.KyuubiReleaseVersion = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetKyuubiServiceId(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.KyuubiServiceId = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetName(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.Name = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetPublicEndpoint(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.PublicEndpoint = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetQueue(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.Queue = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetReleaseVersion(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.ReleaseVersion = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetReplica(v int32) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.Replica = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetSparkConfigs(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.SparkConfigs = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetStartTime(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.StartTime = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) SetState(v string) *ListKyuubiServicesResponseBodyDataKyuubiServices {
	s.State = &v
	return s
}

func (s *ListKyuubiServicesResponseBodyDataKyuubiServices) Validate() error {
	return dara.Validate(s)
}
