// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertsCountResponseBody
	GetCode() *int32
	SetData(v *DescribeAlertsCountResponseBodyData) *DescribeAlertsCountResponseBody
	GetData() *DescribeAlertsCountResponseBodyData
	SetMessage(v string) *DescribeAlertsCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertsCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertsCountResponseBody
	GetSuccess() *bool
}

type DescribeAlertsCountResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeAlertsCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertsCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertsCountResponseBody) GetData() *DescribeAlertsCountResponseBodyData {
	return s.Data
}

func (s *DescribeAlertsCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertsCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertsCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertsCountResponseBody) SetCode(v int32) *DescribeAlertsCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetData(v *DescribeAlertsCountResponseBodyData) *DescribeAlertsCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetMessage(v string) *DescribeAlertsCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetRequestId(v string) *DescribeAlertsCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetSuccess(v bool) *DescribeAlertsCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsCountResponseBodyData struct {
	// The total number of alerts.
	//
	// example:
	//
	// 75
	All      *int64            `json:"All,omitempty" xml:"All,omitempty"`
	CountMap map[string]*int64 `json:"CountMap,omitempty" xml:"CountMap,omitempty"`
	// The number of high-risk alerts.
	//
	// example:
	//
	// 25
	High *int64 `json:"High,omitempty" xml:"High,omitempty"`
	// The number of low-risk alerts.
	//
	// example:
	//
	// 25
	Low *int64 `json:"Low,omitempty" xml:"Low,omitempty"`
	// The number of medium-risk alerts.
	//
	// example:
	//
	// 25
	Medium *int64 `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// The number of connected services.
	//
	// example:
	//
	// 3
	ProductNum *int32 `json:"ProductNum,omitempty" xml:"ProductNum,omitempty"`
}

func (s DescribeAlertsCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponseBodyData) GetAll() *int64 {
	return s.All
}

func (s *DescribeAlertsCountResponseBodyData) GetCountMap() map[string]*int64 {
	return s.CountMap
}

func (s *DescribeAlertsCountResponseBodyData) GetHigh() *int64 {
	return s.High
}

func (s *DescribeAlertsCountResponseBodyData) GetLow() *int64 {
	return s.Low
}

func (s *DescribeAlertsCountResponseBodyData) GetMedium() *int64 {
	return s.Medium
}

func (s *DescribeAlertsCountResponseBodyData) GetProductNum() *int32 {
	return s.ProductNum
}

func (s *DescribeAlertsCountResponseBodyData) SetAll(v int64) *DescribeAlertsCountResponseBodyData {
	s.All = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetCountMap(v map[string]*int64) *DescribeAlertsCountResponseBodyData {
	s.CountMap = v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetHigh(v int64) *DescribeAlertsCountResponseBodyData {
	s.High = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetLow(v int64) *DescribeAlertsCountResponseBodyData {
	s.Low = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetMedium(v int64) *DescribeAlertsCountResponseBodyData {
	s.Medium = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetProductNum(v int32) *DescribeAlertsCountResponseBodyData {
	s.ProductNum = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
