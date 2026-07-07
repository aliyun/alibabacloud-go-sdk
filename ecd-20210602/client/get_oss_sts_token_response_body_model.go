// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssStsTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *GetOssStsTokenResponseBody
	GetAccessKeyId() *string
	SetAccessKeySecret(v string) *GetOssStsTokenResponseBody
	GetAccessKeySecret() *string
	SetBucket(v string) *GetOssStsTokenResponseBody
	GetBucket() *string
	SetObjectKeyPrefix(v string) *GetOssStsTokenResponseBody
	GetObjectKeyPrefix() *string
	SetOssRegion(v string) *GetOssStsTokenResponseBody
	GetOssRegion() *string
	SetRequestId(v string) *GetOssStsTokenResponseBody
	GetRequestId() *string
	SetSecurityToken(v string) *GetOssStsTokenResponseBody
	GetSecurityToken() *string
}

type GetOssStsTokenResponseBody struct {
	// example:
	//
	// STS.NZeNA1kdCm4QPuAJ9kN******
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 9EStV7fgkSQsPuBi576EmNQXLxJGddL2EGyX********
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// prod-wy-*****
	Bucket          *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	ObjectKeyPrefix *string `json:"ObjectKeyPrefix,omitempty" xml:"ObjectKeyPrefix,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	OssRegion *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// CAISvAN1q6Ft5B2yfSjIr5n2Bez81ZRTgqOGZn6FkHBnXf9qgI6apjz2IH*******
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetOssStsTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssStsTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssStsTokenResponseBody) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetOssStsTokenResponseBody) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetOssStsTokenResponseBody) GetBucket() *string {
	return s.Bucket
}

func (s *GetOssStsTokenResponseBody) GetObjectKeyPrefix() *string {
	return s.ObjectKeyPrefix
}

func (s *GetOssStsTokenResponseBody) GetOssRegion() *string {
	return s.OssRegion
}

func (s *GetOssStsTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssStsTokenResponseBody) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetOssStsTokenResponseBody) SetAccessKeyId(v string) *GetOssStsTokenResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetOssStsTokenResponseBody) SetAccessKeySecret(v string) *GetOssStsTokenResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *GetOssStsTokenResponseBody) SetBucket(v string) *GetOssStsTokenResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetOssStsTokenResponseBody) SetObjectKeyPrefix(v string) *GetOssStsTokenResponseBody {
	s.ObjectKeyPrefix = &v
	return s
}

func (s *GetOssStsTokenResponseBody) SetOssRegion(v string) *GetOssStsTokenResponseBody {
	s.OssRegion = &v
	return s
}

func (s *GetOssStsTokenResponseBody) SetRequestId(v string) *GetOssStsTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssStsTokenResponseBody) SetSecurityToken(v string) *GetOssStsTokenResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *GetOssStsTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
