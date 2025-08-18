// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineCodeVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersions(v []*ListRoutineCodeVersionsResponseBodyCodeVersions) *ListRoutineCodeVersionsResponseBody
	GetCodeVersions() []*ListRoutineCodeVersionsResponseBodyCodeVersions
	SetPageNumber(v int64) *ListRoutineCodeVersionsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRoutineCodeVersionsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListRoutineCodeVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRoutineCodeVersionsResponseBody
	GetTotalCount() *int64
}

type ListRoutineCodeVersionsResponseBody struct {
	CodeVersions []*ListRoutineCodeVersionsResponseBodyCodeVersions `json:"CodeVersions,omitempty" xml:"CodeVersions,omitempty" type:"Repeated"`
	PageNumber   *int64                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoutineCodeVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCodeVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutineCodeVersionsResponseBody) GetCodeVersions() []*ListRoutineCodeVersionsResponseBodyCodeVersions {
	return s.CodeVersions
}

func (s *ListRoutineCodeVersionsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRoutineCodeVersionsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRoutineCodeVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutineCodeVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRoutineCodeVersionsResponseBody) SetCodeVersions(v []*ListRoutineCodeVersionsResponseBodyCodeVersions) *ListRoutineCodeVersionsResponseBody {
	s.CodeVersions = v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) SetPageNumber(v int64) *ListRoutineCodeVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) SetPageSize(v int64) *ListRoutineCodeVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) SetRequestId(v string) *ListRoutineCodeVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) SetTotalCount(v int64) *ListRoutineCodeVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRoutineCodeVersionsResponseBodyCodeVersions struct {
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	CodeVersion     *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListRoutineCodeVersionsResponseBodyCodeVersions) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCodeVersionsResponseBodyCodeVersions) GoString() string {
	return s.String()
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetCodeDescription(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.CodeDescription = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetCodeVersion(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.CodeVersion = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetCreateTime(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.CreateTime = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) Validate() error {
	return dara.Validate(s)
}
