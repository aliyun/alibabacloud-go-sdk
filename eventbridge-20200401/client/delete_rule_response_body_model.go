// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRuleResponseBody
	GetSuccess() *bool
}

type DeleteRuleResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// The event rule not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FF942675-F937-549C-A942-EB94FFE28DD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRuleResponseBody) SetCode(v string) *DeleteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRuleResponseBody) SetMessage(v string) *DeleteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleResponseBody) SetSuccess(v bool) *DeleteRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
