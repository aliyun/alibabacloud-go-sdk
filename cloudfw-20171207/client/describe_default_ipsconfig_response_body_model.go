// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultIPSConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBasicRules(v int32) *DescribeDefaultIPSConfigResponseBody
	GetBasicRules() *int32
	SetCtiRules(v int32) *DescribeDefaultIPSConfigResponseBody
	GetCtiRules() *int32
	SetMaxSdl(v int64) *DescribeDefaultIPSConfigResponseBody
	GetMaxSdl() *int64
	SetPatchRules(v int32) *DescribeDefaultIPSConfigResponseBody
	GetPatchRules() *int32
	SetRequestId(v string) *DescribeDefaultIPSConfigResponseBody
	GetRequestId() *string
	SetRuleClass(v int32) *DescribeDefaultIPSConfigResponseBody
	GetRuleClass() *int32
	SetRunMode(v int32) *DescribeDefaultIPSConfigResponseBody
	GetRunMode() *int32
}

type DescribeDefaultIPSConfigResponseBody struct {
	// Indicates whether basic protection is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether threat intelligence is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	CtiRules *int32 `json:"CtiRules,omitempty" xml:"CtiRules,omitempty"`
	// The maximum amount of traffic that can be processed by the sensitive data leak detection feature each day.
	//
	// example:
	//
	// 10
	MaxSdl *int64 `json:"MaxSdl,omitempty" xml:"MaxSdl,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	PatchRules *int32 `json:"PatchRules,omitempty" xml:"PatchRules,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 133173B9-8010-5DF5-8B93-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The level of the rule group for the IPS. Valid values:
	//
	// 	- **1**: loose
	//
	// 	- **2**: medium
	//
	// 	- **3**: strict
	//
	// example:
	//
	// 3
	RuleClass *int32 `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// 	- **1**: block mode
	//
	// 	- **0**: monitor mode
	//
	// example:
	//
	// 0
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeDefaultIPSConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefaultIPSConfigResponseBody) GetBasicRules() *int32 {
	return s.BasicRules
}

func (s *DescribeDefaultIPSConfigResponseBody) GetCtiRules() *int32 {
	return s.CtiRules
}

func (s *DescribeDefaultIPSConfigResponseBody) GetMaxSdl() *int64 {
	return s.MaxSdl
}

func (s *DescribeDefaultIPSConfigResponseBody) GetPatchRules() *int32 {
	return s.PatchRules
}

func (s *DescribeDefaultIPSConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefaultIPSConfigResponseBody) GetRuleClass() *int32 {
	return s.RuleClass
}

func (s *DescribeDefaultIPSConfigResponseBody) GetRunMode() *int32 {
	return s.RunMode
}

func (s *DescribeDefaultIPSConfigResponseBody) SetBasicRules(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.BasicRules = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetCtiRules(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.CtiRules = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetMaxSdl(v int64) *DescribeDefaultIPSConfigResponseBody {
	s.MaxSdl = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetPatchRules(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.PatchRules = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetRequestId(v string) *DescribeDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetRuleClass(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.RuleClass = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetRunMode(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.RunMode = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
