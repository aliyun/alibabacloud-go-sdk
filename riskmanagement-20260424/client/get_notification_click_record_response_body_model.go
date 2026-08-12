// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationClickRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNotificationClickRecordResponseBody
	GetCode() *string
	SetData(v *GetNotificationClickRecordResponseBodyData) *GetNotificationClickRecordResponseBody
	GetData() *GetNotificationClickRecordResponseBodyData
	SetMessage(v string) *GetNotificationClickRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNotificationClickRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNotificationClickRecordResponseBody
	GetSuccess() *bool
}

type GetNotificationClickRecordResponseBody struct {
	// The status code.
	//
	// - **200**: Succeeded.
	//
	// - **Others (400, 500)**: Failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The metadata.
	Data *GetNotificationClickRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request was successful, a success message is returned. If the request failed, the failure reason is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 99D93ED4-D462-5FC5-8518-9BC1C49C7B6C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNotificationClickRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationClickRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotificationClickRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNotificationClickRecordResponseBody) GetData() *GetNotificationClickRecordResponseBodyData {
	return s.Data
}

func (s *GetNotificationClickRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNotificationClickRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNotificationClickRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNotificationClickRecordResponseBody) SetCode(v string) *GetNotificationClickRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetNotificationClickRecordResponseBody) SetData(v *GetNotificationClickRecordResponseBodyData) *GetNotificationClickRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetNotificationClickRecordResponseBody) SetMessage(v string) *GetNotificationClickRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetNotificationClickRecordResponseBody) SetRequestId(v string) *GetNotificationClickRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotificationClickRecordResponseBody) SetSuccess(v bool) *GetNotificationClickRecordResponseBody {
	s.Success = &v
	return s
}

func (s *GetNotificationClickRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNotificationClickRecordResponseBodyData struct {
	// Indicates whether the user clicked cancel.
	//
	// - **true**: Canceled.
	//
	// - **false**: Not canceled.
	//
	// example:
	//
	// false
	UserCancel *string `json:"UserCancel,omitempty" xml:"UserCancel,omitempty"`
	// Indicates whether the user clicked confirm.
	//
	// - **true**: Confirmed.
	//
	// - **false**: Not confirmed.
	//
	// example:
	//
	// true
	UserConfirm *string `json:"UserConfirm,omitempty" xml:"UserConfirm,omitempty"`
}

func (s GetNotificationClickRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationClickRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNotificationClickRecordResponseBodyData) GetUserCancel() *string {
	return s.UserCancel
}

func (s *GetNotificationClickRecordResponseBodyData) GetUserConfirm() *string {
	return s.UserConfirm
}

func (s *GetNotificationClickRecordResponseBodyData) SetUserCancel(v string) *GetNotificationClickRecordResponseBodyData {
	s.UserCancel = &v
	return s
}

func (s *GetNotificationClickRecordResponseBodyData) SetUserConfirm(v string) *GetNotificationClickRecordResponseBodyData {
	s.UserConfirm = &v
	return s
}

func (s *GetNotificationClickRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
