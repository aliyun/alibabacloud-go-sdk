// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaskOssImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *MaskOssImageRequest
	GetBucketName() *string
	SetIsAlwaysUpload(v bool) *MaskOssImageRequest
	GetIsAlwaysUpload() *bool
	SetIsCoverObject(v bool) *MaskOssImageRequest
	GetIsCoverObject() *bool
	SetIsSupportRestore(v bool) *MaskOssImageRequest
	GetIsSupportRestore() *bool
	SetLang(v string) *MaskOssImageRequest
	GetLang() *string
	SetMaskRuleIdList(v string) *MaskOssImageRequest
	GetMaskRuleIdList() *string
	SetObjectKey(v string) *MaskOssImageRequest
	GetObjectKey() *string
	SetServiceRegionId(v string) *MaskOssImageRequest
	GetServiceRegionId() *string
}

type MaskOssImageRequest struct {
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// sddp-api-demo-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// Specifies whether to always upload the processed image.
	//
	// If you set this to false, the image is uploaded only if it matches a masking rule.
	//
	// If you set this to true, the processed image is always uploaded.
	//
	// example:
	//
	// true
	IsAlwaysUpload *bool `json:"IsAlwaysUpload,omitempty" xml:"IsAlwaysUpload,omitempty"`
	IsCoverObject  *bool `json:"IsCoverObject,omitempty" xml:"IsCoverObject,omitempty"`
	// Specifies whether you can restore the original image from the masked version.
	//
	// example:
	//
	// true
	IsSupportRestore *bool `json:"IsSupportRestore,omitempty" xml:"IsSupportRestore,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// A comma-separated list of masking rule IDs.
	//
	// The following rule IDs correspond to the listed sensitive data types:
	//
	// 3000: Images that contain ID card information (Chinese mainland)
	//
	// 3009: Images that contain license plate information (Chinese mainland)
	//
	// 3002: Images that contain faces
	//
	// 1002: Names (Simplified Chinese)
	//
	// 1003: Addresses (Chinese mainland)
	//
	// 4003: Unified Social Credit Code
	//
	// 63009: Images that contain eyes
	//
	// This parameter is required.
	//
	// example:
	//
	// 3000,3002
	MaskRuleIdList *string `json:"MaskRuleIdList,omitempty" xml:"MaskRuleIdList,omitempty"`
	// The full key of the object stored in OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// dir1/test.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The region where the bucket is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s MaskOssImageRequest) String() string {
	return dara.Prettify(s)
}

func (s MaskOssImageRequest) GoString() string {
	return s.String()
}

func (s *MaskOssImageRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *MaskOssImageRequest) GetIsAlwaysUpload() *bool {
	return s.IsAlwaysUpload
}

func (s *MaskOssImageRequest) GetIsCoverObject() *bool {
	return s.IsCoverObject
}

func (s *MaskOssImageRequest) GetIsSupportRestore() *bool {
	return s.IsSupportRestore
}

func (s *MaskOssImageRequest) GetLang() *string {
	return s.Lang
}

func (s *MaskOssImageRequest) GetMaskRuleIdList() *string {
	return s.MaskRuleIdList
}

func (s *MaskOssImageRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *MaskOssImageRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *MaskOssImageRequest) SetBucketName(v string) *MaskOssImageRequest {
	s.BucketName = &v
	return s
}

func (s *MaskOssImageRequest) SetIsAlwaysUpload(v bool) *MaskOssImageRequest {
	s.IsAlwaysUpload = &v
	return s
}

func (s *MaskOssImageRequest) SetIsCoverObject(v bool) *MaskOssImageRequest {
	s.IsCoverObject = &v
	return s
}

func (s *MaskOssImageRequest) SetIsSupportRestore(v bool) *MaskOssImageRequest {
	s.IsSupportRestore = &v
	return s
}

func (s *MaskOssImageRequest) SetLang(v string) *MaskOssImageRequest {
	s.Lang = &v
	return s
}

func (s *MaskOssImageRequest) SetMaskRuleIdList(v string) *MaskOssImageRequest {
	s.MaskRuleIdList = &v
	return s
}

func (s *MaskOssImageRequest) SetObjectKey(v string) *MaskOssImageRequest {
	s.ObjectKey = &v
	return s
}

func (s *MaskOssImageRequest) SetServiceRegionId(v string) *MaskOssImageRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *MaskOssImageRequest) Validate() error {
	return dara.Validate(s)
}
