// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainGroups(v *DescribeDomainGroupsResponseBodyDomainGroups) *DescribeDomainGroupsResponseBody
	GetDomainGroups() *DescribeDomainGroupsResponseBodyDomainGroups
	SetPageNumber(v int64) *DescribeDomainGroupsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainGroupsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDomainGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDomainGroupsResponseBody
	GetTotalCount() *int64
}

type DescribeDomainGroupsResponseBody struct {
	// The domain name groups.
	DomainGroups *DescribeDomainGroupsResponseBodyDomainGroups `json:"DomainGroups,omitempty" xml:"DomainGroups,omitempty" type:"Struct"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsResponseBody) GetDomainGroups() *DescribeDomainGroupsResponseBodyDomainGroups {
	return s.DomainGroups
}

func (s *DescribeDomainGroupsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainGroupsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainGroupsResponseBody) SetDomainGroups(v *DescribeDomainGroupsResponseBodyDomainGroups) *DescribeDomainGroupsResponseBody {
	s.DomainGroups = v
	return s
}

func (s *DescribeDomainGroupsResponseBody) SetPageNumber(v int64) *DescribeDomainGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainGroupsResponseBody) SetPageSize(v int64) *DescribeDomainGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainGroupsResponseBody) SetRequestId(v string) *DescribeDomainGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainGroupsResponseBody) SetTotalCount(v int64) *DescribeDomainGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainGroupsResponseBody) Validate() error {
	if s.DomainGroups != nil {
		if err := s.DomainGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainGroupsResponseBodyDomainGroups struct {
	DomainGroup []*DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup `json:"DomainGroup,omitempty" xml:"DomainGroup,omitempty" type:"Repeated"`
}

func (s DescribeDomainGroupsResponseBodyDomainGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainGroupsResponseBodyDomainGroups) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsResponseBodyDomainGroups) GetDomainGroup() []*DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup {
	return s.DomainGroup
}

func (s *DescribeDomainGroupsResponseBodyDomainGroups) SetDomainGroup(v []*DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) *DescribeDomainGroupsResponseBodyDomainGroups {
	s.DomainGroup = v
	return s
}

func (s *DescribeDomainGroupsResponseBodyDomainGroups) Validate() error {
	if s.DomainGroup != nil {
		for _, item := range s.DomainGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup struct {
	// The number of domain name groups.
	//
	// example:
	//
	// 2
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The ID of the domain name group. Valid values:
	//
	// 	- defaultGroup: the default group
	//
	// 	- If an empty string is returned, it indicates the group that contains all domain names.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the domain name group.
	//
	// example:
	//
	// MyGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) SetDomainCount(v int64) *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup {
	s.DomainCount = &v
	return s
}

func (s *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) SetGroupId(v string) *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) SetGroupName(v string) *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeDomainGroupsResponseBodyDomainGroupsDomainGroup) Validate() error {
	return dara.Validate(s)
}
