// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsOcrOssInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetSmsOcrOssInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetSmsOcrOssInfoResponseBody
	GetCode() *string
	SetMessage(v string) *GetSmsOcrOssInfoResponseBody
	GetMessage() *string
	SetModel(v *GetSmsOcrOssInfoResponseBodyModel) *GetSmsOcrOssInfoResponseBody
	GetModel() *GetSmsOcrOssInfoResponseBodyModel
	SetRequestId(v string) *GetSmsOcrOssInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSmsOcrOssInfoResponseBody
	GetSuccess() *bool
}

type GetSmsOcrOssInfoResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetSmsOcrOssInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 25D5AFDE-xxxx-132E-8909-1FDC071DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSmsOcrOssInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmsOcrOssInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsOcrOssInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetSmsOcrOssInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmsOcrOssInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmsOcrOssInfoResponseBody) GetModel() *GetSmsOcrOssInfoResponseBodyModel {
	return s.Model
}

func (s *GetSmsOcrOssInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmsOcrOssInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSmsOcrOssInfoResponseBody) SetAccessDeniedDetail(v string) *GetSmsOcrOssInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBody) SetCode(v string) *GetSmsOcrOssInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBody) SetMessage(v string) *GetSmsOcrOssInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBody) SetModel(v *GetSmsOcrOssInfoResponseBodyModel) *GetSmsOcrOssInfoResponseBody {
	s.Model = v
	return s
}

func (s *GetSmsOcrOssInfoResponseBody) SetRequestId(v string) *GetSmsOcrOssInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBody) SetSuccess(v bool) *GetSmsOcrOssInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmsOcrOssInfoResponseBodyModel struct {
	// example:
	//
	// bypFNbG******
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 示例值示例值
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 1741521339
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// http://***.oss-cn-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wMy0wOVQxMTo1NTozOS4wMDFaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0seyJidWNrZXQiOiJhbGljb20tZmMtbWVkaWEifSxbImVxIiwiJGtleSIsIjEwMDAwMDM1ODA4MjA2M1wv********
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// QvNTGC9DSLTeByP+ZWW******
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// 1000********001
	StartPath *string `json:"StartPath,omitempty" xml:"StartPath,omitempty"`
}

func (s GetSmsOcrOssInfoResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetSmsOcrOssInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetSmsOcrOssInfoResponseBodyModel) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetSmsOcrOssInfoResponseBodyModel) GetBucket() *string {
	return s.Bucket
}

func (s *GetSmsOcrOssInfoResponseBodyModel) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetSmsOcrOssInfoResponseBodyModel) GetHost() *string {
	return s.Host
}

func (s *GetSmsOcrOssInfoResponseBodyModel) GetPolicy() *string {
	return s.Policy
}

func (s *GetSmsOcrOssInfoResponseBodyModel) GetSignature() *string {
	return s.Signature
}

func (s *GetSmsOcrOssInfoResponseBodyModel) GetStartPath() *string {
	return s.StartPath
}

func (s *GetSmsOcrOssInfoResponseBodyModel) SetAccessKeyId(v string) *GetSmsOcrOssInfoResponseBodyModel {
	s.AccessKeyId = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBodyModel) SetBucket(v string) *GetSmsOcrOssInfoResponseBodyModel {
	s.Bucket = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBodyModel) SetExpireTime(v string) *GetSmsOcrOssInfoResponseBodyModel {
	s.ExpireTime = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBodyModel) SetHost(v string) *GetSmsOcrOssInfoResponseBodyModel {
	s.Host = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBodyModel) SetPolicy(v string) *GetSmsOcrOssInfoResponseBodyModel {
	s.Policy = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBodyModel) SetSignature(v string) *GetSmsOcrOssInfoResponseBodyModel {
	s.Signature = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBodyModel) SetStartPath(v string) *GetSmsOcrOssInfoResponseBodyModel {
	s.StartPath = &v
	return s
}

func (s *GetSmsOcrOssInfoResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
