// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTopicResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTopicResponseBody
	GetSuccess() *bool
}

type UpdateTopicResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****F68692
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTopicResponseBody) SetCode(v string) *UpdateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTopicResponseBody) SetHttpStatusCode(v int32) *UpdateTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTopicResponseBody) SetMessage(v string) *UpdateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTopicResponseBody) SetRequestId(v string) *UpdateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTopicResponseBody) SetSuccess(v bool) *UpdateTopicResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
