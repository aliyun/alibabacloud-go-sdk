// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdsFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DeleteCdsFileRequest
	GetCdsId() *string
	SetEndUserId(v string) *DeleteCdsFileRequest
	GetEndUserId() *string
	SetFileId(v string) *DeleteCdsFileRequest
	GetFileId() *string
	SetGroupId(v string) *DeleteCdsFileRequest
	GetGroupId() *string
	SetRegionId(v string) *DeleteCdsFileRequest
	GetRegionId() *string
}

type DeleteCdsFileRequest struct {
	// The ID of the enterprise drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-066224****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the user who uses the network disk.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID. You can call the [ListCdsFiles](https://help.aliyun.com/document_detail/2247622.html) operation to query the ID of the file.
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

func (s DeleteCdsFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdsFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteCdsFileRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DeleteCdsFileRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DeleteCdsFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteCdsFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteCdsFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCdsFileRequest) SetCdsId(v string) *DeleteCdsFileRequest {
	s.CdsId = &v
	return s
}

func (s *DeleteCdsFileRequest) SetEndUserId(v string) *DeleteCdsFileRequest {
	s.EndUserId = &v
	return s
}

func (s *DeleteCdsFileRequest) SetFileId(v string) *DeleteCdsFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteCdsFileRequest) SetGroupId(v string) *DeleteCdsFileRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteCdsFileRequest) SetRegionId(v string) *DeleteCdsFileRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCdsFileRequest) Validate() error {
	return dara.Validate(s)
}
