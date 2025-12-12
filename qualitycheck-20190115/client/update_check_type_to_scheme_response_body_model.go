// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckTypeToSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCheckTypeToSchemeResponseBody
	GetCode() *string
	SetData(v int64) *UpdateCheckTypeToSchemeResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateCheckTypeToSchemeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCheckTypeToSchemeResponseBody
	GetMessage() *string
	SetMessages(v *UpdateCheckTypeToSchemeResponseBodyMessages) *UpdateCheckTypeToSchemeResponseBody
	GetMessages() *UpdateCheckTypeToSchemeResponseBodyMessages
	SetRequestId(v string) *UpdateCheckTypeToSchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCheckTypeToSchemeResponseBody
	GetSuccess() *bool
}

type UpdateCheckTypeToSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 4
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateCheckTypeToSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCheckTypeToSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckTypeToSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCheckTypeToSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCheckTypeToSchemeResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateCheckTypeToSchemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCheckTypeToSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCheckTypeToSchemeResponseBody) GetMessages() *UpdateCheckTypeToSchemeResponseBodyMessages {
	return s.Messages
}

func (s *UpdateCheckTypeToSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCheckTypeToSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetCode(v string) *UpdateCheckTypeToSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetData(v int64) *UpdateCheckTypeToSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetHttpStatusCode(v int32) *UpdateCheckTypeToSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetMessage(v string) *UpdateCheckTypeToSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetMessages(v *UpdateCheckTypeToSchemeResponseBodyMessages) *UpdateCheckTypeToSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetRequestId(v string) *UpdateCheckTypeToSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) SetSuccess(v bool) *UpdateCheckTypeToSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCheckTypeToSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateCheckTypeToSchemeResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckTypeToSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateCheckTypeToSchemeResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *UpdateCheckTypeToSchemeResponseBodyMessages) SetMessage(v []*string) *UpdateCheckTypeToSchemeResponseBodyMessages {
	s.Message = v
	return s
}

func (s *UpdateCheckTypeToSchemeResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
