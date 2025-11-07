// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOssUploadToken(v *DescribeVerifyTokenResponseBodyOssUploadToken) *DescribeVerifyTokenResponseBody
	GetOssUploadToken() *DescribeVerifyTokenResponseBodyOssUploadToken
	SetRequestId(v string) *DescribeVerifyTokenResponseBody
	GetRequestId() *string
	SetVerifyPageUrl(v string) *DescribeVerifyTokenResponseBody
	GetVerifyPageUrl() *string
	SetVerifyToken(v string) *DescribeVerifyTokenResponseBody
	GetVerifyToken() *string
}

type DescribeVerifyTokenResponseBody struct {
	// OSS upload token information.
	OssUploadToken *DescribeVerifyTokenResponseBodyOssUploadToken `json:"OssUploadToken,omitempty" xml:"OssUploadToken,omitempty" type:"Struct"`
	// The ID of this request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The entry link for the original H5 verification scheme, which has been discontinued and no longer supports new integrations. If you need to integrate an H5 verification scheme, it is recommended to use the [PC or mobile H5 web integration solution](https://help.aliyun.com/document_detail/173779.html) of financial-grade real-person authentication.
	//
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	VerifyPageUrl *string `json:"VerifyPageUrl,omitempty" xml:"VerifyPageUrl,omitempty"`
	// The token for this verification, used to link various interfaces in the verification request, valid for 30 minutes.
	//
	// example:
	//
	// c302c0797679457685410ee51a5ba375
	VerifyToken *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s DescribeVerifyTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseBody) GetOssUploadToken() *DescribeVerifyTokenResponseBodyOssUploadToken {
	return s.OssUploadToken
}

func (s *DescribeVerifyTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyTokenResponseBody) GetVerifyPageUrl() *string {
	return s.VerifyPageUrl
}

func (s *DescribeVerifyTokenResponseBody) GetVerifyToken() *string {
	return s.VerifyToken
}

func (s *DescribeVerifyTokenResponseBody) SetOssUploadToken(v *DescribeVerifyTokenResponseBodyOssUploadToken) *DescribeVerifyTokenResponseBody {
	s.OssUploadToken = v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetRequestId(v string) *DescribeVerifyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetVerifyPageUrl(v string) *DescribeVerifyTokenResponseBody {
	s.VerifyPageUrl = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) SetVerifyToken(v string) *DescribeVerifyTokenResponseBody {
	s.VerifyToken = &v
	return s
}

func (s *DescribeVerifyTokenResponseBody) Validate() error {
	if s.OssUploadToken != nil {
		if err := s.OssUploadToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyTokenResponseBodyOssUploadToken struct {
	// OSS file storage bucket.
	//
	// example:
	//
	// cloudauth-zhangjiakou-external
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Access endpoint.
	//
	// example:
	//
	// https://oss-cn-zhangjiakou.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// Expiration time. Expressed in timestamp format, unit: milliseconds.
	//
	// example:
	//
	// 1582636610000
	Expired *int64 `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The key required for file upload.
	//
	// example:
	//
	// STS.NU8rUBj****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// File storage path.
	//
	// example:
	//
	// prod/RdNLC@Ox2n-1s7NMt
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The secret required for file upload.
	//
	// example:
	//
	// FwmnyoqT8dHj7nJLuM67T****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The token required for file upload.
	//
	// example:
	//
	// uWia500nTS5knZaDzq4/KqpvhcLnO****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeVerifyTokenResponseBodyOssUploadToken) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyTokenResponseBodyOssUploadToken) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) GetEndPoint() *string {
	return s.EndPoint
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) GetExpired() *int64 {
	return s.Expired
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) GetKey() *string {
	return s.Key
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) GetPath() *string {
	return s.Path
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) GetSecret() *string {
	return s.Secret
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) GetToken() *string {
	return s.Token
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetBucket(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Bucket = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetEndPoint(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.EndPoint = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetExpired(v int64) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Expired = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetKey(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Key = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetPath(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Path = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetSecret(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Secret = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) SetToken(v string) *DescribeVerifyTokenResponseBodyOssUploadToken {
	s.Token = &v
	return s
}

func (s *DescribeVerifyTokenResponseBodyOssUploadToken) Validate() error {
	return dara.Validate(s)
}
