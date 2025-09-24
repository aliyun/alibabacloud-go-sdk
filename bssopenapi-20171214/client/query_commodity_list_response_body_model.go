// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommodityListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCommodityListResponseBody
	GetCode() *string
	SetData(v *QueryCommodityListResponseBodyData) *QueryCommodityListResponseBody
	GetData() *QueryCommodityListResponseBodyData
	SetMessage(v string) *QueryCommodityListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCommodityListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCommodityListResponseBody
	GetSuccess() *bool
}

type QueryCommodityListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// NotApplicable
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryCommodityListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CC706AAC-75A6-55B5-9AB7-7D171C6C7655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the information about the service was queried.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCommodityListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCommodityListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCommodityListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCommodityListResponseBody) GetData() *QueryCommodityListResponseBodyData {
	return s.Data
}

func (s *QueryCommodityListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCommodityListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCommodityListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCommodityListResponseBody) SetCode(v string) *QueryCommodityListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCommodityListResponseBody) SetData(v *QueryCommodityListResponseBodyData) *QueryCommodityListResponseBody {
	s.Data = v
	return s
}

func (s *QueryCommodityListResponseBody) SetMessage(v string) *QueryCommodityListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCommodityListResponseBody) SetRequestId(v string) *QueryCommodityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCommodityListResponseBody) SetSuccess(v bool) *QueryCommodityListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCommodityListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCommodityListResponseBodyData struct {
	// The information about the service.
	CommodityList []*QueryCommodityListResponseBodyDataCommodityList `json:"CommodityList,omitempty" xml:"CommodityList,omitempty" type:"Repeated"`
}

func (s QueryCommodityListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCommodityListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCommodityListResponseBodyData) GetCommodityList() []*QueryCommodityListResponseBodyDataCommodityList {
	return s.CommodityList
}

func (s *QueryCommodityListResponseBodyData) SetCommodityList(v []*QueryCommodityListResponseBodyDataCommodityList) *QueryCommodityListResponseBodyData {
	s.CommodityList = v
	return s
}

func (s *QueryCommodityListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryCommodityListResponseBodyDataCommodityList struct {
	// The payment type. Valid values: POSTPAY (pay-as-you-go) and PREPAY (subscription).
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The code of the service, which is the same as that on the Billing Management page.
	//
	// example:
	//
	// ecs
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// ECS (Pay-As-You-Go)
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
}

func (s QueryCommodityListResponseBodyDataCommodityList) String() string {
	return dara.Prettify(s)
}

func (s QueryCommodityListResponseBodyDataCommodityList) GoString() string {
	return s.String()
}

func (s *QueryCommodityListResponseBodyDataCommodityList) GetChargeType() *string {
	return s.ChargeType
}

func (s *QueryCommodityListResponseBodyDataCommodityList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryCommodityListResponseBodyDataCommodityList) GetCommodityName() *string {
	return s.CommodityName
}

func (s *QueryCommodityListResponseBodyDataCommodityList) SetChargeType(v string) *QueryCommodityListResponseBodyDataCommodityList {
	s.ChargeType = &v
	return s
}

func (s *QueryCommodityListResponseBodyDataCommodityList) SetCommodityCode(v string) *QueryCommodityListResponseBodyDataCommodityList {
	s.CommodityCode = &v
	return s
}

func (s *QueryCommodityListResponseBodyDataCommodityList) SetCommodityName(v string) *QueryCommodityListResponseBodyDataCommodityList {
	s.CommodityName = &v
	return s
}

func (s *QueryCommodityListResponseBodyDataCommodityList) Validate() error {
	return dara.Validate(s)
}
