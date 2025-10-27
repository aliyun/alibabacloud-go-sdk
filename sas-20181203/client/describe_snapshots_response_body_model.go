// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeSnapshotsResponseBodyPageInfo) *DescribeSnapshotsResponseBody
	GetPageInfo() *DescribeSnapshotsResponseBodyPageInfo
	SetRequestId(v string) *DescribeSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody
	GetSnapshots() []*DescribeSnapshotsResponseBodySnapshots
}

type DescribeSnapshotsResponseBody struct {
	// The pagination information.
	PageInfo *DescribeSnapshotsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB393****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the backup snapshots.
	Snapshots []*DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) GetPageInfo() *DescribeSnapshotsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotsResponseBody) GetSnapshots() []*DescribeSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *DescribeSnapshotsResponseBody) SetPageInfo(v *DescribeSnapshotsResponseBodyPageInfo) *DescribeSnapshotsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeSnapshotsResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnapshotsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The token that is used to initiate the next call.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAGYXFWIAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM1Mzc3Njc4MzA2ODY5NmI2YTY1Nzg2NTcxNjE2NDc4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnapshotsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSnapshotsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSnapshotsResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnapshotsResponseBodyPageInfo) SetCount(v int32) *DescribeSnapshotsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeSnapshotsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeSnapshotsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSnapshotsResponseBodyPageInfo) SetNextToken(v string) *DescribeSnapshotsResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponseBodyPageInfo) SetPageSize(v int32) *DescribeSnapshotsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodyPageInfo) SetTotalCount(v int32) *DescribeSnapshotsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnapshotsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeSnapshotsResponseBodySnapshots struct {
	// The actual data amount of backup snapshots after duplicates are removed. Unit: bytes.
	//
	// example:
	//
	// 686188****
	ActualBytes *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The actual number of backup objects.
	//
	// >  This parameter is available only for file backup.
	//
	// example:
	//
	// 123
	ActualItems *int64 `json:"ActualItems,omitempty" xml:"ActualItems,omitempty"`
	// The actual amount of data that is generated by incremental backup. Unit: bytes.
	//
	// example:
	//
	// 800
	BytesDone *int64 `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	// The total data amount of the data source. Unit: bytes.
	//
	// example:
	//
	// 3484541815****
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// This parameter is returned only if the value of the **SourceType*	- parameter is **ECS_FILE**. This parameter indicates the ID of the Hybrid Backup Recovery (HBR) agent.
	//
	// example:
	//
	// c-000dbefaw9f7gnbw****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The version of the anti-ransomware agent.
	//
	// example:
	//
	// 2.10.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The time when the backup snapshot was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1646793988
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The file that records the information about backup failures, including the information about partially completed backup tasks.
	//
	// example:
	//
	// s-0008ndhgrflh55i5****.csv
	ErrorFile *string `json:"ErrorFile,omitempty" xml:"ErrorFile,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-2ze78zfakirgh1yl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of backup objects.
	//
	// >  This parameter is available only for file backup.
	//
	// example:
	//
	// 6
	ItemsDone *int64 `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	// The total number of objects in the data source.
	//
	// >  This parameter is available only for file backup.
	//
	// example:
	//
	// 7
	ItemsTotal *int64 `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	// The ID of the backup task.
	//
	// example:
	//
	// job-000a2q5vg6awgo01****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The hash value of the parent backup snapshot.
	//
	// example:
	//
	// a0181606689c9562f092b3190bddb8a62bb5a24784424ba2102bc7fe92ae****
	ParentSnapshotHash *string `json:"ParentSnapshotHash,omitempty" xml:"ParentSnapshotHash,omitempty"`
	// This parameter is returned only if the value of the **SourceType*	- parameter is **ECS_FILE**. This parameter indicates the path to the backup files.
	//
	// example:
	//
	// ["/home"]
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is returned only if the value of the **SourceType*	- parameter is **NAS**. This parameter indicates the path to the backup files.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The ID of the backup plan.
	//
	// example:
	//
	// plan-000ee8gh2ljelsnb****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the region in which backup snapshot is stored.
	//
	// example:
	//
	// us-east-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The retention period of the backup snapshot.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The hash value of the backup snapshot.
	//
	// example:
	//
	// 9ee47cf12351e4ddecce8c12f4957d3946cd96fbe24cd4ab264c7200839d****
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The ID of the backup snapshot.
	//
	// example:
	//
	// s-0003ahfuqpjdztsg****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: ECS files
	//
	// 	- **OSS**: Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: File Storage NAS file systems
	//
	// 	- **OTS_TABLE**: Tablestore instances
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of backup tasks. Valid values:
	//
	// 	- **COMPLETE**: complete
	//
	// 	- **PARTIAL_COMPLETE**: partial complete
	//
	// 	- **FAILED**: failed
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// d332af48-1269-4a55-a6db-8543a80f****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the backup vault that stores the backup snapshot.
	//
	// example:
	//
	// v-000ccok3zmw7fbzz****
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetActualBytes() *int64 {
	return s.ActualBytes
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetActualItems() *int64 {
	return s.ActualItems
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetErrorFile() *string {
	return s.ErrorFile
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetItemsDone() *int64 {
	return s.ItemsDone
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetItemsTotal() *int64 {
	return s.ItemsTotal
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetJobId() *string {
	return s.JobId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetParentSnapshotHash() *string {
	return s.ParentSnapshotHash
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetPath() *string {
	return s.Path
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetPaths() []*string {
	return s.Paths
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetActualBytes(v int64) *DescribeSnapshotsResponseBodySnapshots {
	s.ActualBytes = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetActualItems(v int64) *DescribeSnapshotsResponseBodySnapshots {
	s.ActualItems = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetBytesDone(v int64) *DescribeSnapshotsResponseBodySnapshots {
	s.BytesDone = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetBytesTotal(v int64) *DescribeSnapshotsResponseBodySnapshots {
	s.BytesTotal = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetClientId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.ClientId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetClientVersion(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.ClientVersion = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetCreatedTime(v int64) *DescribeSnapshotsResponseBodySnapshots {
	s.CreatedTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetErrorFile(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.ErrorFile = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetInstanceId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetItemsDone(v int64) *DescribeSnapshotsResponseBodySnapshots {
	s.ItemsDone = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetItemsTotal(v int64) *DescribeSnapshotsResponseBodySnapshots {
	s.ItemsTotal = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetJobId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.JobId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetParentSnapshotHash(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.ParentSnapshotHash = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetPath(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Path = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetPaths(v []*string) *DescribeSnapshotsResponseBodySnapshots {
	s.Paths = v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetPlanId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.PlanId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRegionId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRetention(v int64) *DescribeSnapshotsResponseBodySnapshots {
	s.Retention = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotHash(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetUuid(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Uuid = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetVaultId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.VaultId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) Validate() error {
	return dara.Validate(s)
}
