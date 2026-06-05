// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDataUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeComfyUserDataUploadUrlResponseBody
	GetCode() *int64
	SetExpiredTime(v string) *DescribeComfyUserDataUploadUrlResponseBody
	GetExpiredTime() *string
	SetMessage(v string) *DescribeComfyUserDataUploadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeComfyUserDataUploadUrlResponseBody
	GetRequestId() *string
	SetUploadUrl(v string) *DescribeComfyUserDataUploadUrlResponseBody
	GetUploadUrl() *string
}

type DescribeComfyUserDataUploadUrlResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2026-08-28T16:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// oss-bucket.oss-cn-shanghai-cloudspe.aliyuncs.com/path/userid/test.png?x-oss-date=20260602T093618Z&x-oss-expires=29&x-oss-security-token=xxxxx%2F68FY2NI6nyNgueynMjneKjKXD6l5tS5h3S2%xxxxx%xxxx%xxxxxxx%2Fb%xxxx%sdffeeeaasdf%xxx%xx%xx%xxxx%2BeASOv2N8q%xxx%2B6XBxllroojFliSTfsqGI2YMHpoRwJmyfXK32BQPb2SvQ0AM23soq%2FspAI4f9vtFG0yv9fKWMw%xxxxxxxxx%xxxxxxxx%3D%3D&x-oss-signature-version=OSS4-HMAC&x-oss-credential=STS.xxxxx%2F20260dfa202%2Fcn-shanghai-cloud%2Foss%2Faliyun_v4_request&x-oss-signature=xxxxxxxxxxxxxx
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s DescribeComfyUserDataUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDataUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) SetCode(v int64) *DescribeComfyUserDataUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) SetExpiredTime(v string) *DescribeComfyUserDataUploadUrlResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) SetMessage(v string) *DescribeComfyUserDataUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) SetRequestId(v string) *DescribeComfyUserDataUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) SetUploadUrl(v string) *DescribeComfyUserDataUploadUrlResponseBody {
	s.UploadUrl = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
