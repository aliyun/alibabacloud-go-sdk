// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketLifecycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetBucketLifecycleResponseBody
	GetRequestId() *string
	SetRule(v []*GetBucketLifecycleResponseBodyRule) *GetBucketLifecycleResponseBody
	GetRule() []*GetBucketLifecycleResponseBodyRule
}

type GetBucketLifecycleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A2583E8B-B930-4F59-ADC0-0E209A90C46E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about the rule.
	Rule []*GetBucketLifecycleResponseBodyRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s GetBucketLifecycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLifecycleResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketLifecycleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBucketLifecycleResponseBody) GetRule() []*GetBucketLifecycleResponseBodyRule {
	return s.Rule
}

func (s *GetBucketLifecycleResponseBody) SetRequestId(v string) *GetBucketLifecycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBucketLifecycleResponseBody) SetRule(v []*GetBucketLifecycleResponseBodyRule) *GetBucketLifecycleResponseBody {
	s.Rule = v
	return s
}

func (s *GetBucketLifecycleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBucketLifecycleResponseBodyRule struct {
	// The expiration time.
	Expiration *GetBucketLifecycleResponseBodyRuleExpiration `json:"Expiration,omitempty" xml:"Expiration,omitempty" type:"Struct"`
	// The unique ID of the rule.
	//
	// example:
	//
	// 1
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// The prefix that is applied to the rule.
	//
	// example:
	//
	// image
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **Enabled**: The rule is periodically executed.
	//
	// 	- **Disabled**: The rule is ignored.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBucketLifecycleResponseBodyRule) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLifecycleResponseBodyRule) GoString() string {
	return s.String()
}

func (s *GetBucketLifecycleResponseBodyRule) GetExpiration() *GetBucketLifecycleResponseBodyRuleExpiration {
	return s.Expiration
}

func (s *GetBucketLifecycleResponseBodyRule) GetID() *string {
	return s.ID
}

func (s *GetBucketLifecycleResponseBodyRule) GetPrefix() *string {
	return s.Prefix
}

func (s *GetBucketLifecycleResponseBodyRule) GetStatus() *string {
	return s.Status
}

func (s *GetBucketLifecycleResponseBodyRule) SetExpiration(v *GetBucketLifecycleResponseBodyRuleExpiration) *GetBucketLifecycleResponseBodyRule {
	s.Expiration = v
	return s
}

func (s *GetBucketLifecycleResponseBodyRule) SetID(v string) *GetBucketLifecycleResponseBodyRule {
	s.ID = &v
	return s
}

func (s *GetBucketLifecycleResponseBodyRule) SetPrefix(v string) *GetBucketLifecycleResponseBodyRule {
	s.Prefix = &v
	return s
}

func (s *GetBucketLifecycleResponseBodyRule) SetStatus(v string) *GetBucketLifecycleResponseBodyRule {
	s.Status = &v
	return s
}

func (s *GetBucketLifecycleResponseBodyRule) Validate() error {
	return dara.Validate(s)
}

type GetBucketLifecycleResponseBodyRuleExpiration struct {
	// The expiration date.
	//
	// example:
	//
	// yyy-MM-DDThh:mm:ssZ
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	// The validity period, in days.
	//
	// example:
	//
	// 5
	Days *string `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s GetBucketLifecycleResponseBodyRuleExpiration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLifecycleResponseBodyRuleExpiration) GoString() string {
	return s.String()
}

func (s *GetBucketLifecycleResponseBodyRuleExpiration) GetCreatedBeforeDate() *string {
	return s.CreatedBeforeDate
}

func (s *GetBucketLifecycleResponseBodyRuleExpiration) GetDays() *string {
	return s.Days
}

func (s *GetBucketLifecycleResponseBodyRuleExpiration) SetCreatedBeforeDate(v string) *GetBucketLifecycleResponseBodyRuleExpiration {
	s.CreatedBeforeDate = &v
	return s
}

func (s *GetBucketLifecycleResponseBodyRuleExpiration) SetDays(v string) *GetBucketLifecycleResponseBodyRuleExpiration {
	s.Days = &v
	return s
}

func (s *GetBucketLifecycleResponseBodyRuleExpiration) Validate() error {
	return dara.Validate(s)
}
