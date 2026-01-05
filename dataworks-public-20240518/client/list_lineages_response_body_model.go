// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListLineagesResponseBodyPagingInfo) *ListLineagesResponseBody
	GetPagingInfo() *ListLineagesResponseBodyPagingInfo
	SetRequestId(v string) *ListLineagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLineagesResponseBody
	GetSuccess() *bool
}

type ListLineagesResponseBody struct {
	PagingInfo *ListLineagesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLineagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLineagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLineagesResponseBody) GetPagingInfo() *ListLineagesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListLineagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLineagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLineagesResponseBody) SetPagingInfo(v *ListLineagesResponseBodyPagingInfo) *ListLineagesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListLineagesResponseBody) SetRequestId(v string) *ListLineagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLineagesResponseBody) SetSuccess(v bool) *ListLineagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLineagesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLineagesResponseBodyPagingInfo struct {
	Lineages []*ListLineagesResponseBodyPagingInfoLineages `json:"Lineages,omitempty" xml:"Lineages,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLineagesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListLineagesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListLineagesResponseBodyPagingInfo) GetLineages() []*ListLineagesResponseBodyPagingInfoLineages {
	return s.Lineages
}

func (s *ListLineagesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLineagesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLineagesResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLineagesResponseBodyPagingInfo) SetLineages(v []*ListLineagesResponseBodyPagingInfoLineages) *ListLineagesResponseBodyPagingInfo {
	s.Lineages = v
	return s
}

func (s *ListLineagesResponseBodyPagingInfo) SetPageNumber(v int32) *ListLineagesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListLineagesResponseBodyPagingInfo) SetPageSize(v int32) *ListLineagesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListLineagesResponseBodyPagingInfo) SetTotalCount(v int64) *ListLineagesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListLineagesResponseBodyPagingInfo) Validate() error {
	if s.Lineages != nil {
		for _, item := range s.Lineages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLineagesResponseBodyPagingInfoLineages struct {
	DstEntity     *LineageEntity         `json:"DstEntity,omitempty" xml:"DstEntity,omitempty"`
	Relationships []*LineageRelationship `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
	SrcEntity     *LineageEntity         `json:"SrcEntity,omitempty" xml:"SrcEntity,omitempty"`
}

func (s ListLineagesResponseBodyPagingInfoLineages) String() string {
	return dara.Prettify(s)
}

func (s ListLineagesResponseBodyPagingInfoLineages) GoString() string {
	return s.String()
}

func (s *ListLineagesResponseBodyPagingInfoLineages) GetDstEntity() *LineageEntity {
	return s.DstEntity
}

func (s *ListLineagesResponseBodyPagingInfoLineages) GetRelationships() []*LineageRelationship {
	return s.Relationships
}

func (s *ListLineagesResponseBodyPagingInfoLineages) GetSrcEntity() *LineageEntity {
	return s.SrcEntity
}

func (s *ListLineagesResponseBodyPagingInfoLineages) SetDstEntity(v *LineageEntity) *ListLineagesResponseBodyPagingInfoLineages {
	s.DstEntity = v
	return s
}

func (s *ListLineagesResponseBodyPagingInfoLineages) SetRelationships(v []*LineageRelationship) *ListLineagesResponseBodyPagingInfoLineages {
	s.Relationships = v
	return s
}

func (s *ListLineagesResponseBodyPagingInfoLineages) SetSrcEntity(v *LineageEntity) *ListLineagesResponseBodyPagingInfoLineages {
	s.SrcEntity = v
	return s
}

func (s *ListLineagesResponseBodyPagingInfoLineages) Validate() error {
	if s.DstEntity != nil {
		if err := s.DstEntity.Validate(); err != nil {
			return err
		}
	}
	if s.Relationships != nil {
		for _, item := range s.Relationships {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SrcEntity != nil {
		if err := s.SrcEntity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
