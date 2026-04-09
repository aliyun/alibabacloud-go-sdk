// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainInfos(v *ListDomainsResponseBodyDomainInfos) *ListDomainsResponseBody
	GetDomainInfos() *ListDomainsResponseBodyDomainInfos
	SetPageNumber(v int64) *ListDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDomainsResponseBody
	GetTotalCount() *int64
}

type ListDomainsResponseBody struct {
	DomainInfos *ListDomainsResponseBodyDomainInfos `json:"DomainInfos,omitempty" xml:"DomainInfos,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-7427E531****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 18
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) GetDomainInfos() *ListDomainsResponseBodyDomainInfos {
	return s.DomainInfos
}

func (s *ListDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDomainsResponseBody) SetDomainInfos(v *ListDomainsResponseBodyDomainInfos) *ListDomainsResponseBody {
	s.DomainInfos = v
	return s
}

func (s *ListDomainsResponseBody) SetPageNumber(v int64) *ListDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsResponseBody) SetPageSize(v int64) *ListDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainsResponseBody) SetTotalCount(v int64) *ListDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDomainsResponseBody) Validate() error {
	if s.DomainInfos != nil {
		if err := s.DomainInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDomainsResponseBodyDomainInfos struct {
	DomainInfo []*ListDomainsResponseBodyDomainInfosDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Repeated"`
}

func (s ListDomainsResponseBodyDomainInfos) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBodyDomainInfos) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomainInfos) GetDomainInfo() []*ListDomainsResponseBodyDomainInfosDomainInfo {
	return s.DomainInfo
}

func (s *ListDomainsResponseBodyDomainInfos) SetDomainInfo(v []*ListDomainsResponseBodyDomainInfosDomainInfo) *ListDomainsResponseBodyDomainInfos {
	s.DomainInfo = v
	return s
}

func (s *ListDomainsResponseBodyDomainInfos) Validate() error {
	if s.DomainInfo != nil {
		for _, item := range s.DomainInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainsResponseBodyDomainInfosDomainInfo struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ResolveHttpAes  *int64  `json:"ResolveHttpAes,omitempty" xml:"ResolveHttpAes,omitempty"`
	ResolveHttpsAes *int64  `json:"ResolveHttpsAes,omitempty" xml:"ResolveHttpsAes,omitempty"`
	Resolved        *int64  `json:"Resolved,omitempty" xml:"Resolved,omitempty"`
	Resolved6       *int64  `json:"Resolved6,omitempty" xml:"Resolved6,omitempty"`
	ResolvedDoh     *int64  `json:"ResolvedDoh,omitempty" xml:"ResolvedDoh,omitempty"`
	ResolvedHttps   *int64  `json:"ResolvedHttps,omitempty" xml:"ResolvedHttps,omitempty"`
	ResolvedHttps6  *int64  `json:"ResolvedHttps6,omitempty" xml:"ResolvedHttps6,omitempty"`
	TimeModified    *int64  `json:"TimeModified,omitempty" xml:"TimeModified,omitempty"`
}

func (s ListDomainsResponseBodyDomainInfosDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBodyDomainInfosDomainInfo) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetResolveHttpAes() *int64 {
	return s.ResolveHttpAes
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetResolveHttpsAes() *int64 {
	return s.ResolveHttpsAes
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetResolved() *int64 {
	return s.Resolved
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetResolved6() *int64 {
	return s.Resolved6
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetResolvedDoh() *int64 {
	return s.ResolvedDoh
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetResolvedHttps() *int64 {
	return s.ResolvedHttps
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetResolvedHttps6() *int64 {
	return s.ResolvedHttps6
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) GetTimeModified() *int64 {
	return s.TimeModified
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetDomainName(v string) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.DomainName = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolveHttpAes(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.ResolveHttpAes = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolveHttpsAes(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.ResolveHttpsAes = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolved(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.Resolved = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolved6(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.Resolved6 = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolvedDoh(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.ResolvedDoh = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolvedHttps(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.ResolvedHttps = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolvedHttps6(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.ResolvedHttps6 = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetTimeModified(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.TimeModified = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) Validate() error {
	return dara.Validate(s)
}
