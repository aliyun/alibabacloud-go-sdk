// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultIPSConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasicRules(v int32) *ModifyDefaultIPSConfigRequest
	GetBasicRules() *int32
	SetCtiRules(v int32) *ModifyDefaultIPSConfigRequest
	GetCtiRules() *int32
	SetLang(v string) *ModifyDefaultIPSConfigRequest
	GetLang() *string
	SetMaxSdl(v int64) *ModifyDefaultIPSConfigRequest
	GetMaxSdl() *int64
	SetPatchRules(v int32) *ModifyDefaultIPSConfigRequest
	GetPatchRules() *int32
	SetRuleClass(v int32) *ModifyDefaultIPSConfigRequest
	GetRuleClass() *int32
	SetRunMode(v int32) *ModifyDefaultIPSConfigRequest
	GetRunMode() *int32
}

type ModifyDefaultIPSConfigRequest struct {
	// Specifies whether to enable basic protection. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Specifies whether to enable threat intelligence. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	CtiRules *int32 `json:"CtiRules,omitempty" xml:"CtiRules,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum amount of traffic that can be processed by the sensitive data leak detection feature each day.
	//
	// example:
	//
	// 100
	MaxSdl *int64 `json:"MaxSdl,omitempty" xml:"MaxSdl,omitempty"`
	// Specifies whether to enable virtual patching. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	PatchRules *int32 `json:"PatchRules,omitempty" xml:"PatchRules,omitempty"`
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
	// 1
	RuleClass *int32 `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// 	- **1**: block mode
	//
	// 	- **0**: monitor mode
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s ModifyDefaultIPSConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefaultIPSConfigRequest) GetBasicRules() *int32 {
	return s.BasicRules
}

func (s *ModifyDefaultIPSConfigRequest) GetCtiRules() *int32 {
	return s.CtiRules
}

func (s *ModifyDefaultIPSConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyDefaultIPSConfigRequest) GetMaxSdl() *int64 {
	return s.MaxSdl
}

func (s *ModifyDefaultIPSConfigRequest) GetPatchRules() *int32 {
	return s.PatchRules
}

func (s *ModifyDefaultIPSConfigRequest) GetRuleClass() *int32 {
	return s.RuleClass
}

func (s *ModifyDefaultIPSConfigRequest) GetRunMode() *int32 {
	return s.RunMode
}

func (s *ModifyDefaultIPSConfigRequest) SetBasicRules(v int32) *ModifyDefaultIPSConfigRequest {
	s.BasicRules = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetCtiRules(v int32) *ModifyDefaultIPSConfigRequest {
	s.CtiRules = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetLang(v string) *ModifyDefaultIPSConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetMaxSdl(v int64) *ModifyDefaultIPSConfigRequest {
	s.MaxSdl = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetPatchRules(v int32) *ModifyDefaultIPSConfigRequest {
	s.PatchRules = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetRuleClass(v int32) *ModifyDefaultIPSConfigRequest {
	s.RuleClass = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetRunMode(v int32) *ModifyDefaultIPSConfigRequest {
	s.RunMode = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) Validate() error {
	return dara.Validate(s)
}
