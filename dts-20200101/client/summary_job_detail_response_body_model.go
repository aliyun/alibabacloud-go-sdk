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
	// The ID of the data migration or data synchronization task.
	//
	// example:
	//
	// l3m1213ye7l****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned information about the migrated or synchronized objects in arrays.
	//
	// >  The arrays are in the following format: [{"key":"Function","state":5,"totalCount":22},{"key":"Procedure","state":5,"totalCount":26},{"key":"Table","state":0,"totalCount":68},{"key":"View","state":5,"totalCount":100}].
	ProgressSummaryDetails []*SummaryJobDetailResponseBodyProgressSummaryDetails `json:"ProgressSummaryDetails,omitempty" xml:"ProgressSummaryDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9033138C-5AB3-5EB7-BA78-43131F19297C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
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
	// The type of migrated or synchronized object. Valid values: **Table**, **Constraint**, **Index**, **View**, **Materialize View**, **Type**, **Synonym**, **Trigger**, **Function**, **Procedure**, **Package**, **Default**, **Rule**, **PlanGuide**, and **Sequence**.
	//
	// example:
	//
	// Table
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The state of the data migration or data synchronization task. Valid values:
	//
	// 	- **0**: The task was complete.
	//
	// 	- **1**: The task was waiting to start.
	//
	// 	- **2**: The task was being initialized.
	//
	// 	- **3**: The task was in progress.
	//
	// 	- **4**: An error occurred.
	//
	// 	- **5**: The task failed.
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// The total number of migrated or synchronized objects.
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
