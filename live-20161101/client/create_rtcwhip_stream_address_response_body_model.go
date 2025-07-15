// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRTCWhipStreamAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRTCWhipStreamAddressResponseBody
	GetRequestId() *string
	SetWhipAddress(v string) *CreateRTCWhipStreamAddressResponseBody
	GetWhipAddress() *string
}

type CreateRTCWhipStreamAddressResponseBody struct {
	// example:
	//
	// 58E7**D4-xxxx-xxxx-xxxx-6B5**6Cxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// https://xxxxxx.media-sh.xxxxx.com/xxxxxxxxxxxx/3723a3xxxxxxxxx223c606b***5f7a2bc7c56ea5cdd0xxxxe?auth_key=17495xxxxx-xxxx-0-f013003067c78c4053f9cd0xxxxxxx&qqzr=H4sIAAAAAAAC_6pWSlayUkrOMxxxxxxrPyCxQ0lFKQTCLlKyUjM2Nj**NDQ1TEy2Mjc3xxxxxxxxxxqRYAAAD__xxxxx__xxxxxxxx
	WhipAddress *string `json:"WhipAddress,omitempty" xml:"WhipAddress,omitempty"`
}

func (s CreateRTCWhipStreamAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRTCWhipStreamAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRTCWhipStreamAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRTCWhipStreamAddressResponseBody) GetWhipAddress() *string {
	return s.WhipAddress
}

func (s *CreateRTCWhipStreamAddressResponseBody) SetRequestId(v string) *CreateRTCWhipStreamAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRTCWhipStreamAddressResponseBody) SetWhipAddress(v string) *CreateRTCWhipStreamAddressResponseBody {
	s.WhipAddress = &v
	return s
}

func (s *CreateRTCWhipStreamAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
