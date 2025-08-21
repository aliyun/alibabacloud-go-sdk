// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketLifecycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *DeleteBucketLifecycleRequest
	GetBucketName() *string
	SetRuleId(v string) *DeleteBucketLifecycleRequest
	GetRuleId() *string
}

type DeleteBucketLifecycleRequest struct {
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The ID of the rule. If this parameter is not specified, all rules are removed.
	//
	// example:
	//
	// 1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteBucketLifecycleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketLifecycleRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketLifecycleRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *DeleteBucketLifecycleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteBucketLifecycleRequest) SetBucketName(v string) *DeleteBucketLifecycleRequest {
	s.BucketName = &v
	return s
}

func (s *DeleteBucketLifecycleRequest) SetRuleId(v string) *DeleteBucketLifecycleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteBucketLifecycleRequest) Validate() error {
	return dara.Validate(s)
}
