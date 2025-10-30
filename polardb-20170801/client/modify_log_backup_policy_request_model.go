// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedLogPolicies(v []*ModifyLogBackupPolicyRequestAdvancedLogPolicies) *ModifyLogBackupPolicyRequest
	GetAdvancedLogPolicies() []*ModifyLogBackupPolicyRequestAdvancedLogPolicies
	SetDBClusterId(v string) *ModifyLogBackupPolicyRequest
	GetDBClusterId() *string
	SetLogBackupAnotherRegionRegion(v string) *ModifyLogBackupPolicyRequest
	GetLogBackupAnotherRegionRegion() *string
	SetLogBackupAnotherRegionRetentionPeriod(v string) *ModifyLogBackupPolicyRequest
	GetLogBackupAnotherRegionRetentionPeriod() *string
	SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyRequest
	GetLogBackupRetentionPeriod() *string
	SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLogBackupPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLogBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type ModifyLogBackupPolicyRequest struct {
	AdvancedLogPolicies []*ModifyLogBackupPolicyRequestAdvancedLogPolicies `json:"AdvancedLogPolicies,omitempty" xml:"AdvancedLogPolicies,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the information of all clusters that are deployed in a specific region, such as the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region in which you want to store cross-region log backups. For information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// cn-hangzhou
	LogBackupAnotherRegionRegion *string `json:"LogBackupAnotherRegionRegion,omitempty" xml:"LogBackupAnotherRegionRegion,omitempty"`
	// The retention period of cross-region log backups. Valid values:
	//
	// 	- **0**: The cross-region backup feature is disabled.
	//
	// 	- **30 to 7300**: Cross-region log backups are retained for 30 to 7,300 days.
	//
	// 	- **-1**: The log backups are permanently retained.
	//
	// >  When you create a cluster, the default value of this parameter is **0**.
	//
	// example:
	//
	// 30
	LogBackupAnotherRegionRetentionPeriod *string `json:"LogBackupAnotherRegionRetentionPeriod,omitempty" xml:"LogBackupAnotherRegionRetentionPeriod,omitempty"`
	// The retention period of the log backups. Valid values:
	//
	// 	- 3 to 7300: The log backups are retained for 3 to 7,300 days.
	//
	// 	- \\-1: The log backups are permanently retained.
	//
	// example:
	//
	// 3
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLogBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyRequest) GetAdvancedLogPolicies() []*ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	return s.AdvancedLogPolicies
}

func (s *ModifyLogBackupPolicyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyLogBackupPolicyRequest) GetLogBackupAnotherRegionRegion() *string {
	return s.LogBackupAnotherRegionRegion
}

func (s *ModifyLogBackupPolicyRequest) GetLogBackupAnotherRegionRetentionPeriod() *string {
	return s.LogBackupAnotherRegionRetentionPeriod
}

func (s *ModifyLogBackupPolicyRequest) GetLogBackupRetentionPeriod() *string {
	return s.LogBackupRetentionPeriod
}

func (s *ModifyLogBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLogBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLogBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLogBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLogBackupPolicyRequest) SetAdvancedLogPolicies(v []*ModifyLogBackupPolicyRequestAdvancedLogPolicies) *ModifyLogBackupPolicyRequest {
	s.AdvancedLogPolicies = v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetDBClusterId(v string) *ModifyLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetLogBackupAnotherRegionRegion(v string) *ModifyLogBackupPolicyRequest {
	s.LogBackupAnotherRegionRegion = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetLogBackupAnotherRegionRetentionPeriod(v string) *ModifyLogBackupPolicyRequest {
	s.LogBackupAnotherRegionRetentionPeriod = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) Validate() error {
	if s.AdvancedLogPolicies != nil {
		for _, item := range s.AdvancedLogPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyLogBackupPolicyRequestAdvancedLogPolicies struct {
	ActionType        *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	DestRegion        *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType          *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	EnableLogBackup   *int32  `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	LogRetentionType  *string `json:"LogRetentionType,omitempty" xml:"LogRetentionType,omitempty"`
	LogRetentionValue *string `json:"LogRetentionValue,omitempty" xml:"LogRetentionValue,omitempty"`
	PolicyId          *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SrcRegion         *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType           *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s ModifyLogBackupPolicyRequestAdvancedLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogBackupPolicyRequestAdvancedLogPolicies) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetEnableLogBackup() *int32 {
	return s.EnableLogBackup
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetLogRetentionType() *string {
	return s.LogRetentionType
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetLogRetentionValue() *string {
	return s.LogRetentionValue
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetActionType(v string) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetDestRegion(v string) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetDestType(v string) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetEnableLogBackup(v int32) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetLogRetentionType(v string) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.LogRetentionType = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetLogRetentionValue(v string) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.LogRetentionValue = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetPolicyId(v string) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetSrcRegion(v string) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) SetSrcType(v string) *ModifyLogBackupPolicyRequestAdvancedLogPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyLogBackupPolicyRequestAdvancedLogPolicies) Validate() error {
	return dara.Validate(s)
}
