// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateNetworkCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateNetworkCidrBlockResponseBody
	GetRequestId() *string
}

type UnassociateNetworkCidrBlockResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateNetworkCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateNetworkCidrBlockResponseBody) SetRequestId(v string) *UnassociateNetworkCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateNetworkCidrBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
