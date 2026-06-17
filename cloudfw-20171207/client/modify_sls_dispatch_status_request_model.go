// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySlsDispatchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchValue(v string) *ModifySlsDispatchStatusRequest
	GetDispatchValue() *string
	SetEnableStatus(v bool) *ModifySlsDispatchStatusRequest
	GetEnableStatus() *bool
	SetFilterKeys(v string) *ModifySlsDispatchStatusRequest
	GetFilterKeys() *string
	SetNewRegionId(v string) *ModifySlsDispatchStatusRequest
	GetNewRegionId() *string
	SetSite(v string) *ModifySlsDispatchStatusRequest
	GetSite() *string
}

type ModifySlsDispatchStatusRequest struct {
	// The key for the log category. Valid values:
	//
	// **internet_log**
	//
	// **vpc_firewall_log**
	//
	// **nat_firewall_log**
	//
	// **ipv6_firewall_log**
	//
	// **dns_firewall_log**.
	//
	// example:
	//
	// internet_log
	DispatchValue *string `json:"DispatchValue,omitempty" xml:"DispatchValue,omitempty"`
	// Specifies whether to deliver logs. A value of \\`true\\` enables delivery, and \\`false\\` disables it.
	//
	// example:
	//
	// true
	EnableStatus *bool `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The supported filter conditions. Valid values:
	//
	// **attack**
	//
	// **acl**
	//
	// **other**.
	//
	// example:
	//
	// attack,acl
	FilterKeys *string `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-shanghai
	NewRegionId *string `json:"NewRegionId,omitempty" xml:"NewRegionId,omitempty"`
	// The site to modify. If the log version is 1, leave this parameter empty or set it to \\`global\\`. If the log version is 2, set this parameter to \\`cn\\` or \\`intl\\`.
	//
	// example:
	//
	// cn
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s ModifySlsDispatchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySlsDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifySlsDispatchStatusRequest) GetDispatchValue() *string {
	return s.DispatchValue
}

func (s *ModifySlsDispatchStatusRequest) GetEnableStatus() *bool {
	return s.EnableStatus
}

func (s *ModifySlsDispatchStatusRequest) GetFilterKeys() *string {
	return s.FilterKeys
}

func (s *ModifySlsDispatchStatusRequest) GetNewRegionId() *string {
	return s.NewRegionId
}

func (s *ModifySlsDispatchStatusRequest) GetSite() *string {
	return s.Site
}

func (s *ModifySlsDispatchStatusRequest) SetDispatchValue(v string) *ModifySlsDispatchStatusRequest {
	s.DispatchValue = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) SetEnableStatus(v bool) *ModifySlsDispatchStatusRequest {
	s.EnableStatus = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) SetFilterKeys(v string) *ModifySlsDispatchStatusRequest {
	s.FilterKeys = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) SetNewRegionId(v string) *ModifySlsDispatchStatusRequest {
	s.NewRegionId = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) SetSite(v string) *ModifySlsDispatchStatusRequest {
	s.Site = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) Validate() error {
	return dara.Validate(s)
}
