// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsBySourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainsList(v *DescribeDomainsBySourceResponseBodyDomainsList) *DescribeDomainsBySourceResponseBody
	GetDomainsList() *DescribeDomainsBySourceResponseBodyDomainsList
	SetRequestId(v string) *DescribeDomainsBySourceResponseBody
	GetRequestId() *string
	SetSources(v string) *DescribeDomainsBySourceResponseBody
	GetSources() *string
}

type DescribeDomainsBySourceResponseBody struct {
	DomainsList *DescribeDomainsBySourceResponseBodyDomainsList `json:"DomainsList,omitempty" xml:"DomainsList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B0F074E5-A1AC-4B32-8EA2-6F450410D1E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The origin servers.
	//
	// example:
	//
	// example.com,aliyundoc.com
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s DescribeDomainsBySourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBody) GetDomainsList() *DescribeDomainsBySourceResponseBodyDomainsList {
	return s.DomainsList
}

func (s *DescribeDomainsBySourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainsBySourceResponseBody) GetSources() *string {
	return s.Sources
}

func (s *DescribeDomainsBySourceResponseBody) SetDomainsList(v *DescribeDomainsBySourceResponseBodyDomainsList) *DescribeDomainsBySourceResponseBody {
	s.DomainsList = v
	return s
}

func (s *DescribeDomainsBySourceResponseBody) SetRequestId(v string) *DescribeDomainsBySourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBody) SetSources(v string) *DescribeDomainsBySourceResponseBody {
	s.Sources = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBody) Validate() error {
	if s.DomainsList != nil {
		if err := s.DomainsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainsBySourceResponseBodyDomainsList struct {
	DomainsData []*DescribeDomainsBySourceResponseBodyDomainsListDomainsData `json:"DomainsData,omitempty" xml:"DomainsData,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsList) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsList) GetDomainsData() []*DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	return s.DomainsData
}

func (s *DescribeDomainsBySourceResponseBodyDomainsList) SetDomainsData(v []*DescribeDomainsBySourceResponseBodyDomainsListDomainsData) *DescribeDomainsBySourceResponseBodyDomainsList {
	s.DomainsData = v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsList) Validate() error {
	if s.DomainsData != nil {
		for _, item := range s.DomainsData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsData struct {
	DomainInfos *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos `json:"DomainInfos,omitempty" xml:"DomainInfos,omitempty" type:"Struct"`
	Domains     *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains     `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	Source      *string                                                               `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsData) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) GetDomainInfos() *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos {
	return s.DomainInfos
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) GetDomains() *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains {
	return s.Domains
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) GetSource() *string {
	return s.Source
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetDomainInfos(v *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.DomainInfos = v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetDomains(v *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.Domains = v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetSource(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.Source = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) Validate() error {
	if s.DomainInfos != nil {
		if err := s.DomainInfos.Validate(); err != nil {
			return err
		}
	}
	if s.Domains != nil {
		if err := s.Domains.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos struct {
	DomainInfo []*DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo `json:"domainInfo,omitempty" xml:"domainInfo,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) GetDomainInfo() []*DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	return s.DomainInfo
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) SetDomainInfo(v []*DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos {
	s.DomainInfo = v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) Validate() error {
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

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo struct {
	CdnType     *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainCname *string `json:"DomainCname,omitempty" xml:"DomainCname,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GetCdnType() *string {
	return s.CdnType
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GetDomainCname() *string {
	return s.DomainCname
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetCdnType(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.CdnType = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetCreateTime(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetDomainCname(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.DomainCname = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetDomainName(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetStatus(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.Status = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetUpdateTime(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains struct {
	DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) GetDomainNames() []*string {
	return s.DomainNames
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) SetDomainNames(v []*string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains {
	s.DomainNames = v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) Validate() error {
	return dara.Validate(s)
}
