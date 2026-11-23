// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQwenEduApiActivationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQwenEduApiActivationStatusResponseBody
	GetCode() *string
	SetData(v *GetQwenEduApiActivationStatusResponseBodyData) *GetQwenEduApiActivationStatusResponseBody
	GetData() *GetQwenEduApiActivationStatusResponseBodyData
	SetMessage(v string) *GetQwenEduApiActivationStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQwenEduApiActivationStatusResponseBody
	GetRequestId() *string
}

type GetQwenEduApiActivationStatusResponseBody struct {
	// The API status or POP error code.
	//
	// example:
	//
	// OK
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The activation result.
	Data *GetQwenEduApiActivationStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The **POP request ID**. This field is not returned if signature verification or identity authentication fails.
	//
	// example:
	//
	// 5C5E3F0A-2B0E-4F7B-9C55-0E1A6A0E2F11
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQwenEduApiActivationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQwenEduApiActivationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetQwenEduApiActivationStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQwenEduApiActivationStatusResponseBody) GetData() *GetQwenEduApiActivationStatusResponseBodyData {
	return s.Data
}

func (s *GetQwenEduApiActivationStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQwenEduApiActivationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQwenEduApiActivationStatusResponseBody) SetCode(v string) *GetQwenEduApiActivationStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetQwenEduApiActivationStatusResponseBody) SetData(v *GetQwenEduApiActivationStatusResponseBodyData) *GetQwenEduApiActivationStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetQwenEduApiActivationStatusResponseBody) SetMessage(v string) *GetQwenEduApiActivationStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetQwenEduApiActivationStatusResponseBody) SetRequestId(v string) *GetQwenEduApiActivationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQwenEduApiActivationStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQwenEduApiActivationStatusResponseBodyData struct {
	// The **failure code**. For more information, see 3.4.
	//
	// example:
	//
	// ACTIVATION_FAILED
	FailCode *string `json:"failCode,omitempty" xml:"failCode,omitempty"`
	// The **failure description**.
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

func (s GetQwenEduApiActivationStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQwenEduApiActivationStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQwenEduApiActivationStatusResponseBodyData) GetFailCode() *string {
	return s.FailCode
}

func (s *GetQwenEduApiActivationStatusResponseBodyData) GetFailMessage() *string {
	return s.FailMessage
}

func (s *GetQwenEduApiActivationStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetQwenEduApiActivationStatusResponseBodyData) SetFailCode(v string) *GetQwenEduApiActivationStatusResponseBodyData {
	s.FailCode = &v
	return s
}

func (s *GetQwenEduApiActivationStatusResponseBodyData) SetFailMessage(v string) *GetQwenEduApiActivationStatusResponseBodyData {
	s.FailMessage = &v
	return s
}

func (s *GetQwenEduApiActivationStatusResponseBodyData) SetStatus(v string) *GetQwenEduApiActivationStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQwenEduApiActivationStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
