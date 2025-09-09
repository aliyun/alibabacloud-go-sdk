// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptVpcPeerConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AcceptVpcPeerConnectionRequest
	GetClientToken() *string
	SetDryRun(v bool) *AcceptVpcPeerConnectionRequest
	GetDryRun() *bool
	SetInstanceId(v string) *AcceptVpcPeerConnectionRequest
	GetInstanceId() *string
	SetResourceGroupId(v string) *AcceptVpcPeerConnectionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AcceptVpcPeerConnectionRequest
	GetResourceOwnerAccount() *string
	SetTag(v []*AcceptVpcPeerConnectionRequestTag) *AcceptVpcPeerConnectionRequest
	GetTag() []*AcceptVpcPeerConnectionRequestTag
}

type AcceptVpcPeerConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the VPC peering connection to be accepted by the accepter VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcc-guzvyqlj0n6e10****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a resource group?](https://help.aliyun.com/document_detail/94475.html)
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The tags.
	Tag []*AcceptVpcPeerConnectionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AcceptVpcPeerConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *AcceptVpcPeerConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AcceptVpcPeerConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AcceptVpcPeerConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AcceptVpcPeerConnectionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AcceptVpcPeerConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AcceptVpcPeerConnectionRequest) GetTag() []*AcceptVpcPeerConnectionRequestTag {
	return s.Tag
}

func (s *AcceptVpcPeerConnectionRequest) SetClientToken(v string) *AcceptVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetDryRun(v bool) *AcceptVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetInstanceId(v string) *AcceptVpcPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetResourceGroupId(v string) *AcceptVpcPeerConnectionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetResourceOwnerAccount(v string) *AcceptVpcPeerConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) SetTag(v []*AcceptVpcPeerConnectionRequestTag) *AcceptVpcPeerConnectionRequest {
	s.Tag = v
	return s
}

func (s *AcceptVpcPeerConnectionRequest) Validate() error {
	return dara.Validate(s)
}

type AcceptVpcPeerConnectionRequestTag struct {
	// The tag key. You must specify at least one tag key and at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `acs:` or `aliyun` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You must specify at least one tag value and can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AcceptVpcPeerConnectionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AcceptVpcPeerConnectionRequestTag) GoString() string {
	return s.String()
}

func (s *AcceptVpcPeerConnectionRequestTag) GetKey() *string {
	return s.Key
}

func (s *AcceptVpcPeerConnectionRequestTag) GetValue() *string {
	return s.Value
}

func (s *AcceptVpcPeerConnectionRequestTag) SetKey(v string) *AcceptVpcPeerConnectionRequestTag {
	s.Key = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequestTag) SetValue(v string) *AcceptVpcPeerConnectionRequestTag {
	s.Value = &v
	return s
}

func (s *AcceptVpcPeerConnectionRequestTag) Validate() error {
	return dara.Validate(s)
}
