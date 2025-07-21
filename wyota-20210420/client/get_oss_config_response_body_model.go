// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOssConfigResponseBody
	GetCode() *string
	SetData(v *GetOssConfigResponseBodyData) *GetOssConfigResponseBody
	GetData() *GetOssConfigResponseBodyData
	SetMessage(v string) *GetOssConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOssConfigResponseBody
	GetRequestId() *string
}

type GetOssConfigResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOssConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOssConfigResponseBody) GetData() *GetOssConfigResponseBodyData {
	return s.Data
}

func (s *GetOssConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOssConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssConfigResponseBody) SetCode(v string) *GetOssConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssConfigResponseBody) SetData(v *GetOssConfigResponseBodyData) *GetOssConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetOssConfigResponseBody) SetMessage(v string) *GetOssConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetOssConfigResponseBody) SetRequestId(v string) *GetOssConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOssConfigResponseBodyData struct {
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	EndPoint      *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	OssPolicy     *string `json:"OssPolicy,omitempty" xml:"OssPolicy,omitempty"`
	OssSignature  *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetOssConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOssConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssConfigResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetOssConfigResponseBodyData) GetEndPoint() *string {
	return s.EndPoint
}

func (s *GetOssConfigResponseBodyData) GetOssPolicy() *string {
	return s.OssPolicy
}

func (s *GetOssConfigResponseBodyData) GetOssSignature() *string {
	return s.OssSignature
}

func (s *GetOssConfigResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetOssConfigResponseBodyData) SetAccessKeyId(v string) *GetOssConfigResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetOssConfigResponseBodyData) SetEndPoint(v string) *GetOssConfigResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *GetOssConfigResponseBodyData) SetOssPolicy(v string) *GetOssConfigResponseBodyData {
	s.OssPolicy = &v
	return s
}

func (s *GetOssConfigResponseBodyData) SetOssSignature(v string) *GetOssConfigResponseBodyData {
	s.OssSignature = &v
	return s
}

func (s *GetOssConfigResponseBodyData) SetSecurityToken(v string) *GetOssConfigResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetOssConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
