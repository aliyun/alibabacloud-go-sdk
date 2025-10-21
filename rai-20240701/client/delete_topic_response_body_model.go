// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTopicResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTopicResponseBody
	GetSuccess() *bool
}

type DeleteTopicResponseBody struct {
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
}

func (s DeleteTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTopicResponseBody) SetCode(v string) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetHttpStatusCode(v int32) *DeleteTopicResponseBody {
	s.HttpStatusCode = &v
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

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
