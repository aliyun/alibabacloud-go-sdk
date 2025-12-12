// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignReviewerBySessionGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AssignReviewerBySessionGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AssignReviewerBySessionGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AssignReviewerBySessionGroupResponseBody
	GetMessage() *string
	SetMessages(v *AssignReviewerBySessionGroupResponseBodyMessages) *AssignReviewerBySessionGroupResponseBody
	GetMessages() *AssignReviewerBySessionGroupResponseBodyMessages
	SetRequestId(v string) *AssignReviewerBySessionGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssignReviewerBySessionGroupResponseBody
	GetSuccess() *bool
}

type AssignReviewerBySessionGroupResponseBody struct {
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
	Message  *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *AssignReviewerBySessionGroupResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignReviewerBySessionGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignReviewerBySessionGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AssignReviewerBySessionGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AssignReviewerBySessionGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AssignReviewerBySessionGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssignReviewerBySessionGroupResponseBody) GetMessages() *AssignReviewerBySessionGroupResponseBodyMessages {
	return s.Messages
}

func (s *AssignReviewerBySessionGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignReviewerBySessionGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssignReviewerBySessionGroupResponseBody) SetCode(v string) *AssignReviewerBySessionGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetHttpStatusCode(v int32) *AssignReviewerBySessionGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetMessage(v string) *AssignReviewerBySessionGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetMessages(v *AssignReviewerBySessionGroupResponseBodyMessages) *AssignReviewerBySessionGroupResponseBody {
	s.Messages = v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetRequestId(v string) *AssignReviewerBySessionGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) SetSuccess(v bool) *AssignReviewerBySessionGroupResponseBody {
	s.Success = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssignReviewerBySessionGroupResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s AssignReviewerBySessionGroupResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s AssignReviewerBySessionGroupResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *AssignReviewerBySessionGroupResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *AssignReviewerBySessionGroupResponseBodyMessages) SetMessage(v []*string) *AssignReviewerBySessionGroupResponseBodyMessages {
	s.Message = v
	return s
}

func (s *AssignReviewerBySessionGroupResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
