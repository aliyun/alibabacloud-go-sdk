// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *CreateAclPolicyResponseBody
	GetGatewayId() *string
	SetMessage(v string) *CreateAclPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAclPolicyResponseBody
	GetRequestId() *string
}

type CreateAclPolicyResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully add acl policy for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAclPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateAclPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAclPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAclPolicyResponseBody) SetGatewayId(v string) *CreateAclPolicyResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateAclPolicyResponseBody) SetMessage(v string) *CreateAclPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAclPolicyResponseBody) SetRequestId(v string) *CreateAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAclPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
