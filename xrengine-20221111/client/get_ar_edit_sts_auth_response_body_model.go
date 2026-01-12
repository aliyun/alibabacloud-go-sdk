// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditStsAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetArEditStsAuthResponseBody
	GetCode() *string
	SetData(v *GetArEditStsAuthResponseBodyData) *GetArEditStsAuthResponseBody
	GetData() *GetArEditStsAuthResponseBodyData
	SetErrorName(v string) *GetArEditStsAuthResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *GetArEditStsAuthResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *GetArEditStsAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetArEditStsAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetArEditStsAuthResponseBody
	GetSuccess() *bool
}

type GetArEditStsAuthResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetArEditStsAuthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                           `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                            `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetArEditStsAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArEditStsAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GetArEditStsAuthResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArEditStsAuthResponseBody) GetData() *GetArEditStsAuthResponseBodyData {
	return s.Data
}

func (s *GetArEditStsAuthResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *GetArEditStsAuthResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetArEditStsAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetArEditStsAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArEditStsAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetArEditStsAuthResponseBody) SetCode(v string) *GetArEditStsAuthResponseBody {
	s.Code = &v
	return s
}

func (s *GetArEditStsAuthResponseBody) SetData(v *GetArEditStsAuthResponseBodyData) *GetArEditStsAuthResponseBody {
	s.Data = v
	return s
}

func (s *GetArEditStsAuthResponseBody) SetErrorName(v string) *GetArEditStsAuthResponseBody {
	s.ErrorName = &v
	return s
}

func (s *GetArEditStsAuthResponseBody) SetHttpCode(v int32) *GetArEditStsAuthResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetArEditStsAuthResponseBody) SetMessage(v string) *GetArEditStsAuthResponseBody {
	s.Message = &v
	return s
}

func (s *GetArEditStsAuthResponseBody) SetRequestId(v string) *GetArEditStsAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArEditStsAuthResponseBody) SetSuccess(v bool) *GetArEditStsAuthResponseBody {
	s.Success = &v
	return s
}

func (s *GetArEditStsAuthResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetArEditStsAuthResponseBodyData struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	EditPath        *string `json:"EditPath,omitempty" xml:"EditPath,omitempty"`
	Expiration      *int64  `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Map3DPath       *string `json:"Map3DPath,omitempty" xml:"Map3DPath,omitempty"`
	OssBucket       *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssRegion       *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	PublishPath     *string `json:"PublishPath,omitempty" xml:"PublishPath,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetArEditStsAuthResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetArEditStsAuthResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetArEditStsAuthResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetArEditStsAuthResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetArEditStsAuthResponseBodyData) GetEditPath() *string {
	return s.EditPath
}

func (s *GetArEditStsAuthResponseBodyData) GetExpiration() *int64 {
	return s.Expiration
}

func (s *GetArEditStsAuthResponseBodyData) GetMap3DPath() *string {
	return s.Map3DPath
}

func (s *GetArEditStsAuthResponseBodyData) GetOssBucket() *string {
	return s.OssBucket
}

func (s *GetArEditStsAuthResponseBodyData) GetOssRegion() *string {
	return s.OssRegion
}

func (s *GetArEditStsAuthResponseBodyData) GetPublishPath() *string {
	return s.PublishPath
}

func (s *GetArEditStsAuthResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetArEditStsAuthResponseBodyData) SetAccessKeyId(v string) *GetArEditStsAuthResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) SetAccessKeySecret(v string) *GetArEditStsAuthResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) SetEditPath(v string) *GetArEditStsAuthResponseBodyData {
	s.EditPath = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) SetExpiration(v int64) *GetArEditStsAuthResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) SetMap3DPath(v string) *GetArEditStsAuthResponseBodyData {
	s.Map3DPath = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) SetOssBucket(v string) *GetArEditStsAuthResponseBodyData {
	s.OssBucket = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) SetOssRegion(v string) *GetArEditStsAuthResponseBodyData {
	s.OssRegion = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) SetPublishPath(v string) *GetArEditStsAuthResponseBodyData {
	s.PublishPath = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) SetSecurityToken(v string) *GetArEditStsAuthResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetArEditStsAuthResponseBodyData) Validate() error {
	return dara.Validate(s)
}
