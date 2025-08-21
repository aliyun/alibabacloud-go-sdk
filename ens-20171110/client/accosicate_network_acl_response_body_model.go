// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccosicateNetworkAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AccosicateNetworkAclResponseBody
	GetRequestId() *string
}

type AccosicateNetworkAclResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AccosicateNetworkAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccosicateNetworkAclResponseBody) GoString() string {
	return s.String()
}

func (s *AccosicateNetworkAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccosicateNetworkAclResponseBody) SetRequestId(v string) *AccosicateNetworkAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccosicateNetworkAclResponseBody) Validate() error {
	return dara.Validate(s)
}
