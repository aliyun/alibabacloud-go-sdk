// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRuleSnapshotRequest
	GetLang() *string
	SetRegId(v string) *DescribeRuleSnapshotRequest
	GetRegId() *string
	SetRuleId(v string) *DescribeRuleSnapshotRequest
	GetRuleId() *string
	SetSnapshotVersion(v string) *DescribeRuleSnapshotRequest
	GetSnapshotVersion() *string
}

type DescribeRuleSnapshotRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 101544
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Snapshot version.
	//
	// example:
	//
	// 10
	SnapshotVersion *string `json:"snapshotVersion,omitempty" xml:"snapshotVersion,omitempty"`
}

func (s DescribeRuleSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleSnapshotRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRuleSnapshotRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRuleSnapshotRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleSnapshotRequest) GetSnapshotVersion() *string {
	return s.SnapshotVersion
}

func (s *DescribeRuleSnapshotRequest) SetLang(v string) *DescribeRuleSnapshotRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleSnapshotRequest) SetRegId(v string) *DescribeRuleSnapshotRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRuleSnapshotRequest) SetRuleId(v string) *DescribeRuleSnapshotRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleSnapshotRequest) SetSnapshotVersion(v string) *DescribeRuleSnapshotRequest {
	s.SnapshotVersion = &v
	return s
}

func (s *DescribeRuleSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
