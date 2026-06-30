// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterMulticastDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyTransitRouterMulticastDomainRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyTransitRouterMulticastDomainRequest
	GetDryRun() *bool
	SetOptions(v *ModifyTransitRouterMulticastDomainRequestOptions) *ModifyTransitRouterMulticastDomainRequest
	GetOptions() *ModifyTransitRouterMulticastDomainRequestOptions
	SetOwnerAccount(v string) *ModifyTransitRouterMulticastDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyTransitRouterMulticastDomainRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyTransitRouterMulticastDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTransitRouterMulticastDomainRequest
	GetResourceOwnerId() *int64
	SetTransitRouterMulticastDomainDescription(v string) *ModifyTransitRouterMulticastDomainRequest
	GetTransitRouterMulticastDomainDescription() *string
	SetTransitRouterMulticastDomainId(v string) *ModifyTransitRouterMulticastDomainRequest
	GetTransitRouterMulticastDomainId() *string
	SetTransitRouterMulticastDomainName(v string) *ModifyTransitRouterMulticastDomainRequest
	GetTransitRouterMulticastDomainName() *string
}

type ModifyTransitRouterMulticastDomainRequest struct {
	// A client token that ensures the idempotence of the request.
	//
	// Generate a unique token on your client for each request. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run. The system checks the required parameters, request format, and service limits. If the check fails, an error message is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends the request. If the request passes the check, the name and description of the multicast domain are modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The feature options of the multicast domain.
	Options              *ModifyTransitRouterMulticastDomainRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	OwnerAccount         *string                                           `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                           `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new description of the multicast domain.
	//
	// The description can be empty or 1 to 256 characters long. It cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// desctest
	TransitRouterMulticastDomainDescription *string `json:"TransitRouterMulticastDomainDescription,omitempty" xml:"TransitRouterMulticastDomainDescription,omitempty"`
	// The ID of the multicast domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-mcast-domain-40cwj0rgzgdtam****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
	// The new name of the multicast domain.
	//
	// The name can be empty or 1 to 128 characters long. It cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// nametest
	TransitRouterMulticastDomainName *string `json:"TransitRouterMulticastDomainName,omitempty" xml:"TransitRouterMulticastDomainName,omitempty"`
}

func (s ModifyTransitRouterMulticastDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterMulticastDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetOptions() *ModifyTransitRouterMulticastDomainRequestOptions {
	return s.Options
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetTransitRouterMulticastDomainDescription() *string {
	return s.TransitRouterMulticastDomainDescription
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *ModifyTransitRouterMulticastDomainRequest) GetTransitRouterMulticastDomainName() *string {
	return s.TransitRouterMulticastDomainName
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetClientToken(v string) *ModifyTransitRouterMulticastDomainRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetDryRun(v bool) *ModifyTransitRouterMulticastDomainRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetOptions(v *ModifyTransitRouterMulticastDomainRequestOptions) *ModifyTransitRouterMulticastDomainRequest {
	s.Options = v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetOwnerAccount(v string) *ModifyTransitRouterMulticastDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetOwnerId(v int64) *ModifyTransitRouterMulticastDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetResourceOwnerAccount(v string) *ModifyTransitRouterMulticastDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetResourceOwnerId(v int64) *ModifyTransitRouterMulticastDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetTransitRouterMulticastDomainDescription(v string) *ModifyTransitRouterMulticastDomainRequest {
	s.TransitRouterMulticastDomainDescription = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetTransitRouterMulticastDomainId(v string) *ModifyTransitRouterMulticastDomainRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) SetTransitRouterMulticastDomainName(v string) *ModifyTransitRouterMulticastDomainRequest {
	s.TransitRouterMulticastDomainName = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequest) Validate() error {
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyTransitRouterMulticastDomainRequestOptions struct {
	// Specifies whether to enable the Internet Group Management Protocol (IGMP) feature for the multicast domain. When this feature is enabled, hosts can use IGMP to dynamically join or leave multicast groups. Set the value to **enable**.
	//
	// > - The IGMP feature is in public preview. To use this feature, contact your account manager.
	//
	// >
	//
	// > - You cannot disable the IGMP feature after it is enabled.
	//
	// example:
	//
	// enable
	Igmpv2Support *string `json:"Igmpv2Support,omitempty" xml:"Igmpv2Support,omitempty"`
	// example:
	//
	// enable
	StrictSourceControl *string `json:"StrictSourceControl,omitempty" xml:"StrictSourceControl,omitempty"`
}

func (s ModifyTransitRouterMulticastDomainRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterMulticastDomainRequestOptions) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterMulticastDomainRequestOptions) GetIgmpv2Support() *string {
	return s.Igmpv2Support
}

func (s *ModifyTransitRouterMulticastDomainRequestOptions) GetStrictSourceControl() *string {
	return s.StrictSourceControl
}

func (s *ModifyTransitRouterMulticastDomainRequestOptions) SetIgmpv2Support(v string) *ModifyTransitRouterMulticastDomainRequestOptions {
	s.Igmpv2Support = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequestOptions) SetStrictSourceControl(v string) *ModifyTransitRouterMulticastDomainRequestOptions {
	s.StrictSourceControl = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainRequestOptions) Validate() error {
	return dara.Validate(s)
}
