// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoConfirmIPList(v string) *GetOriginProtectionResponseBody
	GetAutoConfirmIPList() *string
	SetCurrentIPWhitelist(v *GetOriginProtectionResponseBodyCurrentIPWhitelist) *GetOriginProtectionResponseBody
	GetCurrentIPWhitelist() *GetOriginProtectionResponseBodyCurrentIPWhitelist
	SetDiffIPWhitelist(v *GetOriginProtectionResponseBodyDiffIPWhitelist) *GetOriginProtectionResponseBody
	GetDiffIPWhitelist() *GetOriginProtectionResponseBodyDiffIPWhitelist
	SetLatestIPWhitelist(v *GetOriginProtectionResponseBodyLatestIPWhitelist) *GetOriginProtectionResponseBody
	GetLatestIPWhitelist() *GetOriginProtectionResponseBodyLatestIPWhitelist
	SetNeedUpdate(v bool) *GetOriginProtectionResponseBody
	GetNeedUpdate() *bool
	SetOriginConverge(v string) *GetOriginProtectionResponseBody
	GetOriginConverge() *string
	SetOriginProtection(v string) *GetOriginProtectionResponseBody
	GetOriginProtection() *string
	SetRegionalCurrentIPWhitelist(v *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) *GetOriginProtectionResponseBody
	GetRegionalCurrentIPWhitelist() *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist
	SetRegionalDiffIPWhitelist(v *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) *GetOriginProtectionResponseBody
	GetRegionalDiffIPWhitelist() *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist
	SetRegionalLatestIPWhitelist(v *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) *GetOriginProtectionResponseBody
	GetRegionalLatestIPWhitelist() *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist
	SetRequestId(v string) *GetOriginProtectionResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetOriginProtectionResponseBody
	GetSiteId() *int64
}

type GetOriginProtectionResponseBody struct {
	// example:
	//
	// off
	AutoConfirmIPList *string `json:"AutoConfirmIPList,omitempty" xml:"AutoConfirmIPList,omitempty"`
	// The IP whitelist for origin protection used by the website.
	CurrentIPWhitelist *GetOriginProtectionResponseBodyCurrentIPWhitelist `json:"CurrentIPWhitelist,omitempty" xml:"CurrentIPWhitelist,omitempty" type:"Struct"`
	// The IP whitelist for origin protection that has been updated.
	DiffIPWhitelist *GetOriginProtectionResponseBodyDiffIPWhitelist `json:"DiffIPWhitelist,omitempty" xml:"DiffIPWhitelist,omitempty" type:"Struct"`
	// The latest IP whitelist for origin protection.
	LatestIPWhitelist *GetOriginProtectionResponseBodyLatestIPWhitelist `json:"LatestIPWhitelist,omitempty" xml:"LatestIPWhitelist,omitempty" type:"Struct"`
	// Indicates whether the IP whitelist for origin protection needs to be updated. If the currently used IP whitelist is different from the latest IP whitelist, it needs to be updated, and the value is true.
	//
	// 	- true: The update is required.
	//
	// 	- false: No update is required.
	//
	// example:
	//
	// true
	NeedUpdate *bool `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	// Indicates whether IP convergence is enabled.
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	OriginConverge *string `json:"OriginConverge,omitempty" xml:"OriginConverge,omitempty"`
	// Indicates whether origin protection is enabled.
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	OriginProtection           *string                                                    `json:"OriginProtection,omitempty" xml:"OriginProtection,omitempty"`
	RegionalCurrentIPWhitelist *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist `json:"RegionalCurrentIPWhitelist,omitempty" xml:"RegionalCurrentIPWhitelist,omitempty" type:"Struct"`
	RegionalDiffIPWhitelist    *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist    `json:"RegionalDiffIPWhitelist,omitempty" xml:"RegionalDiffIPWhitelist,omitempty" type:"Struct"`
	RegionalLatestIPWhitelist  *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist  `json:"RegionalLatestIPWhitelist,omitempty" xml:"RegionalLatestIPWhitelist,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBody) GetAutoConfirmIPList() *string {
	return s.AutoConfirmIPList
}

