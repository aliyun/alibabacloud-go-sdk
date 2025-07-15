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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid Values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, `DRYRUN.SUCCESS` is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the VBR is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VBR failover group.
	//
	// example:
	//
	// vbrha-sa1sxheuxtd98****
	VbrHaId *string `json:"VbrHaId,omitempty" xml:"VbrHaId,omitempty"`
	// The VBR ID.
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
