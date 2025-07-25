// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVpcIpamServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenVpcIpamServiceResponseBody
	GetCode() *string
	SetMessage(v string) *OpenVpcIpamServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenVpcIpamServiceResponseBody
	GetRequestId() *string
}

type OpenVpcIpamServiceResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Information returned upon successful IPAM activation.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 3F814C37-B032-5477-AF5A-2925D0593CD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenVpcIpamServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenVpcIpamServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenVpcIpamServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenVpcIpamServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenVpcIpamServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenVpcIpamServiceResponseBody) SetCode(v string) *OpenVpcIpamServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenVpcIpamServiceResponseBody) SetMessage(v string) *OpenVpcIpamServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenVpcIpamServiceResponseBody) SetRequestId(v string) *OpenVpcIpamServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenVpcIpamServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
