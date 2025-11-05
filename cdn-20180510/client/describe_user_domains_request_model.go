// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdnType(v string) *DescribeUserDomainsRequest
	GetCdnType() *string
	SetChangeEndTime(v string) *DescribeUserDomainsRequest
	GetChangeEndTime() *string
	SetChangeStartTime(v string) *DescribeUserDomainsRequest
	GetChangeStartTime() *string
	SetCheckDomainShow(v bool) *DescribeUserDomainsRequest
	GetCheckDomainShow() *bool
	SetCoverage(v string) *DescribeUserDomainsRequest
	GetCoverage() *string
	SetDomainName(v string) *DescribeUserDomainsRequest
	GetDomainName() *string
	SetDomainSearchType(v string) *DescribeUserDomainsRequest
	GetDomainSearchType() *string
	SetDomainStatus(v string) *DescribeUserDomainsRequest
	GetDomainStatus() *string
	SetOwnerId(v int64) *DescribeUserDomainsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeUserDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUserDomainsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeUserDomainsRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *DescribeUserDomainsRequest
	GetSecurityToken() *string
	SetSource(v string) *DescribeUserDomainsRequest
	GetSource() *string
	SetTag(v []*DescribeUserDomainsRequestTag) *DescribeUserDomainsRequest
	GetTag() []*DescribeUserDomainsRequestTag
}

type DescribeUserDomainsRequest struct {
	// The type of workload accelerated by Alibaba Cloud CDN. Separate types with commas (,). Valid values:
	//
	// 	- **web**: images and small files
	//
	// 	- **download**: large files
	//
	// 	- **video**: on-demand video and audio streaming
	//
	// If you do not set this parameter, all service types are queried.
	//
	// example:
	//
	// download,web,video
	CdnType *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-10-10T12:14:58Z
	ChangeEndTime *string `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-10-10T12:14:55Z
	ChangeStartTime *string `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	// Specifies whether to display domain names that are under review, failed the review, or failed to be configured. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CheckDomainShow *bool `json:"CheckDomainShow,omitempty" xml:"CheckDomainShow,omitempty"`
	// The acceleration region. By default, all acceleration regions are queried. Valid values:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **global**: global
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The accelerated domain. If you do not set this parameter, all domain names that match the conditions are returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The search mode. Valid values:
	//
	// 	- **fuzzy_match**: fuzzy match
	//
	// 	- **pre_match**: prefix match
	//
	// 	- **suf_match**: suffix match
	//
	// 	- **full_match*	- (default): exact match
	//
	// > If you specify the domain names to query but do not set the DomainSearchType parameter, the exact match mode is used.
	//
	// example:
	//
	// fuzzy_match
	DomainSearchType *string `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// 	- **configuring**
	//
	// 	- **configure_failed**
	//
	// 	- **checking**
	//
	// 	- **check_failed**
	//
	// 	- **stopping**
	//
	// 	- **deleting**
	//
	// If you do not set this parameter, domain names in all states are queried.
	//
	// example:
	//
	// configure_failed
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: **1*	- to **100000**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1 to 500**. Default value: **20**. Maximum value: **500**.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. By default, all IDs are queried.
	//
	// example:
	//
	// abcd1234abcd1234
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The information about the origin server.
	//
	// example:
	//
	// example.source.com
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of tags. Maximum number of elements in the list: 20
	Tag []*DescribeUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeUserDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsRequest) GetCdnType() *string {
	return s.CdnType
}

func (s *DescribeUserDomainsRequest) GetChangeEndTime() *string {
	return s.ChangeEndTime
}

func (s *DescribeUserDomainsRequest) GetChangeStartTime() *string {
	return s.ChangeStartTime
}

func (s *DescribeUserDomainsRequest) GetCheckDomainShow() *bool {
	return s.CheckDomainShow
}

func (s *DescribeUserDomainsRequest) GetCoverage() *string {
	return s.Coverage
}

func (s *DescribeUserDomainsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUserDomainsRequest) GetDomainSearchType() *string {
	return s.DomainSearchType
}

func (s *DescribeUserDomainsRequest) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeUserDomainsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeUserDomainsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUserDomainsRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeUserDomainsRequest) GetTag() []*DescribeUserDomainsRequestTag {
	return s.Tag
}

func (s *DescribeUserDomainsRequest) SetCdnType(v string) *DescribeUserDomainsRequest {
	s.CdnType = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetChangeEndTime(v string) *DescribeUserDomainsRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetChangeStartTime(v string) *DescribeUserDomainsRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetCheckDomainShow(v bool) *DescribeUserDomainsRequest {
	s.CheckDomainShow = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetCoverage(v string) *DescribeUserDomainsRequest {
	s.Coverage = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainName(v string) *DescribeUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainSearchType(v string) *DescribeUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainStatus(v string) *DescribeUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetOwnerId(v int64) *DescribeUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetPageNumber(v int32) *DescribeUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetPageSize(v int32) *DescribeUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetResourceGroupId(v string) *DescribeUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetSecurityToken(v string) *DescribeUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetSource(v string) *DescribeUserDomainsRequest {
	s.Source = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetTag(v []*DescribeUserDomainsRequestTag) *DescribeUserDomainsRequest {
	s.Tag = v
	return s
}

func (s *DescribeUserDomainsRequest) Validate() error {
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

type DescribeUserDomainsRequestTag struct {
	// The key of a tag.
	//
	// By default, all tag keys are queried.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// By default, all tag values are queried.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeUserDomainsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeUserDomainsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeUserDomainsRequestTag) SetKey(v string) *DescribeUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeUserDomainsRequestTag) SetValue(v string) *DescribeUserDomainsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeUserDomainsRequestTag) Validate() error {
	return dara.Validate(s)
}
