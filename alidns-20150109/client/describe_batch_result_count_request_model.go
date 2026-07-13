// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchResultCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchType(v string) *DescribeBatchResultCountRequest
	GetBatchType() *string
	SetLang(v string) *DescribeBatchResultCountRequest
	GetLang() *string
	SetTaskId(v int64) *DescribeBatchResultCountRequest
	GetTaskId() *int64
}

type DescribeBatchResultCountRequest struct {
	// The type of the batch operation. Valid values:
	//
	// - **DOMAIN_ADD**: Batch add domain names.
	//
	// - **DOMAIN_DEL**: Batch delete domain names.
	//
	// - **RR_ADD**: Batch add DNS records.
	//
	// - **RR_DEL**: Batch delete DNS records.
	//
	// > If you leave this parameter empty, no filter is applied.
	//
	// example:
	//
	// DOMAIN_ADD
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The task ID.
	//
	// > If you specify a task ID, the results of that task are returned. If you leave this parameter empty, the results of the most recent task are returned.
	//
	// example:
	//
	// 1****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeBatchResultCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchResultCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultCountRequest) GetBatchType() *string {
	return s.BatchType
}

func (s *DescribeBatchResultCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBatchResultCountRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeBatchResultCountRequest) SetBatchType(v string) *DescribeBatchResultCountRequest {
	s.BatchType = &v
	return s
}

func (s *DescribeBatchResultCountRequest) SetLang(v string) *DescribeBatchResultCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBatchResultCountRequest) SetTaskId(v int64) *DescribeBatchResultCountRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeBatchResultCountRequest) Validate() error {
	return dara.Validate(s)
}
