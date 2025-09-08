// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteCustomizeRuleResponseBody
	GetCode() *int32
	SetData(v int32) *DeleteCustomizeRuleResponseBody
	GetData() *int32
	SetMessage(v string) *DeleteCustomizeRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomizeRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomizeRuleResponseBody
	GetSuccess() *bool
}

type DeleteCustomizeRuleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomizeRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteCustomizeRuleResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteCustomizeRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomizeRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomizeRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomizeRuleResponseBody) SetCode(v int32) *DeleteCustomizeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetData(v int32) *DeleteCustomizeRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetMessage(v string) *DeleteCustomizeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetRequestId(v string) *DeleteCustomizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetSuccess(v bool) *DeleteCustomizeRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
