// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecurityProxyResponseBody
	GetRequestId() *string
}

type DeleteSecurityProxyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7447795A-39AB-52CB-8F92-128DF4898F36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityProxyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityProxyResponseBody) SetRequestId(v string) *DeleteSecurityProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
