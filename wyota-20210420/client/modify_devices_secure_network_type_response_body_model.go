// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDevicesSecureNetworkTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDevicesSecureNetworkTypeResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyDevicesSecureNetworkTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDevicesSecureNetworkTypeResponseBody
	GetRequestId() *string
}

type ModifyDevicesSecureNetworkTypeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDevicesSecureNetworkTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDevicesSecureNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) SetCode(v string) *ModifyDevicesSecureNetworkTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) SetMessage(v string) *ModifyDevicesSecureNetworkTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) SetRequestId(v string) *ModifyDevicesSecureNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
