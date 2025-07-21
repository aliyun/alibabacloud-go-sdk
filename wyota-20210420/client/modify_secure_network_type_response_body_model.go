// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecureNetworkTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySecureNetworkTypeResponseBody
	GetCode() *string
	SetMessage(v string) *ModifySecureNetworkTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySecureNetworkTypeResponseBody
	GetRequestId() *string
}

type ModifySecureNetworkTypeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecureNetworkTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecureNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecureNetworkTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySecureNetworkTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySecureNetworkTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecureNetworkTypeResponseBody) SetCode(v string) *ModifySecureNetworkTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySecureNetworkTypeResponseBody) SetMessage(v string) *ModifySecureNetworkTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySecureNetworkTypeResponseBody) SetRequestId(v string) *ModifySecureNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecureNetworkTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
