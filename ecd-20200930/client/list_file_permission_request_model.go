// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *ListFilePermissionRequest
	GetCdsId() *string
	SetEndUserId(v string) *ListFilePermissionRequest
	GetEndUserId() *string
	SetFileId(v string) *ListFilePermissionRequest
	GetFileId() *string
	SetGroupId(v string) *ListFilePermissionRequest
	GetGroupId() *string
	SetRegionId(v string) *ListFilePermissionRequest
	GetRegionId() *string
}

type ListFilePermissionRequest struct {
	// The ID of the enterprise drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-346063****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the user who uses the drive.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID. You can call the [ListCdsFiles](https://help.aliyun.com/document_detail/2247622.html) operation to get the file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6333e553a133ce21e6f747cf948bb9ef95d7****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the team space.
	//
	// example:
	//
	// cg-c3acvkkbsfkte****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFilePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFilePermissionRequest) GoString() string {
	return s.String()
}

func (s *ListFilePermissionRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *ListFilePermissionRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListFilePermissionRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListFilePermissionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListFilePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListFilePermissionRequest) SetCdsId(v string) *ListFilePermissionRequest {
	s.CdsId = &v
	return s
}

func (s *ListFilePermissionRequest) SetEndUserId(v string) *ListFilePermissionRequest {
	s.EndUserId = &v
	return s
}

func (s *ListFilePermissionRequest) SetFileId(v string) *ListFilePermissionRequest {
	s.FileId = &v
	return s
}

func (s *ListFilePermissionRequest) SetGroupId(v string) *ListFilePermissionRequest {
	s.GroupId = &v
	return s
}

func (s *ListFilePermissionRequest) SetRegionId(v string) *ListFilePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ListFilePermissionRequest) Validate() error {
	return dara.Validate(s)
}
