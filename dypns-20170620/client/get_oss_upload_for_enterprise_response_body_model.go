// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadForEnterpriseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOssUploadForEnterpriseResponseBody
	GetCode() *string
	SetData(v *GetOssUploadForEnterpriseResponseBodyData) *GetOssUploadForEnterpriseResponseBody
	GetData() *GetOssUploadForEnterpriseResponseBodyData
	SetMessage(v string) *GetOssUploadForEnterpriseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOssUploadForEnterpriseResponseBody
	GetRequestId() *string
}

type GetOssUploadForEnterpriseResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOssUploadForEnterpriseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssUploadForEnterpriseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadForEnterpriseResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssUploadForEnterpriseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOssUploadForEnterpriseResponseBody) GetData() *GetOssUploadForEnterpriseResponseBodyData {
	return s.Data
}

func (s *GetOssUploadForEnterpriseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOssUploadForEnterpriseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssUploadForEnterpriseResponseBody) SetCode(v string) *GetOssUploadForEnterpriseResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBody) SetData(v *GetOssUploadForEnterpriseResponseBodyData) *GetOssUploadForEnterpriseResponseBody {
	s.Data = v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBody) SetMessage(v string) *GetOssUploadForEnterpriseResponseBody {
	s.Message = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBody) SetRequestId(v string) *GetOssUploadForEnterpriseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssUploadForEnterpriseResponseBodyData struct {
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Expire      *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetOssUploadForEnterpriseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadForEnterpriseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssUploadForEnterpriseResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetOssUploadForEnterpriseResponseBodyData) GetExpire() *string {
	return s.Expire
}

func (s *GetOssUploadForEnterpriseResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetOssUploadForEnterpriseResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetOssUploadForEnterpriseResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetOssUploadForEnterpriseResponseBodyData) SetAccessKeyId(v string) *GetOssUploadForEnterpriseResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBodyData) SetExpire(v string) *GetOssUploadForEnterpriseResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBodyData) SetHost(v string) *GetOssUploadForEnterpriseResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBodyData) SetPolicy(v string) *GetOssUploadForEnterpriseResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBodyData) SetSignature(v string) *GetOssUploadForEnterpriseResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
