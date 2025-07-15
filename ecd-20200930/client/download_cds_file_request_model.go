// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadCdsFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DownloadCdsFileRequest
	GetCdsId() *string
	SetEndUserId(v string) *DownloadCdsFileRequest
	GetEndUserId() *string
	SetFileId(v string) *DownloadCdsFileRequest
	GetFileId() *string
	SetGroupId(v string) *DownloadCdsFileRequest
	GetGroupId() *string
	SetRegionId(v string) *DownloadCdsFileRequest
	GetRegionId() *string
}

type DownloadCdsFileRequest struct {
	// The enterprise drive ID.
	//
	// example:
	//
	// cn-hangzhou+cds-643267****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// user****
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 63f3257b68b018170b194d87b875512d108f****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The team ID.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DownloadCdsFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadCdsFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadCdsFileRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DownloadCdsFileRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DownloadCdsFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DownloadCdsFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DownloadCdsFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DownloadCdsFileRequest) SetCdsId(v string) *DownloadCdsFileRequest {
	s.CdsId = &v
	return s
}

func (s *DownloadCdsFileRequest) SetEndUserId(v string) *DownloadCdsFileRequest {
	s.EndUserId = &v
	return s
}

func (s *DownloadCdsFileRequest) SetFileId(v string) *DownloadCdsFileRequest {
	s.FileId = &v
	return s
}

func (s *DownloadCdsFileRequest) SetGroupId(v string) *DownloadCdsFileRequest {
	s.GroupId = &v
	return s
}

func (s *DownloadCdsFileRequest) SetRegionId(v string) *DownloadCdsFileRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadCdsFileRequest) Validate() error {
	return dara.Validate(s)
}
