// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCloudAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableCloudAccountPrivilegeRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *DisableCloudAccountPrivilegeRequest
	GetCloudAccountId() *string
	SetInstanceId(v string) *DisableCloudAccountPrivilegeRequest
	GetInstanceId() *string
}

type DisableCloudAccountPrivilegeRequest struct {
	// Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac_mkv7rgt4d7i4u7zqtzev2mxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableCloudAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCloudAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *DisableCloudAccountPrivilegeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableCloudAccountPrivilegeRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *DisableCloudAccountPrivilegeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableCloudAccountPrivilegeRequest) SetClientToken(v string) *DisableCloudAccountPrivilegeRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableCloudAccountPrivilegeRequest) SetCloudAccountId(v string) *DisableCloudAccountPrivilegeRequest {
	s.CloudAccountId = &v
	return s
}

func (s *DisableCloudAccountPrivilegeRequest) SetInstanceId(v string) *DisableCloudAccountPrivilegeRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableCloudAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
