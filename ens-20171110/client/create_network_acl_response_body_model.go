// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclId(v string) *CreateNetworkAclResponseBody
	GetNetworkAclId() *string
	SetRequestId(v string) *CreateNetworkAclResponseBody
	GetRequestId() *string
}

type CreateNetworkAclResponseBody struct {
	// The ID of the network ACL.
	//
	// example:
	//
	// nacl-5p1fg655nh68xyz9i****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBody) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *CreateNetworkAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkAclResponseBody) SetNetworkAclId(v string) *CreateNetworkAclResponseBody {
	s.NetworkAclId = &v
	return s
}

func (s *CreateNetworkAclResponseBody) SetRequestId(v string) *CreateNetworkAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkAclResponseBody) Validate() error {
	return dara.Validate(s)
}
