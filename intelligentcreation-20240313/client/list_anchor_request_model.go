// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnchorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnchorCategory(v string) *ListAnchorRequest
	GetAnchorCategory() *string
	SetAnchorId(v string) *ListAnchorRequest
	GetAnchorId() *string
	SetAnchorType(v string) *ListAnchorRequest
	GetAnchorType() *string
	SetCoverRate(v string) *ListAnchorRequest
	GetCoverRate() *string
	SetDigitalHumanType(v string) *ListAnchorRequest
	GetDigitalHumanType() *string
	SetPageNumber(v int32) *ListAnchorRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAnchorRequest
	GetPageSize() *int32
	SetResSpecType(v string) *ListAnchorRequest
	GetResSpecType() *string
	SetUseScene(v string) *ListAnchorRequest
	GetUseScene() *string
}

type ListAnchorRequest struct {
	AnchorCategory *string `json:"anchorCategory,omitempty" xml:"anchorCategory,omitempty"`
	AnchorId       *string `json:"anchorId,omitempty" xml:"anchorId,omitempty"`
	// example:
	//
	// PUBLIC_MODEL
	AnchorType *string `json:"anchorType,omitempty" xml:"anchorType,omitempty"`
	// example:
	//
	// 9:16
	CoverRate *string `json:"coverRate,omitempty" xml:"coverRate,omitempty"`
	// example:
	//
	// staticTransparency
	DigitalHumanType *string `json:"digitalHumanType,omitempty" xml:"digitalHumanType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// offlineSynthesis
	UseScene *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
}

func (s ListAnchorRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnchorRequest) GoString() string {
	return s.String()
}

func (s *ListAnchorRequest) GetAnchorCategory() *string {
	return s.AnchorCategory
}

func (s *ListAnchorRequest) GetAnchorId() *string {
	return s.AnchorId
}

func (s *ListAnchorRequest) GetAnchorType() *string {
	return s.AnchorType
}

func (s *ListAnchorRequest) GetCoverRate() *string {
	return s.CoverRate
}

func (s *ListAnchorRequest) GetDigitalHumanType() *string {
	return s.DigitalHumanType
}

func (s *ListAnchorRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAnchorRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAnchorRequest) GetResSpecType() *string {
	return s.ResSpecType
}

func (s *ListAnchorRequest) GetUseScene() *string {
	return s.UseScene
}

func (s *ListAnchorRequest) SetAnchorCategory(v string) *ListAnchorRequest {
	s.AnchorCategory = &v
	return s
}

func (s *ListAnchorRequest) SetAnchorId(v string) *ListAnchorRequest {
	s.AnchorId = &v
	return s
}

func (s *ListAnchorRequest) SetAnchorType(v string) *ListAnchorRequest {
	s.AnchorType = &v
	return s
}

func (s *ListAnchorRequest) SetCoverRate(v string) *ListAnchorRequest {
	s.CoverRate = &v
	return s
}

func (s *ListAnchorRequest) SetDigitalHumanType(v string) *ListAnchorRequest {
	s.DigitalHumanType = &v
	return s
}

func (s *ListAnchorRequest) SetPageNumber(v int32) *ListAnchorRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAnchorRequest) SetPageSize(v int32) *ListAnchorRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnchorRequest) SetResSpecType(v string) *ListAnchorRequest {
	s.ResSpecType = &v
	return s
}

func (s *ListAnchorRequest) SetUseScene(v string) *ListAnchorRequest {
	s.UseScene = &v
	return s
}

func (s *ListAnchorRequest) Validate() error {
	return dara.Validate(s)
}
