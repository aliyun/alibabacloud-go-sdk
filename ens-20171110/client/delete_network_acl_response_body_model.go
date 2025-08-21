// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkAclResponseBody
	GetRequestId() *string
}

type DeleteNetworkAclResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkAclResponseBody) SetRequestId(v string) *DeleteNetworkAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkAclResponseBody) Validate() error {
	return dara.Validate(s)
}
