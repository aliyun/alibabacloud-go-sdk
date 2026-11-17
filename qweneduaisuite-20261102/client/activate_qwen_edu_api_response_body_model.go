// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateQwenEduApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ActivateQwenEduApiResponseBody
	GetCode() *string
	SetData(v *ActivateQwenEduApiResponseBodyData) *ActivateQwenEduApiResponseBody
	GetData() *ActivateQwenEduApiResponseBodyData
	SetMessage(v string) *ActivateQwenEduApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *ActivateQwenEduApiResponseBody
	GetRequestId() *string
}

type ActivateQwenEduApiResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *ActivateQwenEduApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ActivateQwenEduApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateQwenEduApiResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateQwenEduApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *ActivateQwenEduApiResponseBody) GetData() *ActivateQwenEduApiResponseBodyData {
	return s.Data
}

func (s *ActivateQwenEduApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ActivateQwenEduApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateQwenEduApiResponseBody) SetCode(v string) *ActivateQwenEduApiResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateQwenEduApiResponseBody) SetData(v *ActivateQwenEduApiResponseBodyData) *ActivateQwenEduApiResponseBody {
	s.Data = v
	return s
}

func (s *ActivateQwenEduApiResponseBody) SetMessage(v string) *ActivateQwenEduApiResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateQwenEduApiResponseBody) SetRequestId(v string) *ActivateQwenEduApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateQwenEduApiResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ActivateQwenEduApiResponseBodyData struct {
	FailMessage *string `json:"failMessage,omitempty" xml:"failMessage,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ActivateQwenEduApiResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ActivateQwenEduApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *ActivateQwenEduApiResponseBodyData) GetFailMessage() *string {
	return s.FailMessage
}

func (s *ActivateQwenEduApiResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ActivateQwenEduApiResponseBodyData) SetFailMessage(v string) *ActivateQwenEduApiResponseBodyData {
	s.FailMessage = &v
	return s
}

func (s *ActivateQwenEduApiResponseBodyData) SetStatus(v string) *ActivateQwenEduApiResponseBodyData {
	s.Status = &v
	return s
}

func (s *ActivateQwenEduApiResponseBodyData) Validate() error {
	return dara.Validate(s)
}
