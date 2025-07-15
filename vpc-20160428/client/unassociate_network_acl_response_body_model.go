// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateNetworkAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateNetworkAclResponseBody
	GetRequestId() *string
}

type UnassociateNetworkAclResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AD024BAA-2D91-48FD-810B-8FB7489B6EE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateNetworkAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkAclResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateNetworkAclResponseBody) SetRequestId(v string) *UnassociateNetworkAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateNetworkAclResponseBody) Validate() error {
	return dara.Validate(s)
}
