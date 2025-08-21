// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateNetworkAclRequest
	GetDescription() *string
	SetNetworkAclName(v string) *CreateNetworkAclRequest
	GetNetworkAclName() *string
}

type CreateNetworkAclRequest struct {
	// The description of the network ACL.
	//
	// The description must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// This is my NetworkAcl.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Enter a name for the network ACL.
	//
	// The name must be 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// acl-1
	NetworkAclName *string `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
}

func (s CreateNetworkAclRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkAclRequest) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *CreateNetworkAclRequest) SetDescription(v string) *CreateNetworkAclRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkAclRequest) SetNetworkAclName(v string) *CreateNetworkAclRequest {
	s.NetworkAclName = &v
	return s
}

func (s *CreateNetworkAclRequest) Validate() error {
	return dara.Validate(s)
}
