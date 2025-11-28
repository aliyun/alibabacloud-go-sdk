// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainPicAvatarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTrainPicAvatarResponseBody
	GetCode() *string
	SetData(v *CreateTrainPicAvatarResponseBodyData) *CreateTrainPicAvatarResponseBody
	GetData() *CreateTrainPicAvatarResponseBodyData
	SetHttpStatusCode(v int32) *CreateTrainPicAvatarResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTrainPicAvatarResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTrainPicAvatarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTrainPicAvatarResponseBody
	GetSuccess() *bool
}

type CreateTrainPicAvatarResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateTrainPicAvatarResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateTrainPicAvatarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainPicAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrainPicAvatarResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTrainPicAvatarResponseBody) GetData() *CreateTrainPicAvatarResponseBodyData {
	return s.Data
}

func (s *CreateTrainPicAvatarResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTrainPicAvatarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTrainPicAvatarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrainPicAvatarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTrainPicAvatarResponseBody) SetCode(v string) *CreateTrainPicAvatarResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTrainPicAvatarResponseBody) SetData(v *CreateTrainPicAvatarResponseBodyData) *CreateTrainPicAvatarResponseBody {
	s.Data = v
	return s
}

func (s *CreateTrainPicAvatarResponseBody) SetHttpStatusCode(v int32) *CreateTrainPicAvatarResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTrainPicAvatarResponseBody) SetMessage(v string) *CreateTrainPicAvatarResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTrainPicAvatarResponseBody) SetRequestId(v string) *CreateTrainPicAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrainPicAvatarResponseBody) SetSuccess(v bool) *CreateTrainPicAvatarResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTrainPicAvatarResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTrainPicAvatarResponseBodyData struct {
	// example:
	//
	// M1AguofmMxaoUQsuSPQ3j0ng
	AvatarId *string `json:"avatarId,omitempty" xml:"avatarId,omitempty"`
	// example:
	//
	// 1200
	ExpectedCompletionTime *int32 `json:"expectedCompletionTime,omitempty" xml:"expectedCompletionTime,omitempty"`
	Pass                   *bool  `json:"pass,omitempty" xml:"pass,omitempty"`
}

func (s CreateTrainPicAvatarResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainPicAvatarResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTrainPicAvatarResponseBodyData) GetAvatarId() *string {
	return s.AvatarId
}

func (s *CreateTrainPicAvatarResponseBodyData) GetExpectedCompletionTime() *int32 {
	return s.ExpectedCompletionTime
}

func (s *CreateTrainPicAvatarResponseBodyData) GetPass() *bool {
	return s.Pass
}

func (s *CreateTrainPicAvatarResponseBodyData) SetAvatarId(v string) *CreateTrainPicAvatarResponseBodyData {
	s.AvatarId = &v
	return s
}

func (s *CreateTrainPicAvatarResponseBodyData) SetExpectedCompletionTime(v int32) *CreateTrainPicAvatarResponseBodyData {
	s.ExpectedCompletionTime = &v
	return s
}

func (s *CreateTrainPicAvatarResponseBodyData) SetPass(v bool) *CreateTrainPicAvatarResponseBodyData {
	s.Pass = &v
	return s
}

func (s *CreateTrainPicAvatarResponseBodyData) Validate() error {
	return dara.Validate(s)
}
