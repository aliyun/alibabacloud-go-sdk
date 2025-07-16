// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachGatewayDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *AttachGatewayDomainResponseBody
	GetGatewayId() *string
	SetMessage(v string) *AttachGatewayDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *AttachGatewayDomainResponseBody
	GetRequestId() *string
}

type AttachGatewayDomainResponseBody struct {
	// The ID of the private gateway. To obtain the private gateway ID, see the GatewayId parameter in the response parameters of the [ListGateway](https://apiworkbench.aliyun-inc.com/document/eas/2021-07-01/ListGateway?spm=openapi-amp.newDocPublishment.0.0.765e281fL2IcjJ\\&ampEnv=online) operation.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The error message.
	//
	// example:
	//
	// Successfully update custom endpoint for gateway gw-1uhcqmsc7x22******
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachGatewayDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachGatewayDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AttachGatewayDomainResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *AttachGatewayDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AttachGatewayDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachGatewayDomainResponseBody) SetGatewayId(v string) *AttachGatewayDomainResponseBody {
	s.GatewayId = &v
	return s
}

func (s *AttachGatewayDomainResponseBody) SetMessage(v string) *AttachGatewayDomainResponseBody {
	s.Message = &v
	return s
}

func (s *AttachGatewayDomainResponseBody) SetRequestId(v string) *AttachGatewayDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachGatewayDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
