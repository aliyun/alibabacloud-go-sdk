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
	// The ID of the data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******3b-4d18-395c-8106-ff21a6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The keyword of the query.
	//
	// 	- You can specify a template ID for an exact match.
	//
	// 	- You can also specify a template name for a fuzzy match.
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
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The sorting order of the templates by creation time. Default value: desc. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of edge transcoding. Valid values:
	//
	// 	- **common**: standard transcoding and Narrowband HD™ 1.0 transcoding.
	//
	// 	- **nbhd-2**: Narrowband HD™ 2.0 transcoding.
	//
	// 	- **ultra-hd**: ultra-high definition transcoding.
	//
	// >  If you do not specify this parameter, the query result is filtered based on the type of edge transcoding on which you are granted permissions.
	//
	// example:
	//
	// nbhd-2
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The video encoding format. Valid values:
	//
	// 	- H.264
	//
	// 	- H.265
	//
	// >  If you do not specify this parameter, the query result is filtered based on the video encoding format on which you are granted permissions.
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
