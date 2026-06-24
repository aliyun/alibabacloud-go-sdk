// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdvancedSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAdvancedSettingResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateAdvancedSettingResponseBody
	GetResult() *bool
}

type UpdateAdvancedSettingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// - true: The garbage collector configuration is changed.
	//
	// - false: The garbage collector configuration failed to be changed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAdvancedSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdvancedSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAdvancedSettingResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateAdvancedSettingResponseBody) SetRequestId(v string) *UpdateAdvancedSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdvancedSettingResponseBody) SetResult(v bool) *UpdateAdvancedSettingResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateAdvancedSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
