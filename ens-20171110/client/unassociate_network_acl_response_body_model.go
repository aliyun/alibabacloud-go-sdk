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
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
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
