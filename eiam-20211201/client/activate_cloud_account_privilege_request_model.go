// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateCloudAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ActivateCloudAccountPrivilegeRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *ActivateCloudAccountPrivilegeRequest
	GetCloudAccountId() *string
	SetInstanceId(v string) *ActivateCloudAccountPrivilegeRequest
	GetInstanceId() *string
}

type ActivateCloudAccountPrivilegeRequest struct {
	// Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. The ClientToken supports only ASCII characters and can be up to 64 characters in length.
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

func (s ActivateCloudAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateCloudAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ActivateCloudAccountPrivilegeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ActivateCloudAccountPrivilegeRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ActivateCloudAccountPrivilegeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ActivateCloudAccountPrivilegeRequest) SetClientToken(v string) *ActivateCloudAccountPrivilegeRequest {
	s.ClientToken = &v
	return s
}

func (s *ActivateCloudAccountPrivilegeRequest) SetCloudAccountId(v string) *ActivateCloudAccountPrivilegeRequest {
	s.CloudAccountId = &v
	return s
}

func (s *ActivateCloudAccountPrivilegeRequest) SetInstanceId(v string) *ActivateCloudAccountPrivilegeRequest {
	s.InstanceId = &v
	return s
}

func (s *ActivateCloudAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
