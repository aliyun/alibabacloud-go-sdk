// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemeTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSchemeTaskConfigResponseBody
	GetCode() *string
	SetData(v int64) *CreateSchemeTaskConfigResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateSchemeTaskConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSchemeTaskConfigResponseBody
	GetMessage() *string
	SetMessages(v *CreateSchemeTaskConfigResponseBodyMessages) *CreateSchemeTaskConfigResponseBody
	GetMessages() *CreateSchemeTaskConfigResponseBodyMessages
	SetRequestId(v string) *CreateSchemeTaskConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSchemeTaskConfigResponseBody
	GetSuccess() *bool
}

type CreateSchemeTaskConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *CreateSchemeTaskConfigResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSchemeTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchemeTaskConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSchemeTaskConfigResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateSchemeTaskConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSchemeTaskConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSchemeTaskConfigResponseBody) GetMessages() *CreateSchemeTaskConfigResponseBodyMessages {
	return s.Messages
}

func (s *CreateSchemeTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSchemeTaskConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSchemeTaskConfigResponseBody) SetCode(v string) *CreateSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetData(v int64) *CreateSchemeTaskConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetHttpStatusCode(v int32) *CreateSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetMessage(v string) *CreateSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetMessages(v *CreateSchemeTaskConfigResponseBodyMessages) *CreateSchemeTaskConfigResponseBody {
	s.Messages = v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetRequestId(v string) *CreateSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) SetSuccess(v bool) *CreateSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSchemeTaskConfigResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSchemeTaskConfigResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s CreateSchemeTaskConfigResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemeTaskConfigResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateSchemeTaskConfigResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *CreateSchemeTaskConfigResponseBodyMessages) SetMessage(v []*string) *CreateSchemeTaskConfigResponseBodyMessages {
	s.Message = v
	return s
}

func (s *CreateSchemeTaskConfigResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
