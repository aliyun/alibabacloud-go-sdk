// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRuleTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventRuleTargetsResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteEventRuleTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEventRuleTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventRuleTargetsResponseBody
	GetSuccess() *bool
}

type DeleteEventRuleTargetsResponseBody struct {
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 7ADD7EFB-7555-4EC1-A3D9-F9955C189CCF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRuleTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleTargetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventRuleTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventRuleTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventRuleTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventRuleTargetsResponseBody) SetCode(v string) *DeleteEventRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRuleTargetsResponseBody) SetMessage(v string) *DeleteEventRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventRuleTargetsResponseBody) SetRequestId(v string) *DeleteEventRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRuleTargetsResponseBody) SetSuccess(v bool) *DeleteEventRuleTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventRuleTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}
