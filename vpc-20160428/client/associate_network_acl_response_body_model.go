// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateNetworkAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateNetworkAclResponseBody
	GetRequestId() *string
}

type AssociateNetworkAclResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4CF20CC7-D1FC-425B-A15B-DF7C8E2131A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateNetworkAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkAclResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateNetworkAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateNetworkAclResponseBody) SetRequestId(v string) *AssociateNetworkAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateNetworkAclResponseBody) Validate() error {
	return dara.Validate(s)
}
