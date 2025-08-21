// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketLifecycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSameActionOverlap(v string) *PutBucketLifecycleRequest
	GetAllowSameActionOverlap() *string
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
	// Specifies whether to allow overlapped prefixes. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AllowSameActionOverlap *string `json:"AllowSameActionOverlap,omitempty" xml:"AllowSameActionOverlap,omitempty"`
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The expiration time. EOS executes a lifecycle rule for objects that were last updated before the expiration time.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  ExpirationDays and CreateBeforeDate are mutually exclusive.
	//
	// example:
	//
	// 2023-10-12T05:45:00Z
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	// The number of days from when the objects were last modified to when the lifecycle rule takes effect. The value must be a positive integer that is greater than 0.
	//
	// >  ExpirationDays and CreateBeforeDate are mutually exclusive.
	//
	// example:
	//
	// 5
	ExpirationDays *int64 `json:"ExpirationDays,omitempty" xml:"ExpirationDays,omitempty"`
	// The prefix of a object name. The prefix must be unique.
	//
	// 	- If you specify a prefix, the rule applies only to objects in the bucket that match the prefix.
	//
	// 	- If you do not specify a prefix, the rule applies to all objects in the bucket.
	//
	// example:
	//
	// image
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The unique ID of the rule. The ID of a lifecycle rule can be up to 255 bytes in length.
	//
	// 	- You do not need to configure this parameter when you create a rule. The system automatically generates a unique ID.
	//
	// 	- When you update a rule, you need to specify this parameter. Make sure that the rule specified by RuleId exists. Otherwise, an error occurs.
	//
	// example:
	//
	// 1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
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

func (s *PutBucketLifecycleRequest) GetAllowSameActionOverlap() *string {
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

func (s *PutBucketLifecycleRequest) SetAllowSameActionOverlap(v string) *PutBucketLifecycleRequest {
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
