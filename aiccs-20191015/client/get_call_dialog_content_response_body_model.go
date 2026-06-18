// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallDialogContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetCallDialogContentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetCallDialogContentResponseBody
	GetCode() *string
	SetData(v *GetCallDialogContentResponseBodyData) *GetCallDialogContentResponseBody
	GetData() *GetCallDialogContentResponseBodyData
	SetMessage(v string) *GetCallDialogContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCallDialogContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCallDialogContentResponseBody
	GetSuccess() *bool
}

type GetCallDialogContentResponseBody struct {
	// Details about the access denial. Returned only when RAM authentication fails.
	//
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetCallDialogContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FB0B0481-F13E-16E0-8A7A-1AD2FXXXEF55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCallDialogContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCallDialogContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallDialogContentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetCallDialogContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCallDialogContentResponseBody) GetData() *GetCallDialogContentResponseBodyData {
	return s.Data
}

func (s *GetCallDialogContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCallDialogContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCallDialogContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallDialogContentResponseBody) SetAccessDeniedDetail(v string) *GetCallDialogContentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCallDialogContentResponseBody) SetCode(v string) *GetCallDialogContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallDialogContentResponseBody) SetData(v *GetCallDialogContentResponseBodyData) *GetCallDialogContentResponseBody {
	s.Data = v
	return s
}

func (s *GetCallDialogContentResponseBody) SetMessage(v string) *GetCallDialogContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallDialogContentResponseBody) SetRequestId(v string) *GetCallDialogContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallDialogContentResponseBody) SetSuccess(v bool) *GetCallDialogContentResponseBody {
	s.Success = &v
	return s
}

func (s *GetCallDialogContentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCallDialogContentResponseBodyData struct {
	// The call ID.
	//
	// example:
	//
	// 123456^123478
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The call status.
	//
	// > Valid values:
	//
	// >
	//
	// > - `0`: Not connected
	//
	// >
	//
	// > - `1`: Connected
	//
	// >
	//
	// > - `2`: Disconnected
	//
	// example:
	//
	// 2
	CallStatus *int64 `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// The dialog content.
	//
	// example:
	//
	// [{\\"content\\":\\"您好。\\",\\"role\\":\\"assistant\\"},{\\"content\\":\\"不用了。\\",\\"role\\":\\"user\\"},{\\"content\\":\\"呃，不用了，再见，谢谢。\\",\\"role\\":\\"user\\"}]
	DialogContent *string `json:"DialogContent,omitempty" xml:"DialogContent,omitempty"`
}

func (s GetCallDialogContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCallDialogContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCallDialogContentResponseBodyData) GetCallId() *string {
	return s.CallId
}

func (s *GetCallDialogContentResponseBodyData) GetCallStatus() *int64 {
	return s.CallStatus
}

func (s *GetCallDialogContentResponseBodyData) GetDialogContent() *string {
	return s.DialogContent
}

func (s *GetCallDialogContentResponseBodyData) SetCallId(v string) *GetCallDialogContentResponseBodyData {
	s.CallId = &v
	return s
}

func (s *GetCallDialogContentResponseBodyData) SetCallStatus(v int64) *GetCallDialogContentResponseBodyData {
	s.CallStatus = &v
	return s
}

func (s *GetCallDialogContentResponseBodyData) SetDialogContent(v string) *GetCallDialogContentResponseBodyData {
	s.DialogContent = &v
	return s
}

func (s *GetCallDialogContentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
