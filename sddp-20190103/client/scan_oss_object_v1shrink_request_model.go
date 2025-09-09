// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanOssObjectV1ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ScanOssObjectV1ShrinkRequest
	GetBucketName() *string
	SetLang(v string) *ScanOssObjectV1ShrinkRequest
	GetLang() *string
	SetObjectKeyListShrink(v string) *ScanOssObjectV1ShrinkRequest
	GetObjectKeyListShrink() *string
	SetServiceRegionId(v string) *ScanOssObjectV1ShrinkRequest
	GetServiceRegionId() *string
	SetTemplateId(v int64) *ScanOssObjectV1ShrinkRequest
	GetTemplateId() *int64
}

type ScanOssObjectV1ShrinkRequest struct {
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// sddp-api-demo-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The objects in the OSS bucket that you want to scan. You can specify up to 50 objects at a time.
	//
	// This parameter is required.
	ObjectKeyListShrink *string `json:"ObjectKeyList,omitempty" xml:"ObjectKeyList,omitempty"`
	// The ID of the region in which the OSS bucket is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific classification template.
	//
	// >  You can call the **DescribeCategoryTemplateList*	- operation to query industry-specific classification templates. If you do not specify this parameter, the system automatically uses the main template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ScanOssObjectV1ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanOssObjectV1ShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScanOssObjectV1ShrinkRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ScanOssObjectV1ShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ScanOssObjectV1ShrinkRequest) GetObjectKeyListShrink() *string {
	return s.ObjectKeyListShrink
}

func (s *ScanOssObjectV1ShrinkRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *ScanOssObjectV1ShrinkRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ScanOssObjectV1ShrinkRequest) SetBucketName(v string) *ScanOssObjectV1ShrinkRequest {
	s.BucketName = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) SetLang(v string) *ScanOssObjectV1ShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) SetObjectKeyListShrink(v string) *ScanOssObjectV1ShrinkRequest {
	s.ObjectKeyListShrink = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) SetServiceRegionId(v string) *ScanOssObjectV1ShrinkRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) SetTemplateId(v int64) *ScanOssObjectV1ShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
