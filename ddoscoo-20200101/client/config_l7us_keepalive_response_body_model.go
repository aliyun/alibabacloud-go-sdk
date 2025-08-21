// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7UsKeepaliveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigL7UsKeepaliveResponseBody
	GetRequestId() *string
}

type ConfigL7UsKeepaliveResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6D48AED0-41DB-5D9B-B484-3B6AAD312AD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigL7UsKeepaliveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7UsKeepaliveResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigL7UsKeepaliveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigL7UsKeepaliveResponseBody) SetRequestId(v string) *ConfigL7UsKeepaliveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigL7UsKeepaliveResponseBody) Validate() error {
	return dara.Validate(s)
}
