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
	// The API status or POP error code.
	//
	// example:
	//
	// OK
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The activation result. This parameter is null when an error occurs.
	Data *ActivateQwenEduApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The POP request ID. This parameter is not returned when signature or identity verification fails.
	//
	// example:
	//
	// 5C5E3F0A-2B0E-4F7B-9C55-0E1A6A0E2F11
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
	// The failure description.
	//
	// example:
	//
	// Service activation failed. Try again later
	FailMessage *string `json:"failMessage,omitempty" xml:"failMessage,omitempty"`
	// The activation status. Valid values:
	//
	// - ACTIVE: Activated.
	//
	// - ACTIVATING: Activating.
	//
	// - FAILED: Activation failed.
	//
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
