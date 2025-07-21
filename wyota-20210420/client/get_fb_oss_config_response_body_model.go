// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbOssConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFbOssConfigResponseBody
	GetCode() *string
	SetData(v *GetFbOssConfigResponseBodyData) *GetFbOssConfigResponseBody
	GetData() *GetFbOssConfigResponseBodyData
	SetMessage(v string) *GetFbOssConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFbOssConfigResponseBody
	GetRequestId() *string
}

type GetFbOssConfigResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetFbOssConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFbOssConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFbOssConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetFbOssConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFbOssConfigResponseBody) GetData() *GetFbOssConfigResponseBodyData {
	return s.Data
}

func (s *GetFbOssConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFbOssConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFbOssConfigResponseBody) SetCode(v string) *GetFbOssConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetFbOssConfigResponseBody) SetData(v *GetFbOssConfigResponseBodyData) *GetFbOssConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetFbOssConfigResponseBody) SetMessage(v string) *GetFbOssConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetFbOssConfigResponseBody) SetRequestId(v string) *GetFbOssConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFbOssConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFbOssConfigResponseBodyData struct {
	AccessKeyId  *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	EndPoint     *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	OssPolicy    *string `json:"OssPolicy,omitempty" xml:"OssPolicy,omitempty"`
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetFbOssConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFbOssConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFbOssConfigResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetFbOssConfigResponseBodyData) GetEndPoint() *string {
	return s.EndPoint
}

func (s *GetFbOssConfigResponseBodyData) GetOssPolicy() *string {
	return s.OssPolicy
}

func (s *GetFbOssConfigResponseBodyData) GetOssSignature() *string {
	return s.OssSignature
}

func (s *GetFbOssConfigResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetFbOssConfigResponseBodyData) SetAccessKeyId(v string) *GetFbOssConfigResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) SetEndPoint(v string) *GetFbOssConfigResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) SetOssPolicy(v string) *GetFbOssConfigResponseBodyData {
	s.OssPolicy = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) SetOssSignature(v string) *GetFbOssConfigResponseBodyData {
	s.OssSignature = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) SetSessionId(v string) *GetFbOssConfigResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetFbOssConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
