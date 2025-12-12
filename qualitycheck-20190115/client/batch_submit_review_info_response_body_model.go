// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSubmitReviewInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchSubmitReviewInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *BatchSubmitReviewInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchSubmitReviewInfoResponseBody
	GetMessage() *string
	SetMessages(v *BatchSubmitReviewInfoResponseBodyMessages) *BatchSubmitReviewInfoResponseBody
	GetMessages() *BatchSubmitReviewInfoResponseBodyMessages
	SetRequestId(v string) *BatchSubmitReviewInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchSubmitReviewInfoResponseBody
	GetSuccess() *bool
}

type BatchSubmitReviewInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *BatchSubmitReviewInfoResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchSubmitReviewInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitReviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSubmitReviewInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchSubmitReviewInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchSubmitReviewInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchSubmitReviewInfoResponseBody) GetMessages() *BatchSubmitReviewInfoResponseBodyMessages {
	return s.Messages
}

func (s *BatchSubmitReviewInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSubmitReviewInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchSubmitReviewInfoResponseBody) SetCode(v string) *BatchSubmitReviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetHttpStatusCode(v int32) *BatchSubmitReviewInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetMessage(v string) *BatchSubmitReviewInfoResponseBody {
	s.Message = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetMessages(v *BatchSubmitReviewInfoResponseBodyMessages) *BatchSubmitReviewInfoResponseBody {
	s.Messages = v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetRequestId(v string) *BatchSubmitReviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) SetSuccess(v bool) *BatchSubmitReviewInfoResponseBody {
	s.Success = &v
	return s
}

func (s *BatchSubmitReviewInfoResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchSubmitReviewInfoResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s BatchSubmitReviewInfoResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitReviewInfoResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *BatchSubmitReviewInfoResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *BatchSubmitReviewInfoResponseBodyMessages) SetMessage(v []*string) *BatchSubmitReviewInfoResponseBodyMessages {
	s.Message = v
	return s
}

func (s *BatchSubmitReviewInfoResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
