// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHostAccountCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialType(v string) *ResetHostAccountCredentialRequest
	GetCredentialType() *string
	SetHostAccountId(v string) *ResetHostAccountCredentialRequest
	GetHostAccountId() *string
	SetInstanceId(v string) *ResetHostAccountCredentialRequest
	GetInstanceId() *string
	SetRegionId(v string) *ResetHostAccountCredentialRequest
	GetRegionId() *string
}

type ResetHostAccountCredentialRequest struct {
	// The type of the logon credential that you want to delete. Valid values:
	//
	// 	- **Password**
	//
	// 	- **PrivateKey**
	//
	// This parameter is required.
	//
	// example:
	//
	// Password
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The ID of the host account for which the logon credential is to be deleted.
	//
	// >  You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the ID of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the bastion host from which you want to delete the logon credential for the host account.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host from which you want to delete the logon credential for the host account.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetHostAccountCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetHostAccountCredentialRequest) GoString() string {
	return s.String()
}

func (s *ResetHostAccountCredentialRequest) GetCredentialType() *string {
	return s.CredentialType
}

func (s *ResetHostAccountCredentialRequest) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *ResetHostAccountCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetHostAccountCredentialRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetHostAccountCredentialRequest) SetCredentialType(v string) *ResetHostAccountCredentialRequest {
	s.CredentialType = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetHostAccountId(v string) *ResetHostAccountCredentialRequest {
	s.HostAccountId = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetInstanceId(v string) *ResetHostAccountCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetRegionId(v string) *ResetHostAccountCredentialRequest {
	s.RegionId = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) Validate() error {
	return dara.Validate(s)
}
