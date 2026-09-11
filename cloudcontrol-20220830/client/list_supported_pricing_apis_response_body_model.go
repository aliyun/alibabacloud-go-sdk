// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedPricingApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSupportedPricingApisResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupportedPricingApisResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSupportedPricingApisResponseBody
	GetRequestId() *string
	SetSupportedApis(v []*ListSupportedPricingApisResponseBodySupportedApis) *ListSupportedPricingApisResponseBody
	GetSupportedApis() []*ListSupportedPricingApisResponseBodySupportedApis
}

type ListSupportedPricingApisResponseBody struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The array of OpenAPI triplets that support price inquiry. The triplets are sorted in alphabetical order by popCode, popVersion, and apiName.
	SupportedApis []*ListSupportedPricingApisResponseBodySupportedApis `json:"supportedApis,omitempty" xml:"supportedApis,omitempty" type:"Repeated"`
}

func (s ListSupportedPricingApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedPricingApisResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportedPricingApisResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupportedPricingApisResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupportedPricingApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupportedPricingApisResponseBody) GetSupportedApis() []*ListSupportedPricingApisResponseBodySupportedApis {
	return s.SupportedApis
}

func (s *ListSupportedPricingApisResponseBody) SetMaxResults(v int32) *ListSupportedPricingApisResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSupportedPricingApisResponseBody) SetNextToken(v string) *ListSupportedPricingApisResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupportedPricingApisResponseBody) SetRequestId(v string) *ListSupportedPricingApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportedPricingApisResponseBody) SetSupportedApis(v []*ListSupportedPricingApisResponseBodySupportedApis) *ListSupportedPricingApisResponseBody {
	s.SupportedApis = v
	return s
}

func (s *ListSupportedPricingApisResponseBody) Validate() error {
	if s.SupportedApis != nil {
		for _, item := range s.SupportedApis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSupportedPricingApisResponseBodySupportedApis struct {
	// The OpenAPI name in PascalCase, such as RunInstances.
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// The POP product code, such as Ecs, Rds, or Alb. This value corresponds to the popCode field used in price inquiry requests.
	PopCode *string `json:"popCode,omitempty" xml:"popCode,omitempty"`
	// The OpenAPI version number, such as 2014-05-26.
	PopVersion *string `json:"popVersion,omitempty" xml:"popVersion,omitempty"`
}

func (s ListSupportedPricingApisResponseBodySupportedApis) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedPricingApisResponseBodySupportedApis) GoString() string {
	return s.String()
}

func (s *ListSupportedPricingApisResponseBodySupportedApis) GetApiName() *string {
	return s.ApiName
}

func (s *ListSupportedPricingApisResponseBodySupportedApis) GetPopCode() *string {
	return s.PopCode
}

func (s *ListSupportedPricingApisResponseBodySupportedApis) GetPopVersion() *string {
	return s.PopVersion
}

func (s *ListSupportedPricingApisResponseBodySupportedApis) SetApiName(v string) *ListSupportedPricingApisResponseBodySupportedApis {
	s.ApiName = &v
	return s
}

func (s *ListSupportedPricingApisResponseBodySupportedApis) SetPopCode(v string) *ListSupportedPricingApisResponseBodySupportedApis {
	s.PopCode = &v
	return s
}

func (s *ListSupportedPricingApisResponseBodySupportedApis) SetPopVersion(v string) *ListSupportedPricingApisResponseBodySupportedApis {
	s.PopVersion = &v
	return s
}

func (s *ListSupportedPricingApisResponseBodySupportedApis) Validate() error {
	return dara.Validate(s)
}
