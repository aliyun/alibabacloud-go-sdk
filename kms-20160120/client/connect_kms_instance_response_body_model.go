// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectKmsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConnectKmsInstanceResponseBody
	GetRequestId() *string
}

type ConnectKmsInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// d3eca5c8-a856-4347-8eb6-e1898c3fda2e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConnectKmsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConnectKmsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConnectKmsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConnectKmsInstanceResponseBody) SetRequestId(v string) *ConnectKmsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConnectKmsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
