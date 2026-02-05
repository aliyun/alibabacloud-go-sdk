// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyType(v []*string) *ListServiceApplyRequest
	GetApplyType() []*string
	SetEndDate(v int64) *ListServiceApplyRequest
	GetEndDate() *int64
	SetPageNum(v int32) *ListServiceApplyRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListServiceApplyRequest
	GetPageSize() *int32
	SetProductCode(v int32) *ListServiceApplyRequest
	GetProductCode() *int32
	SetStartDate(v int64) *ListServiceApplyRequest
	GetStartDate() *int64
	SetStatus(v string) *ListServiceApplyRequest
	GetStatus() *string
}

type ListServiceApplyRequest struct {
	ApplyType   []*string `json:"applyType,omitempty" xml:"applyType,omitempty" type:"Repeated"`
	EndDate     *int64    `json:"endDate,omitempty" xml:"endDate,omitempty"`
	PageNum     *int32    `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize    *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProductCode *int32    `json:"productCode,omitempty" xml:"productCode,omitempty"`
	StartDate   *int64    `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Status      *string   `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyRequest) GoString() string {
	return s.String()
}

func (s *ListServiceApplyRequest) GetApplyType() []*string {
	return s.ApplyType
}

func (s *ListServiceApplyRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListServiceApplyRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListServiceApplyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServiceApplyRequest) GetProductCode() *int32 {
	return s.ProductCode
}

func (s *ListServiceApplyRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListServiceApplyRequest) GetStatus() *string {
	return s.Status
}

func (s *ListServiceApplyRequest) SetApplyType(v []*string) *ListServiceApplyRequest {
	s.ApplyType = v
	return s
}

func (s *ListServiceApplyRequest) SetEndDate(v int64) *ListServiceApplyRequest {
	s.EndDate = &v
	return s
}

func (s *ListServiceApplyRequest) SetPageNum(v int32) *ListServiceApplyRequest {
	s.PageNum = &v
	return s
}

func (s *ListServiceApplyRequest) SetPageSize(v int32) *ListServiceApplyRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceApplyRequest) SetProductCode(v int32) *ListServiceApplyRequest {
	s.ProductCode = &v
	return s
}

func (s *ListServiceApplyRequest) SetStartDate(v int64) *ListServiceApplyRequest {
	s.StartDate = &v
	return s
}

func (s *ListServiceApplyRequest) SetStatus(v string) *ListServiceApplyRequest {
	s.Status = &v
	return s
}

func (s *ListServiceApplyRequest) Validate() error {
	return dara.Validate(s)
}
