// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteCdsFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *CompleteCdsFileRequest
	GetCdsId() *string
	SetEndUserId(v string) *CompleteCdsFileRequest
	GetEndUserId() *string
	SetFileId(v string) *CompleteCdsFileRequest
	GetFileId() *string
	SetGroupId(v string) *CompleteCdsFileRequest
	GetGroupId() *string
	SetRegionId(v string) *CompleteCdsFileRequest
	GetRegionId() *string
	SetUploadId(v string) *CompleteCdsFileRequest
	GetUploadId() *string
}

type CompleteCdsFileRequest struct {
	// The ID of the enterprise drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+cds-465878****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The username. You must specify at least one of `EndUserId` and `GroupId`.
	//
	// example:
	//
	// test0
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 635a316c94f40f35f5354da29b2aee88c9d1****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the team space. You must specify at least one of `EndUserId` and `GroupId`.
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
	// The ID of the file upload task. Callable interface [CreateCdsFile](https://help.aliyun.com/document_detail/2247619.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 6C48B55A1FAC4E1A9E0579059514****
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s CompleteCdsFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteCdsFileRequest) GoString() string {
	return s.String()
}

func (s *CompleteCdsFileRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *CompleteCdsFileRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CompleteCdsFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *CompleteCdsFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CompleteCdsFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CompleteCdsFileRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *CompleteCdsFileRequest) SetCdsId(v string) *CompleteCdsFileRequest {
	s.CdsId = &v
	return s
}

func (s *CompleteCdsFileRequest) SetEndUserId(v string) *CompleteCdsFileRequest {
	s.EndUserId = &v
	return s
}

func (s *CompleteCdsFileRequest) SetFileId(v string) *CompleteCdsFileRequest {
	s.FileId = &v
	return s
}

func (s *CompleteCdsFileRequest) SetGroupId(v string) *CompleteCdsFileRequest {
	s.GroupId = &v
	return s
}

func (s *CompleteCdsFileRequest) SetRegionId(v string) *CompleteCdsFileRequest {
	s.RegionId = &v
	return s
}

func (s *CompleteCdsFileRequest) SetUploadId(v string) *CompleteCdsFileRequest {
	s.UploadId = &v
	return s
}

func (s *CompleteCdsFileRequest) Validate() error {
	return dara.Validate(s)
}
