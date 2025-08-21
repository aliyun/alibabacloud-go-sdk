// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteILMPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteILMPolicyResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteILMPolicyResponseBody
	GetResult() *bool
}

type DeleteILMPolicyResponseBody struct {
	// example:
	//
	// 694FDC20-0FDD-47C4-B921-BFF902FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteILMPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteILMPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteILMPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteILMPolicyResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteILMPolicyResponseBody) SetRequestId(v string) *DeleteILMPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteILMPolicyResponseBody) SetResult(v bool) *DeleteILMPolicyResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteILMPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
