// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodUserDomainsRequest
	GetDomainName() *string
	SetDomainSearchType(v string) *DescribeVodUserDomainsRequest
	GetDomainSearchType() *string
	SetDomainStatus(v string) *DescribeVodUserDomainsRequest
	GetDomainStatus() *string
	SetOwnerId(v int64) *DescribeVodUserDomainsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVodUserDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVodUserDomainsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeVodUserDomainsRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeVodUserDomainsRequestTag) *DescribeVodUserDomainsRequest
	GetTag() []*DescribeVodUserDomainsRequestTag
}

type DescribeVodUserDomainsRequest struct {
	// The domain name. The value of this parameter is used as a filter condition for a fuzzy match.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The search method. Valid values:
	//
	// 	- **fuzzy_match*	- (default): fuzzy match.
	//
	// 	- **pre_match**: prefix match
	//
	// 	- **suf_match**: suffix match
	//
	// 	- **full_match**: exact match
	//
	// example:
	//
	// fuzzy_match
	DomainSearchType *string `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	// The status of the domain name. Value values:
	//
	// 	- **online**: indicates that the domain name is enabled.
	//
	// 	- **offline**: indicates that the domain name is disabled.
	//
	// 	- **configuring**: indicates that the domain name is being configured.
	//
	// 	- **configure_failed**: indicates that the domain name failed to be configured.
	//
	// 	- **checking**: indicates that the domain name is under review.
	//
	// 	- **check_failed**: indicates that the domain name failed the review.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**. Maximum value: **50**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 20
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tags.
	Tag []*DescribeVodUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVodUserDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodUserDomainsRequest) GetDomainSearchType() *string {
	return s.DomainSearchType
}

func (s *DescribeVodUserDomainsRequest) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeVodUserDomainsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodUserDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVodUserDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVodUserDomainsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodUserDomainsRequest) GetTag() []*DescribeVodUserDomainsRequestTag {
	return s.Tag
}

func (s *DescribeVodUserDomainsRequest) SetDomainName(v string) *DescribeVodUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetDomainSearchType(v string) *DescribeVodUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetDomainStatus(v string) *DescribeVodUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetOwnerId(v int64) *DescribeVodUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetPageNumber(v int32) *DescribeVodUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetPageSize(v int32) *DescribeVodUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetSecurityToken(v string) *DescribeVodUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetTag(v []*DescribeVodUserDomainsRequestTag) *DescribeVodUserDomainsRequest {
	s.Tag = v
	return s
}

func (s *DescribeVodUserDomainsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodUserDomainsRequestTag struct {
	// The key of tag N. Valid values of N: **1*	- to **20**.
	//
	// By default, all tag keys are queried.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: **1*	- to **20**.
	//
	// By default, all tag values are queried.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodUserDomainsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVodUserDomainsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVodUserDomainsRequestTag) SetKey(v string) *DescribeVodUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVodUserDomainsRequestTag) SetValue(v string) *DescribeVodUserDomainsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVodUserDomainsRequestTag) Validate() error {
	return dara.Validate(s)
}
