// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTopicResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTopicResponseBody
	GetSuccess() *bool
	SetTopicId(v int64) *CreateTopicResponseBody
	GetTopicId() *int64
}

type CreateTopicResponseBody struct {
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
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 217
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTopicResponseBody) GetTopicId() *int64 {
	return s.TopicId
}

func (s *CreateTopicResponseBody) SetCode(v string) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetHttpStatusCode(v int32) *CreateTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTopicResponseBody) SetMessage(v string) *CreateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTopicResponseBody) SetTopicId(v int64) *CreateTopicResponseBody {
	s.TopicId = &v
	return s
}

func (s *CreateTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
