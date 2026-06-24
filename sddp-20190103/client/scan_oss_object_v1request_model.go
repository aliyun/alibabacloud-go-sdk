// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanOssObjectV1Request interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ScanOssObjectV1Request
	GetBucketName() *string
	SetLang(v string) *ScanOssObjectV1Request
	GetLang() *string
	SetObjectKeyList(v []*string) *ScanOssObjectV1Request
	GetObjectKeyList() []*string
	SetServiceRegionId(v string) *ScanOssObjectV1Request
	GetServiceRegionId() *string
	SetTemplateId(v int64) *ScanOssObjectV1Request
	GetTemplateId() *int64
}

type ScanOssObjectV1Request struct {
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// sddp-api-demo-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// A list of objects to scan in the OSS bucket. You can specify up to 50 ObjectKeys.
	//
	// This parameter is required.
	ObjectKeyList []*string `json:"ObjectKeyList,omitempty" xml:"ObjectKeyList,omitempty" type:"Repeated"`
	// The ID of the region where the OSS bucket is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific template.
	//
	// > Call **DescribeCategoryTemplateList*	- to get a list of templates. If you do not specify this parameter, the default active template is used for the scan.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ScanOssObjectV1Request) String() string {
	return dara.Prettify(s)
}

func (s ScanOssObjectV1Request) GoString() string {
	return s.String()
}

func (s *ScanOssObjectV1Request) GetBucketName() *string {
	return s.BucketName
}

func (s *ScanOssObjectV1Request) GetLang() *string {
	return s.Lang
}

func (s *ScanOssObjectV1Request) GetObjectKeyList() []*string {
	return s.ObjectKeyList
}

func (s *ScanOssObjectV1Request) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *ScanOssObjectV1Request) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ScanOssObjectV1Request) SetBucketName(v string) *ScanOssObjectV1Request {
	s.BucketName = &v
	return s
}

func (s *ScanOssObjectV1Request) SetLang(v string) *ScanOssObjectV1Request {
	s.Lang = &v
	return s
}

func (s *ScanOssObjectV1Request) SetObjectKeyList(v []*string) *ScanOssObjectV1Request {
	s.ObjectKeyList = v
	return s
}

func (s *ScanOssObjectV1Request) SetServiceRegionId(v string) *ScanOssObjectV1Request {
	s.ServiceRegionId = &v
	return s
}

func (s *ScanOssObjectV1Request) SetTemplateId(v int64) *ScanOssObjectV1Request {
	s.TemplateId = &v
	return s
}

func (s *ScanOssObjectV1Request) Validate() error {
	return dara.Validate(s)
}
