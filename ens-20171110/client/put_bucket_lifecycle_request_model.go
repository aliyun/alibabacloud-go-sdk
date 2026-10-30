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
	// Specifies whether to allow overlapping prefixes. Valid values:
	//
	// - **true**: Overlapping prefixes are allowed.
	//
	// - **false*	- (default): Overlapping prefixes are not allowed.
	//
	// example:
	//
	// false
	AllowSameActionOverlap *bool `json:"AllowSameActionOverlap,omitempty" xml:"AllowSameActionOverlap,omitempty"`
	// The bucket name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The expiration date. The system performs lifecycle operations on objects that were last modified before this date.
	//
	// The date must be specified in the ISO 8601 standard and in UTC. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// >ExpirationDays and CreatedBeforeDate are mutually exclusive. You must specify one of them.
	//
	// example:
	//
	// 2023-10-12T05:45:00Z
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	// The number of days after the last modification time of an object before the lifecycle rule takes effect. The value must be a positive integer greater than 0.
	//
	// >ExpirationDays and CreatedBeforeDate are mutually exclusive. You must specify one of them.
	//
	// example:
	//
	// 5
	ExpirationDays *int64 `json:"ExpirationDays,omitempty" xml:"ExpirationDays,omitempty"`
	// The prefix to which the rule applies. You cannot specify the same prefix for multiple rules.
	//
	// - If you specify a prefix, the rule applies only to objects whose names match the prefix in the bucket.
	//
	// - If you leave the prefix empty, the rule applies to all objects in the bucket.
	//
	// example:
	//
	// image
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The unique ID that identifies the rule. The ID can be up to 255 bytes in length.
	//
	// - You do not need to specify this parameter when you create a rule. The system automatically generates a unique ID.
	//
	// - You must specify the RuleId when you update an existing rule. The rule corresponding to the RuleId must exist. Otherwise, an error is returned.
	//
	// example:
	//
	// b8f93xxxxx4881xxxxxc71d991
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the rule. Valid values:
	//
	// - **Enabled**: The rule is executed periodically.
	//
	// - **Disabled**: The rule is ignored.
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
