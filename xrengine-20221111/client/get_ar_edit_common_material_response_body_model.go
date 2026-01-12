// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditCommonMaterialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetArEditCommonMaterialResponseBody
	GetCode() *string
	SetData(v *GetArEditCommonMaterialResponseBodyData) *GetArEditCommonMaterialResponseBody
	GetData() *GetArEditCommonMaterialResponseBodyData
	SetErrorName(v string) *GetArEditCommonMaterialResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *GetArEditCommonMaterialResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *GetArEditCommonMaterialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetArEditCommonMaterialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetArEditCommonMaterialResponseBody
	GetSuccess() *bool
}

type GetArEditCommonMaterialResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetArEditCommonMaterialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                                  `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetArEditCommonMaterialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArEditCommonMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *GetArEditCommonMaterialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArEditCommonMaterialResponseBody) GetData() *GetArEditCommonMaterialResponseBodyData {
	return s.Data
}

func (s *GetArEditCommonMaterialResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *GetArEditCommonMaterialResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetArEditCommonMaterialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetArEditCommonMaterialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArEditCommonMaterialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetArEditCommonMaterialResponseBody) SetCode(v string) *GetArEditCommonMaterialResponseBody {
	s.Code = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBody) SetData(v *GetArEditCommonMaterialResponseBodyData) *GetArEditCommonMaterialResponseBody {
	s.Data = v
	return s
}

func (s *GetArEditCommonMaterialResponseBody) SetErrorName(v string) *GetArEditCommonMaterialResponseBody {
	s.ErrorName = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBody) SetHttpCode(v int32) *GetArEditCommonMaterialResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBody) SetMessage(v string) *GetArEditCommonMaterialResponseBody {
	s.Message = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBody) SetRequestId(v string) *GetArEditCommonMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBody) SetSuccess(v bool) *GetArEditCommonMaterialResponseBody {
	s.Success = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetArEditCommonMaterialResponseBodyData struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Expiration      *int64  `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	OssBucket       *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssPath         *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	OssRegion       *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetArEditCommonMaterialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetArEditCommonMaterialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetArEditCommonMaterialResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetArEditCommonMaterialResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetArEditCommonMaterialResponseBodyData) GetExpiration() *int64 {
	return s.Expiration
}

func (s *GetArEditCommonMaterialResponseBodyData) GetOssBucket() *string {
	return s.OssBucket
}

func (s *GetArEditCommonMaterialResponseBodyData) GetOssPath() *string {
	return s.OssPath
}

func (s *GetArEditCommonMaterialResponseBodyData) GetOssRegion() *string {
	return s.OssRegion
}

func (s *GetArEditCommonMaterialResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetArEditCommonMaterialResponseBodyData) SetAccessKeyId(v string) *GetArEditCommonMaterialResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBodyData) SetAccessKeySecret(v string) *GetArEditCommonMaterialResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBodyData) SetExpiration(v int64) *GetArEditCommonMaterialResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBodyData) SetOssBucket(v string) *GetArEditCommonMaterialResponseBodyData {
	s.OssBucket = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBodyData) SetOssPath(v string) *GetArEditCommonMaterialResponseBodyData {
	s.OssPath = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBodyData) SetOssRegion(v string) *GetArEditCommonMaterialResponseBodyData {
	s.OssRegion = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBodyData) SetSecurityToken(v string) *GetArEditCommonMaterialResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetArEditCommonMaterialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