func (s *GetOriginProtectionResponseBody) GetCurrentIPWhitelist() *GetOriginProtectionResponseBodyCurrentIPWhitelist {
	return s.CurrentIPWhitelist
}

func (s *GetOriginProtectionResponseBody) GetDiffIPWhitelist() *GetOriginProtectionResponseBodyDiffIPWhitelist {
	return s.DiffIPWhitelist
}

func (s *GetOriginProtectionResponseBody) GetLatestIPWhitelist() *GetOriginProtectionResponseBodyLatestIPWhitelist {
	return s.LatestIPWhitelist
}

func (s *GetOriginProtectionResponseBody) GetNeedUpdate() *bool {
	return s.NeedUpdate
}

func (s *GetOriginProtectionResponseBody) GetOriginConverge() *string {
	return s.OriginConverge
}

func (s *GetOriginProtectionResponseBody) GetOriginProtection() *string {
	return s.OriginProtection
}

func (s *GetOriginProtectionResponseBody) GetRegionalCurrentIPWhitelist() *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist {
	return s.RegionalCurrentIPWhitelist
}

func (s *GetOriginProtectionResponseBody) GetRegionalDiffIPWhitelist() *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist {
	return s.RegionalDiffIPWhitelist
}

func (s *GetOriginProtectionResponseBody) GetRegionalLatestIPWhitelist() *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist {
	return s.RegionalLatestIPWhitelist
}

func (s *GetOriginProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginProtectionResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginProtectionResponseBody) SetAutoConfirmIPList(v string) *GetOriginProtectionResponseBody {
	s.AutoConfirmIPList = &v
	return s
}

func (s *GetOriginProtectionResponseBody) SetCurrentIPWhitelist(v *GetOriginProtectionResponseBodyCurrentIPWhitelist) *GetOriginProtectionResponseBody {
	s.CurrentIPWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBody) SetDiffIPWhitelist(v *GetOriginProtectionResponseBodyDiffIPWhitelist) *GetOriginProtectionResponseBody {
	s.DiffIPWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBody) SetLatestIPWhitelist(v *GetOriginProtectionResponseBodyLatestIPWhitelist) *GetOriginProtectionResponseBody {
	s.LatestIPWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBody) SetNeedUpdate(v bool) *GetOriginProtectionResponseBody {
	s.NeedUpdate = &v
	return s
}

func (s *GetOriginProtectionResponseBody) SetOriginConverge(v string) *GetOriginProtectionResponseBody {
	s.OriginConverge = &v
	return s
}

func (s *GetOriginProtectionResponseBody) SetOriginProtection(v string) *GetOriginProtectionResponseBody {
	s.OriginProtection = &v
	return s
}

