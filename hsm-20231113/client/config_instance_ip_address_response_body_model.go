// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigInstanceIpAddressResponseBody
	GetRequestId() *string
}

type ConfigInstanceIpAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigInstanceIpAddressResponseBody) SetRequestId(v string) *ConfigInstanceIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigInstanceIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
