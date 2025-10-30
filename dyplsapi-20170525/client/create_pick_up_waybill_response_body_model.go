// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePickUpWaybillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreatePickUpWaybillResponseBodyData) *CreatePickUpWaybillResponseBody
	GetData() *CreatePickUpWaybillResponseBodyData
	SetHttpStatusCode(v int32) *CreatePickUpWaybillResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePickUpWaybillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePickUpWaybillResponseBody
	GetRequestId() *string
}

type CreatePickUpWaybillResponseBody struct {
	// The returned data.
	Data *CreatePickUpWaybillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FC30594-3841-43AD-9008-03393BCB5CD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePickUpWaybillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillResponseBody) GetData() *CreatePickUpWaybillResponseBodyData {
	return s.Data
}

func (s *CreatePickUpWaybillResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePickUpWaybillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePickUpWaybillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePickUpWaybillResponseBody) SetData(v *CreatePickUpWaybillResponseBodyData) *CreatePickUpWaybillResponseBody {
	s.Data = v
	return s
}

func (s *CreatePickUpWaybillResponseBody) SetHttpStatusCode(v int32) *CreatePickUpWaybillResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePickUpWaybillResponseBody) SetMessage(v string) *CreatePickUpWaybillResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePickUpWaybillResponseBody) SetRequestId(v string) *CreatePickUpWaybillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePickUpWaybillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePickUpWaybillResponseBodyData struct {
	// The code of the courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// none
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The pickup code.
	//
	// example:
	//
	// 3524
	GotCode *string `json:"GotCode,omitempty" xml:"GotCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 77312345629****
	MailNo *string `json:"MailNo,omitempty" xml:"MailNo,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePickUpWaybillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillResponseBodyData) GetCpCode() *string {
	return s.CpCode
}

func (s *CreatePickUpWaybillResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreatePickUpWaybillResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreatePickUpWaybillResponseBodyData) GetGotCode() *string {
	return s.GotCode
}

func (s *CreatePickUpWaybillResponseBodyData) GetMailNo() *string {
	return s.MailNo
}

func (s *CreatePickUpWaybillResponseBodyData) GetSuccess() *string {
	return s.Success
}

func (s *CreatePickUpWaybillResponseBodyData) SetCpCode(v string) *CreatePickUpWaybillResponseBodyData {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetErrorCode(v string) *CreatePickUpWaybillResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetErrorMsg(v string) *CreatePickUpWaybillResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetGotCode(v string) *CreatePickUpWaybillResponseBodyData {
	s.GotCode = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetMailNo(v string) *CreatePickUpWaybillResponseBodyData {
	s.MailNo = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) SetSuccess(v string) *CreatePickUpWaybillResponseBodyData {
	s.Success = &v
	return s
}

func (s *CreatePickUpWaybillResponseBodyData) Validate() error {
	return dara.Validate(s)
}
