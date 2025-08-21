// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketLifecycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *GetBucketLifecycleRequest
	GetBucketName() *string
	SetRuleId(v string) *GetBucketLifecycleRequest
	GetRuleId() *string
}

type GetBucketLifecycleRequest struct {
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// tese
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The ID of the rule that you want to query. If this parameter is not specified, all rules are returned.
	//
	// example:
	//
	// 1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetBucketLifecycleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLifecycleRequest) GoString() string {
	return s.String()
}

func (s *GetBucketLifecycleRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *GetBucketLifecycleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetBucketLifecycleRequest) SetBucketName(v string) *GetBucketLifecycleRequest {
	s.BucketName = &v
	return s
}

func (s *GetBucketLifecycleRequest) SetRuleId(v string) *GetBucketLifecycleRequest {
	s.RuleId = &v
	return s
}

func (s *GetBucketLifecycleRequest) Validate() error {
	return dara.Validate(s)
}
