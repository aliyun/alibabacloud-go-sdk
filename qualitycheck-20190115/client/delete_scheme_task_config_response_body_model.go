// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemeTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSchemeTaskConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteSchemeTaskConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSchemeTaskConfigResponseBody
	GetMessage() *string
	SetMessages(v *DeleteSchemeTaskConfigResponseBodyMessages) *DeleteSchemeTaskConfigResponseBody
	GetMessages() *DeleteSchemeTaskConfigResponseBodyMessages
	SetRequestId(v string) *DeleteSchemeTaskConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSchemeTaskConfigResponseBody
	GetSuccess() *bool
}

type DeleteSchemeTaskConfigResponseBody struct {
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
	Messages *DeleteSchemeTaskConfigResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSchemeTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchemeTaskConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSchemeTaskConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSchemeTaskConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSchemeTaskConfigResponseBody) GetMessages() *DeleteSchemeTaskConfigResponseBodyMessages {
	return s.Messages
}

func (s *DeleteSchemeTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchemeTaskConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSchemeTaskConfigResponseBody) SetCode(v string) *DeleteSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetHttpStatusCode(v int32) *DeleteSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetMessage(v string) *DeleteSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetMessages(v *DeleteSchemeTaskConfigResponseBodyMessages) *DeleteSchemeTaskConfigResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetRequestId(v string) *DeleteSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) SetSuccess(v bool) *DeleteSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSchemeTaskConfigResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s DeleteSchemeTaskConfigResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemeTaskConfigResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *DeleteSchemeTaskConfigResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *DeleteSchemeTaskConfigResponseBodyMessages) SetMessage(v []*string) *DeleteSchemeTaskConfigResponseBodyMessages {
	s.Message = v
	return s
}

func (s *DeleteSchemeTaskConfigResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
