// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteTopicResponseBody
	GetCode() *int64
	SetData(v map[string]interface{}) *DeleteTopicResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *DeleteTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTopicResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteTopicResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteTopicResponseBody
	GetSuccess() *bool
}

type DeleteTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteTopicResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *DeleteTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTopicResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTopicResponseBody) SetCode(v int64) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetData(v map[string]interface{}) *DeleteTopicResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTopicResponseBody) SetMessage(v string) *DeleteTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTopicResponseBody) SetRequestId(v string) *DeleteTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTopicResponseBody) SetStatus(v string) *DeleteTopicResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
