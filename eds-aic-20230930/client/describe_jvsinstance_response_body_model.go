// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJVSInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeJVSInstanceResponseBodyData) *DescribeJVSInstanceResponseBody
	GetData() []*DescribeJVSInstanceResponseBodyData
	SetMaxResults(v int32) *DescribeJVSInstanceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeJVSInstanceResponseBody
	GetNextToken() *string
	SetPendingUpgradeCount(v int32) *DescribeJVSInstanceResponseBody
	GetPendingUpgradeCount() *int32
	SetRequestId(v string) *DescribeJVSInstanceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeJVSInstanceResponseBody
	GetTotalCount() *int32
}

type DescribeJVSInstanceResponseBody struct {
	// The returned result object.
	Data []*DescribeJVSInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that indicates the current position from which to start reading. An empty value indicates reading from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PendingUpgradeCount *int32  `json:"PendingUpgradeCount,omitempty" xml:"PendingUpgradeCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeJVSInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBody) GetData() []*DescribeJVSInstanceResponseBodyData {
	return s.Data
}

func (s *DescribeJVSInstanceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeJVSInstanceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeJVSInstanceResponseBody) GetPendingUpgradeCount() *int32 {
	return s.PendingUpgradeCount
}

func (s *DescribeJVSInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJVSInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeJVSInstanceResponseBody) SetData(v []*DescribeJVSInstanceResponseBodyData) *DescribeJVSInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetMaxResults(v int32) *DescribeJVSInstanceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetNextToken(v string) *DescribeJVSInstanceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetPendingUpgradeCount(v int32) *DescribeJVSInstanceResponseBody {
	s.PendingUpgradeCount = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetRequestId(v string) *DescribeJVSInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) SetTotalCount(v int32) *DescribeJVSInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeJVSInstanceResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJVSInstanceResponseBodyData struct {
	AgentVersion *DescribeJVSInstanceResponseBodyDataAgentVersion `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2026-04-10T01:31:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The credit quota configuration. Subsequent quota configurations overwrite previous configurations.
	CreditConfig []*DescribeJVSInstanceResponseBodyDataCreditConfig `json:"CreditConfig,omitempty" xml:"CreditConfig,omitempty" type:"Repeated"`
	// The expiration time.
	//
	// example:
	//
	// 2026-04-10T01:31:32Z
	ExpireTime      *string                                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstalledSkills []*DescribeJVSInstanceResponseBodyDataInstalledSkills `json:"InstalledSkills,omitempty" xml:"InstalledSkills,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is not supported.
	//
	// example:
	//
	// null
	JvsPackageId *string `json:"JvsPackageId,omitempty" xml:"JvsPackageId,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-04-10T01:31:32Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The instance status.
	//
	// example:
	//
	// RUNNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The used credits.
	UsedCredit []*DescribeJVSInstanceResponseBodyDataUsedCredit `json:"UsedCredit,omitempty" xml:"UsedCredit,omitempty" type:"Repeated"`
}

func (s DescribeJVSInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBodyData) GetAgentVersion() *DescribeJVSInstanceResponseBodyDataAgentVersion {
	return s.AgentVersion
}

func (s *DescribeJVSInstanceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeJVSInstanceResponseBodyData) GetCreditConfig() []*DescribeJVSInstanceResponseBodyDataCreditConfig {
	return s.CreditConfig
}

func (s *DescribeJVSInstanceResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeJVSInstanceResponseBodyData) GetInstalledSkills() []*DescribeJVSInstanceResponseBodyDataInstalledSkills {
	return s.InstalledSkills
}

func (s *DescribeJVSInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeJVSInstanceResponseBodyData) GetJvsPackageId() *string {
	return s.JvsPackageId
}

func (s *DescribeJVSInstanceResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeJVSInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeJVSInstanceResponseBodyData) GetUsedCredit() []*DescribeJVSInstanceResponseBodyDataUsedCredit {
	return s.UsedCredit
}

