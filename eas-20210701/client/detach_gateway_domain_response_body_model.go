// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGatewayDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *DetachGatewayDomainResponseBody
	GetGatewayId() *string
	SetMessage(v string) *DetachGatewayDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachGatewayDomainResponseBody
	GetRequestId() *string
}

type DetachGatewayDomainResponseBody struct {
	// The ID of the private gateway. To obtain the private gateway ID, see the GatewayId parameter in the response parameters of the [ListGateway](https://apiworkbench.aliyun-inc.com/document/eas/2021-07-01/ListGateway?spm=openapi-amp.newDocPublishment.0.0.765e281fL2IcjJ\\&ampEnv=online) operation.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// Successfully delete custom endpoint for gateway gw-1uhcqmsc7x22******
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachGatewayDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachGatewayDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DetachGatewayDomainResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DetachGatewayDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachGatewayDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachGatewayDomainResponseBody) SetGatewayId(v string) *DetachGatewayDomainResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DetachGatewayDomainResponseBody) SetMessage(v string) *DetachGatewayDomainResponseBody {
	s.Message = &v
	return s
}

func (s *DetachGatewayDomainResponseBody) SetRequestId(v string) *DetachGatewayDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachGatewayDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
