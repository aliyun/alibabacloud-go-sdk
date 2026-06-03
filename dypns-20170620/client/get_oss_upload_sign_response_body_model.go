// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOssUploadSignResponseBody
	GetCode() *string
	SetData(v *GetOssUploadSignResponseBodyData) *GetOssUploadSignResponseBody
	GetData() *GetOssUploadSignResponseBodyData
	SetMessage(v string) *GetOssUploadSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOssUploadSignResponseBody
	GetRequestId() *string
}

type GetOssUploadSignResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOssUploadSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssUploadSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssUploadSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOssUploadSignResponseBody) GetData() *GetOssUploadSignResponseBodyData {
	return s.Data
}

func (s *GetOssUploadSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOssUploadSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssUploadSignResponseBody) SetCode(v string) *GetOssUploadSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssUploadSignResponseBody) SetData(v *GetOssUploadSignResponseBodyData) *GetOssUploadSignResponseBody {
	s.Data = v
	return s
}

func (s *GetOssUploadSignResponseBody) SetMessage(v string) *GetOssUploadSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetOssUploadSignResponseBody) SetRequestId(v string) *GetOssUploadSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssUploadSignResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssUploadSignResponseBodyData struct {
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Expire      *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetOssUploadSignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssUploadSignResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetOssUploadSignResponseBodyData) GetExpire() *string {
	return s.Expire
}

func (s *GetOssUploadSignResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetOssUploadSignResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetOssUploadSignResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetOssUploadSignResponseBodyData) SetAccessKeyId(v string) *GetOssUploadSignResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetOssUploadSignResponseBodyData) SetExpire(v string) *GetOssUploadSignResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetOssUploadSignResponseBodyData) SetHost(v string) *GetOssUploadSignResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetOssUploadSignResponseBodyData) SetPolicy(v string) *GetOssUploadSignResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetOssUploadSignResponseBodyData) SetSignature(v string) *GetOssUploadSignResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetOssUploadSignResponseBodyData) Validate() error {
	return dara.Validate(s)
}
