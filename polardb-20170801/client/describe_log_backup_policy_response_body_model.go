// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedLogPolicies(v *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies) *DescribeLogBackupPolicyResponseBody
	GetAdvancedLogPolicies() *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies
	SetEnableBackupLog(v int32) *DescribeLogBackupPolicyResponseBody
	GetEnableBackupLog() *int32
	SetLogBackupAnotherRegionRegion(v string) *DescribeLogBackupPolicyResponseBody
	GetLogBackupAnotherRegionRegion() *string
	SetLogBackupAnotherRegionRetentionPeriod(v string) *DescribeLogBackupPolicyResponseBody
	GetLogBackupAnotherRegionRetentionPeriod() *string
	SetLogBackupRetentionPeriod(v int32) *DescribeLogBackupPolicyResponseBody
	GetLogBackupRetentionPeriod() *int32
	SetRequestId(v string) *DescribeLogBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeLogBackupPolicyResponseBody struct {
	AdvancedLogPolicies *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies `json:"AdvancedLogPolicies,omitempty" xml:"AdvancedLogPolicies,omitempty" type:"Struct"`
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// 	- 0: The log backup feature is disabled.
	//
	// 	- 1: The log backup feature is enabled. By default, the log backup feature is enabled and cannot be disabled.
	//
	// example:
	//
	// 1
	EnableBackupLog *int32 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The region in which you want to store cross-region log backups. For more information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// cn-beijing
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
	// 0
	LogBackupAnotherRegionRetentionPeriod *string `json:"LogBackupAnotherRegionRetentionPeriod,omitempty" xml:"LogBackupAnotherRegionRetentionPeriod,omitempty"`
	// The retention period of the log backups. Valid values:
	//
	// 	- 3 to 7300: The log backups are retained for 3 to 7,300 days.
	//
	// 	- \\-1: The log backups are permanently retained.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *int32 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 62EE0051-102B-488D-9C79-D607B8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyResponseBody) GetAdvancedLogPolicies() *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies {
	return s.AdvancedLogPolicies
}

func (s *DescribeLogBackupPolicyResponseBody) GetEnableBackupLog() *int32 {
	return s.EnableBackupLog
}

func (s *DescribeLogBackupPolicyResponseBody) GetLogBackupAnotherRegionRegion() *string {
	return s.LogBackupAnotherRegionRegion
}

func (s *DescribeLogBackupPolicyResponseBody) GetLogBackupAnotherRegionRetentionPeriod() *string {
	return s.LogBackupAnotherRegionRetentionPeriod
}

func (s *DescribeLogBackupPolicyResponseBody) GetLogBackupRetentionPeriod() *int32 {
	return s.LogBackupRetentionPeriod
}

func (s *DescribeLogBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogBackupPolicyResponseBody) SetAdvancedLogPolicies(v *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies) *DescribeLogBackupPolicyResponseBody {
	s.AdvancedLogPolicies = v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetEnableBackupLog(v int32) *DescribeLogBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetLogBackupAnotherRegionRegion(v string) *DescribeLogBackupPolicyResponseBody {
	s.LogBackupAnotherRegionRegion = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetLogBackupAnotherRegionRetentionPeriod(v string) *DescribeLogBackupPolicyResponseBody {
	s.LogBackupAnotherRegionRetentionPeriod = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeLogBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetRequestId(v string) *DescribeLogBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) Validate() error {
	if s.AdvancedLogPolicies != nil {
		if err := s.AdvancedLogPolicies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies struct {
	AdvancedLogPolicy []*DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy `json:"AdvancedLogPolicy,omitempty" xml:"AdvancedLogPolicy,omitempty" type:"Repeated"`
}

func (s DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies) GetAdvancedLogPolicy() []*DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	return s.AdvancedLogPolicy
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies) SetAdvancedLogPolicy(v []*DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies {
	s.AdvancedLogPolicy = v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPolicies) Validate() error {
	if s.AdvancedLogPolicy != nil {
		for _, item := range s.AdvancedLogPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy struct {
	DestRegion        *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType          *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	EnableLogBackup   *int32  `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	LogRetentionType  *string `json:"LogRetentionType,omitempty" xml:"LogRetentionType,omitempty"`
	LogRetentionValue *string `json:"LogRetentionValue,omitempty" xml:"LogRetentionValue,omitempty"`
	PolicyId          *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SrcRegion         *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType           *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetDestType() *string {
	return s.DestType
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetEnableLogBackup() *int32 {
	return s.EnableLogBackup
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetLogRetentionType() *string {
	return s.LogRetentionType
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetLogRetentionValue() *string {
	return s.LogRetentionValue
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetDestRegion(v string) *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.DestRegion = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetDestType(v string) *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.DestType = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetEnableLogBackup(v int32) *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.EnableLogBackup = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetLogRetentionType(v string) *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.LogRetentionType = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetLogRetentionValue(v string) *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.LogRetentionValue = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetPolicyId(v string) *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.PolicyId = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetSrcRegion(v string) *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.SrcRegion = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) SetSrcType(v string) *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy {
	s.SrcType = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBodyAdvancedLogPoliciesAdvancedLogPolicy) Validate() error {
	return dara.Validate(s)
}
