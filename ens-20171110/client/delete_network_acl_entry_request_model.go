// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAclEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclEntryId(v string) *DeleteNetworkAclEntryRequest
	GetNetworkAclEntryId() *string
}

type DeleteNetworkAclEntryRequest struct {
	// The ID of the network ACL for which you want to delete a rule.
	//
	// example:
	//
	// nae-5****
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
}

func (s DeleteNetworkAclEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAclEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAclEntryRequest) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *DeleteNetworkAclEntryRequest) SetNetworkAclEntryId(v string) *DeleteNetworkAclEntryRequest {
	s.NetworkAclEntryId = &v
	return s
}

func (s *DeleteNetworkAclEntryRequest) Validate() error {
	return dara.Validate(s)
}
