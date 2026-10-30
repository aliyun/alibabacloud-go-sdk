// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCopilotSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCopilotSessionResponseBody
	GetCode() *string
	SetData(v *DeleteCopilotSessionResponseBodyData) *DeleteCopilotSessionResponseBody
	GetData() *DeleteCopilotSessionResponseBodyData
	SetMessage(v string) *DeleteCopilotSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCopilotSessionResponseBody
	GetRequestId() *string
}

type DeleteCopilotSessionResponseBody struct {
	// example:
	//
	// Success
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeleteCopilotSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5BC79EF5-798B-5499-9683-4E14EBBBF712
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteCopilotSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCopilotSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCopilotSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCopilotSessionResponseBody) GetData() *DeleteCopilotSessionResponseBodyData {
	return s.Data
}

func (s *DeleteCopilotSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCopilotSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCopilotSessionResponseBody) SetCode(v string) *DeleteCopilotSessionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCopilotSessionResponseBody) SetData(v *DeleteCopilotSessionResponseBodyData) *DeleteCopilotSessionResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCopilotSessionResponseBody) SetMessage(v string) *DeleteCopilotSessionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCopilotSessionResponseBody) SetRequestId(v string) *DeleteCopilotSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCopilotSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCopilotSessionResponseBodyData struct {
	// example:
	//
	// true
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// c356bc310e4345c3a602a2912322ad21
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s DeleteCopilotSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteCopilotSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCopilotSessionResponseBodyData) GetDeleted() *bool {
	return s.Deleted
}

func (s *DeleteCopilotSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteCopilotSessionResponseBodyData) SetDeleted(v bool) *DeleteCopilotSessionResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *DeleteCopilotSessionResponseBodyData) SetSessionId(v string) *DeleteCopilotSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *DeleteCopilotSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
