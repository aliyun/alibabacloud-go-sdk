// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSummaryJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SummaryJobDetailResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SummaryJobDetailResponseBody
	GetHttpStatusCode() *int32
	SetJobId(v string) *SummaryJobDetailResponseBody
	GetJobId() *string
	SetProgressSummaryDetails(v []*SummaryJobDetailResponseBodyProgressSummaryDetails) *SummaryJobDetailResponseBody
	GetProgressSummaryDetails() []*SummaryJobDetailResponseBodyProgressSummaryDetails
	SetRequestId(v string) *SummaryJobDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SummaryJobDetailResponseBody
	GetSuccess() *bool
}

type SummaryJobDetailResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the data migration or synchronization task.
	//
	// example:
	//
	// l3m1213ye7l****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The array of migration object information.
	//
	// > The array is returned in the following format: [{"key":"Function","state":5,"totalCount":22},{"key":"Procedure","state":5,"totalCount":26},{"key":"Table","state":0,"totalCount":68},{"key":"View","state":5,"totalCount":100}].
	ProgressSummaryDetails []*SummaryJobDetailResponseBodyProgressSummaryDetails `json:"ProgressSummaryDetails,omitempty" xml:"ProgressSummaryDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9033138C-5AB3-5EB7-BA78-43131F19297C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SummaryJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SummaryJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *SummaryJobDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *SummaryJobDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SummaryJobDetailResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SummaryJobDetailResponseBody) GetProgressSummaryDetails() []*SummaryJobDetailResponseBodyProgressSummaryDetails {
	return s.ProgressSummaryDetails
}

func (s *SummaryJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SummaryJobDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SummaryJobDetailResponseBody) SetCode(v string) *SummaryJobDetailResponseBody {
	s.Code = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetHttpStatusCode(v int32) *SummaryJobDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetJobId(v string) *SummaryJobDetailResponseBody {
	s.JobId = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetProgressSummaryDetails(v []*SummaryJobDetailResponseBodyProgressSummaryDetails) *SummaryJobDetailResponseBody {
	s.ProgressSummaryDetails = v
	return s
}

func (s *SummaryJobDetailResponseBody) SetRequestId(v string) *SummaryJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetSuccess(v bool) *SummaryJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *SummaryJobDetailResponseBody) Validate() error {
	if s.ProgressSummaryDetails != nil {
		for _, item := range s.ProgressSummaryDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SummaryJobDetailResponseBodyProgressSummaryDetails struct {
	// The object type of the migration object. Valid values: **Table**, **Constraint**, **Index**, **View**, **Materialize View**, **Type*	- (user-defined type), **Synonym**, **Trigger**, **Function**, **Procedure*	- (stored procedure), **Package**, **Default**, **Rule**, **PlanGuide*	- (execute plan), and **Sequence**.
	//
	// example:
	//
	// Table
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The migration status. Valid values:
	//
	// - **0**: finish (completed).
	//
	// - **1**: catched (waiting for synchronization).
	//
	// - **2**: init (initializing).
	//
	// - **3**: running (synchronizing).
	//
	// - **4**: warning (error).
	//
	// - **5**: failed (failed).
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// The total number of migration objects.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SummaryJobDetailResponseBodyProgressSummaryDetails) String() string {
	return dara.Prettify(s)
}

func (s SummaryJobDetailResponseBodyProgressSummaryDetails) GoString() string {
	return s.String()
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) GetKey() *string {
	return s.Key
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) GetState() *int32 {
	return s.State
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) SetKey(v string) *SummaryJobDetailResponseBodyProgressSummaryDetails {
	s.Key = &v
	return s
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) SetState(v int32) *SummaryJobDetailResponseBodyProgressSummaryDetails {
	s.State = &v
	return s
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) SetTotalCount(v int64) *SummaryJobDetailResponseBodyProgressSummaryDetails {
	s.TotalCount = &v
	return s
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) Validate() error {
	return dara.Validate(s)
}
