// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchResultDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchType(v string) *DescribeBatchResultDetailRequest
	GetBatchType() *string
	SetLang(v string) *DescribeBatchResultDetailRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeBatchResultDetailRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBatchResultDetailRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeBatchResultDetailRequest
	GetStatus() *string
	SetTaskId(v int64) *DescribeBatchResultDetailRequest
	GetTaskId() *int64
}

type DescribeBatchResultDetailRequest struct {
	// The type of the batch operation. Valid values:
	//
	// 	- **DOMAIN_ADD**: adds domain names in batches.
	//
	// 	- **DOMAIN_DEL**: deletes domain names in batches.
	//
	// 	- **RR_ADD**: adds Domain Name System (DNS) records in batches.
	//
	// 	- **RR_DEL**: deletes DNS records in batches.
	//
	// >  Do not perform filtering when this field is empty.
	//
	// example:
	//
	// DOMAIN_ADD
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The language of the content within the request and response. Default: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The execution result. If you do not specify this parameter, all results are returned.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 83618818
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeBatchResultDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchResultDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultDetailRequest) GetBatchType() *string {
	return s.BatchType
}

func (s *DescribeBatchResultDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBatchResultDetailRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBatchResultDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBatchResultDetailRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeBatchResultDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeBatchResultDetailRequest) SetBatchType(v string) *DescribeBatchResultDetailRequest {
	s.BatchType = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetLang(v string) *DescribeBatchResultDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetPageNumber(v int32) *DescribeBatchResultDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetPageSize(v int32) *DescribeBatchResultDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetStatus(v string) *DescribeBatchResultDetailRequest {
	s.Status = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) SetTaskId(v int64) *DescribeBatchResultDetailRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeBatchResultDetailRequest) Validate() error {
	return dara.Validate(s)
}
