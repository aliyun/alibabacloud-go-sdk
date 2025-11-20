// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOtsTableSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountRoleName(v string) *DescribeOtsTableSnapshotsRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *DescribeOtsTableSnapshotsRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *DescribeOtsTableSnapshotsRequest
	GetCrossAccountUserId() *int64
	SetEndTime(v int64) *DescribeOtsTableSnapshotsRequest
	GetEndTime() *int64
	SetLimit(v int32) *DescribeOtsTableSnapshotsRequest
	GetLimit() *int32
	SetNextToken(v string) *DescribeOtsTableSnapshotsRequest
	GetNextToken() *string
	SetOtsInstances(v []*DescribeOtsTableSnapshotsRequestOtsInstances) *DescribeOtsTableSnapshotsRequest
	GetOtsInstances() []*DescribeOtsTableSnapshotsRequestOtsInstances
	SetSnapshotIds(v []*string) *DescribeOtsTableSnapshotsRequest
	GetSnapshotIds() []*string
	SetStartTime(v int64) *DescribeOtsTableSnapshotsRequest
	GetStartTime() *int64
}

type DescribeOtsTableSnapshotsRequest struct {
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether data is backed up within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- SELF_ACCOUNT: Data is backed up within the same Alibaba Cloud account.
	//
	// 	- CROSS_ACCOUNT: Data is backed up across Alibaba Cloud accounts.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The UID of the source account used for cross-account backup.
	//
	// example:
	//
	// 144015xxxxx98732
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The end time of the backup. The value must be a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652068250881
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of rows that you want the current query to return.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that is required to obtain the next page of backup snapshots.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The Tablestore instances that are backed up.
	OtsInstances []*DescribeOtsTableSnapshotsRequestOtsInstances `json:"OtsInstances,omitempty" xml:"OtsInstances,omitempty" type:"Repeated"`
	// The snapshot IDs.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty" type:"Repeated"`
	// The start time of the backup. The value must be a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1611109271630
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOtsTableSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOtsTableSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeOtsTableSnapshotsRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribeOtsTableSnapshotsRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeOtsTableSnapshotsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeOtsTableSnapshotsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeOtsTableSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOtsTableSnapshotsRequest) GetOtsInstances() []*DescribeOtsTableSnapshotsRequestOtsInstances {
	return s.OtsInstances
}

func (s *DescribeOtsTableSnapshotsRequest) GetSnapshotIds() []*string {
	return s.SnapshotIds
}

func (s *DescribeOtsTableSnapshotsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeOtsTableSnapshotsRequest) SetCrossAccountRoleName(v string) *DescribeOtsTableSnapshotsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetCrossAccountType(v string) *DescribeOtsTableSnapshotsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetCrossAccountUserId(v int64) *DescribeOtsTableSnapshotsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetEndTime(v int64) *DescribeOtsTableSnapshotsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetLimit(v int32) *DescribeOtsTableSnapshotsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetNextToken(v string) *DescribeOtsTableSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetOtsInstances(v []*DescribeOtsTableSnapshotsRequestOtsInstances) *DescribeOtsTableSnapshotsRequest {
	s.OtsInstances = v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetSnapshotIds(v []*string) *DescribeOtsTableSnapshotsRequest {
	s.SnapshotIds = v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetStartTime(v int64) *DescribeOtsTableSnapshotsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) Validate() error {
	if s.OtsInstances != nil {
		for _, item := range s.OtsInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOtsTableSnapshotsRequestOtsInstances struct {
	// The name of the Tablestore instance.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The names of the tables in the Tablestore instance.
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s DescribeOtsTableSnapshotsRequestOtsInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeOtsTableSnapshotsRequestOtsInstances) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsRequestOtsInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeOtsTableSnapshotsRequestOtsInstances) GetTableNames() []*string {
	return s.TableNames
}

func (s *DescribeOtsTableSnapshotsRequestOtsInstances) SetInstanceName(v string) *DescribeOtsTableSnapshotsRequestOtsInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequestOtsInstances) SetTableNames(v []*string) *DescribeOtsTableSnapshotsRequestOtsInstances {
	s.TableNames = v
	return s
}

func (s *DescribeOtsTableSnapshotsRequestOtsInstances) Validate() error {
	return dara.Validate(s)
}
