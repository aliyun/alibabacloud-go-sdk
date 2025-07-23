// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListUserResourcesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListUserResourcesResponseBody
	GetCode() *string
	SetData(v []*ListUserResourcesResponseBodyData) *ListUserResourcesResponseBody
	GetData() []*ListUserResourcesResponseBodyData
	SetHttpStatusCode(v int32) *ListUserResourcesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUserResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUserResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListUserResourcesResponseBody
	GetSuccess() *string
}

type ListUserResourcesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListUserResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 65308A66-8764-53EE-8D4A-201E86CA88C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListUserResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserResourcesResponseBody) GetData() []*ListUserResourcesResponseBodyData {
	return s.Data
}

func (s *ListUserResourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUserResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserResourcesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListUserResourcesResponseBody) SetAccessDeniedDetail(v string) *ListUserResourcesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetCode(v string) *ListUserResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetData(v []*ListUserResourcesResponseBodyData) *ListUserResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserResourcesResponseBody) SetHttpStatusCode(v int32) *ListUserResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetMessage(v string) *ListUserResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetRequestId(v string) *ListUserResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetSuccess(v string) *ListUserResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserResourcesResponseBodyData struct {
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// brainindustrial_simupostpaid_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// 2024-12-21
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 12345ea3cff446e8837078c2baffbe83
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 20240902
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// ""
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusMsg *string `json:"statusMsg,omitempty" xml:"statusMsg,omitempty"`
}

func (s ListUserResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListUserResourcesResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListUserResourcesResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *ListUserResourcesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserResourcesResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *ListUserResourcesResponseBodyData) GetStartDate() *string {
	return s.StartDate
}

func (s *ListUserResourcesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListUserResourcesResponseBodyData) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *ListUserResourcesResponseBodyData) SetChargeType(v string) *ListUserResourcesResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetCommodityCode(v string) *ListUserResourcesResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetEndDate(v string) *ListUserResourcesResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetInstanceId(v string) *ListUserResourcesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetRegion(v string) *ListUserResourcesResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetStartDate(v string) *ListUserResourcesResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetStatus(v string) *ListUserResourcesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetStatusMsg(v string) *ListUserResourcesResponseBodyData {
	s.StatusMsg = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
