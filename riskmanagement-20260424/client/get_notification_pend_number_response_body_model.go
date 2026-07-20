// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationPendNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNotificationPendNumberResponseBody
	GetCode() *string
	SetData(v int64) *GetNotificationPendNumberResponseBody
	GetData() *int64
	SetMessage(v string) *GetNotificationPendNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNotificationPendNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNotificationPendNumberResponseBody
	GetSuccess() *bool
}

type GetNotificationPendNumberResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 5
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BEE90F8C-EDC2-5394-953B-D07A121612B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNotificationPendNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationPendNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotificationPendNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNotificationPendNumberResponseBody) GetData() *int64 {
	return s.Data
}

func (s *GetNotificationPendNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNotificationPendNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNotificationPendNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNotificationPendNumberResponseBody) SetCode(v string) *GetNotificationPendNumberResponseBody {
	s.Code = &v
	return s
}

func (s *GetNotificationPendNumberResponseBody) SetData(v int64) *GetNotificationPendNumberResponseBody {
	s.Data = &v
	return s
}

func (s *GetNotificationPendNumberResponseBody) SetMessage(v string) *GetNotificationPendNumberResponseBody {
	s.Message = &v
	return s
}

func (s *GetNotificationPendNumberResponseBody) SetRequestId(v string) *GetNotificationPendNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotificationPendNumberResponseBody) SetSuccess(v bool) *GetNotificationPendNumberResponseBody {
	s.Success = &v
	return s
}

func (s *GetNotificationPendNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
