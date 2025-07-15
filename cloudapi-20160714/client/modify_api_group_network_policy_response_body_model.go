// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupNetworkPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApiGroupNetworkPolicyResponseBody
	GetRequestId() *string
}

type ModifyApiGroupNetworkPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4E707B25-5119-5ACF-9D26-7D2A2762F05C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiGroupNetworkPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupNetworkPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupNetworkPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApiGroupNetworkPolicyResponseBody) SetRequestId(v string) *ModifyApiGroupNetworkPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
