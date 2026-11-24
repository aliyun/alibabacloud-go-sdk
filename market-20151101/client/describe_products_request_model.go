// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeProductsRequestFilter) *DescribeProductsRequest
	GetFilter() []*DescribeProductsRequestFilter
	SetPageNumber(v int32) *DescribeProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeProductsRequest
	GetPageSize() *int32
	SetSearchTerm(v string) *DescribeProductsRequest
	GetSearchTerm() *string
}

type DescribeProductsRequest struct {
	// The list of filter conditions.
	Filter []*DescribeProductsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of products to display per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// Image
	SearchTerm *string `json:"SearchTerm,omitempty" xml:"SearchTerm,omitempty"`
}

func (s DescribeProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductsRequest) GetFilter() []*DescribeProductsRequestFilter {
	return s.Filter
}

func (s *DescribeProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProductsRequest) GetSearchTerm() *string {
	return s.SearchTerm
}

func (s *DescribeProductsRequest) SetFilter(v []*DescribeProductsRequestFilter) *DescribeProductsRequest {
	s.Filter = v
	return s
}

func (s *DescribeProductsRequest) SetPageNumber(v int32) *DescribeProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProductsRequest) SetPageSize(v int32) *DescribeProductsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProductsRequest) SetSearchTerm(v string) *DescribeProductsRequest {
	s.SearchTerm = &v
	return s
}

func (s *DescribeProductsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProductsRequestFilter struct {
	// The category of the search tag. Valid values:
	//
	// 	- sort: the sorting method
	//
	// 	- categoryId: the category ID
	//
	// 	- productType: the product type
	//
	// 	- aliyunUid: the Alibaba Cloud account ID of the service provider that owns the product.
	//
	// example:
	//
	// categoryId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the search label:
	//
	//
	// - sort
	//
	//   - user_count-desc: number of transactions in the last 180 days
	//
	//   - created_on-desc: creation time
	//
	//   - price-desc: price
	//
	//   - score-desc: rating
	//
	// - categoryId
	//
	//   - Software Marketplace  53366009
	//
	//     + Runtime Environment  53448001
	//
	//     - Management & Monitoring  53690006
	//
	//     - Website Building System  53616009
	//
	//     - Application Development  55530001
	//
	//     - Database  56024006
	//
	//     - Server Software  56014009
	//
	//     - Network Feature Software  56368007
	//
	//     - Operating System  57742013
	//
	//   - Service Marketplace  52734001
	//
	//     - Data Migration  52738004
	//
	//     - Environment Configuration  52746001
	//
	//     - Troubleshooting  52740002
	//
	//     - Security O&M  52732002
	//
	//     - Custom Development  56082003
	//
	//     - Training & Certification  57252001
	//
	//     - Dedicated Line Access  57392001
	//
	//     - Enterprise Services  56838014
	//
	//     - Security Services  57004003
	//
	//   - Website Building Marketplace  52738001
	//
	//     - Enterprise Display  52738005
	//
	//     - E-commerce  52750001
	//
	//     - Mobile Website  52752001
	//
	//     - Design  52732003
	//
	//     - Information Portal  52744002
	//
	//     - After-sales Service  52744003
	//
	//     - Community Forum  55586021
	//
	//     - Mobile Marketing  55514022
	//
	//     - Template Website  56598032
	//
	//     - Website Hosting  57342011
	//
	//   - Cloud Security Marketplace  56764045
	//
	//     + Host Security  56832023
	//
	//     + Application Security  56846020
	//
	//     + Data Security  56824015
	//
	//     + Security Management  56830014
	//
	//     + Network Security  56820014
	//
	//   - Enterprise Applications  56832009
	//
	//     - Office Management  56778013
	//
	//     - Financial Management  56764034
	//
	//     - HR Management  56780006
	//
	//     - Sales Management  56842010
	//
	//     - Advertising & Marketing  56842011
	//
	//     - Foreign Trade  56790007
	//
	//     - Cloud Communications  57602001
	//
	//     - Supply Chain Management  57604001
	//
	//     - Data Services  57606001
	//
	//   - Solutions  56848023
	//
	//     - E-commerce  56848024
	//
	//     - Finance  56820017
	//
	//     - Logistics  56776025
	//
	//     - Government  56842031
	//
	//     - Multimedia  56784018
	//
	//     - IoT  56794026
	//
	//     - Big Data  56792022
	//
	//   - API Marketplace  56956004
	//
	//     - E-commerce  56928004
	//
	//     - Finance & Investment  57000002
	//
	//     - Transportation & Geography  57002002
	//
	//     - Enterprise Management  56928005
	//
	//     - Public Affairs  57002003
	//
	//     - Meteorology & Water Resources  57096001
	//
	//     - Artificial Intelligence  57124001
	//
	//     - Lifestyle Services  57126001
	//
	//   - IoT Marketplace  201194001
	//
	//     - Integrated Systems  201196001
	//
	//     - Smart Hardware  201198001
	//
	//     - Modules  201200001
	//
	//     - Sensors  201192002
	//
	//     - Chips  201202001
	//
	//     - IoT Applications  201196002
	//
	//     - IoT Services  201196003
	//
	//   - Data Intelligence  201204005
	//
	//     - Data Analytics & Visualization  201190008
	//
	//     - Public Opinion Analysis  201198004
	//
	//     - E-commerce & Marketing  201202006
	//
	//     - Artificial Intelligence  201204006
	//
	//     - Meteorology & Transportation  201214006
	//
	//     - Finance & Accounting  201208013
	//
	// - productType:
	//
	//   - APP: application
	//
	//   - SERVICE: service
	//
	//   - MIRROR: image
	//
	//   - DOWNLOAD: download
	//
	//   - API_SERVICE: API.
	//
	// example:
	//
	// 53366009
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeProductsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeProductsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeProductsRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeProductsRequestFilter) SetKey(v string) *DescribeProductsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeProductsRequestFilter) SetValue(v string) *DescribeProductsRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeProductsRequestFilter) Validate() error {
	return dara.Validate(s)
}
