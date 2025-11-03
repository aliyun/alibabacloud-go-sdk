// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdsFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *ModifyCdsFileRequest
	GetCdsId() *string
	SetConflictPolicy(v string) *ModifyCdsFileRequest
	GetConflictPolicy() *string
	SetEndUserId(v string) *ModifyCdsFileRequest
	GetEndUserId() *string
	SetFileId(v string) *ModifyCdsFileRequest
	GetFileId() *string
	SetFileName(v string) *ModifyCdsFileRequest
	GetFileName() *string
	SetGroupId(v string) *ModifyCdsFileRequest
	GetGroupId() *string
	SetRegionId(v string) *ModifyCdsFileRequest
	GetRegionId() *string
}

type ModifyCdsFileRequest struct {
	// The enterprise drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-64326*****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The processing policy when a file with the same name appears.
	//
	// Valid values:
	//
	// 	- refuse: If you want to create a file that uses the same name as an existing file in the cloud, the system denies your request and returns the details of the existing file.
	//
	// 	- auto_rename: automatically renames a file if the file has the same name as an existing file in the cloud. By default, the current point in time is appended to the end of the original file name. Example: xxx20240102_150405.
	//
	// 	- ignore: allows the file to be with the same name.
	//
	// 	- over_write: After you create a file that uses the same name as an existing file in the cloud, the new file overwrites the existing file.
	//
	// example:
	//
	// ignore
	ConflictPolicy *string `json:"ConflictPolicy,omitempty" xml:"ConflictPolicy,omitempty"`
	// The ID of the user who uses the network disk.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the file. You can call the [ListCdsFiles](https://help.aliyun.com/document_detail/2247622.html) operation to query the ID of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6333e553a133ce21e6f747cf948bb9ef95d7****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The ID of the team space.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
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

func (s ModifyCdsFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdsFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdsFileRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *ModifyCdsFileRequest) GetConflictPolicy() *string {
	return s.ConflictPolicy
}

func (s *ModifyCdsFileRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ModifyCdsFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *ModifyCdsFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *ModifyCdsFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyCdsFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCdsFileRequest) SetCdsId(v string) *ModifyCdsFileRequest {
	s.CdsId = &v
	return s
}

func (s *ModifyCdsFileRequest) SetConflictPolicy(v string) *ModifyCdsFileRequest {
	s.ConflictPolicy = &v
	return s
}

func (s *ModifyCdsFileRequest) SetEndUserId(v string) *ModifyCdsFileRequest {
	s.EndUserId = &v
	return s
}

func (s *ModifyCdsFileRequest) SetFileId(v string) *ModifyCdsFileRequest {
	s.FileId = &v
	return s
}

func (s *ModifyCdsFileRequest) SetFileName(v string) *ModifyCdsFileRequest {
	s.FileName = &v
	return s
}

func (s *ModifyCdsFileRequest) SetGroupId(v string) *ModifyCdsFileRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyCdsFileRequest) SetRegionId(v string) *ModifyCdsFileRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCdsFileRequest) Validate() error {
	return dara.Validate(s)
}
