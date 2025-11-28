// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmTrainPicAvatarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfirmTrainPicAvatarResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ConfirmTrainPicAvatarResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ConfirmTrainPicAvatarResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfirmTrainPicAvatarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfirmTrainPicAvatarResponseBody
	GetSuccess() *bool
}

type ConfirmTrainPicAvatarResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConfirmTrainPicAvatarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmTrainPicAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmTrainPicAvatarResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfirmTrainPicAvatarResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ConfirmTrainPicAvatarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfirmTrainPicAvatarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmTrainPicAvatarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfirmTrainPicAvatarResponseBody) SetCode(v string) *ConfirmTrainPicAvatarResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmTrainPicAvatarResponseBody) SetHttpStatusCode(v int32) *ConfirmTrainPicAvatarResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfirmTrainPicAvatarResponseBody) SetMessage(v string) *ConfirmTrainPicAvatarResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmTrainPicAvatarResponseBody) SetRequestId(v string) *ConfirmTrainPicAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmTrainPicAvatarResponseBody) SetSuccess(v bool) *ConfirmTrainPicAvatarResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmTrainPicAvatarResponseBody) Validate() error {
	return dara.Validate(s)
}
