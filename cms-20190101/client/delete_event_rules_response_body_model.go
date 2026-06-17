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
	// The status code. A value of 200 indicates that the request was successful.
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
	// The ID of the request. You can use this ID to query logs.
	//
	// example:
	//
	// 45231A42-5A09-5AFF-953C-A5B3D4ED8925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
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
