// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgLabRecordListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUid(v string) *QueryOrgLabRecordListRequest
	GetAliyunUid() *string
	SetLabId(v string) *QueryOrgLabRecordListRequest
	GetLabId() *string
	SetOrderBy(v string) *QueryOrgLabRecordListRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *QueryOrgLabRecordListRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *QueryOrgLabRecordListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *QueryOrgLabRecordListRequest
	GetPageSize() *int32
}

type QueryOrgLabRecordListRequest struct {
	// This parameter is required.
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// This parameter is required.
	LabId          *string `json:"LabId,omitempty" xml:"LabId,omitempty"`
	OrderBy        *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	PageIndex      *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryOrgLabRecordListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgLabRecordListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgLabRecordListRequest) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *QueryOrgLabRecordListRequest) GetLabId() *string {
	return s.LabId
}

func (s *QueryOrgLabRecordListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryOrgLabRecordListRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *QueryOrgLabRecordListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *QueryOrgLabRecordListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryOrgLabRecordListRequest) SetAliyunUid(v string) *QueryOrgLabRecordListRequest {
	s.AliyunUid = &v
	return s
}

func (s *QueryOrgLabRecordListRequest) SetLabId(v string) *QueryOrgLabRecordListRequest {
	s.LabId = &v
	return s
}

func (s *QueryOrgLabRecordListRequest) SetOrderBy(v string) *QueryOrgLabRecordListRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryOrgLabRecordListRequest) SetOrderDirection(v string) *QueryOrgLabRecordListRequest {
	s.OrderDirection = &v
	return s
}

func (s *QueryOrgLabRecordListRequest) SetPageIndex(v int32) *QueryOrgLabRecordListRequest {
	s.PageIndex = &v
	return s
}

func (s *QueryOrgLabRecordListRequest) SetPageSize(v int32) *QueryOrgLabRecordListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrgLabRecordListRequest) Validate() error {
	return dara.Validate(s)
}
