// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPublicIpAddressPoolServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenPublicIpAddressPoolServiceResponseBody
	GetCode() *string
	SetMessage(v string) *OpenPublicIpAddressPoolServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenPublicIpAddressPoolServiceResponseBody
	GetRequestId() *string
}

type OpenPublicIpAddressPoolServiceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 28CF47AB-B6C0-5FA2-80C7-2B37726A92CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenPublicIpAddressPoolServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenPublicIpAddressPoolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenPublicIpAddressPoolServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenPublicIpAddressPoolServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenPublicIpAddressPoolServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenPublicIpAddressPoolServiceResponseBody) SetCode(v string) *OpenPublicIpAddressPoolServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceResponseBody) SetMessage(v string) *OpenPublicIpAddressPoolServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceResponseBody) SetRequestId(v string) *OpenPublicIpAddressPoolServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
