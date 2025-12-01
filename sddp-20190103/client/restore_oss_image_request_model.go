// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreOssImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *RestoreOssImageRequest
	GetBucket() *string
	SetLang(v string) *RestoreOssImageRequest
	GetLang() *string
	SetObjectKey(v string) *RestoreOssImageRequest
	GetObjectKey() *string
	SetServiceRegionId(v string) *RestoreOssImageRequest
	GetServiceRegionId() *string
	SetTargetObjectKey(v string) *RestoreOssImageRequest
	GetTargetObjectKey() *string
}

type RestoreOssImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-sddp
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aliyun_dsc_desensitization/dir1/test.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// example:
	//
	// aliyun_dsc_original /dir1/test.png
	TargetObjectKey *string `json:"TargetObjectKey,omitempty" xml:"TargetObjectKey,omitempty"`
}

func (s RestoreOssImageRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreOssImageRequest) GoString() string {
	return s.String()
}

func (s *RestoreOssImageRequest) GetBucket() *string {
	return s.Bucket
}

func (s *RestoreOssImageRequest) GetLang() *string {
	return s.Lang
}

func (s *RestoreOssImageRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *RestoreOssImageRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *RestoreOssImageRequest) GetTargetObjectKey() *string {
	return s.TargetObjectKey
}

func (s *RestoreOssImageRequest) SetBucket(v string) *RestoreOssImageRequest {
	s.Bucket = &v
	return s
}

func (s *RestoreOssImageRequest) SetLang(v string) *RestoreOssImageRequest {
	s.Lang = &v
	return s
}

func (s *RestoreOssImageRequest) SetObjectKey(v string) *RestoreOssImageRequest {
	s.ObjectKey = &v
	return s
}

func (s *RestoreOssImageRequest) SetServiceRegionId(v string) *RestoreOssImageRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *RestoreOssImageRequest) SetTargetObjectKey(v string) *RestoreOssImageRequest {
	s.TargetObjectKey = &v
	return s
}

func (s *RestoreOssImageRequest) Validate() error {
	return dara.Validate(s)
}
