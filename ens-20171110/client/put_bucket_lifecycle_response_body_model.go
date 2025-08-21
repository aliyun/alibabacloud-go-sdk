// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketLifecycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutBucketLifecycleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *PutBucketLifecycleResponseBody
	GetRuleId() *string
}

type PutBucketLifecycleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 62373E71-5521-4620-8AAB-133CCE49357A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// b8f93xxxxx4881xxxxxc71d991
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s PutBucketLifecycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutBucketLifecycleResponseBody) GoString() string {
	return s.String()
}

func (s *PutBucketLifecycleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutBucketLifecycleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *PutBucketLifecycleResponseBody) SetRequestId(v string) *PutBucketLifecycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutBucketLifecycleResponseBody) SetRuleId(v string) *PutBucketLifecycleResponseBody {
	s.RuleId = &v
	return s
}

func (s *PutBucketLifecycleResponseBody) Validate() error {
	return dara.Validate(s)
}
