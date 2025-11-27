// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualificationOssInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetQualificationOssInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetQualificationOssInfoResponseBody
	GetCode() *string
	SetData(v *GetQualificationOssInfoResponseBodyData) *GetQualificationOssInfoResponseBody
	GetData() *GetQualificationOssInfoResponseBodyData
	SetMessage(v string) *GetQualificationOssInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualificationOssInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualificationOssInfoResponseBody
	GetSuccess() *bool
}

type GetQualificationOssInfoResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetQualificationOssInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 25D5AFDE-8EBC-132E-8909-1FDC071DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualificationOssInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualificationOssInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualificationOssInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetQualificationOssInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualificationOssInfoResponseBody) GetData() *GetQualificationOssInfoResponseBodyData {
	return s.Data
}

func (s *GetQualificationOssInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualificationOssInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualificationOssInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualificationOssInfoResponseBody) SetAccessDeniedDetail(v string) *GetQualificationOssInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetQualificationOssInfoResponseBody) SetCode(v string) *GetQualificationOssInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualificationOssInfoResponseBody) SetData(v *GetQualificationOssInfoResponseBodyData) *GetQualificationOssInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetQualificationOssInfoResponseBody) SetMessage(v string) *GetQualificationOssInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualificationOssInfoResponseBody) SetRequestId(v string) *GetQualificationOssInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualificationOssInfoResponseBody) SetSuccess(v bool) *GetQualificationOssInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualificationOssInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualificationOssInfoResponseBodyData struct {
	// ak
	//
	// example:
	//
	// bypFNbG******
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// 过期时间
	//
	// example:
	//
	// 1741521339
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// 域名
	//
	// example:
	//
	// http://***.oss-cn-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 策略
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wMy0wOVQxMTo1NTozOS4wMDFaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0seyJidWNrZXQiOiJhbGljb20tZmMtbWVkaWEifSxbImVxIiwiJGtleSIsIjEwMDAwMDM1ODA4MjA2M1wv********
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 签名
	//
	// example:
	//
	// QvNTGC9DSLTeByP+ZWW******
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 前缀
	//
	// example:
	//
	// 1000********001
	StartPath *string `json:"StartPath,omitempty" xml:"StartPath,omitempty"`
}

func (s GetQualificationOssInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualificationOssInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualificationOssInfoResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetQualificationOssInfoResponseBodyData) GetExpire() *int64 {
	return s.Expire
}

func (s *GetQualificationOssInfoResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetQualificationOssInfoResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetQualificationOssInfoResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetQualificationOssInfoResponseBodyData) GetStartPath() *string {
	return s.StartPath
}

func (s *GetQualificationOssInfoResponseBodyData) SetAccessKeyId(v string) *GetQualificationOssInfoResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetQualificationOssInfoResponseBodyData) SetExpire(v int64) *GetQualificationOssInfoResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetQualificationOssInfoResponseBodyData) SetHost(v string) *GetQualificationOssInfoResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetQualificationOssInfoResponseBodyData) SetPolicy(v string) *GetQualificationOssInfoResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetQualificationOssInfoResponseBodyData) SetSignature(v string) *GetQualificationOssInfoResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetQualificationOssInfoResponseBodyData) SetStartPath(v string) *GetQualificationOssInfoResponseBodyData {
	s.StartPath = &v
	return s
}

func (s *GetQualificationOssInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
