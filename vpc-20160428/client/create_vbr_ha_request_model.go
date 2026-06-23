// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVbrHaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVbrHaRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVbrHaRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateVbrHaRequest
	GetDryRun() *bool
	SetName(v string) *CreateVbrHaRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateVbrHaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVbrHaRequest
	GetOwnerId() *int64
	SetPeerVbrId(v string) *CreateVbrHaRequest
	GetPeerVbrId() *string
	SetRegionId(v string) *CreateVbrHaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVbrHaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVbrHaRequest
	GetResourceOwnerId() *int64
	SetVbrId(v string) *CreateVbrHaRequest
	GetVbrId() *string
}

type CreateVbrHaRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the VBR failover group.
	//
	// The description must be 2 to 256 characters in length and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// VBRHa
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and instance status. If the check fails, the corresponding error is returned. If the check succeeds, `DRYRUN.SUCCESS` is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the check, the instance is started.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the VBR failover group.
	//
	// example:
	//
	// VBRHa
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the other VBR in the VBR failover group.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp12mw1f8k3jgygk9****
	PeerVbrId *string `json:"PeerVbrId,omitempty" xml:"PeerVbrId,omitempty"`
	// The region ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The instance ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp1jcg5cmxjbl9xgc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s CreateVbrHaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVbrHaRequest) GoString() string {
	return s.String()
}

func (s *CreateVbrHaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVbrHaRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVbrHaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVbrHaRequest) GetName() *string {
	return s.Name
}

func (s *CreateVbrHaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVbrHaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVbrHaRequest) GetPeerVbrId() *string {
	return s.PeerVbrId
}

func (s *CreateVbrHaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVbrHaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVbrHaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVbrHaRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *CreateVbrHaRequest) SetClientToken(v string) *CreateVbrHaRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVbrHaRequest) SetDescription(v string) *CreateVbrHaRequest {
	s.Description = &v
	return s
}

func (s *CreateVbrHaRequest) SetDryRun(v bool) *CreateVbrHaRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVbrHaRequest) SetName(v string) *CreateVbrHaRequest {
	s.Name = &v
	return s
}

func (s *CreateVbrHaRequest) SetOwnerAccount(v string) *CreateVbrHaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVbrHaRequest) SetOwnerId(v int64) *CreateVbrHaRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVbrHaRequest) SetPeerVbrId(v string) *CreateVbrHaRequest {
	s.PeerVbrId = &v
	return s
}

func (s *CreateVbrHaRequest) SetRegionId(v string) *CreateVbrHaRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVbrHaRequest) SetResourceOwnerAccount(v string) *CreateVbrHaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVbrHaRequest) SetResourceOwnerId(v int64) *CreateVbrHaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVbrHaRequest) SetVbrId(v string) *CreateVbrHaRequest {
	s.VbrId = &v
	return s
}

func (s *CreateVbrHaRequest) Validate() error {
	return dara.Validate(s)
}
