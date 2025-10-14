// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategyAvailableConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainAddrPools(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetDomainAddrPools() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools
	SetIpv4AddrPools(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetIpv4AddrPools() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools
	SetIpv6AddrPools(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetIpv6AddrPools() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools
	SetLines(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetLines() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines
	SetRequestId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetRequestId() *string
	SetSelectedDomainLines(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetSelectedDomainLines() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines
	SetSelectedIpv4Lines(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetSelectedIpv4Lines() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines
	SetSelectedIpv6Lines(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetSelectedIpv6Lines() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines
	SetSuggestSetDefaultLine(v bool) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody
	GetSuggestSetDefaultLine() *bool
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBody struct {
	// The available address pools of the domain name type.
	DomainAddrPools *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools `json:"DomainAddrPools,omitempty" xml:"DomainAddrPools,omitempty" type:"Struct"`
	// The available address pools of the IPv4 type.
	Ipv4AddrPools *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools `json:"Ipv4AddrPools,omitempty" xml:"Ipv4AddrPools,omitempty" type:"Struct"`
	// The available address pools of the IPv6 type.
	Ipv6AddrPools *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools `json:"Ipv6AddrPools,omitempty" xml:"Ipv6AddrPools,omitempty" type:"Struct"`
	// The source regions.
	Lines *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C2851BA9-CE56-49AF-8D12-4FC6A49EE688
	RequestId           *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SelectedDomainLines *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines `json:"SelectedDomainLines,omitempty" xml:"SelectedDomainLines,omitempty" type:"Struct"`
	SelectedIpv4Lines   *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines   `json:"SelectedIpv4Lines,omitempty" xml:"SelectedIpv4Lines,omitempty" type:"Struct"`
	SelectedIpv6Lines   *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines   `json:"SelectedIpv6Lines,omitempty" xml:"SelectedIpv6Lines,omitempty" type:"Struct"`
	// Indicates whether we recommend that you set the source region to global.
	//
	// example:
	//
	// true
	SuggestSetDefaultLine *bool `json:"SuggestSetDefaultLine,omitempty" xml:"SuggestSetDefaultLine,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetDomainAddrPools() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools {
	return s.DomainAddrPools
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetIpv4AddrPools() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools {
	return s.Ipv4AddrPools
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetIpv6AddrPools() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools {
	return s.Ipv6AddrPools
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetLines() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines {
	return s.Lines
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetSelectedDomainLines() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines {
	return s.SelectedDomainLines
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetSelectedIpv4Lines() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines {
	return s.SelectedIpv4Lines
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetSelectedIpv6Lines() *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines {
	return s.SelectedIpv6Lines
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) GetSuggestSetDefaultLine() *bool {
	return s.SuggestSetDefaultLine
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetDomainAddrPools(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.DomainAddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetIpv4AddrPools(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.Ipv4AddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetIpv6AddrPools(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.Ipv6AddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetLines(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.Lines = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetRequestId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetSelectedDomainLines(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.SelectedDomainLines = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetSelectedIpv4Lines(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.SelectedIpv4Lines = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetSelectedIpv6Lines(v *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.SelectedIpv6Lines = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) SetSuggestSetDefaultLine(v bool) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody {
	s.SuggestSetDefaultLine = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBody) Validate() error {
	if s.DomainAddrPools != nil {
		if err := s.DomainAddrPools.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv4AddrPools != nil {
		if err := s.Ipv4AddrPools.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv6AddrPools != nil {
		if err := s.Ipv6AddrPools.Validate(); err != nil {
			return err
		}
	}
	if s.Lines != nil {
		if err := s.Lines.Validate(); err != nil {
			return err
		}
	}
	if s.SelectedDomainLines != nil {
		if err := s.SelectedDomainLines.Validate(); err != nil {
			return err
		}
	}
	if s.SelectedIpv4Lines != nil {
		if err := s.SelectedIpv4Lines.Validate(); err != nil {
			return err
		}
	}
	if s.SelectedIpv6Lines != nil {
		if err := s.SelectedIpv6Lines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools struct {
	DomainAddrPool []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool `json:"DomainAddrPool,omitempty" xml:"DomainAddrPool,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) GetDomainAddrPool() []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool {
	return s.DomainAddrPool
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) SetDomainAddrPool(v []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools {
	s.DomainAddrPool = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPools) Validate() error {
	if s.DomainAddrPool != nil {
		for _, item := range s.DomainAddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool struct {
	// The number of addresses in the address pool.
	//
	// example:
	//
	// 1
	AddrCount *int32 `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// pool1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// testpool
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) GetId() *string {
	return s.Id
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) SetId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool {
	s.Id = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) SetName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyDomainAddrPoolsDomainAddrPool) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools struct {
	Ipv4AddrPool []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool `json:"Ipv4AddrPool,omitempty" xml:"Ipv4AddrPool,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) GetIpv4AddrPool() []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool {
	return s.Ipv4AddrPool
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) SetIpv4AddrPool(v []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools {
	s.Ipv4AddrPool = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPools) Validate() error {
	if s.Ipv4AddrPool != nil {
		for _, item := range s.Ipv4AddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool struct {
	// The number of addresses in the address pool.
	//
	// example:
	//
	// 1
	AddrCount *int32 `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// pool1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// testpool
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) GetId() *string {
	return s.Id
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) SetId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool {
	s.Id = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) SetName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv4AddrPoolsIpv4AddrPool) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools struct {
	Ipv6AddrPool []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool `json:"Ipv6AddrPool,omitempty" xml:"Ipv6AddrPool,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) GetIpv6AddrPool() []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool {
	return s.Ipv6AddrPool
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) SetIpv6AddrPool(v []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools {
	s.Ipv6AddrPool = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPools) Validate() error {
	if s.Ipv6AddrPool != nil {
		for _, item := range s.Ipv6AddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool struct {
	// The number of addresses in the address pool.
	//
	// example:
	//
	// 1
	AddrCount *int32 `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// pool1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// testpool
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) GetId() *string {
	return s.Id
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) SetId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool {
	s.Id = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) SetName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyIpv6AddrPoolsIpv6AddrPool) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines struct {
	Line []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) GetLine() []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	return s.Line
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) SetLine(v []*DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines {
	s.Line = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLines) Validate() error {
	if s.Line != nil {
		for _, item := range s.Line {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine struct {
	// The parent line code of the source region. Leave it blank if no parent line exists.
	//
	// example:
	//
	// telecom
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	// The line name of the source region.
	//
	// example:
	//
	// default
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the source region group.
	//
	// example:
	//
	// global
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The line code of the source region.
	//
	// example:
	//
	// default
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The code of the source region group.
	//
	// example:
	//
	// global
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetFatherCode() *string {
	return s.FatherCode
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) GetLineName() *string {
	return s.LineName
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetFatherCode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.FatherCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetGroupCode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.GroupCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetGroupName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetLineCode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) SetLineName(v string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodyLinesLine) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines struct {
	SelectedDomainLine []*string `json:"SelectedDomainLine,omitempty" xml:"SelectedDomainLine,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines) GetSelectedDomainLine() []*string {
	return s.SelectedDomainLine
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines) SetSelectedDomainLine(v []*string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines {
	s.SelectedDomainLine = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedDomainLines) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines struct {
	SelectedIpv4Line []*string `json:"SelectedIpv4Line,omitempty" xml:"SelectedIpv4Line,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines) GetSelectedIpv4Line() []*string {
	return s.SelectedIpv4Line
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines) SetSelectedIpv4Line(v []*string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines {
	s.SelectedIpv4Line = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv4Lines) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines struct {
	SelectedIpv6Line []*string `json:"SelectedIpv6Line,omitempty" xml:"SelectedIpv6Line,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines) GetSelectedIpv6Line() []*string {
	return s.SelectedIpv6Line
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines) SetSelectedIpv6Line(v []*string) *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines {
	s.SelectedIpv6Line = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigResponseBodySelectedIpv6Lines) Validate() error {
	return dara.Validate(s)
}
