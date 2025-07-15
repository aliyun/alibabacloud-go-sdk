// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAIStudioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeLiveAIStudioRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLiveAIStudioRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLiveAIStudioRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLiveAIStudioRequest
	GetRegionId() *string
	SetStudioId(v string) *DescribeLiveAIStudioRequest
	GetStudioId() *string
	SetStudioName(v string) *DescribeLiveAIStudioRequest
	GetStudioName() *string
}

type DescribeLiveAIStudioRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 50.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the virtual studio template that you want to query. This parameter is optional.
	//
	// example:
	//
	// dbe61b87-db9a-448f-8757-a875edb3f944
	StudioId *string `json:"StudioId,omitempty" xml:"StudioId,omitempty"`
	// The name of the virtual studio template.
	//
	// example:
	//
	// sub02
	StudioName *string `json:"StudioName,omitempty" xml:"StudioName,omitempty"`
}

func (s DescribeLiveAIStudioRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIStudioRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIStudioRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveAIStudioRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLiveAIStudioRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveAIStudioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveAIStudioRequest) GetStudioId() *string {
	return s.StudioId
}

func (s *DescribeLiveAIStudioRequest) GetStudioName() *string {
	return s.StudioName
}

func (s *DescribeLiveAIStudioRequest) SetOwnerId(v int64) *DescribeLiveAIStudioRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveAIStudioRequest) SetPageNumber(v int32) *DescribeLiveAIStudioRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveAIStudioRequest) SetPageSize(v int32) *DescribeLiveAIStudioRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveAIStudioRequest) SetRegionId(v string) *DescribeLiveAIStudioRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveAIStudioRequest) SetStudioId(v string) *DescribeLiveAIStudioRequest {
	s.StudioId = &v
	return s
}

func (s *DescribeLiveAIStudioRequest) SetStudioName(v string) *DescribeLiveAIStudioRequest {
	s.StudioName = &v
	return s
}

func (s *DescribeLiveAIStudioRequest) Validate() error {
	return dara.Validate(s)
}
