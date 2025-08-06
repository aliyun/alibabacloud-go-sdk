// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnactivatedLicenseOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUnactivatedLicenseOrderResponseBody
	GetCode() *string
	SetData(v []*GetUnactivatedLicenseOrderResponseBodyData) *GetUnactivatedLicenseOrderResponseBody
	GetData() []*GetUnactivatedLicenseOrderResponseBodyData
	SetHttpStatusCode(v int32) *GetUnactivatedLicenseOrderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUnactivatedLicenseOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUnactivatedLicenseOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUnactivatedLicenseOrderResponseBody
	GetSuccess() *bool
}

type GetUnactivatedLicenseOrderResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*GetUnactivatedLicenseOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUnactivatedLicenseOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUnactivatedLicenseOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnactivatedLicenseOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUnactivatedLicenseOrderResponseBody) GetData() []*GetUnactivatedLicenseOrderResponseBodyData {
	return s.Data
}

func (s *GetUnactivatedLicenseOrderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUnactivatedLicenseOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUnactivatedLicenseOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUnactivatedLicenseOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUnactivatedLicenseOrderResponseBody) SetCode(v string) *GetUnactivatedLicenseOrderResponseBody {
	s.Code = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBody) SetData(v []*GetUnactivatedLicenseOrderResponseBodyData) *GetUnactivatedLicenseOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBody) SetHttpStatusCode(v int32) *GetUnactivatedLicenseOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBody) SetMessage(v string) *GetUnactivatedLicenseOrderResponseBody {
	s.Message = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBody) SetRequestId(v string) *GetUnactivatedLicenseOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBody) SetSuccess(v bool) *GetUnactivatedLicenseOrderResponseBody {
	s.Success = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUnactivatedLicenseOrderResponseBodyData struct {
	BeginOn        *string                                                     `json:"BeginOn,omitempty" xml:"BeginOn,omitempty"`
	ContractNo     *string                                                     `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	ExpiredOn      *string                                                     `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
	InstanceId     *string                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LicenseConfigs []*GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs `json:"LicenseConfigs,omitempty" xml:"LicenseConfigs,omitempty" type:"Repeated"`
	OrderType      *string                                                     `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s GetUnactivatedLicenseOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUnactivatedLicenseOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) GetBeginOn() *string {
	return s.BeginOn
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) GetContractNo() *string {
	return s.ContractNo
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) GetLicenseConfigs() []*GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	return s.LicenseConfigs
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) GetOrderType() *string {
	return s.OrderType
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) SetBeginOn(v string) *GetUnactivatedLicenseOrderResponseBodyData {
	s.BeginOn = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) SetContractNo(v string) *GetUnactivatedLicenseOrderResponseBodyData {
	s.ContractNo = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) SetExpiredOn(v string) *GetUnactivatedLicenseOrderResponseBodyData {
	s.ExpiredOn = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) SetInstanceId(v string) *GetUnactivatedLicenseOrderResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) SetLicenseConfigs(v []*GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) *GetUnactivatedLicenseOrderResponseBodyData {
	s.LicenseConfigs = v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) SetOrderType(v string) *GetUnactivatedLicenseOrderResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs struct {
	BusinessType    *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	FeatureIds      *string `json:"FeatureIds,omitempty" xml:"FeatureIds,omitempty"`
	IsTrial         *bool   `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	SdkId           *string `json:"SdkId,omitempty" xml:"SdkId,omitempty"`
	SdkName         *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	Subscription    *string `json:"Subscription,omitempty" xml:"Subscription,omitempty"`
	SubscriptionImp *string `json:"SubscriptionImp,omitempty" xml:"SubscriptionImp,omitempty"`
	SubscriptionPkg *string `json:"SubscriptionPkg,omitempty" xml:"SubscriptionPkg,omitempty"`
}

func (s GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GoString() string {
	return s.String()
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GetFeatureIds() *string {
	return s.FeatureIds
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GetSdkId() *string {
	return s.SdkId
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GetSdkName() *string {
	return s.SdkName
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GetSubscription() *string {
	return s.Subscription
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GetSubscriptionImp() *string {
	return s.SubscriptionImp
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) GetSubscriptionPkg() *string {
	return s.SubscriptionPkg
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) SetBusinessType(v string) *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	s.BusinessType = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) SetFeatureIds(v string) *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	s.FeatureIds = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) SetIsTrial(v bool) *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	s.IsTrial = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) SetSdkId(v string) *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	s.SdkId = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) SetSdkName(v string) *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	s.SdkName = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) SetSubscription(v string) *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	s.Subscription = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) SetSubscriptionImp(v string) *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	s.SubscriptionImp = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) SetSubscriptionPkg(v string) *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs {
	s.SubscriptionPkg = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponseBodyDataLicenseConfigs) Validate() error {
	return dara.Validate(s)
}
