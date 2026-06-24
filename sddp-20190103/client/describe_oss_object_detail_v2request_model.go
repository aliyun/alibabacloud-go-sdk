// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssObjectDetailV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *DescribeOssObjectDetailV2Request
	GetBucketName() *string
	SetId(v string) *DescribeOssObjectDetailV2Request
	GetId() *string
	SetLang(v string) *DescribeOssObjectDetailV2Request
	GetLang() *string
	SetObjectKey(v string) *DescribeOssObjectDetailV2Request
	GetObjectKey() *string
	SetServiceRegionId(v string) *DescribeOssObjectDetailV2Request
	GetServiceRegionId() *string
	SetTemplateId(v int64) *DescribeOssObjectDetailV2Request
	GetTemplateId() *int64
}

type DescribeOssObjectDetailV2Request struct {
	// The name of the bucket.
	//
	// example:
	//
	// sddp-api-scan-demo
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The unique ID of the OSS object.
	//
	// > Call the [DescribeOssObjects](https://help.aliyun.com/document_detail/410152.html) operation to obtain the ID.
	//
	// example:
	//
	// 12300
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **en_us**. Valid values:
	//
	// - **en_us**: Simplified Chinese
	//
	// - **en_us**: English (US)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The full name of the object stored in OSS.
	//
	// example:
	//
	// dir1/test.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The ID of the region where the service is enabled. This is also the region where the bucket is located.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific template.
	//
	// > Call the [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html) operation to obtain the ID of the industry-specific template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeOssObjectDetailV2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectDetailV2Request) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailV2Request) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeOssObjectDetailV2Request) GetId() *string {
	return s.Id
}

func (s *DescribeOssObjectDetailV2Request) GetLang() *string {
	return s.Lang
}

func (s *DescribeOssObjectDetailV2Request) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *DescribeOssObjectDetailV2Request) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *DescribeOssObjectDetailV2Request) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeOssObjectDetailV2Request) SetBucketName(v string) *DescribeOssObjectDetailV2Request {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetId(v string) *DescribeOssObjectDetailV2Request {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetLang(v string) *DescribeOssObjectDetailV2Request {
	s.Lang = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetObjectKey(v string) *DescribeOssObjectDetailV2Request {
	s.ObjectKey = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetServiceRegionId(v string) *DescribeOssObjectDetailV2Request {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetTemplateId(v int64) *DescribeOssObjectDetailV2Request {
	s.TemplateId = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) Validate() error {
	return dara.Validate(s)
}
