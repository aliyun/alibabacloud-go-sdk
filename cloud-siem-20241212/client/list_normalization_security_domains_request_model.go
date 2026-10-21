// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationSecurityDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationSecurityDomainsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationSecurityDomainsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationSecurityDomainsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListNormalizationSecurityDomainsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationSecurityDomainsRequest
	GetRoleFor() *int64
}

type ListNormalizationSecurityDomainsRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required. If you do not specify this parameter, the service returns a 400 error: "The specified Lang can not be empty." Explicitly pass a language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of results to return when you use the NextToken-based pagination method. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. You do not need to specify this parameter for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the threat detection and response data management center is located. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: The assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to the view of another member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListNormalizationSecurityDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationSecurityDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationSecurityDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationSecurityDomainsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationSecurityDomainsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationSecurityDomainsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationSecurityDomainsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationSecurityDomainsRequest) SetLang(v string) *ListNormalizationSecurityDomainsRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationSecurityDomainsRequest) SetMaxResults(v int32) *ListNormalizationSecurityDomainsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationSecurityDomainsRequest) SetNextToken(v string) *ListNormalizationSecurityDomainsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationSecurityDomainsRequest) SetRegionId(v string) *ListNormalizationSecurityDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationSecurityDomainsRequest) SetRoleFor(v int64) *ListNormalizationSecurityDomainsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationSecurityDomainsRequest) Validate() error {
	return dara.Validate(s)
}
