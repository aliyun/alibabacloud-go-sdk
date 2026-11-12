// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlans(v []*ListPlansResponseBodyPlans) *ListPlansResponseBody
	GetPlans() []*ListPlansResponseBodyPlans
	SetRequestId(v string) *ListPlansResponseBody
	GetRequestId() *string
}

type ListPlansResponseBody struct {
	// The plan information.
	Plans []*ListPlansResponseBodyPlans `json:"Plans,omitempty" xml:"Plans,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlansResponseBody) GetPlans() []*ListPlansResponseBodyPlans {
	return s.Plans
}

func (s *ListPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPlansResponseBody) SetPlans(v []*ListPlansResponseBodyPlans) *ListPlansResponseBody {
	s.Plans = v
	return s
}

func (s *ListPlansResponseBody) SetRequestId(v string) *ListPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPlansResponseBody) Validate() error {
	if s.Plans != nil {
		for _, item := range s.Plans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPlansResponseBodyPlans struct {
	// The peak bandwidth. Unit: Mbps.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The number of credits.
	//
	// example:
	//
	// 12500
	Credits *int32 `json:"Credits,omitempty" xml:"Credits,omitempty"`
	// The currency unit. Valid values:
	//
	// - CNY: Chinese Yuan.
	//
	// - USD: US Dollar.
	//
	// > The currency unit is CNY on the China site and USD on the international site.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The disk capacity. Unit: GiB.
	//
	// example:
	//
	// 40
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// - ESSD: Enterprise SSD (ESSD) at performance level (PL) 0.
	//
	// - SSD: standard SSD.
	//
	// - CLOUD_EFFICIENCY: ultra cloud disk.
	//
	// example:
	//
	// ESSD
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The monthly data transfer quota. Unit: GB.
	//
	// example:
	//
	// 400
	Flow *int32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The public network line type. Valid values:
	//
	// - **BGP*	- (default): BGP line.
	//
	// - **BGP_NCO**: BGP (outside the Chinese mainland optimized) line.
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
	// The memory size. Unit: GiB.
	//
	// example:
	//
	// 1
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The monthly price of the simple application server plan.
	//
	// example:
	//
	// 60
	OriginPrice *string `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	// The plan ID.
	//
	// example:
	//
	// swas.s2.c2m1s40b3t04
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The plan type. Valid values:
	//
	// - NORMAL: general-purpose.
	//
	// - CAPACITY: storage-optimized.
	//
	// - PREVIOUS: previous generation.
	//
	// - MULTI_IP: multiple public IP addresses.
	//
	// - INTERNATIONAL: international.
	//
	// > For more information, see [Instance families](https://help.aliyun.com/document_detail/2861773.html).
	//
	// example:
	//
	// NORMAL
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The number of public IP addresses.
	//
	// example:
	//
	// 1
	PublicIpNum *string `json:"PublicIpNum,omitempty" xml:"PublicIpNum,omitempty"`
	// The operating system type supported by the plan.
	//
	// example:
	//
	// ["Linux","Windows"]
	SupportPlatform *string `json:"SupportPlatform,omitempty" xml:"SupportPlatform,omitempty"`
	// The tags.
	Tags []*ListPlansResponseBodyPlansTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The number of tokens.
	//
	// example:
	//
	// 300
	Tokens *int32 `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
}

func (s ListPlansResponseBodyPlans) String() string {
	return dara.Prettify(s)
}

func (s ListPlansResponseBodyPlans) GoString() string {
	return s.String()
}

func (s *ListPlansResponseBodyPlans) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListPlansResponseBodyPlans) GetCore() *int32 {
	return s.Core
}

func (s *ListPlansResponseBodyPlans) GetCredits() *int32 {
	return s.Credits
}

func (s *ListPlansResponseBodyPlans) GetCurrency() *string {
	return s.Currency
}

func (s *ListPlansResponseBodyPlans) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *ListPlansResponseBodyPlans) GetDiskType() *string {
	return s.DiskType
}

func (s *ListPlansResponseBodyPlans) GetFlow() *int32 {
	return s.Flow
}

func (s *ListPlansResponseBodyPlans) GetIspType() *string {
	return s.IspType
}

func (s *ListPlansResponseBodyPlans) GetMemory() *float32 {
	return s.Memory
}

func (s *ListPlansResponseBodyPlans) GetOriginPrice() *string {
	return s.OriginPrice
}

func (s *ListPlansResponseBodyPlans) GetPlanId() *string {
	return s.PlanId
}

func (s *ListPlansResponseBodyPlans) GetPlanType() *string {
	return s.PlanType
}

func (s *ListPlansResponseBodyPlans) GetPublicIpNum() *string {
	return s.PublicIpNum
}

func (s *ListPlansResponseBodyPlans) GetSupportPlatform() *string {
	return s.SupportPlatform
}

func (s *ListPlansResponseBodyPlans) GetTags() []*ListPlansResponseBodyPlansTags {
	return s.Tags
}

func (s *ListPlansResponseBodyPlans) GetTokens() *int32 {
	return s.Tokens
}

func (s *ListPlansResponseBodyPlans) SetBandwidth(v int32) *ListPlansResponseBodyPlans {
	s.Bandwidth = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetCore(v int32) *ListPlansResponseBodyPlans {
	s.Core = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetCredits(v int32) *ListPlansResponseBodyPlans {
	s.Credits = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetCurrency(v string) *ListPlansResponseBodyPlans {
	s.Currency = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetDiskSize(v int32) *ListPlansResponseBodyPlans {
	s.DiskSize = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetDiskType(v string) *ListPlansResponseBodyPlans {
	s.DiskType = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetFlow(v int32) *ListPlansResponseBodyPlans {
	s.Flow = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetIspType(v string) *ListPlansResponseBodyPlans {
	s.IspType = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetMemory(v float32) *ListPlansResponseBodyPlans {
	s.Memory = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetOriginPrice(v string) *ListPlansResponseBodyPlans {
	s.OriginPrice = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetPlanId(v string) *ListPlansResponseBodyPlans {
	s.PlanId = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetPlanType(v string) *ListPlansResponseBodyPlans {
	s.PlanType = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetPublicIpNum(v string) *ListPlansResponseBodyPlans {
	s.PublicIpNum = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetSupportPlatform(v string) *ListPlansResponseBodyPlans {
	s.SupportPlatform = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetTags(v []*ListPlansResponseBodyPlansTags) *ListPlansResponseBodyPlans {
	s.Tags = v
	return s
}

func (s *ListPlansResponseBodyPlans) SetTokens(v int32) *ListPlansResponseBodyPlans {
	s.Tokens = &v
	return s
}

func (s *ListPlansResponseBodyPlans) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPlansResponseBodyPlansTags struct {
	// The Chinese title.
	//
	// example:
	//
	// BGP
	CnTitle *string `json:"CnTitle,omitempty" xml:"CnTitle,omitempty"`
	// The color.
	//
	// example:
	//
	// grey
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The English title.
	//
	// example:
	//
	// BGP
	EnTitle *string `json:"EnTitle,omitempty" xml:"EnTitle,omitempty"`
}

func (s ListPlansResponseBodyPlansTags) String() string {
	return dara.Prettify(s)
}

func (s ListPlansResponseBodyPlansTags) GoString() string {
	return s.String()
}

func (s *ListPlansResponseBodyPlansTags) GetCnTitle() *string {
	return s.CnTitle
}

func (s *ListPlansResponseBodyPlansTags) GetColor() *string {
	return s.Color
}

func (s *ListPlansResponseBodyPlansTags) GetEnTitle() *string {
	return s.EnTitle
}

func (s *ListPlansResponseBodyPlansTags) SetCnTitle(v string) *ListPlansResponseBodyPlansTags {
	s.CnTitle = &v
	return s
}

func (s *ListPlansResponseBodyPlansTags) SetColor(v string) *ListPlansResponseBodyPlansTags {
	s.Color = &v
	return s
}

func (s *ListPlansResponseBodyPlansTags) SetEnTitle(v string) *ListPlansResponseBodyPlansTags {
	s.EnTitle = &v
	return s
}

func (s *ListPlansResponseBodyPlansTags) Validate() error {
	return dara.Validate(s)
}
