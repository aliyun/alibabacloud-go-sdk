// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchChangeTableOwnerStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetBatchChangeTableOwnerStatusResponseBodyData) *GetBatchChangeTableOwnerStatusResponseBody
	GetData() *GetBatchChangeTableOwnerStatusResponseBodyData
	SetRequestId(v string) *GetBatchChangeTableOwnerStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBatchChangeTableOwnerStatusResponseBody
	GetSuccess() *bool
}

type GetBatchChangeTableOwnerStatusResponseBody struct {
	Data *GetBatchChangeTableOwnerStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 01A017D3-207E-582C-A683-BE991E54051D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBatchChangeTableOwnerStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchChangeTableOwnerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchChangeTableOwnerStatusResponseBody) GetData() *GetBatchChangeTableOwnerStatusResponseBodyData {
	return s.Data
}

func (s *GetBatchChangeTableOwnerStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchChangeTableOwnerStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBatchChangeTableOwnerStatusResponseBody) SetData(v *GetBatchChangeTableOwnerStatusResponseBodyData) *GetBatchChangeTableOwnerStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBody) SetRequestId(v string) *GetBatchChangeTableOwnerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBody) SetSuccess(v bool) *GetBatchChangeTableOwnerStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchChangeTableOwnerStatusResponseBodyData struct {
	// example:
	//
	// 524257_openapi-req-abc123
	BatchId *string                                                  `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	Details []*GetBatchChangeTableOwnerStatusResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// example:
	//
	// 5
	OngoingCount *int32 `json:"OngoingCount,omitempty" xml:"OngoingCount,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetBatchChangeTableOwnerStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBatchChangeTableOwnerStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) GetBatchId() *string {
	return s.BatchId
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) GetDetails() []*GetBatchChangeTableOwnerStatusResponseBodyDataDetails {
	return s.Details
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) GetOngoingCount() *int32 {
	return s.OngoingCount
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) SetBatchId(v string) *GetBatchChangeTableOwnerStatusResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) SetDetails(v []*GetBatchChangeTableOwnerStatusResponseBodyDataDetails) *GetBatchChangeTableOwnerStatusResponseBodyData {
	s.Details = v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) SetFailedCount(v int32) *GetBatchChangeTableOwnerStatusResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) SetOngoingCount(v int32) *GetBatchChangeTableOwnerStatusResponseBodyData {
	s.OngoingCount = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) SetStatus(v string) *GetBatchChangeTableOwnerStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) SetSuccessCount(v int32) *GetBatchChangeTableOwnerStatusResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) SetTotalCount(v int32) *GetBatchChangeTableOwnerStatusResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyData) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchChangeTableOwnerStatusResponseBodyDataDetails struct {
	// example:
	//
	// Forbidden: You are not a member of this project
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// maxcompute-table:123:project_a::table_1
	TableMetaEntityId *string `json:"TableMetaEntityId,omitempty" xml:"TableMetaEntityId,omitempty"`
}

func (s GetBatchChangeTableOwnerStatusResponseBodyDataDetails) String() string {
	return dara.Prettify(s)
}

func (s GetBatchChangeTableOwnerStatusResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyDataDetails) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyDataDetails) GetStatus() *string {
	return s.Status
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyDataDetails) GetTableMetaEntityId() *string {
	return s.TableMetaEntityId
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyDataDetails) SetErrorMessage(v string) *GetBatchChangeTableOwnerStatusResponseBodyDataDetails {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyDataDetails) SetStatus(v string) *GetBatchChangeTableOwnerStatusResponseBodyDataDetails {
	s.Status = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyDataDetails) SetTableMetaEntityId(v string) *GetBatchChangeTableOwnerStatusResponseBodyDataDetails {
	s.TableMetaEntityId = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponseBodyDataDetails) Validate() error {
	return dara.Validate(s)
}
