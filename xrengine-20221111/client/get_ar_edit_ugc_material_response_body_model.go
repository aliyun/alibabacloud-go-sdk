// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditUgcMaterialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetArEditUgcMaterialResponseBody
	GetCode() *string
	SetData(v *GetArEditUgcMaterialResponseBodyData) *GetArEditUgcMaterialResponseBody
	GetData() *GetArEditUgcMaterialResponseBodyData
	SetErrorName(v string) *GetArEditUgcMaterialResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *GetArEditUgcMaterialResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *GetArEditUgcMaterialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetArEditUgcMaterialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetArEditUgcMaterialResponseBody
	GetSuccess() *bool
}

type GetArEditUgcMaterialResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetArEditUgcMaterialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                               `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetArEditUgcMaterialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArEditUgcMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *GetArEditUgcMaterialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArEditUgcMaterialResponseBody) GetData() *GetArEditUgcMaterialResponseBodyData {
	return s.Data
}

func (s *GetArEditUgcMaterialResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *GetArEditUgcMaterialResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetArEditUgcMaterialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetArEditUgcMaterialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArEditUgcMaterialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetArEditUgcMaterialResponseBody) SetCode(v string) *GetArEditUgcMaterialResponseBody {
	s.Code = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBody) SetData(v *GetArEditUgcMaterialResponseBodyData) *GetArEditUgcMaterialResponseBody {
	s.Data = v
	return s
}

func (s *GetArEditUgcMaterialResponseBody) SetErrorName(v string) *GetArEditUgcMaterialResponseBody {
	s.ErrorName = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBody) SetHttpCode(v int32) *GetArEditUgcMaterialResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBody) SetMessage(v string) *GetArEditUgcMaterialResponseBody {
	s.Message = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBody) SetRequestId(v string) *GetArEditUgcMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBody) SetSuccess(v bool) *GetArEditUgcMaterialResponseBody {
	s.Success = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetArEditUgcMaterialResponseBodyData struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Expiration      *int64  `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	OssBucket       *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssPath         *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	OssRegion       *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetArEditUgcMaterialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetArEditUgcMaterialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetArEditUgcMaterialResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetArEditUgcMaterialResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetArEditUgcMaterialResponseBodyData) GetExpiration() *int64 {
	return s.Expiration
}

func (s *GetArEditUgcMaterialResponseBodyData) GetOssBucket() *string {
	return s.OssBucket
}

func (s *GetArEditUgcMaterialResponseBodyData) GetOssPath() *string {
	return s.OssPath
}

func (s *GetArEditUgcMaterialResponseBodyData) GetOssRegion() *string {
	return s.OssRegion
}

func (s *GetArEditUgcMaterialResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetArEditUgcMaterialResponseBodyData) SetAccessKeyId(v string) *GetArEditUgcMaterialResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBodyData) SetAccessKeySecret(v string) *GetArEditUgcMaterialResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBodyData) SetExpiration(v int64) *GetArEditUgcMaterialResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBodyData) SetOssBucket(v string) *GetArEditUgcMaterialResponseBodyData {
	s.OssBucket = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBodyData) SetOssPath(v string) *GetArEditUgcMaterialResponseBodyData {
	s.OssPath = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBodyData) SetOssRegion(v string) *GetArEditUgcMaterialResponseBodyData {
	s.OssRegion = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBodyData) SetSecurityToken(v string) *GetArEditUgcMaterialResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetArEditUgcMaterialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
