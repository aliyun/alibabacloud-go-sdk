// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiVersion(v string) *DescribeSnapshotsRequest
	GetApiVersion() *string
	SetCurrentPage(v int32) *DescribeSnapshotsRequest
	GetCurrentPage() *int32
	SetIsAliYunEcs(v string) *DescribeSnapshotsRequest
	GetIsAliYunEcs() *string
	SetMachineRegion(v string) *DescribeSnapshotsRequest
	GetMachineRegion() *string
	SetMachineRemark(v string) *DescribeSnapshotsRequest
	GetMachineRemark() *string
	SetNextToken(v string) *DescribeSnapshotsRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribeSnapshotsRequest
	GetPageSize() *int32
	SetStatusList(v string) *DescribeSnapshotsRequest
	GetStatusList() *string
	SetUuid(v string) *DescribeSnapshotsRequest
	GetUuid() *string
}

type DescribeSnapshotsRequest struct {
	// The version of the anti-ransomware policy. Valid values:
	//
	// 	- **1.0.0**
	//
	// 	- **2.0.0**
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0.0
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the server is an Elastic Compute Service (ECS) instance. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	IsAliYunEcs *string `json:"IsAliYunEcs,omitempty" xml:"IsAliYunEcs,omitempty"`
	// The region in which the server resides.
	//
	// >  If the Uuid parameter is not specified, this parameter is required.
	//
	// example:
	//
	// us-east-1
	MachineRegion *string `json:"MachineRegion,omitempty" xml:"MachineRegion,omitempty"`
	// The name or IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	MachineRemark *string `json:"MachineRemark,omitempty" xml:"MachineRemark,omitempty"`
	// The starting position of the query. If this parameter is left empty, the query starts from the beginning.
	//
	// >  If you call the operation for the first time, you do not need to specify the parameter. The response to the first call contains the token that can be used for the second call. Each subsequent response contains the token that can be used for the next call.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAB4SwmEAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM0NzY3YTZjNjI3NjZmNmU3MjcxNjk3NDY5MzY3MjY4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of backup snapshots from which data can be restored. Valid values:
	//
	// 	- **COMPLETE**: complete
	//
	// 	- **PARTIAL_COMPLETE**: partial complete
	//
	// example:
	//
	// ["COMPLETE"]
	StatusList *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// The UUID of the server.
	//
	// >  You can call the [DescribeBackupPolicy](~~DescribeBackupPolicy~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 061d8042-59ff-416e-bc33-294a1cf5****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *DescribeSnapshotsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSnapshotsRequest) GetIsAliYunEcs() *string {
	return s.IsAliYunEcs
}

func (s *DescribeSnapshotsRequest) GetMachineRegion() *string {
	return s.MachineRegion
}

func (s *DescribeSnapshotsRequest) GetMachineRemark() *string {
	return s.MachineRemark
}

func (s *DescribeSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotsRequest) GetStatusList() *string {
	return s.StatusList
}

func (s *DescribeSnapshotsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSnapshotsRequest) SetApiVersion(v string) *DescribeSnapshotsRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetCurrentPage(v int32) *DescribeSnapshotsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetIsAliYunEcs(v string) *DescribeSnapshotsRequest {
	s.IsAliYunEcs = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMachineRegion(v string) *DescribeSnapshotsRequest {
	s.MachineRegion = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMachineRemark(v string) *DescribeSnapshotsRequest {
	s.MachineRemark = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageSize(v int32) *DescribeSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetStatusList(v string) *DescribeSnapshotsRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetUuid(v string) *DescribeSnapshotsRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
