// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7RsPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigL7RsPolicyResponseBody
	GetRequestId() *string
}

type ConfigL7RsPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigL7RsPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7RsPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigL7RsPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigL7RsPolicyResponseBody) SetRequestId(v string) *ConfigL7RsPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigL7RsPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
