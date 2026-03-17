// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseCodes(v []*ListEnterpriseCodeResponseBodyEnterpriseCodes) *ListEnterpriseCodeResponseBody
	GetEnterpriseCodes() []*ListEnterpriseCodeResponseBodyEnterpriseCodes
	SetMaxResults(v int32) *ListEnterpriseCodeResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEnterpriseCodeResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEnterpriseCodeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEnterpriseCodeResponseBody
	GetTotalCount() *int32
}

type ListEnterpriseCodeResponseBody struct {
	// The information about enterprise codes.
	EnterpriseCodes []*ListEnterpriseCodeResponseBodyEnterpriseCodes `json:"EnterpriseCodes,omitempty" xml:"EnterpriseCodes,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for returning the next page when the data is returned in more than one page.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A57EF84-D587-4CF9-B0C8-307488BF52C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEnterpriseCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseCodeResponseBody) GetEnterpriseCodes() []*ListEnterpriseCodeResponseBodyEnterpriseCodes {
	return s.EnterpriseCodes
}

func (s *ListEnterpriseCodeResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEnterpriseCodeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEnterpriseCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnterpriseCodeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEnterpriseCodeResponseBody) SetEnterpriseCodes(v []*ListEnterpriseCodeResponseBodyEnterpriseCodes) *ListEnterpriseCodeResponseBody {
	s.EnterpriseCodes = v
	return s
}

func (s *ListEnterpriseCodeResponseBody) SetMaxResults(v int32) *ListEnterpriseCodeResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEnterpriseCodeResponseBody) SetNextToken(v string) *ListEnterpriseCodeResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEnterpriseCodeResponseBody) SetRequestId(v string) *ListEnterpriseCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseCodeResponseBody) SetTotalCount(v int32) *ListEnterpriseCodeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEnterpriseCodeResponseBody) Validate() error {
	if s.EnterpriseCodes != nil {
		for _, item := range s.EnterpriseCodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnterpriseCodeResponseBodyEnterpriseCodes struct {
	// The enterprise code.
	//
	// example:
	//
	// 12P**
	EnterpriseCode *string `json:"EnterpriseCode,omitempty" xml:"EnterpriseCode,omitempty"`
	// Indicates whether the enterprise code is the default one.
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
}

func (s ListEnterpriseCodeResponseBodyEnterpriseCodes) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseCodeResponseBodyEnterpriseCodes) GoString() string {
	return s.String()
}

func (s *ListEnterpriseCodeResponseBodyEnterpriseCodes) GetEnterpriseCode() *string {
	return s.EnterpriseCode
}

func (s *ListEnterpriseCodeResponseBodyEnterpriseCodes) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListEnterpriseCodeResponseBodyEnterpriseCodes) SetEnterpriseCode(v string) *ListEnterpriseCodeResponseBodyEnterpriseCodes {
	s.EnterpriseCode = &v
	return s
}

func (s *ListEnterpriseCodeResponseBodyEnterpriseCodes) SetIsDefault(v bool) *ListEnterpriseCodeResponseBodyEnterpriseCodes {
	s.IsDefault = &v
	return s
}

func (s *ListEnterpriseCodeResponseBodyEnterpriseCodes) Validate() error {
	return dara.Validate(s)
}
