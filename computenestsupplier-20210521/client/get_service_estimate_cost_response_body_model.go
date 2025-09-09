// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEstimateCostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommodity(v map[string]interface{}) *GetServiceEstimateCostResponseBody
	GetCommodity() map[string]interface{}
	SetRequestId(v string) *GetServiceEstimateCostResponseBody
	GetRequestId() *string
	SetResources(v map[string]interface{}) *GetServiceEstimateCostResponseBody
	GetResources() map[string]interface{}
}

type GetServiceEstimateCostResponseBody struct {
	// The subscription duration information about the purchase order of Alibaba Cloud Marketplace.
	//
	// example:
	//
	// {\\"PayPeriodUnit\\":Month,\\"PayPeriod\\":1}
	Commodity map[string]interface{} `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E73F09DC-6C13-5CB1-A10F-7A4E125ABD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources.
	//
	// example:
	//
	// {
	//
	//       "ECSInstances":{
	//
	//         "Type":"ALIYUN::ECS::InstanceGroup",
	//
	//         "Success":true,
	//
	//         "Result":{
	//
	//           "Order":{
	//
	//             "Currency":"CNY",
	//
	//             "RuleIds":[
	//
	//               1752723
	//
	//             ],
	//
	//             "DetailInfos":{
	//
	//               "ResourcePriceModel":[
	//
	//                 {
	//
	//                   "Resource":"bandwidth",
	//
	//                   "TradeAmount":0.0,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.0,
	//
	//                   "DiscountAmount":0.0
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"image",
	//
	//                   "TradeAmount":0.0,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.0,
	//
	//                   "DiscountAmount":0.0
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"instanceType",
	//
	//                   "TradeAmount":0.006966,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.45,
	//
	//                   "DiscountAmount":0.443034
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"systemDisk",
	//
	//                   "TradeAmount":0.000867,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.056,
	//
	//                   "DiscountAmount":0.055133
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"dataDisk",
	//
	//                   "TradeAmount":0.002167,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.14,
	//
	//                   "DiscountAmount":0.137833
	//
	//                 }
	//
	//               ]
	//
	//             }
	Resources map[string]interface{} `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s GetServiceEstimateCostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponseBody) GetCommodity() map[string]interface{} {
	return s.Commodity
}

func (s *GetServiceEstimateCostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceEstimateCostResponseBody) GetResources() map[string]interface{} {
	return s.Resources
}

func (s *GetServiceEstimateCostResponseBody) SetCommodity(v map[string]interface{}) *GetServiceEstimateCostResponseBody {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostResponseBody) SetRequestId(v string) *GetServiceEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceEstimateCostResponseBody) SetResources(v map[string]interface{}) *GetServiceEstimateCostResponseBody {
	s.Resources = v
	return s
}

func (s *GetServiceEstimateCostResponseBody) Validate() error {
	return dara.Validate(s)
}
