// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigUdpReflectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigUdpReflectResponseBody
	GetRequestId() *string
}

type ConfigUdpReflectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9EC62E89-BD30-4FCD-9CB8-FA53865FF0D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigUdpReflectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigUdpReflectResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigUdpReflectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigUdpReflectResponseBody) SetRequestId(v string) *ConfigUdpReflectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigUdpReflectResponseBody) Validate() error {
	return dara.Validate(s)
}
