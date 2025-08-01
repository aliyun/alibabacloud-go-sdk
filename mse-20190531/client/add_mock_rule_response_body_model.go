// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMockRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddMockRuleResponseBody
	GetCode() *int32
	SetData(v *AddMockRuleResponseBodyData) *AddMockRuleResponseBody
	GetData() *AddMockRuleResponseBodyData
	SetHttpStatusCode(v int32) *AddMockRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddMockRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddMockRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddMockRuleResponseBody
	GetSuccess() *bool
}

type AddMockRuleResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data *AddMockRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMockRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMockRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddMockRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddMockRuleResponseBody) GetData() *AddMockRuleResponseBodyData {
	return s.Data
}

func (s *AddMockRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddMockRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddMockRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMockRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMockRuleResponseBody) SetCode(v int32) *AddMockRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddMockRuleResponseBody) SetData(v *AddMockRuleResponseBodyData) *AddMockRuleResponseBody {
	s.Data = v
	return s
}

func (s *AddMockRuleResponseBody) SetHttpStatusCode(v int32) *AddMockRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddMockRuleResponseBody) SetMessage(v string) *AddMockRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddMockRuleResponseBody) SetRequestId(v string) *AddMockRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMockRuleResponseBody) SetSuccess(v bool) *AddMockRuleResponseBody {
	s.Success = &v
	return s
}

func (s *AddMockRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddMockRuleResponseBodyData struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 178432728867xxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the consumer application.
	//
	// example:
	//
	// hkhon1po62@a000601b265xxxx
	ConsumerAppId *string `json:"ConsumerAppId,omitempty" xml:"ConsumerAppId,omitempty"`
	// The name of the consumer application.
	//
	// example:
	//
	// demo-xxxx
	ConsumerAppName *string `json:"ConsumerAppName,omitempty" xml:"ConsumerAppName,omitempty"`
	// Indicates whether the mock rule is enabled.
	//
	// 	- `true`: enabled
	//
	// 	- `false`: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The description.
	//
	// example:
	//
	// desc
	ExtraJson *string `json:"ExtraJson,omitempty" xml:"ExtraJson,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 275
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The mock type. Valid values:
	//
	// 	- \\- `[unk]0[unk]`: desktop client
	//
	// 	- \\- `[unk]1[unk]`: mobile client
	//
	// example:
	//
	// 0
	MockType *int64 `json:"MockType,omitempty" xml:"MockType,omitempty"`
	// The name.
	//
	// example:
	//
	// mse-bc1a29b0-160230875****-reg-center-0-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The ID of the service provider application.
	//
	// example:
	//
	// hkhon1po62@a000601b265xxxx
	ProviderAppId *string `json:"ProviderAppId,omitempty" xml:"ProviderAppId,omitempty"`
	// The name of the service provider application.
	//
	// example:
	//
	// demo-xxxx
	ProviderAppName *string `json:"ProviderAppName,omitempty" xml:"ProviderAppName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The HTTP mock rule.
	//
	// example:
	//
	// [{"oper":"return+json","Path":"/mock","Value":"{\\n \\"date\\": \\"2021-09-10T07:45:12.357+0000\\",\\n \\"name\\": \\"name\\",\\n \\"id\\": \\"1\\"\\n}","Method":"GET","Condition":"AND","Timeout":1,"ArgumentMockItems":[{"type":"param","name":"id","value":"1","cond":"==","operator":"rawvalue"},{"type":"param","name":"name","value":"aliyun","cond":"==","operator":"rawvalue"}]}]
	ScMockItemJson *string `json:"ScMockItemJson,omitempty" xml:"ScMockItemJson,omitempty"`
	// The service source.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s AddMockRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddMockRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddMockRuleResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *AddMockRuleResponseBodyData) GetConsumerAppId() *string {
	return s.ConsumerAppId
}

func (s *AddMockRuleResponseBodyData) GetConsumerAppName() *string {
	return s.ConsumerAppName
}

func (s *AddMockRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *AddMockRuleResponseBodyData) GetExtraJson() *string {
	return s.ExtraJson
}

func (s *AddMockRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *AddMockRuleResponseBodyData) GetMockType() *int64 {
	return s.MockType
}

func (s *AddMockRuleResponseBodyData) GetName() *string {
	return s.Name
}

func (s *AddMockRuleResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *AddMockRuleResponseBodyData) GetProviderAppId() *string {
	return s.ProviderAppId
}

func (s *AddMockRuleResponseBodyData) GetProviderAppName() *string {
	return s.ProviderAppName
}

func (s *AddMockRuleResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *AddMockRuleResponseBodyData) GetScMockItemJson() *string {
	return s.ScMockItemJson
}

func (s *AddMockRuleResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *AddMockRuleResponseBodyData) SetAccountId(v string) *AddMockRuleResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetConsumerAppId(v string) *AddMockRuleResponseBodyData {
	s.ConsumerAppId = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetConsumerAppName(v string) *AddMockRuleResponseBodyData {
	s.ConsumerAppName = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetEnable(v bool) *AddMockRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetExtraJson(v string) *AddMockRuleResponseBodyData {
	s.ExtraJson = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetId(v int64) *AddMockRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetMockType(v int64) *AddMockRuleResponseBodyData {
	s.MockType = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetName(v string) *AddMockRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetNamespaceId(v string) *AddMockRuleResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetProviderAppId(v string) *AddMockRuleResponseBodyData {
	s.ProviderAppId = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetProviderAppName(v string) *AddMockRuleResponseBodyData {
	s.ProviderAppName = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetRegion(v string) *AddMockRuleResponseBodyData {
	s.Region = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetScMockItemJson(v string) *AddMockRuleResponseBodyData {
	s.ScMockItemJson = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetSource(v string) *AddMockRuleResponseBodyData {
	s.Source = &v
	return s
}

func (s *AddMockRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
