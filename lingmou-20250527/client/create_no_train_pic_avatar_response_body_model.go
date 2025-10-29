// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNoTrainPicAvatarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateNoTrainPicAvatarResponseBody
	GetCode() *string
	SetData(v *CreateNoTrainPicAvatarResponseBodyData) *CreateNoTrainPicAvatarResponseBody
	GetData() *CreateNoTrainPicAvatarResponseBodyData
	SetHttpStatusCode(v int32) *CreateNoTrainPicAvatarResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateNoTrainPicAvatarResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNoTrainPicAvatarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNoTrainPicAvatarResponseBody
	GetSuccess() *bool
}

type CreateNoTrainPicAvatarResponseBody struct {
	// example:
	//
	// 200
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateNoTrainPicAvatarResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateNoTrainPicAvatarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNoTrainPicAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNoTrainPicAvatarResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateNoTrainPicAvatarResponseBody) GetData() *CreateNoTrainPicAvatarResponseBodyData {
	return s.Data
}

func (s *CreateNoTrainPicAvatarResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateNoTrainPicAvatarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNoTrainPicAvatarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNoTrainPicAvatarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNoTrainPicAvatarResponseBody) SetCode(v string) *CreateNoTrainPicAvatarResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNoTrainPicAvatarResponseBody) SetData(v *CreateNoTrainPicAvatarResponseBodyData) *CreateNoTrainPicAvatarResponseBody {
	s.Data = v
	return s
}

func (s *CreateNoTrainPicAvatarResponseBody) SetHttpStatusCode(v int32) *CreateNoTrainPicAvatarResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateNoTrainPicAvatarResponseBody) SetMessage(v string) *CreateNoTrainPicAvatarResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNoTrainPicAvatarResponseBody) SetRequestId(v string) *CreateNoTrainPicAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNoTrainPicAvatarResponseBody) SetSuccess(v bool) *CreateNoTrainPicAvatarResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNoTrainPicAvatarResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNoTrainPicAvatarResponseBodyData struct {
	// example:
	//
	// M1ONzwuILu-nPT7pvr6maKvQ
	AvatarId *string `json:"avatarId,omitempty" xml:"avatarId,omitempty"`
	// example:
	//
	// true/false
	Pass *bool `json:"pass,omitempty" xml:"pass,omitempty"`
}

func (s CreateNoTrainPicAvatarResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateNoTrainPicAvatarResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateNoTrainPicAvatarResponseBodyData) GetAvatarId() *string {
	return s.AvatarId
}

func (s *CreateNoTrainPicAvatarResponseBodyData) GetPass() *bool {
	return s.Pass
}

func (s *CreateNoTrainPicAvatarResponseBodyData) SetAvatarId(v string) *CreateNoTrainPicAvatarResponseBodyData {
	s.AvatarId = &v
	return s
}

func (s *CreateNoTrainPicAvatarResponseBodyData) SetPass(v bool) *CreateNoTrainPicAvatarResponseBodyData {
	s.Pass = &v
	return s
}

func (s *CreateNoTrainPicAvatarResponseBodyData) Validate() error {
	return dara.Validate(s)
}
