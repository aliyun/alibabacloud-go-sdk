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
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetNotificationClickRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 99D93ED4-D462-5FC5-8518-9BC1C49C7B6C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// false
	UserCancel *string `json:"UserCancel,omitempty" xml:"UserCancel,omitempty"`
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
