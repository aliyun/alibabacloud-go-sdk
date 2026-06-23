// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBrands(v []*ListBrandsResponseBodyBrands) *ListBrandsResponseBody
	GetBrands() []*ListBrandsResponseBodyBrands
	SetMaxResults(v int64) *ListBrandsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListBrandsResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListBrandsResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListBrandsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListBrandsResponseBody
	GetTotalCount() *int64
}

type ListBrandsResponseBody struct {
	// The list of brand data.
	Brands []*ListBrandsResponseBodyBrands `json:"Brands,omitempty" xml:"Brands,omitempty" type:"Repeated"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBrandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBrandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBrandsResponseBody) GetBrands() []*ListBrandsResponseBodyBrands {
	return s.Brands
}

func (s *ListBrandsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListBrandsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBrandsResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListBrandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBrandsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListBrandsResponseBody) SetBrands(v []*ListBrandsResponseBodyBrands) *ListBrandsResponseBody {
	s.Brands = v
	return s
}

func (s *ListBrandsResponseBody) SetMaxResults(v int64) *ListBrandsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBrandsResponseBody) SetNextToken(v string) *ListBrandsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBrandsResponseBody) SetPreviousToken(v string) *ListBrandsResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListBrandsResponseBody) SetRequestId(v string) *ListBrandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBrandsResponseBody) SetTotalCount(v int64) *ListBrandsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBrandsResponseBody) Validate() error {
	if s.Brands != nil {
		for _, item := range s.Brands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBrandsResponseBodyBrands struct {
	// The brand ID.
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// The brand name.
	//
	// example:
	//
	// Custom Brand
	BrandName *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	// The brand type.
	//
	// example:
	//
	// user_custom
	BrandType *string `json:"BrandType,omitempty" xml:"BrandType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The brand status.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBrandsResponseBodyBrands) String() string {
	return dara.Prettify(s)
}

func (s ListBrandsResponseBodyBrands) GoString() string {
	return s.String()
}

func (s *ListBrandsResponseBodyBrands) GetBrandId() *string {
	return s.BrandId
}

func (s *ListBrandsResponseBodyBrands) GetBrandName() *string {
	return s.BrandName
}

func (s *ListBrandsResponseBodyBrands) GetBrandType() *string {
	return s.BrandType
}

func (s *ListBrandsResponseBodyBrands) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBrandsResponseBodyBrands) GetStatus() *string {
	return s.Status
}

func (s *ListBrandsResponseBodyBrands) SetBrandId(v string) *ListBrandsResponseBodyBrands {
	s.BrandId = &v
	return s
}

func (s *ListBrandsResponseBodyBrands) SetBrandName(v string) *ListBrandsResponseBodyBrands {
	s.BrandName = &v
	return s
}

func (s *ListBrandsResponseBodyBrands) SetBrandType(v string) *ListBrandsResponseBodyBrands {
	s.BrandType = &v
	return s
}

func (s *ListBrandsResponseBodyBrands) SetInstanceId(v string) *ListBrandsResponseBodyBrands {
	s.InstanceId = &v
	return s
}

func (s *ListBrandsResponseBodyBrands) SetStatus(v string) *ListBrandsResponseBodyBrands {
	s.Status = &v
	return s
}

func (s *ListBrandsResponseBodyBrands) Validate() error {
	return dara.Validate(s)
}
