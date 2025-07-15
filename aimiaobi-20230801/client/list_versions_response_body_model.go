// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVersionsResponseBody
	GetCode() *string
	SetData(v []*ListVersionsResponseBodyData) *ListVersionsResponseBody
	GetData() []*ListVersionsResponseBodyData
	SetHttpStatusCode(v int32) *ListVersionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVersionsResponseBody
	GetSuccess() *bool
}

type ListVersionsResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVersionsResponseBody) GetData() []*ListVersionsResponseBodyData {
	return s.Data
}

func (s *ListVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVersionsResponseBody) SetCode(v string) *ListVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVersionsResponseBody) SetData(v []*ListVersionsResponseBodyData) *ListVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListVersionsResponseBody) SetHttpStatusCode(v int32) *ListVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVersionsResponseBody) SetMessage(v string) *ListVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVersionsResponseBody) SetRequestId(v string) *ListVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVersionsResponseBody) SetSuccess(v bool) *ListVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVersionsResponseBodyData struct {
	// example:
	//
	// 43
	ConcurrentCount *int32 `json:"ConcurrentCount,omitempty" xml:"ConcurrentCount,omitempty"`
	// example:
	//
	// 2023-04-23 02:00:34
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 55
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// example:
	//
	// ga-bp12pismsw4v3tzhf62p1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 7
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// CUSTOMIZE
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 13
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// 2023-05-27 04:11:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 65
	UseQuota      *int32  `json:"UseQuota,omitempty" xml:"UseQuota,omitempty"`
	VersionDetail *string `json:"VersionDetail,omitempty" xml:"VersionDetail,omitempty"`
	// example:
	//
	// 试用版
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// example:
	//
	// 87
	VersionStatus *int32 `json:"VersionStatus,omitempty" xml:"VersionStatus,omitempty"`
}

func (s ListVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBodyData) GetConcurrentCount() *int32 {
	return s.ConcurrentCount
}

func (s *ListVersionsResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListVersionsResponseBodyData) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *ListVersionsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVersionsResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListVersionsResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *ListVersionsResponseBodyData) GetQuota() *int32 {
	return s.Quota
}

func (s *ListVersionsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListVersionsResponseBodyData) GetUseQuota() *int32 {
	return s.UseQuota
}

func (s *ListVersionsResponseBodyData) GetVersionDetail() *string {
	return s.VersionDetail
}

func (s *ListVersionsResponseBodyData) GetVersionName() *string {
	return s.VersionName
}

func (s *ListVersionsResponseBodyData) GetVersionStatus() *int32 {
	return s.VersionStatus
}

func (s *ListVersionsResponseBodyData) SetConcurrentCount(v int32) *ListVersionsResponseBodyData {
	s.ConcurrentCount = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetEndTime(v string) *ListVersionsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetInstanceCount(v int32) *ListVersionsResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetInstanceId(v string) *ListVersionsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetOrderId(v int64) *ListVersionsResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetProductType(v string) *ListVersionsResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetQuota(v int32) *ListVersionsResponseBodyData {
	s.Quota = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetStartTime(v string) *ListVersionsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetUseQuota(v int32) *ListVersionsResponseBodyData {
	s.UseQuota = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionDetail(v string) *ListVersionsResponseBodyData {
	s.VersionDetail = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionName(v string) *ListVersionsResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionStatus(v int32) *ListVersionsResponseBodyData {
	s.VersionStatus = &v
	return s
}

func (s *ListVersionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
