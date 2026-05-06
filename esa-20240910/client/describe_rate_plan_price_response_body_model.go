// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRatePlanPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceModel(v *DescribeRatePlanPriceResponseBodyPriceModel) *DescribeRatePlanPriceResponseBody
	GetPriceModel() *DescribeRatePlanPriceResponseBodyPriceModel
	SetRequestId(v string) *DescribeRatePlanPriceResponseBody
	GetRequestId() *string
}

type DescribeRatePlanPriceResponseBody struct {
	PriceModel *DescribeRatePlanPriceResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 50423A7F-A83D-1E24-B80E-86DD25790759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRatePlanPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceResponseBody) GetPriceModel() *DescribeRatePlanPriceResponseBodyPriceModel {
	return s.PriceModel
}

func (s *DescribeRatePlanPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRatePlanPriceResponseBody) SetPriceModel(v *DescribeRatePlanPriceResponseBodyPriceModel) *DescribeRatePlanPriceResponseBody {
	s.PriceModel = v
	return s
}

func (s *DescribeRatePlanPriceResponseBody) SetRequestId(v string) *DescribeRatePlanPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBody) Validate() error {
	if s.PriceModel != nil {
		if err := s.PriceModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRatePlanPriceResponseBodyPriceModel struct {
	RatePlan *DescribeRatePlanPriceResponseBodyPriceModelRatePlan `json:"RatePlan,omitempty" xml:"RatePlan,omitempty" type:"Struct"`
	Rule     *DescribeRatePlanPriceResponseBodyPriceModelRule     `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
}

func (s DescribeRatePlanPriceResponseBodyPriceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceResponseBodyPriceModel) GetRatePlan() *DescribeRatePlanPriceResponseBodyPriceModelRatePlan {
	return s.RatePlan
}

func (s *DescribeRatePlanPriceResponseBodyPriceModel) GetRule() *DescribeRatePlanPriceResponseBodyPriceModelRule {
	return s.Rule
}

func (s *DescribeRatePlanPriceResponseBodyPriceModel) SetRatePlan(v *DescribeRatePlanPriceResponseBodyPriceModelRatePlan) *DescribeRatePlanPriceResponseBodyPriceModel {
	s.RatePlan = v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModel) SetRule(v *DescribeRatePlanPriceResponseBodyPriceModelRule) *DescribeRatePlanPriceResponseBodyPriceModel {
	s.Rule = v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModel) Validate() error {
	if s.RatePlan != nil {
		if err := s.RatePlan.Validate(); err != nil {
			return err
		}
	}
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRatePlanPriceResponseBodyPriceModelRatePlan struct {
	PlanPriceList []*DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList `json:"PlanPriceList,omitempty" xml:"PlanPriceList,omitempty" type:"Repeated"`
}

func (s DescribeRatePlanPriceResponseBodyPriceModelRatePlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceResponseBodyPriceModelRatePlan) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlan) GetPlanPriceList() []*DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	return s.PlanPriceList
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlan) SetPlanPriceList(v []*DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) *DescribeRatePlanPriceResponseBodyPriceModelRatePlan {
	s.PlanPriceList = v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlan) Validate() error {
	if s.PlanPriceList != nil {
		for _, item := range s.PlanPriceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList struct {
	AccelerateType       *string  `json:"AccelerateType,omitempty" xml:"AccelerateType,omitempty"`
	ChargeType           *string  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Coverages            *string  `json:"Coverages,omitempty" xml:"Coverages,omitempty"`
	CrossborderTraffic   *string  `json:"CrossborderTraffic,omitempty" xml:"CrossborderTraffic,omitempty"`
	Currency             *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DcdnPlan             *string  `json:"DcdnPlan,omitempty" xml:"DcdnPlan,omitempty"`
	DiscountPrice        *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	EdgeCompute          *string  `json:"EdgeCompute,omitempty" xml:"EdgeCompute,omitempty"`
	EdgeDdos4Layer       *string  `json:"EdgeDdos4Layer,omitempty" xml:"EdgeDdos4Layer,omitempty"`
	EdgeDdos4LayerIntl   *string  `json:"EdgeDdos4LayerIntl,omitempty" xml:"EdgeDdos4LayerIntl,omitempty"`
	EdgeDdos7Layer       *string  `json:"EdgeDdos7Layer,omitempty" xml:"EdgeDdos7Layer,omitempty"`
	EdgeDdosInstanceCn   *string  `json:"EdgeDdosInstanceCn,omitempty" xml:"EdgeDdosInstanceCn,omitempty"`
	EdgeDdosInstanceIntl *string  `json:"EdgeDdosInstanceIntl,omitempty" xml:"EdgeDdosInstanceIntl,omitempty"`
	EdgeLb4Layer         *string  `json:"EdgeLb4Layer,omitempty" xml:"EdgeLb4Layer,omitempty"`
	EdgeLb4LayerIntl     *string  `json:"EdgeLb4LayerIntl,omitempty" xml:"EdgeLb4LayerIntl,omitempty"`
	EdgeLb7Layer         *string  `json:"EdgeLb7Layer,omitempty" xml:"EdgeLb7Layer,omitempty"`
	EdgeWaf              *string  `json:"EdgeWaf,omitempty" xml:"EdgeWaf,omitempty"`
	EdgeWafInstance      *string  `json:"EdgeWafInstance,omitempty" xml:"EdgeWafInstance,omitempty"`
	Layer4Traffic        *string  `json:"Layer4Traffic,omitempty" xml:"Layer4Traffic,omitempty"`
	Layer4TrafficIntl    *string  `json:"Layer4TrafficIntl,omitempty" xml:"Layer4TrafficIntl,omitempty"`
	PlanName             *string  `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanStatus           *string  `json:"PlanStatus,omitempty" xml:"PlanStatus,omitempty"`
	PlanTraffic          *string  `json:"PlanTraffic,omitempty" xml:"PlanTraffic,omitempty"`
	PlanType             *string  `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	Position             *int32   `json:"Position,omitempty" xml:"Position,omitempty"`
	Price                *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	TotalPrice           *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetAccelerateType() *string {
	return s.AccelerateType
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetCoverages() *string {
	return s.Coverages
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetCrossborderTraffic() *string {
	return s.CrossborderTraffic
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetDcdnPlan() *string {
	return s.DcdnPlan
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeCompute() *string {
	return s.EdgeCompute
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeDdos4Layer() *string {
	return s.EdgeDdos4Layer
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeDdos4LayerIntl() *string {
	return s.EdgeDdos4LayerIntl
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeDdos7Layer() *string {
	return s.EdgeDdos7Layer
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeDdosInstanceCn() *string {
	return s.EdgeDdosInstanceCn
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeDdosInstanceIntl() *string {
	return s.EdgeDdosInstanceIntl
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeLb4Layer() *string {
	return s.EdgeLb4Layer
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeLb4LayerIntl() *string {
	return s.EdgeLb4LayerIntl
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeLb7Layer() *string {
	return s.EdgeLb7Layer
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeWaf() *string {
	return s.EdgeWaf
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetEdgeWafInstance() *string {
	return s.EdgeWafInstance
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetLayer4Traffic() *string {
	return s.Layer4Traffic
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetLayer4TrafficIntl() *string {
	return s.Layer4TrafficIntl
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetPlanStatus() *string {
	return s.PlanStatus
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetPlanTraffic() *string {
	return s.PlanTraffic
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetPlanType() *string {
	return s.PlanType
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetPosition() *int32 {
	return s.Position
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetPrice() *float32 {
	return s.Price
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetAccelerateType(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.AccelerateType = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetChargeType(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.ChargeType = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetCoverages(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.Coverages = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetCrossborderTraffic(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.CrossborderTraffic = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetCurrency(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.Currency = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetDcdnPlan(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.DcdnPlan = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetDiscountPrice(v float32) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeCompute(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeCompute = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeDdos4Layer(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeDdos4Layer = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeDdos4LayerIntl(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeDdos4LayerIntl = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeDdos7Layer(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeDdos7Layer = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeDdosInstanceCn(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeDdosInstanceCn = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeDdosInstanceIntl(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeDdosInstanceIntl = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeLb4Layer(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeLb4Layer = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeLb4LayerIntl(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeLb4LayerIntl = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeLb7Layer(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeLb7Layer = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeWaf(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeWaf = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetEdgeWafInstance(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.EdgeWafInstance = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetLayer4Traffic(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.Layer4Traffic = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetLayer4TrafficIntl(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.Layer4TrafficIntl = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetPlanName(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.PlanName = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetPlanStatus(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.PlanStatus = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetPlanTraffic(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.PlanTraffic = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetPlanType(v string) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.PlanType = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetPosition(v int32) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.Position = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetPrice(v float32) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.Price = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) SetTotalPrice(v float32) *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList {
	s.TotalPrice = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRatePlanPlanPriceList) Validate() error {
	return dara.Validate(s)
}

type DescribeRatePlanPriceResponseBodyPriceModelRule struct {
	RuleList []*DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeRatePlanPriceResponseBodyPriceModelRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceResponseBodyPriceModelRule) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRule) GetRuleList() []*DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList {
	return s.RuleList
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRule) SetRuleList(v []*DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList) *DescribeRatePlanPriceResponseBodyPriceModelRule {
	s.RuleList = v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRule) Validate() error {
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList) SetName(v string) *DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList {
	s.Name = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList) SetRuleDescId(v int64) *DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList {
	s.RuleDescId = &v
	return s
}

func (s *DescribeRatePlanPriceResponseBodyPriceModelRuleRuleList) Validate() error {
	return dara.Validate(s)
}
