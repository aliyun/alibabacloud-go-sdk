// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateTopicConfigResponseBody
	GetCode() *int64
	SetData(v string) *UpdateTopicConfigResponseBody
	GetData() *string
	SetMessage(v string) *UpdateTopicConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTopicConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTopicConfigResponseBody
	GetSuccess() *bool
}

type UpdateTopicConfigResponseBody struct {
	// The HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// []
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0178A3A7-E87B-5E50-A16F-3E62F534****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTopicConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTopicConfigResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateTopicConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateTopicConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTopicConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTopicConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTopicConfigResponseBody) SetCode(v int64) *UpdateTopicConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) SetData(v string) *UpdateTopicConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) SetMessage(v string) *UpdateTopicConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) SetRequestId(v string) *UpdateTopicConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) SetSuccess(v bool) *UpdateTopicConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
