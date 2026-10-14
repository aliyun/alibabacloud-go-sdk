// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketLifecycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSameActionOverlap(v bool) *PutBucketLifecycleRequest
	GetAllowSameActionOverlap() *bool
	SetBucketName(v string) *PutBucketLifecycleRequest
	GetBucketName() *string
	SetCreatedBeforeDate(v string) *PutBucketLifecycleRequest
	GetCreatedBeforeDate() *string
	SetExpirationDays(v int64) *PutBucketLifecycleRequest
	GetExpirationDays() *int64
	SetPrefix(v string) *PutBucketLifecycleRequest
	GetPrefix() *string
	SetRuleId(v string) *PutBucketLifecycleRequest
	GetRuleId() *string
	SetStatus(v string) *PutBucketLifecycleRequest
	GetStatus() *string
}

type PutBucketLifecycleRequest struct {
	// Whether to allow overlapping prefixes. Valid values:
	//
	// - **true**: Overlapping prefixes are allowed.
	//
	// - **false*	- (default): Overlapping prefixes are not allowed.
	//
	// example:
	//
	// false
	AllowSameActionOverlap *bool `json:"AllowSameActionOverlap,omitempty" xml:"AllowSameActionOverlap,omitempty"`
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The date before which objects expire. The system applies the lifecycle rule to objects last modified before this date.
	//
	// Use ISO 8601 format and UTC time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// > ExpirationDays and CreatedBeforeDate are mutually exclusive. Specify one.
	//
	// example:
	//
	// 2023-10-12T05:45:00Z
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	// The number of days after an object\\"s last modification when the lifecycle rule takes effect. Must be a positive integer greater than zero.
	//
	// > ExpirationDays and CreatedBeforeDate are mutually exclusive. Specify one.
	//
	// example:
	//
	// 5
	ExpirationDays *int64 `json:"ExpirationDays,omitempty" xml:"ExpirationDays,omitempty"`
	// The prefix that the rule applies to. You cannot set duplicate prefixes.
	//
	// - If you specify a prefix, the rule applies only to objects in the bucket whose names start with that prefix.
	//
	// - If you leave Prefix empty, the rule applies to all objects in the bucket.
	//
	// example:
	//
	// image
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// A unique ID that identifies the rule. Maximum length is 255 bytes.
	//
	// - Omit this parameter when creating a new rule. The system generates a unique ID automatically.
	//
	// - Specify RuleId when updating an existing rule. The rule must exist. Otherwise, the call fails.
	//
	// example:
	//
	// b8f93xxxxx4881xxxxxc71d991
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the rule. Valid values:
	//
	// - **Enabled**: The system runs the rule on schedule.
	//
	// - **Disabled**: The system ignores the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PutBucketLifecycleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketLifecycleRequest) GoString() string {
	return s.String()
}

func (s *PutBucketLifecycleRequest) GetAllowSameActionOverlap() *bool {
	return s.AllowSameActionOverlap
}

func (s *PutBucketLifecycleRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *PutBucketLifecycleRequest) GetCreatedBeforeDate() *string {
	return s.CreatedBeforeDate
}

func (s *PutBucketLifecycleRequest) GetExpirationDays() *int64 {
	return s.ExpirationDays
}

func (s *PutBucketLifecycleRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *PutBucketLifecycleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *PutBucketLifecycleRequest) GetStatus() *string {
	return s.Status
}

func (s *PutBucketLifecycleRequest) SetAllowSameActionOverlap(v bool) *PutBucketLifecycleRequest {
	s.AllowSameActionOverlap = &v
	return s
}

func (s *PutBucketLifecycleRequest) SetBucketName(v string) *PutBucketLifecycleRequest {
	s.BucketName = &v
	return s
}

func (s *PutBucketLifecycleRequest) SetCreatedBeforeDate(v string) *PutBucketLifecycleRequest {
	s.CreatedBeforeDate = &v
	return s
}

func (s *PutBucketLifecycleRequest) SetExpirationDays(v int64) *PutBucketLifecycleRequest {
	s.ExpirationDays = &v
	return s
}

func (s *PutBucketLifecycleRequest) SetPrefix(v string) *PutBucketLifecycleRequest {
	s.Prefix = &v
	return s
}

func (s *PutBucketLifecycleRequest) SetRuleId(v string) *PutBucketLifecycleRequest {
	s.RuleId = &v
	return s
}

func (s *PutBucketLifecycleRequest) SetStatus(v string) *PutBucketLifecycleRequest {
	s.Status = &v
	return s
}

func (s *PutBucketLifecycleRequest) Validate() error {
	return dara.Validate(s)
}
