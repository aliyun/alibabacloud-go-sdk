// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketRegionId(v string) *CreatePolarClawSkillRequest
	GetBucketRegionId() *string
	SetBucketType(v string) *CreatePolarClawSkillRequest
	GetBucketType() *string
	SetOSSBucket(v string) *CreatePolarClawSkillRequest
	GetOSSBucket() *string
	SetOSSPrefix(v string) *CreatePolarClawSkillRequest
	GetOSSPrefix() *string
	SetRegionId(v string) *CreatePolarClawSkillRequest
	GetRegionId() *string
	SetSkillName(v string) *CreatePolarClawSkillRequest
	GetSkillName() *string
	SetSkillVersion(v string) *CreatePolarClawSkillRequest
	GetSkillVersion() *string
	SetSummary(v string) *CreatePolarClawSkillRequest
	GetSummary() *string
	SetTags(v string) *CreatePolarClawSkillRequest
	GetTags() *string
	SetTitle(v string) *CreatePolarClawSkillRequest
	GetTitle() *string
}

type CreatePolarClawSkillRequest struct {
	// example:
	//
	// cn-hangzhou
	BucketRegionId *string `json:"BucketRegionId,omitempty" xml:"BucketRegionId,omitempty"`
	// example:
	//
	// UserOwned
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss-bucket-test
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// steelStorage
	OSSPrefix *string `json:"OSSPrefix,omitempty" xml:"OSSPrefix,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polarclawhub
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
	// example:
	//
	// 用于管理 PolarDB 技能的示例。
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// [\\"PolarDB\\"]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PolarClawHub
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreatePolarClawSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawSkillRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarClawSkillRequest) GetBucketRegionId() *string {
	return s.BucketRegionId
}

func (s *CreatePolarClawSkillRequest) GetBucketType() *string {
	return s.BucketType
}

func (s *CreatePolarClawSkillRequest) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *CreatePolarClawSkillRequest) GetOSSPrefix() *string {
	return s.OSSPrefix
}

func (s *CreatePolarClawSkillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolarClawSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *CreatePolarClawSkillRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *CreatePolarClawSkillRequest) GetSummary() *string {
	return s.Summary
}

func (s *CreatePolarClawSkillRequest) GetTags() *string {
	return s.Tags
}

func (s *CreatePolarClawSkillRequest) GetTitle() *string {
	return s.Title
}

func (s *CreatePolarClawSkillRequest) SetBucketRegionId(v string) *CreatePolarClawSkillRequest {
	s.BucketRegionId = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetBucketType(v string) *CreatePolarClawSkillRequest {
	s.BucketType = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetOSSBucket(v string) *CreatePolarClawSkillRequest {
	s.OSSBucket = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetOSSPrefix(v string) *CreatePolarClawSkillRequest {
	s.OSSPrefix = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetRegionId(v string) *CreatePolarClawSkillRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetSkillName(v string) *CreatePolarClawSkillRequest {
	s.SkillName = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetSkillVersion(v string) *CreatePolarClawSkillRequest {
	s.SkillVersion = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetSummary(v string) *CreatePolarClawSkillRequest {
	s.Summary = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetTags(v string) *CreatePolarClawSkillRequest {
	s.Tags = &v
	return s
}

func (s *CreatePolarClawSkillRequest) SetTitle(v string) *CreatePolarClawSkillRequest {
	s.Title = &v
	return s
}

func (s *CreatePolarClawSkillRequest) Validate() error {
	return dara.Validate(s)
}
