// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInfoList(v []*DescribeLogStoreInfoResponseBodyInfoList) *DescribeLogStoreInfoResponseBody
	GetInfoList() []*DescribeLogStoreInfoResponseBodyInfoList
	SetLogModifyQuota(v int32) *DescribeLogStoreInfoResponseBody
	GetLogModifyQuota() *int32
	SetLogStoreName(v string) *DescribeLogStoreInfoResponseBody
	GetLogStoreName() *string
	SetLogVersion(v int32) *DescribeLogStoreInfoResponseBody
	GetLogVersion() *int32
	SetProjectName(v string) *DescribeLogStoreInfoResponseBody
	GetProjectName() *string
	SetQuota(v int64) *DescribeLogStoreInfoResponseBody
	GetQuota() *int64
	SetRegionId(v string) *DescribeLogStoreInfoResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeLogStoreInfoResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DescribeLogStoreInfoResponseBody
	GetTaskId() *string
	SetTotalQuota(v int64) *DescribeLogStoreInfoResponseBody
	GetTotalQuota() *int64
	SetTtl(v int32) *DescribeLogStoreInfoResponseBody
	GetTtl() *int32
	SetUsed(v int64) *DescribeLogStoreInfoResponseBody
	GetUsed() *int64
}

type DescribeLogStoreInfoResponseBody struct {
	// The information list.
	InfoList []*DescribeLogStoreInfoResponseBodyInfoList `json:"InfoList,omitempty" xml:"InfoList,omitempty" type:"Repeated"`
	// The number of times the log storage mode can be changed.
	//
	// example:
	//
	// 2
	LogModifyQuota *int32 `json:"LogModifyQuota,omitempty" xml:"LogModifyQuota,omitempty"`
	// The name of the SLS Logstore.
	//
	// example:
	//
	// xxx-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The log version. 1: indicates one Logstore. 2: indicates two Logstores.
	//
	// example:
	//
	// 2
	LogVersion *int32 `json:"LogVersion,omitempty" xml:"LogVersion,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// project-xxx-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The available log storage capacity, in bytes.
	//
	// example:
	//
	// 50000000
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The ID of the region where logs are delivered.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6C3B72B********E95FB0A161
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 132
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total purchased log storage capacity, in bytes.
	//
	// example:
	//
	// 50000000
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The storage duration of logs, in days.
	//
	// example:
	//
	// 20
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The used storage capacity, in bytes.
	//
	// > Statistics from Simple Log Service may be delayed by up to two hours.
	//
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s DescribeLogStoreInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreInfoResponseBody) GetInfoList() []*DescribeLogStoreInfoResponseBodyInfoList {
	return s.InfoList
}

func (s *DescribeLogStoreInfoResponseBody) GetLogModifyQuota() *int32 {
	return s.LogModifyQuota
}

func (s *DescribeLogStoreInfoResponseBody) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeLogStoreInfoResponseBody) GetLogVersion() *int32 {
	return s.LogVersion
}

func (s *DescribeLogStoreInfoResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLogStoreInfoResponseBody) GetQuota() *int64 {
	return s.Quota
}

func (s *DescribeLogStoreInfoResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogStoreInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogStoreInfoResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeLogStoreInfoResponseBody) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *DescribeLogStoreInfoResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeLogStoreInfoResponseBody) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeLogStoreInfoResponseBody) SetInfoList(v []*DescribeLogStoreInfoResponseBodyInfoList) *DescribeLogStoreInfoResponseBody {
	s.InfoList = v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetLogModifyQuota(v int32) *DescribeLogStoreInfoResponseBody {
	s.LogModifyQuota = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetLogStoreName(v string) *DescribeLogStoreInfoResponseBody {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetLogVersion(v int32) *DescribeLogStoreInfoResponseBody {
	s.LogVersion = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetProjectName(v string) *DescribeLogStoreInfoResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetQuota(v int64) *DescribeLogStoreInfoResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetRegionId(v string) *DescribeLogStoreInfoResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetRequestId(v string) *DescribeLogStoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetTaskId(v string) *DescribeLogStoreInfoResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetTotalQuota(v int64) *DescribeLogStoreInfoResponseBody {
	s.TotalQuota = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetTtl(v int32) *DescribeLogStoreInfoResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetUsed(v int64) *DescribeLogStoreInfoResponseBody {
	s.Used = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) Validate() error {
	if s.InfoList != nil {
		for _, item := range s.InfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLogStoreInfoResponseBodyInfoList struct {
	// The name of the SLS Logstore.
	//
	// example:
	//
	// xxx-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The maximum number of shards supported for scaling.
	//
	// example:
	//
	// 4
	MaxSplitShard *int32 `json:"MaxSplitShard,omitempty" xml:"MaxSplitShard,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// cloudfirewall-project-14151892848****-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The storage capacity threshold, in bytes.
	//
	// example:
	//
	// 50000000
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of shards in use.
	//
	// example:
	//
	// 2
	Shard *int32 `json:"Shard,omitempty" xml:"Shard,omitempty"`
	// The location of the Logstore. Valid values: \\`cn\\` for the Chinese mainland and \\`intl\\` for regions outside the Chinese mainland.
	//
	// example:
	//
	// cn
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The storage duration of logs, in days.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The amount of stored logs, in bytes.
	//
	// example:
	//
	// 21852955752
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s DescribeLogStoreInfoResponseBodyInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreInfoResponseBodyInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetMaxSplitShard() *int32 {
	return s.MaxSplitShard
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetQuota() *int64 {
	return s.Quota
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetShard() *int32 {
	return s.Shard
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetSite() *string {
	return s.Site
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetLogStoreName(v string) *DescribeLogStoreInfoResponseBodyInfoList {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetMaxSplitShard(v int32) *DescribeLogStoreInfoResponseBodyInfoList {
	s.MaxSplitShard = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetProjectName(v string) *DescribeLogStoreInfoResponseBodyInfoList {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetQuota(v int64) *DescribeLogStoreInfoResponseBodyInfoList {
	s.Quota = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetRegionId(v string) *DescribeLogStoreInfoResponseBodyInfoList {
	s.RegionId = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetShard(v int32) *DescribeLogStoreInfoResponseBodyInfoList {
	s.Shard = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetSite(v string) *DescribeLogStoreInfoResponseBodyInfoList {
	s.Site = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetTtl(v int32) *DescribeLogStoreInfoResponseBodyInfoList {
	s.Ttl = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) SetUsed(v int64) *DescribeLogStoreInfoResponseBodyInfoList {
	s.Used = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBodyInfoList) Validate() error {
	return dara.Validate(s)
}
