// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
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
	SetRequestId(v string) *GetOriginProtectionResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetOriginProtectionResponseBody
	GetSiteId() *int64
}

type GetOriginProtectionResponseBody struct {
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
	OriginProtection *string `json:"OriginProtection,omitempty" xml:"OriginProtection,omitempty"`
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

func (s *GetOriginProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginProtectionResponseBody) GetSiteId() *int64 {
	return s.SiteId
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
