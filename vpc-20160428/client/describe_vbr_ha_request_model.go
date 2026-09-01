// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVbrHaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeVbrHaRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeVbrHaRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DescribeVbrHaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVbrHaRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVbrHaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVbrHaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVbrHaRequest
	GetResourceOwnerId() *int64
	SetVbrHaId(v string) *DescribeVbrHaRequest
	GetVbrHaId() *string
	SetVbrId(v string) *DescribeVbrHaRequest
	GetVbrId() *string
}

type DescribeVbrHaRequest struct {
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
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and instance status. If the check fails, the corresponding error is returned. If the check succeeds, `DRYRUN.SUCCESS` is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the check, the instance is started.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The VBR failover group instance ID. You must specify at least one of **VbrHaId*	- and **VbrId**. If both are omitted, the service returns MissingParam.VbrHaIdOrVbrId (400).
	//
	// example:
	//
	// vbrha-sa1sxheuxtd98****
	VbrHaId *string `json:"VbrHaId,omitempty" xml:"VbrHaId,omitempty"`
	// The VBR instance ID. You must specify at least one of **VbrId*	- and **VbrHaId**. If both are omitted, the service returns MissingParam.VbrHaIdOrVbrId (400).
	//
	// example:
	//
	// vbr-bp1jcg5cmxjbl9xgc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s DescribeVbrHaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVbrHaRequest) GoString() string {
	return s.String()
}

func (s *DescribeVbrHaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeVbrHaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeVbrHaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVbrHaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVbrHaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVbrHaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVbrHaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVbrHaRequest) GetVbrHaId() *string {
	return s.VbrHaId
}

func (s *DescribeVbrHaRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *DescribeVbrHaRequest) SetClientToken(v string) *DescribeVbrHaRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeVbrHaRequest) SetDryRun(v bool) *DescribeVbrHaRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeVbrHaRequest) SetOwnerAccount(v string) *DescribeVbrHaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVbrHaRequest) SetOwnerId(v int64) *DescribeVbrHaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVbrHaRequest) SetRegionId(v string) *DescribeVbrHaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVbrHaRequest) SetResourceOwnerAccount(v string) *DescribeVbrHaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVbrHaRequest) SetResourceOwnerId(v int64) *DescribeVbrHaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVbrHaRequest) SetVbrHaId(v string) *DescribeVbrHaRequest {
	s.VbrHaId = &v
	return s
}

func (s *DescribeVbrHaRequest) SetVbrId(v string) *DescribeVbrHaRequest {
	s.VbrId = &v
	return s
}

func (s *DescribeVbrHaRequest) Validate() error {
	return dara.Validate(s)
}