func (s *GetOriginProtectionResponseBody) SetRegionalCurrentIPWhitelist(v *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) *GetOriginProtectionResponseBody {
	s.RegionalCurrentIPWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBody) SetRegionalDiffIPWhitelist(v *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) *GetOriginProtectionResponseBody {
	s.RegionalDiffIPWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBody) SetRegionalLatestIPWhitelist(v *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) *GetOriginProtectionResponseBody {
	s.RegionalLatestIPWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBody) SetRequestId(v string) *GetOriginProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginProtectionResponseBody) SetSiteId(v int64) *GetOriginProtectionResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetOriginProtectionResponseBody) Validate() error {
	if s.CurrentIPWhitelist != nil {
		if err := s.CurrentIPWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.DiffIPWhitelist != nil {
		if err := s.DiffIPWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.LatestIPWhitelist != nil {
		if err := s.LatestIPWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.RegionalCurrentIPWhitelist != nil {
		if err := s.RegionalCurrentIPWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.RegionalDiffIPWhitelist != nil {
		if err := s.RegionalDiffIPWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.RegionalLatestIPWhitelist != nil {
		if err := s.RegionalLatestIPWhitelist.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginProtectionResponseBodyCurrentIPWhitelist struct {
	// The IP whitelist for origin protection used by the website, specified as IPv4 addresses or CIDR blocks.
	IPv4 []*string `json:"IPv4,omitempty" xml:"IPv4,omitempty" type:"Repeated"`
	// The IP whitelist for origin protection used by the website, specified as IPv6 addresses or CIDR blocks.
	IPv6 []*string `json:"IPv6,omitempty" xml:"IPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyCurrentIPWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyCurrentIPWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyCurrentIPWhitelist) GetIPv4() []*string {
	return s.IPv4
}

func (s *GetOriginProtectionResponseBodyCurrentIPWhitelist) GetIPv6() []*string {
	return s.IPv6
}

func (s *GetOriginProtectionResponseBodyCurrentIPWhitelist) SetIPv4(v []*string) *GetOriginProtectionResponseBodyCurrentIPWhitelist {
	s.IPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyCurrentIPWhitelist) SetIPv6(v []*string) *GetOriginProtectionResponseBodyCurrentIPWhitelist {
	s.IPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyCurrentIPWhitelist) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyDiffIPWhitelist struct {
	// The new IP whitelist for origin protection.
	AddedIPWhitelist *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist `json:"AddedIPWhitelist,omitempty" xml:"AddedIPWhitelist,omitempty" type:"Struct"`
	// The IP whitelist for origin protection that remains unchanged.
	NoChangeIpWhitelist *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist `json:"NoChangeIpWhitelist,omitempty" xml:"NoChangeIpWhitelist,omitempty" type:"Struct"`
	// The IP whitelist for origin protection that has been deleted.
	RemovedIPWhitelist *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist `json:"RemovedIPWhitelist,omitempty" xml:"RemovedIPWhitelist,omitempty" type:"Struct"`
}

func (s GetOriginProtectionResponseBodyDiffIPWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyDiffIPWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelist) GetAddedIPWhitelist() *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist {
	return s.AddedIPWhitelist
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelist) GetNoChangeIpWhitelist() *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist {
	return s.NoChangeIpWhitelist
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelist) GetRemovedIPWhitelist() *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist {
	return s.RemovedIPWhitelist
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelist) SetAddedIPWhitelist(v *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist) *GetOriginProtectionResponseBodyDiffIPWhitelist {
	s.AddedIPWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelist) SetNoChangeIpWhitelist(v *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist) *GetOriginProtectionResponseBodyDiffIPWhitelist {
	s.NoChangeIpWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelist) SetRemovedIPWhitelist(v *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist) *GetOriginProtectionResponseBodyDiffIPWhitelist {
	s.RemovedIPWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelist) Validate() error {
	if s.AddedIPWhitelist != nil {
		if err := s.AddedIPWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.NoChangeIpWhitelist != nil {
		if err := s.NoChangeIpWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.RemovedIPWhitelist != nil {
		if err := s.RemovedIPWhitelist.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist struct {
	// The IP whitelist for origin protection, specified as IPv4 addresses or CIDR blocks.
	IPv4 []*string `json:"IPv4,omitempty" xml:"IPv4,omitempty" type:"Repeated"`
	// The IP whitelist for origin protection, specified as IPv6 addresses or CIDR blocks.
	IPv6 []*string `json:"IPv6,omitempty" xml:"IPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist) GetIPv4() []*string {
	return s.IPv4
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist) GetIPv6() []*string {
	return s.IPv6
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist) SetIPv4(v []*string) *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist {
	s.IPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist) SetIPv6(v []*string) *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist {
	s.IPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistAddedIPWhitelist) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist struct {
	// The IP whitelist for origin protection, specified as IPv4 addresses or CIDR blocks.
	IPv4 []*string `json:"IPv4,omitempty" xml:"IPv4,omitempty" type:"Repeated"`
	// The IP whitelist for origin protection, specified as IPv6 addresses or CIDR blocks.
	IPv6 []*string `json:"IPv6,omitempty" xml:"IPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist) GetIPv4() []*string {
	return s.IPv4
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist) GetIPv6() []*string {
	return s.IPv6
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist) SetIPv4(v []*string) *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist {
	s.IPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist) SetIPv6(v []*string) *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist {
	s.IPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistNoChangeIpWhitelist) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist struct {
	// The IP whitelist for origin protection, specified as IPv4 addresses or CIDR blocks.
	IPv4 []*string `json:"IPv4,omitempty" xml:"IPv4,omitempty" type:"Repeated"`
	// The IP whitelist for origin protection, specified as IPv6 addresses or CIDR blocks.
	IPv6 []*string `json:"IPv6,omitempty" xml:"IPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist) GetIPv4() []*string {
	return s.IPv4
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist) GetIPv6() []*string {
	return s.IPv6
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist) SetIPv4(v []*string) *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist {
	s.IPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist) SetIPv6(v []*string) *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist {
	s.IPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyDiffIPWhitelistRemovedIPWhitelist) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyLatestIPWhitelist struct {
	// The latest IP whitelist for origin protection, specified as IPv4 addresses or CIDR blocks.
	IPv4 []*string `json:"IPv4,omitempty" xml:"IPv4,omitempty" type:"Repeated"`
	// The latest IP whitelist for origin protection, specified as IPv6 addresses or CIDR blocks.
	IPv6 []*string `json:"IPv6,omitempty" xml:"IPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyLatestIPWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyLatestIPWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyLatestIPWhitelist) GetIPv4() []*string {
	return s.IPv4
}

func (s *GetOriginProtectionResponseBodyLatestIPWhitelist) GetIPv6() []*string {
	return s.IPv6
}

func (s *GetOriginProtectionResponseBodyLatestIPWhitelist) SetIPv4(v []*string) *GetOriginProtectionResponseBodyLatestIPWhitelist {
	s.IPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyLatestIPWhitelist) SetIPv6(v []*string) *GetOriginProtectionResponseBodyLatestIPWhitelist {
	s.IPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyLatestIPWhitelist) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist struct {
	RegionalIPv4 []*GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4 `json:"RegionalIPv4,omitempty" xml:"RegionalIPv4,omitempty" type:"Repeated"`
	RegionalIPv6 []*GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6 `json:"RegionalIPv6,omitempty" xml:"RegionalIPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) GetRegionalIPv4() []*GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4 {
	return s.RegionalIPv4
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) GetRegionalIPv6() []*GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6 {
	return s.RegionalIPv6
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) SetRegionalIPv4(v []*GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4) *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist {
	s.RegionalIPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) SetRegionalIPv6(v []*GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6) *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist {
	s.RegionalIPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelist) Validate() error {
	if s.RegionalIPv4 != nil {
		for _, item := range s.RegionalIPv4 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegionalIPv6 != nil {
		for _, item := range s.RegionalIPv6 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv4) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalCurrentIPWhitelistRegionalIPv6) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelist struct {
	AddedIPRegionWhitelist   *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist   `json:"AddedIPRegionWhitelist,omitempty" xml:"AddedIPRegionWhitelist,omitempty" type:"Struct"`
	NoChangeIpWhitelist      *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist      `json:"NoChangeIpWhitelist,omitempty" xml:"NoChangeIpWhitelist,omitempty" type:"Struct"`
	RemovedIPRegionWhitelist *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist `json:"RemovedIPRegionWhitelist,omitempty" xml:"RemovedIPRegionWhitelist,omitempty" type:"Struct"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) GetAddedIPRegionWhitelist() *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist {
	return s.AddedIPRegionWhitelist
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) GetNoChangeIpWhitelist() *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist {
	return s.NoChangeIpWhitelist
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) GetRemovedIPRegionWhitelist() *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist {
	return s.RemovedIPRegionWhitelist
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) SetAddedIPRegionWhitelist(v *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist {
	s.AddedIPRegionWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) SetNoChangeIpWhitelist(v *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist {
	s.NoChangeIpWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) SetRemovedIPRegionWhitelist(v *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist {
	s.RemovedIPRegionWhitelist = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelist) Validate() error {
	if s.AddedIPRegionWhitelist != nil {
		if err := s.AddedIPRegionWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.NoChangeIpWhitelist != nil {
		if err := s.NoChangeIpWhitelist.Validate(); err != nil {
			return err
		}
	}
	if s.RemovedIPRegionWhitelist != nil {
		if err := s.RemovedIPRegionWhitelist.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist struct {
	RegionalIPv4 []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4 `json:"RegionalIPv4,omitempty" xml:"RegionalIPv4,omitempty" type:"Repeated"`
	RegionalIPv6 []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6 `json:"RegionalIPv6,omitempty" xml:"RegionalIPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist) GetRegionalIPv4() []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4 {
	return s.RegionalIPv4
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist) GetRegionalIPv6() []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6 {
	return s.RegionalIPv6
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist) SetRegionalIPv4(v []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist {
	s.RegionalIPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist) SetRegionalIPv6(v []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist {
	s.RegionalIPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelist) Validate() error {
	if s.RegionalIPv4 != nil {
		for _, item := range s.RegionalIPv4 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegionalIPv6 != nil {
		for _, item := range s.RegionalIPv6 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv4) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistAddedIPRegionWhitelistRegionalIPv6) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist struct {
	RegionalIPv4 []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4 `json:"RegionalIPv4,omitempty" xml:"RegionalIPv4,omitempty" type:"Repeated"`
	RegionalIPv6 []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6 `json:"RegionalIPv6,omitempty" xml:"RegionalIPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist) GetRegionalIPv4() []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4 {
	return s.RegionalIPv4
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist) GetRegionalIPv6() []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6 {
	return s.RegionalIPv6
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist) SetRegionalIPv4(v []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist {
	s.RegionalIPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist) SetRegionalIPv6(v []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist {
	s.RegionalIPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelist) Validate() error {
	if s.RegionalIPv4 != nil {
		for _, item := range s.RegionalIPv4 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegionalIPv6 != nil {
		for _, item := range s.RegionalIPv6 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv4) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistNoChangeIpWhitelistRegionalIPv6) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist struct {
	RegionalIPv4 []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4 `json:"RegionalIPv4,omitempty" xml:"RegionalIPv4,omitempty" type:"Repeated"`
	RegionalIPv6 []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6 `json:"RegionalIPv6,omitempty" xml:"RegionalIPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist) GetRegionalIPv4() []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4 {
	return s.RegionalIPv4
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist) GetRegionalIPv6() []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6 {
	return s.RegionalIPv6
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist) SetRegionalIPv4(v []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist {
	s.RegionalIPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist) SetRegionalIPv6(v []*GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist {
	s.RegionalIPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelist) Validate() error {
	if s.RegionalIPv4 != nil {
		for _, item := range s.RegionalIPv4 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegionalIPv6 != nil {
		for _, item := range s.RegionalIPv6 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv4) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalDiffIPWhitelistRemovedIPRegionWhitelistRegionalIPv6) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalLatestIPWhitelist struct {
	RegionalIPv4 []*GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4 `json:"RegionalIPv4,omitempty" xml:"RegionalIPv4,omitempty" type:"Repeated"`
	RegionalIPv6 []*GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6 `json:"RegionalIPv6,omitempty" xml:"RegionalIPv6,omitempty" type:"Repeated"`
}

func (s GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) GetRegionalIPv4() []*GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4 {
	return s.RegionalIPv4
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) GetRegionalIPv6() []*GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6 {
	return s.RegionalIPv6
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) SetRegionalIPv4(v []*GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4) *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist {
	s.RegionalIPv4 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) SetRegionalIPv6(v []*GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6) *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist {
	s.RegionalIPv6 = v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelist) Validate() error {
	if s.RegionalIPv4 != nil {
		for _, item := range s.RegionalIPv4 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegionalIPv6 != nil {
		for _, item := range s.RegionalIPv6 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv4) Validate() error {
	return dara.Validate(s)
}

type GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6 struct {
	// example:
	//
	// 101.66.250.0/25
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// chinese_mainland
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6) GetCidr() *string {
	return s.Cidr
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6) GetRegion() *string {
	return s.Region
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6) SetCidr(v string) *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6 {
	s.Cidr = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6) SetRegion(v string) *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6 {
	s.Region = &v
	return s
}

func (s *GetOriginProtectionResponseBodyRegionalLatestIPWhitelistRegionalIPv6) Validate() error {
	return dara.Validate(s)
}
