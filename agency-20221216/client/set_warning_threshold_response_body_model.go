// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWarningThresholdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetWarningThresholdResponseBody
	GetCode() *string
	SetMessage(v string) *SetWarningThresholdResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetWarningThresholdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetWarningThresholdResponseBody
	GetSuccess() *bool
}

type SetWarningThresholdResponseBody struct {
	// The result code. Valid values:
	//
	// - 200: OK.
	//
	// - 1109: System error.
	//
	// - 3040: The sub-account is frozen and the operation cannot be completed.
	//
	// - 3044: The alert percentage value is not a number.
	//
	// - 3045: The alert percentage value must be from 1 to 100.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that corresponds to the Code parameter value.
	//
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. A unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the current API call was successful. Valid values: True and False. This does not indicate whether the subsequent business operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetWarningThresholdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetWarningThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetWarningThresholdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetWarningThresholdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetWarningThresholdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetWarningThresholdResponseBody) SetCode(v string) *SetWarningThresholdResponseBody {
	s.Code = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetMessage(v string) *SetWarningThresholdResponseBody {
	s.Message = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetRequestId(v string) *SetWarningThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetSuccess(v bool) *SetWarningThresholdResponseBody {
	s.Success = &v
	return s
}

func (s *SetWarningThresholdResponseBody) Validate() error {
	return dara.Validate(s)
}
