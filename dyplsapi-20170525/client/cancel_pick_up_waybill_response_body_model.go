// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPickUpWaybillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelPickUpWaybillResponseBody
	GetCode() *string
	SetData(v *CancelPickUpWaybillResponseBodyData) *CancelPickUpWaybillResponseBody
	GetData() *CancelPickUpWaybillResponseBodyData
	SetMessage(v string) *CancelPickUpWaybillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelPickUpWaybillResponseBody
	GetRequestId() *string
}

type CancelPickUpWaybillResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CancelPickUpWaybillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CancelPickUpWaybillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelPickUpWaybillResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPickUpWaybillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelPickUpWaybillResponseBody) GetData() *CancelPickUpWaybillResponseBodyData {
	return s.Data
}

func (s *CancelPickUpWaybillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelPickUpWaybillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelPickUpWaybillResponseBody) SetCode(v string) *CancelPickUpWaybillResponseBody {
	s.Code = &v
	return s
}

func (s *CancelPickUpWaybillResponseBody) SetData(v *CancelPickUpWaybillResponseBodyData) *CancelPickUpWaybillResponseBody {
	s.Data = v
	return s
}

func (s *CancelPickUpWaybillResponseBody) SetMessage(v string) *CancelPickUpWaybillResponseBody {
	s.Message = &v
	return s
}

func (s *CancelPickUpWaybillResponseBody) SetRequestId(v string) *CancelPickUpWaybillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPickUpWaybillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelPickUpWaybillResponseBodyData struct {
	// The error code.
	//
	// example:
	//
	// none
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// none
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The cancellation result.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the cancellation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelPickUpWaybillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelPickUpWaybillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelPickUpWaybillResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CancelPickUpWaybillResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CancelPickUpWaybillResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CancelPickUpWaybillResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CancelPickUpWaybillResponseBodyData) SetErrorCode(v string) *CancelPickUpWaybillResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CancelPickUpWaybillResponseBodyData) SetErrorMsg(v string) *CancelPickUpWaybillResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *CancelPickUpWaybillResponseBodyData) SetMessage(v string) *CancelPickUpWaybillResponseBodyData {
	s.Message = &v
	return s
}

func (s *CancelPickUpWaybillResponseBodyData) SetSuccess(v bool) *CancelPickUpWaybillResponseBodyData {
	s.Success = &v
	return s
}

func (s *CancelPickUpWaybillResponseBodyData) Validate() error {
	return dara.Validate(s)
}
