// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclId(v string) *DeleteNetworkAclRequest
	GetNetworkAclId() *string
}

type DeleteNetworkAclRequest struct {
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-bp1lhl0taikrbgnh****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
}

func (s DeleteNetworkAclRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAclRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DeleteNetworkAclRequest) SetNetworkAclId(v string) *DeleteNetworkAclRequest {
	s.NetworkAclId = &v
	return s
}

func (s *DeleteNetworkAclRequest) Validate() error {
	return dara.Validate(s)
}
