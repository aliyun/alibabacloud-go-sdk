// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchChangeTableOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitBatchChangeTableOwnerResponseBodyData) *SubmitBatchChangeTableOwnerResponseBody
	GetData() *SubmitBatchChangeTableOwnerResponseBodyData
	SetRequestId(v string) *SubmitBatchChangeTableOwnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitBatchChangeTableOwnerResponseBody
	GetSuccess() *bool
}

type SubmitBatchChangeTableOwnerResponseBody struct {
	Data *SubmitBatchChangeTableOwnerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitBatchChangeTableOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchChangeTableOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBatchChangeTableOwnerResponseBody) GetData() *SubmitBatchChangeTableOwnerResponseBodyData {
	return s.Data
}

func (s *SubmitBatchChangeTableOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitBatchChangeTableOwnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitBatchChangeTableOwnerResponseBody) SetData(v *SubmitBatchChangeTableOwnerResponseBodyData) *SubmitBatchChangeTableOwnerResponseBody {
	s.Data = v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponseBody) SetRequestId(v string) *SubmitBatchChangeTableOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponseBody) SetSuccess(v bool) *SubmitBatchChangeTableOwnerResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitBatchChangeTableOwnerResponseBodyData struct {
	// example:
	//
	// 524257_xxxxx
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// SUBMITTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SubmitBatchChangeTableOwnerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchChangeTableOwnerResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitBatchChangeTableOwnerResponseBodyData) GetBatchId() *string {
	return s.BatchId
}

func (s *SubmitBatchChangeTableOwnerResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitBatchChangeTableOwnerResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SubmitBatchChangeTableOwnerResponseBodyData) SetBatchId(v string) *SubmitBatchChangeTableOwnerResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponseBodyData) SetStatus(v string) *SubmitBatchChangeTableOwnerResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponseBodyData) SetTotalCount(v int32) *SubmitBatchChangeTableOwnerResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
