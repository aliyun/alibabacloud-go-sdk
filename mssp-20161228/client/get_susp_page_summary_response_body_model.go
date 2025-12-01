// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspPageSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSuspPageSummaryResponseBody
	GetCode() *string
	SetData(v *GetSuspPageSummaryResponseBodyData) *GetSuspPageSummaryResponseBody
	GetData() *GetSuspPageSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetSuspPageSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSuspPageSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSuspPageSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSuspPageSummaryResponseBody
	GetSuccess() *bool
}

type GetSuspPageSummaryResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetSuspPageSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the result returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspPageSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSuspPageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSuspPageSummaryResponseBody) GetData() *GetSuspPageSummaryResponseBodyData {
	return s.Data
}

func (s *GetSuspPageSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSuspPageSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSuspPageSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSuspPageSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSuspPageSummaryResponseBody) SetCode(v string) *GetSuspPageSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetData(v *GetSuspPageSummaryResponseBodyData) *GetSuspPageSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetHttpStatusCode(v int32) *GetSuspPageSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetMessage(v string) *GetSuspPageSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetRequestId(v string) *GetSuspPageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetSuccess(v bool) *GetSuspPageSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSuspPageSummaryResponseBodyData struct {
	// Number of completed items.
	//
	// example:
	//
	// 10
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// Number of items being processed.
	//
	// example:
	//
	// 10
	HandingCount *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	// Number of high-risk items.
	//
	// example:
	//
	// 10
	HighCount *int64 `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	// Number of low-risk items.
	//
	// example:
	//
	// 10
	LowCount *int64 `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	// Number of medium-risk items.
	//
	// example:
	//
	// 10
	MediumCount *int64 `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 30
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Number of unhandled items.
	//
	// example:
	//
	// 10
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetSuspPageSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSuspPageSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponseBodyData) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *GetSuspPageSummaryResponseBodyData) GetHandingCount() *int64 {
	return s.HandingCount
}

func (s *GetSuspPageSummaryResponseBodyData) GetHighCount() *int64 {
	return s.HighCount
}

func (s *GetSuspPageSummaryResponseBodyData) GetLowCount() *int64 {
	return s.LowCount
}

func (s *GetSuspPageSummaryResponseBodyData) GetMediumCount() *int64 {
	return s.MediumCount
}

func (s *GetSuspPageSummaryResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetSuspPageSummaryResponseBodyData) GetWaitHandleCount() *int64 {
	return s.WaitHandleCount
}

func (s *GetSuspPageSummaryResponseBodyData) SetCompletedCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetHandingCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.HandingCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetHighCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.HighCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetLowCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.LowCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetMediumCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.MediumCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetTotalCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
