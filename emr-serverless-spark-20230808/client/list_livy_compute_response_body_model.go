// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLivyComputeResponseBody
	GetCode() *string
	SetData(v *ListLivyComputeResponseBodyData) *ListLivyComputeResponseBody
	GetData() *ListLivyComputeResponseBodyData
	SetMessage(v string) *ListLivyComputeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLivyComputeResponseBody
	GetRequestId() *string
}

type ListLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListLivyComputeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListLivyComputeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivyComputeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLivyComputeResponseBody) GetData() *ListLivyComputeResponseBodyData {
	return s.Data
}

func (s *ListLivyComputeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLivyComputeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLivyComputeResponseBody) SetCode(v string) *ListLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *ListLivyComputeResponseBody) SetData(v *ListLivyComputeResponseBodyData) *ListLivyComputeResponseBody {
	s.Data = v
	return s
}

func (s *ListLivyComputeResponseBody) SetMessage(v string) *ListLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *ListLivyComputeResponseBody) SetRequestId(v string) *ListLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLivyComputeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLivyComputeResponseBodyData struct {
	LivyComputes []*ListLivyComputeResponseBodyDataLivyComputes `json:"livyComputes,omitempty" xml:"livyComputes,omitempty" type:"Repeated"`
}

func (s ListLivyComputeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLivyComputeResponseBodyData) GetLivyComputes() []*ListLivyComputeResponseBodyDataLivyComputes {
	return s.LivyComputes
}

func (s *ListLivyComputeResponseBodyData) SetLivyComputes(v []*ListLivyComputeResponseBodyDataLivyComputes) *ListLivyComputeResponseBodyData {
	s.LivyComputes = v
	return s
}

func (s *ListLivyComputeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListLivyComputeResponseBodyDataLivyComputes struct {
	// example:
	//
	// lc-xxxxxxxxxxxx
	ComputeId *string `json:"computeId,omitempty" xml:"computeId,omitempty"`
	// example:
	//
	// alice
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// emr-spark-livy-gateway-cn-hangzhou.data.aliyun.com/api/v1/workspace/w-xxxxxxxxx/livycompute/lc-xxxxxxxxxxx
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// emr-spark-livy-gateway-cn-hangzhou-internal.aliyun.com/api/v1/workspace/w-xxxxxxxxx/livycompute/lc-xxxxxxxxxxx
	EndpointInner *string `json:"endpointInner,omitempty" xml:"endpointInner,omitempty"`
	// example:
	//
	// 1749456094000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// 1749456094000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListLivyComputeResponseBodyDataLivyComputes) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeResponseBodyDataLivyComputes) GoString() string {
	return s.String()
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetComputeId() *string {
	return s.ComputeId
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetEndpointInner() *string {
	return s.EndpointInner
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetName() *string {
	return s.Name
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetQueueName() *string {
	return s.QueueName
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) GetStatus() *string {
	return s.Status
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetComputeId(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.ComputeId = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetCreatedBy(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.CreatedBy = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetEndpoint(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.Endpoint = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetEndpointInner(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.EndpointInner = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetGmtCreate(v int64) *ListLivyComputeResponseBodyDataLivyComputes {
	s.GmtCreate = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetName(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.Name = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetQueueName(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.QueueName = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetStartTime(v int64) *ListLivyComputeResponseBodyDataLivyComputes {
	s.StartTime = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) SetStatus(v string) *ListLivyComputeResponseBodyDataLivyComputes {
	s.Status = &v
	return s
}

func (s *ListLivyComputeResponseBodyDataLivyComputes) Validate() error {
	return dara.Validate(s)
}
