// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAivppResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAivppResourcesResponseBody
	GetCode() *string
	SetData(v []*ListAivppResourcesResponseBodyData) *ListAivppResourcesResponseBody
	GetData() []*ListAivppResourcesResponseBodyData
	SetMaxResults(v int32) *ListAivppResourcesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAivppResourcesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAivppResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAivppResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListAivppResourcesResponseBody
	GetSuccess() *string
	SetTotalCount(v int32) *ListAivppResourcesResponseBody
	GetTotalCount() *int32
}

type ListAivppResourcesResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAivppResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 37cc36b4344b182d36b45d95ae4ef03952ee5c24733ba461
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 68738E75-43C1-5AE5-9F3A-AFEF576D7B5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAivppResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAivppResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAivppResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAivppResourcesResponseBody) GetData() []*ListAivppResourcesResponseBodyData {
	return s.Data
}

func (s *ListAivppResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAivppResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAivppResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAivppResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAivppResourcesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListAivppResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAivppResourcesResponseBody) SetCode(v string) *ListAivppResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetData(v []*ListAivppResourcesResponseBodyData) *ListAivppResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListAivppResourcesResponseBody) SetMaxResults(v int32) *ListAivppResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetMessage(v string) *ListAivppResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetNextToken(v string) *ListAivppResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetRequestId(v string) *ListAivppResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetSuccess(v string) *ListAivppResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetTotalCount(v int32) *ListAivppResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAivppResourcesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAivppResourcesResponseBodyData struct {
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// 2034-03-09T17:47:11Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// i-bp154xh3gt3adb4xu1ue
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// api
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1
	LeftQuantity *string `json:"LeftQuantity,omitempty" xml:"LeftQuantity,omitempty"`
	// example:
	//
	// 231287932080007
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// example:
	//
	// LoadForecasting
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// example:
	//
	// 2024-09-08T01:16Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAivppResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAivppResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAivppResourcesResponseBodyData) GetDetail() *string {
	return s.Detail
}

func (s *ListAivppResourcesResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListAivppResourcesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAivppResourcesResponseBodyData) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListAivppResourcesResponseBodyData) GetLeftQuantity() *string {
	return s.LeftQuantity
}

func (s *ListAivppResourcesResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *ListAivppResourcesResponseBodyData) GetQuantity() *string {
	return s.Quantity
}

func (s *ListAivppResourcesResponseBodyData) GetSpecification() *string {
	return s.Specification
}

func (s *ListAivppResourcesResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAivppResourcesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAivppResourcesResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListAivppResourcesResponseBodyData) SetDetail(v string) *ListAivppResourcesResponseBodyData {
	s.Detail = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetExpireTime(v string) *ListAivppResourcesResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetInstanceId(v string) *ListAivppResourcesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetInstanceType(v string) *ListAivppResourcesResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetLeftQuantity(v string) *ListAivppResourcesResponseBodyData {
	s.LeftQuantity = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetOrderId(v string) *ListAivppResourcesResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetQuantity(v string) *ListAivppResourcesResponseBodyData {
	s.Quantity = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetSpecification(v string) *ListAivppResourcesResponseBodyData {
	s.Specification = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetStartTime(v string) *ListAivppResourcesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetStatus(v string) *ListAivppResourcesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetUserId(v string) *ListAivppResourcesResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
