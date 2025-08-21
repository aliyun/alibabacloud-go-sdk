// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAclEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclEntryId(v string) *CreateNetworkAclEntryResponseBody
	GetNetworkAclEntryId() *string
	SetRequestId(v string) *CreateNetworkAclEntryResponseBody
	GetRequestId() *string
}

type CreateNetworkAclEntryResponseBody struct {
	// The ID of the network ACL.
	//
	// example:
	//
	// nae-5****
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkAclEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclEntryResponseBody) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *CreateNetworkAclEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkAclEntryResponseBody) SetNetworkAclEntryId(v string) *CreateNetworkAclEntryResponseBody {
	s.NetworkAclEntryId = &v
	return s
}

func (s *CreateNetworkAclEntryResponseBody) SetRequestId(v string) *CreateNetworkAclEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkAclEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
