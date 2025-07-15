// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedApiGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribePurchasedApiGroupResponseBody
	GetDescription() *string
	SetDomains(v *DescribePurchasedApiGroupResponseBodyDomains) *DescribePurchasedApiGroupResponseBody
	GetDomains() *DescribePurchasedApiGroupResponseBodyDomains
	SetGroupId(v string) *DescribePurchasedApiGroupResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribePurchasedApiGroupResponseBody
	GetGroupName() *string
	SetPurchasedTime(v string) *DescribePurchasedApiGroupResponseBody
	GetPurchasedTime() *string
	SetRegionId(v string) *DescribePurchasedApiGroupResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribePurchasedApiGroupResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribePurchasedApiGroupResponseBody
	GetStatus() *string
}

type DescribePurchasedApiGroupResponseBody struct {
	// The description of the API group.
	//
	// example:
	//
	// api group description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of domain names.
	Domains *DescribePurchasedApiGroupResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// The ID of the API group.
	//
	// example:
	//
	// 48977d7b96074966a7c9c2a8872d7e06
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// Weather
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the API group was purchased.
	//
	// example:
	//
	// 2021-12-19T00:00:00
	PurchasedTime *string `json:"PurchasedTime,omitempty" xml:"PurchasedTime,omitempty"`
	// The region where the API group is located.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 61A16D46-EC04-5288-8A18-811B0F536CC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the API group.
	//
	// 	- **NORMAL**: The API group is normal.
	//
	// 	- **DELETE**: The API group is deleted.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePurchasedApiGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribePurchasedApiGroupResponseBody) GetDomains() *DescribePurchasedApiGroupResponseBodyDomains {
	return s.Domains
}

func (s *DescribePurchasedApiGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePurchasedApiGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePurchasedApiGroupResponseBody) GetPurchasedTime() *string {
	return s.PurchasedTime
}

func (s *DescribePurchasedApiGroupResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePurchasedApiGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePurchasedApiGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribePurchasedApiGroupResponseBody) SetDescription(v string) *DescribePurchasedApiGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetDomains(v *DescribePurchasedApiGroupResponseBodyDomains) *DescribePurchasedApiGroupResponseBody {
	s.Domains = v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetGroupId(v string) *DescribePurchasedApiGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetGroupName(v string) *DescribePurchasedApiGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetPurchasedTime(v string) *DescribePurchasedApiGroupResponseBody {
	s.PurchasedTime = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetRegionId(v string) *DescribePurchasedApiGroupResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetRequestId(v string) *DescribePurchasedApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetStatus(v string) *DescribePurchasedApiGroupResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePurchasedApiGroupResponseBodyDomains struct {
	DomainItem []*DescribePurchasedApiGroupResponseBodyDomainsDomainItem `json:"DomainItem,omitempty" xml:"DomainItem,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiGroupResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupResponseBodyDomains) GetDomainItem() []*DescribePurchasedApiGroupResponseBodyDomainsDomainItem {
	return s.DomainItem
}

func (s *DescribePurchasedApiGroupResponseBodyDomains) SetDomainItem(v []*DescribePurchasedApiGroupResponseBodyDomainsDomainItem) *DescribePurchasedApiGroupResponseBodyDomains {
	s.DomainItem = v
	return s
}

func (s *DescribePurchasedApiGroupResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribePurchasedApiGroupResponseBodyDomainsDomainItem struct {
	// The domain name.
	//
	// example:
	//
	// test_domain.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribePurchasedApiGroupResponseBodyDomainsDomainItem) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupResponseBodyDomainsDomainItem) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupResponseBodyDomainsDomainItem) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePurchasedApiGroupResponseBodyDomainsDomainItem) SetDomainName(v string) *DescribePurchasedApiGroupResponseBodyDomainsDomainItem {
	s.DomainName = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBodyDomainsDomainItem) Validate() error {
	return dara.Validate(s)
}
