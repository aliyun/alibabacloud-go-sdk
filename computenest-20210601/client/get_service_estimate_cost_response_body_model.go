// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEstimateCostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommodity(v map[string]*CommodityValue) *GetServiceEstimateCostResponseBody
	GetCommodity() map[string]*CommodityValue
	SetRequestId(v string) *GetServiceEstimateCostResponseBody
	GetRequestId() *string
	SetResources(v map[string]interface{}) *GetServiceEstimateCostResponseBody
	GetResources() map[string]interface{}
}

type GetServiceEstimateCostResponseBody struct {
	// Alibaba Cloud Marketplace purchase order information.
	//
	// example:
	//
	// {\\"cmgj00059839\\": {\\"Result\\": {\\"InquiryType\\": \\"Buy\\", \\"Order\\": {\\"Currency\\": \\"CNY\\", \\"DiscountAmount\\": \\"0.0\\", \\"TradeAmount\\": \\"0.01\\", \\"OriginalAmount\\": \\"0.01\\"}}}}
	Commodity map[string]*CommodityValue `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 08ABBB67-39C9-5EE7-98E5-80486F75CE8D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	//
	// example:
	//
	// { "EcsInstance" : { "Type" : "ALIYUN::ECS::Instance", "Success" : true, "Result" : { "Order" : { "Currency" : "CNY", "RuleIds" : [ "102111101338\\*\\*\\*\\*" ], "DetailInfos" : { "ResourcePriceModel" : [ { "OriginalPrice" : 0, "DiscountPrice" : 0, "SubRules" : { "Rule" : [ ] }, "Resource" : "bandwidth", "TradePrice" : 0 }, { "OriginalPrice" : 0, "DiscountPrice" : 0, "SubRules" : { "Rule" : [ ] }, "Resource" : "image", "TradePrice" : 0 }, { "OriginalPrice" : 0.366666, "DiscountPrice" : 0.249012, "SubRules" : { "Rule" : [ ] }, "Resource" : "instanceType", "TradePrice" : 0.117654 }, { "OriginalPrice" : 0.055555, "DiscountPrice" : 0.037729, "SubRules" : { "Rule" : [ ] }, "Resource" : "systemDisk", "TradePrice" : 0.017826 } ] }, "TradeAmount" : 0.135, "OriginalAmount" : 0.422, "Coupons" : { "Coupon" : [ ] }, "DiscountAmount" : 0.287 }, "OrderSupplement" : { "PriceUnit" : "/Hour", "ChargeType" : "PostPaid", "Quantity" : 1, "PriceType" : "Total" }, "Rules" : { "Rule" : [ { "RuleDescId" : "102111101338\\*\\*\\*\\*", "Name" : "contract discount_multi-billing item discount_3.208750 discount" } ] } } } }
	Resources map[string]interface{} `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s GetServiceEstimateCostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponseBody) GetCommodity() map[string]*CommodityValue {
	return s.Commodity
}

func (s *GetServiceEstimateCostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceEstimateCostResponseBody) GetResources() map[string]interface{} {
	return s.Resources
}

func (s *GetServiceEstimateCostResponseBody) SetCommodity(v map[string]*CommodityValue) *GetServiceEstimateCostResponseBody {
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
