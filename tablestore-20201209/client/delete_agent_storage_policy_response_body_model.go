// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentStoragePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAgentStoragePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAgentStoragePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAgentStoragePolicyResponseBody
	GetRequestId() *string
}

type DeleteAgentStoragePolicyResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgentStoragePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentStoragePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentStoragePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAgentStoragePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentStoragePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentStoragePolicyResponseBody) SetCode(v string) *DeleteAgentStoragePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentStoragePolicyResponseBody) SetMessage(v string) *DeleteAgentStoragePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentStoragePolicyResponseBody) SetRequestId(v string) *DeleteAgentStoragePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentStoragePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
