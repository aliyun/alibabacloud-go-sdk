// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDistillationTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDistillationTemplates(v []*DistillationTemplateSummary) *ListDistillationTemplatesResponseBody
	GetDistillationTemplates() []*DistillationTemplateSummary
	SetPageNumber(v int32) *ListDistillationTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDistillationTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDistillationTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDistillationTemplatesResponseBody
	GetTotalCount() *int32
}

type ListDistillationTemplatesResponseBody struct {
	DistillationTemplates []*DistillationTemplateSummary `json:"DistillationTemplates,omitempty" xml:"DistillationTemplates,omitempty" type:"Repeated"`
	PageNumber            *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDistillationTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDistillationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDistillationTemplatesResponseBody) GetDistillationTemplates() []*DistillationTemplateSummary {
	return s.DistillationTemplates
}

func (s *ListDistillationTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDistillationTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDistillationTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDistillationTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDistillationTemplatesResponseBody) SetDistillationTemplates(v []*DistillationTemplateSummary) *ListDistillationTemplatesResponseBody {
	s.DistillationTemplates = v
	return s
}

func (s *ListDistillationTemplatesResponseBody) SetPageNumber(v int32) *ListDistillationTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDistillationTemplatesResponseBody) SetPageSize(v int32) *ListDistillationTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDistillationTemplatesResponseBody) SetRequestId(v string) *ListDistillationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDistillationTemplatesResponseBody) SetTotalCount(v int32) *ListDistillationTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDistillationTemplatesResponseBody) Validate() error {
	if s.DistillationTemplates != nil {
		for _, item := range s.DistillationTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
