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
	// The list of distillation template summaries, sorted by OrderNumber in ascending order.
	DistillationTemplates []*DistillationTemplateSummary `json:"DistillationTemplates,omitempty" xml:"DistillationTemplates,omitempty" type:"Repeated"`
	// The page number, which echoes the PageNumber value in the request.
	//
	// example:
	//
	// 4
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page, which echoes the PageSize value in the request.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of templates that match the filter conditions.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
