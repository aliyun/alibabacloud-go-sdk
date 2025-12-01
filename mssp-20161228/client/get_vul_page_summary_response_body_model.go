// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulPageSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVulPageSummaryResponseBody
	GetCode() *string
	SetData(v *GetVulPageSummaryResponseBodyData) *GetVulPageSummaryResponseBody
	GetData() *GetVulPageSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetVulPageSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVulPageSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVulPageSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVulPageSummaryResponseBody
	GetSuccess() *bool
}

type GetVulPageSummaryResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetVulPageSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// A3A575C8-80F9-5F04-AA24-CCAC246884A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulPageSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVulPageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVulPageSummaryResponseBody) GetData() *GetVulPageSummaryResponseBodyData {
	return s.Data
}

func (s *GetVulPageSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVulPageSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVulPageSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVulPageSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVulPageSummaryResponseBody) SetCode(v string) *GetVulPageSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetData(v *GetVulPageSummaryResponseBodyData) *GetVulPageSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetHttpStatusCode(v int32) *GetVulPageSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetMessage(v string) *GetVulPageSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetRequestId(v string) *GetVulPageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetSuccess(v bool) *GetVulPageSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVulPageSummaryResponseBodyData struct {
	// Number of completed items.
	//
	// example:
	//
	// 1990
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// Number of items being handled.
	//
	// example:
	//
	// 6
	HandingCount *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	// Number of high-risk items.
	//
	// example:
	//
	// 500
	HighCount *int64 `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	// Number of low-risk items.
	//
	// example:
	//
	// 1000
	LowCount *int64 `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	// Number of medium-risk items.
	//
	// example:
	//
	// 500
	MediumCount *int64 `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 2000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Number of unhandled items.
	//
	// example:
	//
	// 4
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetVulPageSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVulPageSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponseBodyData) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *GetVulPageSummaryResponseBodyData) GetHandingCount() *int64 {
	return s.HandingCount
}

func (s *GetVulPageSummaryResponseBodyData) GetHighCount() *int64 {
	return s.HighCount
}

func (s *GetVulPageSummaryResponseBodyData) GetLowCount() *int64 {
	return s.LowCount
}

func (s *GetVulPageSummaryResponseBodyData) GetMediumCount() *int64 {
	return s.MediumCount
}

func (s *GetVulPageSummaryResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetVulPageSummaryResponseBodyData) GetWaitHandleCount() *int64 {
	return s.WaitHandleCount
}

func (s *GetVulPageSummaryResponseBodyData) SetCompletedCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetHandingCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.HandingCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetHighCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.HighCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetLowCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.LowCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetMediumCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.MediumCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetTotalCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
