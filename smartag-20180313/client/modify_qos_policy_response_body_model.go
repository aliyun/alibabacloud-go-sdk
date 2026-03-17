// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyQosPolicyResponseBody
	GetRequestId() *string
}

type ModifyQosPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 551CD836-9E46-4F2C-A167-B4363180A647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyQosPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQosPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyQosPolicyResponseBody) SetRequestId(v string) *ModifyQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQosPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
