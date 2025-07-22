// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentFunctionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppAgentFunctionStatusResponseBody
	GetRequestId() *string
}

type ModifyAppAgentFunctionStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppAgentFunctionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentFunctionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentFunctionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppAgentFunctionStatusResponseBody) SetRequestId(v string) *ModifyAppAgentFunctionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppAgentFunctionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
