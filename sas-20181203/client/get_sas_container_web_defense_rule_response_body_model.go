// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContainerWebDefenseRule(v *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) *GetSasContainerWebDefenseRuleResponseBody
	GetContainerWebDefenseRule() *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule
	SetRequestId(v string) *GetSasContainerWebDefenseRuleResponseBody
	GetRequestId() *string
}

type GetSasContainerWebDefenseRuleResponseBody struct {
	// The details of the container file defense rule.
	ContainerWebDefenseRule *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule `json:"ContainerWebDefenseRule,omitempty" xml:"ContainerWebDefenseRule,omitempty" type:"Struct"`
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// BA674E4**62FB5133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSasContainerWebDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleResponseBody) GetContainerWebDefenseRule() *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	return s.ContainerWebDefenseRule
}

func (s *GetSasContainerWebDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSasContainerWebDefenseRuleResponseBody) SetContainerWebDefenseRule(v *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) *GetSasContainerWebDefenseRuleResponseBody {
	s.ContainerWebDefenseRule = v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBody) SetRequestId(v string) *GetSasContainerWebDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBody) Validate() error {
	if s.ContainerWebDefenseRule != nil {
		if err := s.ContainerWebDefenseRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule struct {
	// The user ID.
	//
	// example:
	//
	// 1000**0002
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The total number of applications.
	//
	// example:
	//
	// 1
	AppTotalCount *int32 `json:"AppTotalCount,omitempty" xml:"AppTotalCount,omitempty"`
	// The timestamp when the rule was created. Unit: milliseconds.
	//
	// example:
	//
	// 1709173360000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the rule was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1655432638000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of configured file paths.
	PathConfDTOList []*GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList `json:"PathConfDTOList,omitempty" xml:"PathConfDTOList,omitempty" type:"Repeated"`
	// The rule name.
	//
	// example:
	//
	// 防篡改规则
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule status. Valid values:
	//
	// - **1**: Enabled.
	//
	// - **0**: Disabled.
	//
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GetAppTotalCount() *int32 {
	return s.AppTotalCount
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GetId() *int64 {
	return s.Id
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GetPathConfDTOList() []*GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	return s.PathConfDTOList
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GetRuleName() *string {
	return s.RuleName
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) GetRuleStatus() *int32 {
	return s.RuleStatus
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) SetAliUid(v int64) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	s.AliUid = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) SetAppTotalCount(v int32) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	s.AppTotalCount = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) SetGmtCreate(v int64) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	s.GmtCreate = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) SetGmtModified(v int64) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	s.GmtModified = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) SetId(v int64) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	s.Id = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) SetPathConfDTOList(v []*GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	s.PathConfDTOList = v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) SetRuleName(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	s.RuleName = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) SetRuleStatus(v int32) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule {
	s.RuleStatus = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRule) Validate() error {
	if s.PathConfDTOList != nil {
		for _, item := range s.PathConfDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList struct {
	// The backup path.
	//
	// example:
	//
	// /usr/path/
	BackupPath *string `json:"BackupPath,omitempty" xml:"BackupPath,omitempty"`
	// The action to take. Valid values:
	//
	// - **block**: Block.
	//
	// - **audit**: Alert.
	//
	// example:
	//
	// audit
	DefenseMode *string `json:"DefenseMode,omitempty" xml:"DefenseMode,omitempty"`
	// The defense path.
	//
	// example:
	//
	// /test11*
	DefensePath *string `json:"DefensePath,omitempty" xml:"DefensePath,omitempty"`
	// The excluded file.
	//
	// example:
	//
	// file1
	ExcludeFile *string `json:"ExcludeFile,omitempty" xml:"ExcludeFile,omitempty"`
	// The excluded file path.
	//
	// example:
	//
	// /test2/sub1,/test2/sub2
	ExcludeFilePath *string `json:"ExcludeFilePath,omitempty" xml:"ExcludeFilePath,omitempty"`
	// The excluded file type.
	//
	// example:
	//
	// doc
	ExcludeFileType *string `json:"ExcludeFileType,omitempty" xml:"ExcludeFileType,omitempty"`
	// The defense mode. Valid values:
	//
	// - **0**: Basic pattern (whitelist).
	//
	// - **1**: Advanced pattern (blacklist).
	//
	// example:
	//
	// 0
	GuardType *int32 `json:"GuardType,omitempty" xml:"GuardType,omitempty"`
	// The included file.
	//
	// example:
	//
	// webapp
	IncludeFile *string `json:"IncludeFile,omitempty" xml:"IncludeFile,omitempty"`
	// The included file type.
	//
	// example:
	//
	// doc
	IncludeFileType *string `json:"IncludeFileType,omitempty" xml:"IncludeFileType,omitempty"`
	// The list of whitelisted processes.
	ProcessPathList []*string `json:"ProcessPathList,omitempty" xml:"ProcessPathList,omitempty" type:"Repeated"`
}

func (s GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetBackupPath() *string {
	return s.BackupPath
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetDefenseMode() *string {
	return s.DefenseMode
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetDefensePath() *string {
	return s.DefensePath
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetExcludeFile() *string {
	return s.ExcludeFile
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetExcludeFilePath() *string {
	return s.ExcludeFilePath
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetExcludeFileType() *string {
	return s.ExcludeFileType
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetGuardType() *int32 {
	return s.GuardType
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetIncludeFile() *string {
	return s.IncludeFile
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetIncludeFileType() *string {
	return s.IncludeFileType
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) GetProcessPathList() []*string {
	return s.ProcessPathList
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetBackupPath(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.BackupPath = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetDefenseMode(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.DefenseMode = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetDefensePath(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.DefensePath = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetExcludeFile(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.ExcludeFile = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetExcludeFilePath(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.ExcludeFilePath = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetExcludeFileType(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.ExcludeFileType = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetGuardType(v int32) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.GuardType = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetIncludeFile(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.IncludeFile = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetIncludeFileType(v string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.IncludeFileType = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) SetProcessPathList(v []*string) *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList {
	s.ProcessPathList = v
	return s
}

func (s *GetSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRulePathConfDTOList) Validate() error {
	return dara.Validate(s)
}
