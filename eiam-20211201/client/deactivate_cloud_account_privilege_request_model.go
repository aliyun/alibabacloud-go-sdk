// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateCloudAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeactivateCloudAccountPrivilegeRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *DeactivateCloudAccountPrivilegeRequest
	GetCloudAccountId() *string
	SetInstanceId(v string) *DeactivateCloudAccountPrivilegeRequest
	GetInstanceId() *string
}

type DeactivateCloudAccountPrivilegeRequest struct {
	// Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cloud account ID.
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

func (s DeactivateCloudAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactivateCloudAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *DeactivateCloudAccountPrivilegeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeactivateCloudAccountPrivilegeRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *DeactivateCloudAccountPrivilegeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeactivateCloudAccountPrivilegeRequest) SetClientToken(v string) *DeactivateCloudAccountPrivilegeRequest {
	s.ClientToken = &v
	return s
}

func (s *DeactivateCloudAccountPrivilegeRequest) SetCloudAccountId(v string) *DeactivateCloudAccountPrivilegeRequest {
	s.CloudAccountId = &v
	return s
}

func (s *DeactivateCloudAccountPrivilegeRequest) SetInstanceId(v string) *DeactivateCloudAccountPrivilegeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeactivateCloudAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
