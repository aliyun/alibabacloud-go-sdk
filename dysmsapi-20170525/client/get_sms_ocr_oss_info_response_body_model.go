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
	// 访问被拒绝详细信息，只有 RAM 校验失败才会返回此字段。
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码。
	//
	// - 返回 OK 代表请求成功。
	//
	// - 其他错误码，请参见 [API 错误码](https://www.alibabacloud.com/help/en/sms/developer-reference/api-error-codes)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 状态码的描述。
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// OSS配置信息。
	Model *GetSmsOcrOssInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// 本次调用请求的 ID，是由阿里云为该请求生成的唯一标识符，可用于排查和定位问题。
	//
	// example:
	//
	// 25D5AFDE-xxxx-132E-8909-1FDC071DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口是否成功。取值：
	//
	// - true：调用成功。
	//
	// - false：调用失败。
	//
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
	// 签名使用的 AccessKey ID。
	//
	// example:
	//
	// bypFNbG******
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// bucket名称。
	//
	// example:
	//
	// 示例值
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// 过期时间戳，单位：秒。
	//
	// example:
	//
	// 1741521339
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Host 地址。
	//
	// example:
	//
	// http://***.oss-cn-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 签名策略。
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wMy0wOVQxMTo1NTozOS4wMDFaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0seyJidWNrZXQiOiJhbGljb20tZmMtbWVkaWEifSxbImVxIiwiJGtleSIsIjEwMDAwMDM1ODA4MjA2M1wv********
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 根据 AccessKey Secret 和 Policy 计算出的签名信息。调用 OSS API 时，OSS 验证该签名信息，从而确认请求的合法性。
	//
	// example:
	//
	// QvNTGC9DSLTeByP+ZWW******
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 策略路径。
	//
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
