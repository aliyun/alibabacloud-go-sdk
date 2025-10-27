// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientProblemTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeClientProblemTypeResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeClientProblemTypeResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeClientProblemTypeResponseBody
	GetPageSize() *int32
	SetProblemTypes(v []*DescribeClientProblemTypeResponseBodyProblemTypes) *DescribeClientProblemTypeResponseBody
	GetProblemTypes() []*DescribeClientProblemTypeResponseBodyProblemTypes
	SetRequestId(v string) *DescribeClientProblemTypeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeClientProblemTypeResponseBody
	GetTotalCount() *int32
}

type DescribeClientProblemTypeResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The issue types.
	ProblemTypes []*DescribeClientProblemTypeResponseBodyProblemTypes `json:"ProblemTypes,omitempty" xml:"ProblemTypes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53CXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 21
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeClientProblemTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientProblemTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientProblemTypeResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeClientProblemTypeResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeClientProblemTypeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClientProblemTypeResponseBody) GetProblemTypes() []*DescribeClientProblemTypeResponseBodyProblemTypes {
	return s.ProblemTypes
}

func (s *DescribeClientProblemTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientProblemTypeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClientProblemTypeResponseBody) SetCount(v int32) *DescribeClientProblemTypeResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeClientProblemTypeResponseBody) SetCurrentPage(v int32) *DescribeClientProblemTypeResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeClientProblemTypeResponseBody) SetPageSize(v int32) *DescribeClientProblemTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeClientProblemTypeResponseBody) SetProblemTypes(v []*DescribeClientProblemTypeResponseBodyProblemTypes) *DescribeClientProblemTypeResponseBody {
	s.ProblemTypes = v
	return s
}

func (s *DescribeClientProblemTypeResponseBody) SetRequestId(v string) *DescribeClientProblemTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientProblemTypeResponseBody) SetTotalCount(v int32) *DescribeClientProblemTypeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeClientProblemTypeResponseBody) Validate() error {
	if s.ProblemTypes != nil {
		for _, item := range s.ProblemTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClientProblemTypeResponseBodyProblemTypes struct {
	// The description of the issue type.
	//
	// example:
	//
	// other
	ProblemDetail *string `json:"problemDetail,omitempty" xml:"problemDetail,omitempty"`
	// The ID of the issue type.
	//
	// example:
	//
	// 4
	ProblemId *string `json:"problemId,omitempty" xml:"problemId,omitempty"`
	// The name of the issue type.
	//
	// example:
	//
	// high_cpu
	ProblemType *string `json:"problemType,omitempty" xml:"problemType,omitempty"`
}

func (s DescribeClientProblemTypeResponseBodyProblemTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientProblemTypeResponseBodyProblemTypes) GoString() string {
	return s.String()
}

func (s *DescribeClientProblemTypeResponseBodyProblemTypes) GetProblemDetail() *string {
	return s.ProblemDetail
}

func (s *DescribeClientProblemTypeResponseBodyProblemTypes) GetProblemId() *string {
	return s.ProblemId
}

func (s *DescribeClientProblemTypeResponseBodyProblemTypes) GetProblemType() *string {
	return s.ProblemType
}

func (s *DescribeClientProblemTypeResponseBodyProblemTypes) SetProblemDetail(v string) *DescribeClientProblemTypeResponseBodyProblemTypes {
	s.ProblemDetail = &v
	return s
}

func (s *DescribeClientProblemTypeResponseBodyProblemTypes) SetProblemId(v string) *DescribeClientProblemTypeResponseBodyProblemTypes {
	s.ProblemId = &v
	return s
}

func (s *DescribeClientProblemTypeResponseBodyProblemTypes) SetProblemType(v string) *DescribeClientProblemTypeResponseBodyProblemTypes {
	s.ProblemType = &v
	return s
}

func (s *DescribeClientProblemTypeResponseBodyProblemTypes) Validate() error {
	return dara.Validate(s)
}
