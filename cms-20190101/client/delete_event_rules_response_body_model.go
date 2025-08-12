// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventRulesResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteEventRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEventRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventRulesResponseBody
	GetSuccess() *bool
}

type DeleteEventRulesResponseBody struct {
	// The status code.
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E5A72B5B-4F44-438C-B68A-147FD5DC53A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventRulesResponseBody) SetCode(v string) *DeleteEventRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRulesResponseBody) SetMessage(v string) *DeleteEventRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventRulesResponseBody) SetRequestId(v string) *DeleteEventRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRulesResponseBody) SetSuccess(v bool) *DeleteEventRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