func (s *DescribeJVSInstanceResponseBodyData) SetAgentVersion(v *DescribeJVSInstanceResponseBodyDataAgentVersion) *DescribeJVSInstanceResponseBodyData {
	s.AgentVersion = v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetCreateTime(v string) *DescribeJVSInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetCreditConfig(v []*DescribeJVSInstanceResponseBodyDataCreditConfig) *DescribeJVSInstanceResponseBodyData {
	s.CreditConfig = v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetExpireTime(v string) *DescribeJVSInstanceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetInstalledSkills(v []*DescribeJVSInstanceResponseBodyDataInstalledSkills) *DescribeJVSInstanceResponseBodyData {
	s.InstalledSkills = v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetInstanceId(v string) *DescribeJVSInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetJvsPackageId(v string) *DescribeJVSInstanceResponseBodyData {
	s.JvsPackageId = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetModifyTime(v string) *DescribeJVSInstanceResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetStatus(v string) *DescribeJVSInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) SetUsedCredit(v []*DescribeJVSInstanceResponseBodyDataUsedCredit) *DescribeJVSInstanceResponseBodyData {
	s.UsedCredit = v
	return s
}

func (s *DescribeJVSInstanceResponseBodyData) Validate() error {
	if s.AgentVersion != nil {
		if err := s.AgentVersion.Validate(); err != nil {
			return err
		}
	}
	if s.CreditConfig != nil {
		for _, item := range s.CreditConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InstalledSkills != nil {
		for _, item := range s.InstalledSkills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UsedCredit != nil {
		for _, item := range s.UsedCredit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJVSInstanceResponseBodyDataAgentVersion struct {
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeJVSInstanceResponseBodyDataAgentVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBodyDataAgentVersion) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBodyDataAgentVersion) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *DescribeJVSInstanceResponseBodyDataAgentVersion) GetVersion() *string {
	return s.Version
}

func (s *DescribeJVSInstanceResponseBodyDataAgentVersion) SetUpgradeStatus(v string) *DescribeJVSInstanceResponseBodyDataAgentVersion {
	s.UpgradeStatus = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataAgentVersion) SetVersion(v string) *DescribeJVSInstanceResponseBodyDataAgentVersion {
	s.Version = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataAgentVersion) Validate() error {
	return dara.Validate(s)
}

type DescribeJVSInstanceResponseBodyDataCreditConfig struct {
	// The quota limit. Valid values:
	//
	// - 0: not available for use.
	//
	// - >0: the quota is configured based on the numeric value.
	//
	// - -1: unlimited.
	//
	// example:
	//
	// -1
	CreditLimit *int64 `json:"CreditLimit,omitempty" xml:"CreditLimit,omitempty"`
	// The quota period. Valid values:
	//
	// - total: The total usage limit.
	//
	// - month: Monthly. The quota resets based on the resource activation time as one cycle.
	//
	// - day: Daily. The quota resets at 00:00.
	//
	// example:
	//
	// day
	LimitPeriod *string `json:"LimitPeriod,omitempty" xml:"LimitPeriod,omitempty"`
}

func (s DescribeJVSInstanceResponseBodyDataCreditConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBodyDataCreditConfig) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) GetCreditLimit() *int64 {
	return s.CreditLimit
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) GetLimitPeriod() *string {
	return s.LimitPeriod
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) SetCreditLimit(v int64) *DescribeJVSInstanceResponseBodyDataCreditConfig {
	s.CreditLimit = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) SetLimitPeriod(v string) *DescribeJVSInstanceResponseBodyDataCreditConfig {
	s.LimitPeriod = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataCreditConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeJVSInstanceResponseBodyDataInstalledSkills struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IconUrl     *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	InstalledAt *string `json:"InstalledAt,omitempty" xml:"InstalledAt,omitempty"`
	SkillId     *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	SkillName   *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	SkillType   *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
}

func (s DescribeJVSInstanceResponseBodyDataInstalledSkills) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBodyDataInstalledSkills) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) GetDescription() *string {
	return s.Description
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) GetIconUrl() *string {
	return s.IconUrl
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) GetInstalledAt() *string {
	return s.InstalledAt
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) GetSkillId() *string {
	return s.SkillId
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) GetSkillName() *string {
	return s.SkillName
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) GetSkillType() *string {
	return s.SkillType
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) SetDescription(v string) *DescribeJVSInstanceResponseBodyDataInstalledSkills {
	s.Description = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) SetIconUrl(v string) *DescribeJVSInstanceResponseBodyDataInstalledSkills {
	s.IconUrl = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) SetInstalledAt(v string) *DescribeJVSInstanceResponseBodyDataInstalledSkills {
	s.InstalledAt = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) SetSkillId(v string) *DescribeJVSInstanceResponseBodyDataInstalledSkills {
	s.SkillId = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) SetSkillName(v string) *DescribeJVSInstanceResponseBodyDataInstalledSkills {
	s.SkillName = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) SetSkillType(v string) *DescribeJVSInstanceResponseBodyDataInstalledSkills {
	s.SkillType = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataInstalledSkills) Validate() error {
	return dara.Validate(s)
}

type DescribeJVSInstanceResponseBodyDataUsedCredit struct {
	// The number of credits.
	//
	// example:
	//
	// 5
	Credit *int64 `json:"Credit,omitempty" xml:"Credit,omitempty"`
	// The dimension of the current credit.
	//
	// example:
	//
	// day
	LimitPeriod *string `json:"LimitPeriod,omitempty" xml:"LimitPeriod,omitempty"`
}

func (s DescribeJVSInstanceResponseBodyDataUsedCredit) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponseBodyDataUsedCredit) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) GetCredit() *int64 {
	return s.Credit
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) GetLimitPeriod() *string {
	return s.LimitPeriod
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) SetCredit(v int64) *DescribeJVSInstanceResponseBodyDataUsedCredit {
	s.Credit = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) SetLimitPeriod(v string) *DescribeJVSInstanceResponseBodyDataUsedCredit {
	s.LimitPeriod = &v
	return s
}

func (s *DescribeJVSInstanceResponseBodyDataUsedCredit) Validate() error {
	return dara.Validate(s)
}
