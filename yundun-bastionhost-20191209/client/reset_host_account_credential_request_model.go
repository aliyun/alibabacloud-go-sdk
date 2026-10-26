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
	SetProjectId(v int64) *ResetHostAccountCredentialRequest
	GetProjectId() *int64
	SetRegionId(v string) *ResetHostAccountCredentialRequest
	GetRegionId() *string
}

type ResetHostAccountCredentialRequest struct {
	// The type of the logon credential to purge. Valid values:
	//
	// - **Password**: password
	//
	// - **PrivateKey**: SSH private key
	//
	// This parameter is required.
	//
	// example:
	//
	// Password
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The ID of the host account whose logon credentials you want to purge.
	//
	// > You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The instance ID of the Bastionhost instance to which the host account whose logon credentials you want to purge belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-rp640dg****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance to which the host account whose logon credentials you want to purge belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *ResetHostAccountCredentialRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ResetHostAccountCredentialRequest) SetProjectId(v int64) *ResetHostAccountCredentialRequest {
	s.ProjectId = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetRegionId(v string) *ResetHostAccountCredentialRequest {
	s.RegionId = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) Validate() error {
	return dara.Validate(s)
}
