// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySasContainerWebDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPathConfDTOList(v []*ModifySasContainerWebDefenseRuleRequestPathConfDTOList) *ModifySasContainerWebDefenseRuleRequest
	GetPathConfDTOList() []*ModifySasContainerWebDefenseRuleRequestPathConfDTOList
	SetRuleId(v int64) *ModifySasContainerWebDefenseRuleRequest
	GetRuleId() *int64
	SetRuleName(v string) *ModifySasContainerWebDefenseRuleRequest
	GetRuleName() *string
}

type ModifySasContainerWebDefenseRuleRequest struct {
	// The paths that are protected.
	PathConfDTOList []*ModifySasContainerWebDefenseRuleRequestPathConfDTOList `json:"PathConfDTOList,omitempty" xml:"PathConfDTOList,omitempty" type:"Repeated"`
	// The ID of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200634
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// wwwwwww
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifySasContainerWebDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySasContainerWebDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifySasContainerWebDefenseRuleRequest) GetPathConfDTOList() []*ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	return s.PathConfDTOList
}

func (s *ModifySasContainerWebDefenseRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifySasContainerWebDefenseRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifySasContainerWebDefenseRuleRequest) SetPathConfDTOList(v []*ModifySasContainerWebDefenseRuleRequestPathConfDTOList) *ModifySasContainerWebDefenseRuleRequest {
	s.PathConfDTOList = v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequest) SetRuleId(v int64) *ModifySasContainerWebDefenseRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequest) SetRuleName(v string) *ModifySasContainerWebDefenseRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}

type ModifySasContainerWebDefenseRuleRequestPathConfDTOList struct {
	// The backup paths.
	//
	// example:
	//
	// /tmp/test
	BackupPath *string `json:"BackupPath,omitempty" xml:"BackupPath,omitempty"`
	// The prevention mode. Valid values:
	//
	// 	- **block**
	//
	// 	- **audit**
	//
	// This parameter is required.
	//
	// example:
	//
	// audit
	DefenseMode *string `json:"DefenseMode,omitempty" xml:"DefenseMode,omitempty"`
	// The path that is protected.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test/home/
	DefensePath *string `json:"DefensePath,omitempty" xml:"DefensePath,omitempty"`
	// The files that are excluded.
	//
	// example:
	//
	// /usr/test
	ExcludeFile *string `json:"ExcludeFile,omitempty" xml:"ExcludeFile,omitempty"`
	// The paths to the files that are excluded.
	//
	// example:
	//
	// /test/home/qq
	ExcludeFilePath *string `json:"ExcludeFilePath,omitempty" xml:"ExcludeFilePath,omitempty"`
	// The types of the files that are excluded.
	//
	// example:
	//
	// php
	ExcludeFileType *string `json:"ExcludeFileType,omitempty" xml:"ExcludeFileType,omitempty"`
	// The protection mode. Valid values:
	//
	// 	- **0**: basic mode (whitelist)
	//
	// 	- **1**: complex mode (blacklist)
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	GuardType *int32 `json:"GuardType,omitempty" xml:"GuardType,omitempty"`
	// The files that are included.
	//
	// example:
	//
	// /home/admin/test
	IncludeFile *string `json:"IncludeFile,omitempty" xml:"IncludeFile,omitempty"`
	// The type of the files that are included.
	//
	// example:
	//
	// jsp
	IncludeFileType *string `json:"IncludeFileType,omitempty" xml:"IncludeFileType,omitempty"`
	// The path ID.
	//
	// example:
	//
	// 12345678
	PathConfId *int64 `json:"PathConfId,omitempty" xml:"PathConfId,omitempty"`
	// The processes that are added to the whitelist.
	ProcessPathList []*string `json:"ProcessPathList,omitempty" xml:"ProcessPathList,omitempty" type:"Repeated"`
}

func (s ModifySasContainerWebDefenseRuleRequestPathConfDTOList) String() string {
	return dara.Prettify(s)
}

func (s ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GoString() string {
	return s.String()
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetBackupPath() *string {
	return s.BackupPath
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetDefenseMode() *string {
	return s.DefenseMode
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetDefensePath() *string {
	return s.DefensePath
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetExcludeFile() *string {
	return s.ExcludeFile
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetExcludeFilePath() *string {
	return s.ExcludeFilePath
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetExcludeFileType() *string {
	return s.ExcludeFileType
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetGuardType() *int32 {
	return s.GuardType
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetIncludeFile() *string {
	return s.IncludeFile
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetIncludeFileType() *string {
	return s.IncludeFileType
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetPathConfId() *int64 {
	return s.PathConfId
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) GetProcessPathList() []*string {
	return s.ProcessPathList
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetBackupPath(v string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.BackupPath = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetDefenseMode(v string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.DefenseMode = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetDefensePath(v string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.DefensePath = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetExcludeFile(v string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.ExcludeFile = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetExcludeFilePath(v string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.ExcludeFilePath = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetExcludeFileType(v string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.ExcludeFileType = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetGuardType(v int32) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.GuardType = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetIncludeFile(v string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.IncludeFile = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetIncludeFileType(v string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.IncludeFileType = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetPathConfId(v int64) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.PathConfId = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) SetProcessPathList(v []*string) *ModifySasContainerWebDefenseRuleRequestPathConfDTOList {
	s.ProcessPathList = v
	return s
}

func (s *ModifySasContainerWebDefenseRuleRequestPathConfDTOList) Validate() error {
	return dara.Validate(s)
}
