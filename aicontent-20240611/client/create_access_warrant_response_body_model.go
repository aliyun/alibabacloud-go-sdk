// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessWarrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAccessWarrantResponseBodyData) *CreateAccessWarrantResponseBody
	GetData() *CreateAccessWarrantResponseBodyData
	SetErrCode(v string) *CreateAccessWarrantResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateAccessWarrantResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateAccessWarrantResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateAccessWarrantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAccessWarrantResponseBody
	GetSuccess() *bool
}

type CreateAccessWarrantResponseBody struct {
	// example:
	//
	// []
	Data *CreateAccessWarrantResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAccessWarrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessWarrantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessWarrantResponseBody) GetData() *CreateAccessWarrantResponseBodyData {
	return s.Data
}

func (s *CreateAccessWarrantResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateAccessWarrantResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateAccessWarrantResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAccessWarrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessWarrantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAccessWarrantResponseBody) SetData(v *CreateAccessWarrantResponseBodyData) *CreateAccessWarrantResponseBody {
	s.Data = v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetErrCode(v string) *CreateAccessWarrantResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetErrMessage(v string) *CreateAccessWarrantResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetHttpStatusCode(v int32) *CreateAccessWarrantResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetRequestId(v string) *CreateAccessWarrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetSuccess(v bool) *CreateAccessWarrantResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAccessWarrantResponseBodyData struct {
	// example:
	//
	// 1234567890
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// ex2xxxxxxxx
	AccessWarrantId *string `json:"AccessWarrantId,omitempty" xml:"AccessWarrantId,omitempty"`
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"ApplicationAccessId,omitempty" xml:"ApplicationAccessId,omitempty"`
	// example:
	//
	// 1672531200
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1672531200
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1234567890
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateAccessWarrantResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessWarrantResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAccessWarrantResponseBodyData) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateAccessWarrantResponseBodyData) GetAccessWarrantId() *string {
	return s.AccessWarrantId
}

func (s *CreateAccessWarrantResponseBodyData) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *CreateAccessWarrantResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAccessWarrantResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateAccessWarrantResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *CreateAccessWarrantResponseBodyData) SetAccessToken(v string) *CreateAccessWarrantResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetAccessWarrantId(v string) *CreateAccessWarrantResponseBodyData {
	s.AccessWarrantId = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetApplicationAccessId(v string) *CreateAccessWarrantResponseBodyData {
	s.ApplicationAccessId = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetCreateTime(v string) *CreateAccessWarrantResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetExpireTime(v string) *CreateAccessWarrantResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetUserId(v string) *CreateAccessWarrantResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) Validate() error {
	return dara.Validate(s)
}
