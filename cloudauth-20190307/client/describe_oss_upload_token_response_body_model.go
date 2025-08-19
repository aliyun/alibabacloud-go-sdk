// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssUploadTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOssUploadToken(v *DescribeOssUploadTokenResponseBodyOssUploadToken) *DescribeOssUploadTokenResponseBody
	GetOssUploadToken() *DescribeOssUploadTokenResponseBodyOssUploadToken
	SetRequestId(v string) *DescribeOssUploadTokenResponseBody
	GetRequestId() *string
}

type DescribeOssUploadTokenResponseBody struct {
	OssUploadToken *DescribeOssUploadTokenResponseBodyOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" type:"Struct"`
	// example:
	//
	// 2FA2C773-47DB-4156-B1EE-5B047321A939
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOssUploadTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponseBody) GetOssUploadToken() *DescribeOssUploadTokenResponseBodyOssUploadToken {
	return s.OssUploadToken
}

func (s *DescribeOssUploadTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssUploadTokenResponseBody) SetOssUploadToken(v *DescribeOssUploadTokenResponseBodyOssUploadToken) *DescribeOssUploadTokenResponseBody {
	s.OssUploadToken = v
	return s
}

func (s *DescribeOssUploadTokenResponseBody) SetRequestId(v string) *DescribeOssUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOssUploadTokenResponseBodyOssUploadToken struct {
	// example:
	//
	// cloudauth-zhangjiakou-external
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// https://oss-cn-zhangjiakou.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// example:
	//
	// 1582636610000
	Expired *int64 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// STS.NU8rUBj****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// prod/RdNLC@Ox2n-1s7NMt
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// FwmnyoqT8dHj7nJLuM67T****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// example:
	//
	// uWia500nTS5knZaDzq4/KqpvhcLnO****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeOssUploadTokenResponseBodyOssUploadToken) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssUploadTokenResponseBodyOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) GetEndPoint() *string {
	return s.EndPoint
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) GetExpired() *int64 {
	return s.Expired
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) GetKey() *string {
	return s.Key
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) GetPath() *string {
	return s.Path
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) GetSecret() *string {
	return s.Secret
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) GetToken() *string {
	return s.Token
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetBucket(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetEndPoint(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetExpired(v int64) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetKey(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetPath(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetSecret(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) SetToken(v string) *DescribeOssUploadTokenResponseBodyOssUploadToken {
	s.Token = &v
	return s
}

func (s *DescribeOssUploadTokenResponseBodyOssUploadToken) Validate() error {
	return dara.Validate(s)
}
