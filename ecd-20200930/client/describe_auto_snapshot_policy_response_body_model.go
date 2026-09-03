// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicies(v []*DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPolicyResponseBody
	GetAutoSnapshotPolicies() []*DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies
	SetNextToken(v string) *DescribeAutoSnapshotPolicyResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAutoSnapshotPolicyResponseBody
	GetRequestId() *string
}

type DescribeAutoSnapshotPolicyResponseBody struct {
	// The list of automatic snapshot policies.
	AutoSnapshotPolicies []*DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies `json:"AutoSnapshotPolicies,omitempty" xml:"AutoSnapshotPolicies,omitempty" type:"Repeated"`
	// The pagination token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A7F6612E-59CC-59F9-9DD1-91867FCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyResponseBody) GetAutoSnapshotPolicies() []*DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	return s.AutoSnapshotPolicies
}

func (s *DescribeAutoSnapshotPolicyResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAutoSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoSnapshotPolicyResponseBody) SetAutoSnapshotPolicies(v []*DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPolicyResponseBody {
	s.AutoSnapshotPolicies = v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBody) SetNextToken(v string) *DescribeAutoSnapshotPolicyResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBody) SetRequestId(v string) *DescribeAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBody) Validate() error {
	if s.AutoSnapshotPolicies != nil {
		for _, item := range s.AutoSnapshotPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies struct {
	// The creation time. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-mm-ddthh:mm:ssz` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-11T09:14:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The cron expression that specifies the snapshot creation time.
	//
	// example:
	//
	// 0 0 5,7 ? 	- 2/2
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The number of cloud computers to which the snapshot policy is attached.
	//
	// example:
	//
	// 1
	DesktopNum *int32 `json:"DesktopNum,omitempty" xml:"DesktopNum,omitempty"`
	// The cloud disk type.
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-3e3bmfcdkjfl1****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The Policy Name of the automatic snapshot policy.
	//
	// example:
	//
	// snapshot01
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID of the automatic snapshot policy.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The retention period of automatic snapshots, in days. Valid values: 1 to 180.
	//
	// example:
	//
	// 2
	RetentionDays *string `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The status of the automatic snapshot policy.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The points in time at which automatic snapshots are created.
	//
	// The parameter value is a JSON array in the format of `["0", "1", ... "23"]`, with a maximum of 24 time points separated by commas (,).
	//
	// example:
	//
	// ["17","18"]
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetDesktopNum() *int32 {
	return s.DesktopNum
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetRetentionDays() *string {
	return s.RetentionDays
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) GetTimePoints() *string {
	return s.TimePoints
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetCreationTime(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.CreationTime = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetCronExpression(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.CronExpression = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetDesktopNum(v int32) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.DesktopNum = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetDiskType(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.DiskType = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetPolicyId(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetPolicyName(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.PolicyName = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetRegionId(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetRetentionDays(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.RetentionDays = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetStatus(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.Status = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) SetTimePoints(v string) *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies {
	s.TimePoints = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyResponseBodyAutoSnapshotPolicies) Validate() error {
	return dara.Validate(s)
}
