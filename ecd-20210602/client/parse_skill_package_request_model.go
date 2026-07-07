// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseSkillPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOssObjectETag(v string) *ParseSkillPackageRequest
	GetOssObjectETag() *string
	SetOssObjectKey(v string) *ParseSkillPackageRequest
	GetOssObjectKey() *string
}

type ParseSkillPackageRequest struct {
	// The OSS ETag returned after the file is uploaded to OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1D9920C4858A60B70705A8765A******
	OssObjectETag *string `json:"OssObjectETag,omitempty" xml:"OssObjectETag,omitempty"`
	// The OSS path of the skill package.
	//
	// This parameter is required.
	OssObjectKey *string `json:"OssObjectKey,omitempty" xml:"OssObjectKey,omitempty"`
}

func (s ParseSkillPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s ParseSkillPackageRequest) GoString() string {
	return s.String()
}

func (s *ParseSkillPackageRequest) GetOssObjectETag() *string {
	return s.OssObjectETag
}

func (s *ParseSkillPackageRequest) GetOssObjectKey() *string {
	return s.OssObjectKey
}

func (s *ParseSkillPackageRequest) SetOssObjectETag(v string) *ParseSkillPackageRequest {
	s.OssObjectETag = &v
	return s
}

func (s *ParseSkillPackageRequest) SetOssObjectKey(v string) *ParseSkillPackageRequest {
	s.OssObjectKey = &v
	return s
}

func (s *ParseSkillPackageRequest) Validate() error {
	return dara.Validate(s)
}
