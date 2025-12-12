// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSchemeTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSchemeTaskConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateSchemeTaskConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSchemeTaskConfigResponseBody
	GetMessage() *string
	SetMessages(v *UpdateSchemeTaskConfigResponseBodyMessages) *UpdateSchemeTaskConfigResponseBody
	GetMessages() *UpdateSchemeTaskConfigResponseBodyMessages
	SetRequestId(v string) *UpdateSchemeTaskConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSchemeTaskConfigResponseBody
	GetSuccess() *bool
}

type UpdateSchemeTaskConfigResponseBody struct {
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
	Message  *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateSchemeTaskConfigResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSchemeTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSchemeTaskConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSchemeTaskConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSchemeTaskConfigResponseBody) GetMessages() *UpdateSchemeTaskConfigResponseBodyMessages {
	return s.Messages
}

func (s *UpdateSchemeTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSchemeTaskConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSchemeTaskConfigResponseBody) SetCode(v string) *UpdateSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetHttpStatusCode(v int32) *UpdateSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetMessage(v string) *UpdateSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetMessages(v *UpdateSchemeTaskConfigResponseBodyMessages) *UpdateSchemeTaskConfigResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetRequestId(v string) *UpdateSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) SetSuccess(v bool) *UpdateSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSchemeTaskConfigResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateSchemeTaskConfigResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemeTaskConfigResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *UpdateSchemeTaskConfigResponseBodyMessages) SetMessage(v []*string) *UpdateSchemeTaskConfigResponseBodyMessages {
	s.Message = v
	return s
}

func (s *UpdateSchemeTaskConfigResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
