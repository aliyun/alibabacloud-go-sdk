// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *DeleteAclPolicyResponseBody
	GetGatewayId() *string
	SetMessage(v string) *DeleteAclPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAclPolicyResponseBody
	GetRequestId() *string
}

type DeleteAclPolicyResponseBody struct {
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
	// Successfully delete acl policy for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAclPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeleteAclPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAclPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAclPolicyResponseBody) SetGatewayId(v string) *DeleteAclPolicyResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DeleteAclPolicyResponseBody) SetMessage(v string) *DeleteAclPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAclPolicyResponseBody) SetRequestId(v string) *DeleteAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAclPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
