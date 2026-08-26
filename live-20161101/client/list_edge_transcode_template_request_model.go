// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeTranscodeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListEdgeTranscodeTemplateRequest
	GetClusterId() *string
	SetKeyword(v string) *ListEdgeTranscodeTemplateRequest
	GetKeyword() *string
	SetOwnerId(v int64) *ListEdgeTranscodeTemplateRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *ListEdgeTranscodeTemplateRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListEdgeTranscodeTemplateRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListEdgeTranscodeTemplateRequest
	GetRegionId() *string
	SetSortBy(v string) *ListEdgeTranscodeTemplateRequest
	GetSortBy() *string
	SetType(v string) *ListEdgeTranscodeTemplateRequest
	GetType() *string
	SetVideoCodec(v string) *ListEdgeTranscodeTemplateRequest
	GetVideoCodec() *string
}

type ListEdgeTranscodeTemplateRequest struct {
	// The data center ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******3b-4d18-395c-8106-ff21a6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The search keyword. Valid values:
	//
	// - Template ID. Exact match is supported.
	//
	// - Template name. Fuzzy match is supported.
	//
	// example:
	//
	// baseline
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The sorting rule. Templates are sorted by creation time (CreateTime). Default value: desc. Valid values:
	//
	// - desc: descending order.
	//
	// - asc: ascending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The edge transcoding type. Valid values:
	//
	// - **common**: default transcoding (standard + Narrowband HD 1.0).
	//
	// - **nbhd-2**: Narrowband HD 2.0.
	//
	// - **ultra-hd**: ultra-high definition.
	//
	// > If this parameter is not specified, the system displays transcoding templates for the transcoding types that the user has permissions to access.
	//
	// example:
	//
	// nbhd-2
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The video encoding format. Valid values:
	//
	// - H.264.
	//
	// - H.265.
	//
	// > If this parameter is not specified, the system displays transcoding templates for the video encoding formats that the user has permissions to access.
	//
	// example:
	//
	// H.264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
}

func (s ListEdgeTranscodeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeTemplateRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListEdgeTranscodeTemplateRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListEdgeTranscodeTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListEdgeTranscodeTemplateRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListEdgeTranscodeTemplateRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeTranscodeTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEdgeTranscodeTemplateRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListEdgeTranscodeTemplateRequest) GetType() *string {
	return s.Type
}

func (s *ListEdgeTranscodeTemplateRequest) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *ListEdgeTranscodeTemplateRequest) SetClusterId(v string) *ListEdgeTranscodeTemplateRequest {
	s.ClusterId = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) SetKeyword(v string) *ListEdgeTranscodeTemplateRequest {
	s.Keyword = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) SetOwnerId(v int64) *ListEdgeTranscodeTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) SetPageNo(v int32) *ListEdgeTranscodeTemplateRequest {
	s.PageNo = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) SetPageSize(v int32) *ListEdgeTranscodeTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) SetRegionId(v string) *ListEdgeTranscodeTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) SetSortBy(v string) *ListEdgeTranscodeTemplateRequest {
	s.SortBy = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) SetType(v string) *ListEdgeTranscodeTemplateRequest {
	s.Type = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) SetVideoCodec(v string) *ListEdgeTranscodeTemplateRequest {
	s.VideoCodec = &v
	return s
}

func (s *ListEdgeTranscodeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
