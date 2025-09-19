// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSasContainerWebDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPathConfDTOList(v []*AddSasContainerWebDefenseRuleRequestPathConfDTOList) *AddSasContainerWebDefenseRuleRequest
	GetPathConfDTOList() []*AddSasContainerWebDefenseRuleRequestPathConfDTOList
	SetRuleName(v string) *AddSasContainerWebDefenseRuleRequest
	GetRuleName() *string
}

type AddSasContainerWebDefenseRuleRequest struct {
	// The paths that you want to protect.
	PathConfDTOList []*AddSasContainerWebDefenseRuleRequestPathConfDTOList `json:"PathConfDTOList,omitempty" xml:"PathConfDTOList,omitempty" type:"Repeated"`
	// The name of the rule.
	//
	// example:
	//
	// test-2020
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s AddSasContainerWebDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSasContainerWebDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *AddSasContainerWebDefenseRuleRequest) GetPathConfDTOList() []*AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	return s.PathConfDTOList
}

func (s *AddSasContainerWebDefenseRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *AddSasContainerWebDefenseRuleRequest) SetPathConfDTOList(v []*AddSasContainerWebDefenseRuleRequestPathConfDTOList) *AddSasContainerWebDefenseRuleRequest {
	s.PathConfDTOList = v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequest) SetRuleName(v string) *AddSasContainerWebDefenseRuleRequest {
	s.RuleName = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}

type AddSasContainerWebDefenseRuleRequestPathConfDTOList struct {
	// The backup path.
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
	// example:
	//
	// audit
	DefenseMode *string `json:"DefenseMode,omitempty" xml:"DefenseMode,omitempty"`
	// The path that you want to protect.
	//
	// This parameter is required.
	//
	// example:
	//
	// /usr/test/
	DefensePath *string `json:"DefensePath,omitempty" xml:"DefensePath,omitempty"`
	// The file that you want to exclude.
	//
	// example:
	//
	// /usr/test/aa
	ExcludeFile *string `json:"ExcludeFile,omitempty" xml:"ExcludeFile,omitempty"`
	// The path to the file that you want to exclude.
	//
	// example:
	//
	// /usr/test/tt
	ExcludeFilePath *string `json:"ExcludeFilePath,omitempty" xml:"ExcludeFilePath,omitempty"`
	// The type of the file that you want to exclude.
	//
	// example:
	//
	// jsp
	ExcludeFileType *string `json:"ExcludeFileType,omitempty" xml:"ExcludeFileType,omitempty"`
	// The protecion mode. Valid values:
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
	// The file that you want to include.
	//
	// example:
	//
	// /usr/test/t1
	IncludeFile *string `json:"IncludeFile,omitempty" xml:"IncludeFile,omitempty"`
	// The type of the file that you want to include.
	//
	// example:
	//
	// *.jsp
	IncludeFileType *string `json:"IncludeFileType,omitempty" xml:"IncludeFileType,omitempty"`
	// The processes that you want to add to the whitelist.
	//
	// This parameter is required.
	ProcessPathList []*string `json:"ProcessPathList,omitempty" xml:"ProcessPathList,omitempty" type:"Repeated"`
}

func (s AddSasContainerWebDefenseRuleRequestPathConfDTOList) String() string {
	return dara.Prettify(s)
}

func (s AddSasContainerWebDefenseRuleRequestPathConfDTOList) GoString() string {
	return s.String()
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetBackupPath() *string {
	return s.BackupPath
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetDefenseMode() *string {
	return s.DefenseMode
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetDefensePath() *string {
	return s.DefensePath
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetExcludeFile() *string {
	return s.ExcludeFile
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetExcludeFilePath() *string {
	return s.ExcludeFilePath
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetExcludeFileType() *string {
	return s.ExcludeFileType
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetGuardType() *int32 {
	return s.GuardType
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetIncludeFile() *string {
	return s.IncludeFile
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetIncludeFileType() *string {
	return s.IncludeFileType
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) GetProcessPathList() []*string {
	return s.ProcessPathList
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetBackupPath(v string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.BackupPath = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetDefenseMode(v string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.DefenseMode = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetDefensePath(v string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.DefensePath = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetExcludeFile(v string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.ExcludeFile = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetExcludeFilePath(v string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.ExcludeFilePath = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetExcludeFileType(v string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.ExcludeFileType = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetGuardType(v int32) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.GuardType = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetIncludeFile(v string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.IncludeFile = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetIncludeFileType(v string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.IncludeFileType = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) SetProcessPathList(v []*string) *AddSasContainerWebDefenseRuleRequestPathConfDTOList {
	s.ProcessPathList = v
	return s
}

func (s *AddSasContainerWebDefenseRuleRequestPathConfDTOList) Validate() error {
	return dara.Validate(s)
}
