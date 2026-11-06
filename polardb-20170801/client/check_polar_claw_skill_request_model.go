// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckPolarClawSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketRegionId(v string) *CheckPolarClawSkillRequest
	GetBucketRegionId() *string
	SetBucketType(v string) *CheckPolarClawSkillRequest
	GetBucketType() *string
	SetOSSBucket(v string) *CheckPolarClawSkillRequest
	GetOSSBucket() *string
	SetRegionId(v string) *CheckPolarClawSkillRequest
	GetRegionId() *string
	SetSkillName(v string) *CheckPolarClawSkillRequest
	GetSkillName() *string
	SetSkillVersion(v string) *CheckPolarClawSkillRequest
	GetSkillVersion() *string
	SetSummary(v string) *CheckPolarClawSkillRequest
	GetSummary() *string
	SetTags(v string) *CheckPolarClawSkillRequest
	GetTags() *string
	SetTitle(v string) *CheckPolarClawSkillRequest
	GetTitle() *string
}

type CheckPolarClawSkillRequest struct {
	// The region ID of the OSS bucket that stores the skill files.
	//
	// example:
	//
	// cn-hangzhou
	BucketRegionId *string `json:"BucketRegionId,omitempty" xml:"BucketRegionId,omitempty"`
	// The type of the bucket used for skill files. UserOwned indicates a user-owned bucket. Hosted indicates a hosted bucket. You must have the required permissions to use a hosted bucket.
	//
	// example:
	//
	// UserOwned
	BucketType *string `json:"BucketType,omitempty" xml:"BucketType,omitempty"`
	// The name of the OSS bucket to be disassociated.
	//
	// example:
	//
	// oss-bucket-test
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The skill name.
	//
	// This parameter is required.
	//
	// example:
	//
	// polarclawhub
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The skill version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
	// The brief description of the skill.
	//
	// example:
	//
	// An example used to manage PolarDB skills
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The tag details.
	//
	// example:
	//
	// [\\"PolarDB\\",\\"programming\\"]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title.
	//
	// example:
	//
	// PolarClawHub
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CheckPolarClawSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckPolarClawSkillRequest) GoString() string {
	return s.String()
}

func (s *CheckPolarClawSkillRequest) GetBucketRegionId() *string {
	return s.BucketRegionId
}

func (s *CheckPolarClawSkillRequest) GetBucketType() *string {
	return s.BucketType
}

func (s *CheckPolarClawSkillRequest) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *CheckPolarClawSkillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckPolarClawSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *CheckPolarClawSkillRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *CheckPolarClawSkillRequest) GetSummary() *string {
	return s.Summary
}

func (s *CheckPolarClawSkillRequest) GetTags() *string {
	return s.Tags
}

func (s *CheckPolarClawSkillRequest) GetTitle() *string {
	return s.Title
}

func (s *CheckPolarClawSkillRequest) SetBucketRegionId(v string) *CheckPolarClawSkillRequest {
	s.BucketRegionId = &v
	return s
}

func (s *CheckPolarClawSkillRequest) SetBucketType(v string) *CheckPolarClawSkillRequest {
	s.BucketType = &v
	return s
}

func (s *CheckPolarClawSkillRequest) SetOSSBucket(v string) *CheckPolarClawSkillRequest {
	s.OSSBucket = &v
	return s
}

func (s *CheckPolarClawSkillRequest) SetRegionId(v string) *CheckPolarClawSkillRequest {
	s.RegionId = &v
	return s
}

func (s *CheckPolarClawSkillRequest) SetSkillName(v string) *CheckPolarClawSkillRequest {
	s.SkillName = &v
	return s
}

func (s *CheckPolarClawSkillRequest) SetSkillVersion(v string) *CheckPolarClawSkillRequest {
	s.SkillVersion = &v
	return s
}

func (s *CheckPolarClawSkillRequest) SetSummary(v string) *CheckPolarClawSkillRequest {
	s.Summary = &v
	return s
}

func (s *CheckPolarClawSkillRequest) SetTags(v string) *CheckPolarClawSkillRequest {
	s.Tags = &v
	return s
}

func (s *CheckPolarClawSkillRequest) SetTitle(v string) *CheckPolarClawSkillRequest {
	s.Title = &v
	return s
}

func (s *CheckPolarClawSkillRequest) Validate() error {
	return dara.Validate(s)
}
