// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPausePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPausePolicyResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyPausePolicyResponseBody
	GetResult() map[string]interface{}
}

type ModifyPausePolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0B1FF998-BB8D-5182-BFC0-E471AA77095A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyPausePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPausePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPausePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPausePolicyResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyPausePolicyResponseBody) SetRequestId(v string) *ModifyPausePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPausePolicyResponseBody) SetResult(v map[string]interface{}) *ModifyPausePolicyResponseBody {
	s.Result = v
	return s
}

func (s *ModifyPausePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
