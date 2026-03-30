// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PreviewVoiceResponseBody
	GetCode() *string
	SetData(v string) *PreviewVoiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *PreviewVoiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PreviewVoiceResponseBody
	GetMessage() *string
	SetParams(v []*string) *PreviewVoiceResponseBody
	GetParams() []*string
	SetRequestId(v string) *PreviewVoiceResponseBody
	GetRequestId() *string
}

type PreviewVoiceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://cab-config-pre.oss-cn-hangzhou.aliyuncs.com/TTSDEMO/05B9FBF3-3681-10FF-9C24-340A93F3A25F.wav?Expires=1774795253&OSSAccessKeyId=STS.NYGg2ejEjYqySx3EsuRutagbd&Signature=L7oa%2F7s%2FeVwBxpkn3SvKfr8uXB0%3D&security-token=CAISzwJ1q6Ft5B2yfSjIr5ryLIjRh5pL7rOSUV6CoXMgXvpYjqLJhjz2IHhMfnlvB%2BgYsfU2m2xR5%2FYclrp6SJtIXleCZtF94oxN9h2gb4fb42Jqag%2B%2F08%2FLI3OaLjKm9u2wCryLYbGwU%2FOpbE%2B%2B5U0X6LDmdDKkckW4OJmS8%2FBOZcgWWQ%2FKBlgvRq0hRG1YpdQdKGHaONu0LxfumRCwNkdzvRdmgm4NgsbWgO%2Fks0GG3ASmlrFF%2B9mufMb5M%2FMBZskvD42Hu8VtbbfE3SJq7BxHybx7lqQs%2B02c5onHUwYPu0vZYrOLroQ%2BfFFjHKMzEetPq%2F7ylPI9ofDamIXxxAarin3kufQeLmrJ4LwneIvBXr5RHd5wa2rbWAEsmLNBEhL2EJMKtT476hcbIAuUI3bC5F%2BkxOHp9i6ErImtRWbLssUUla4R5TGOWbLJWzkTH93xuRqAAapuIRuRt0d2Odr1hsaYukMd42UkNapdTrehzmXeR6lyv1jlLmkAHve9Cbl9N5bO3A96FSlEfjHksQBWG0CEXRm3jLW41bpR00dgnM6gpOj7lRW2z33L0dTtaRw79X3%2BUqz3gv9md5QvoaVi1jnr%2FcFRNxbjl7DI39pdcGlTI2lqIAA%3D
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreviewVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *PreviewVoiceResponseBody) GetData() *string {
	return s.Data
}

func (s *PreviewVoiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PreviewVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreviewVoiceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *PreviewVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewVoiceResponseBody) SetCode(v string) *PreviewVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *PreviewVoiceResponseBody) SetData(v string) *PreviewVoiceResponseBody {
	s.Data = &v
	return s
}

func (s *PreviewVoiceResponseBody) SetHttpStatusCode(v int32) *PreviewVoiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PreviewVoiceResponseBody) SetMessage(v string) *PreviewVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *PreviewVoiceResponseBody) SetParams(v []*string) *PreviewVoiceResponseBody {
	s.Params = v
	return s
}

func (s *PreviewVoiceResponseBody) SetRequestId(v string) *PreviewVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
