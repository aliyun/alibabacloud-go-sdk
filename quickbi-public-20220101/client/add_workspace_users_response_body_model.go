// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddWorkspaceUsersResponseBody
	GetRequestId() *string
	SetResult(v *AddWorkspaceUsersResponseBodyResult) *AddWorkspaceUsersResponseBody
	GetResult() *AddWorkspaceUsersResponseBodyResult
	SetSuccess(v bool) *AddWorkspaceUsersResponseBody
	GetSuccess() *bool
}

type AddWorkspaceUsersResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 7AAB95D7-2E11-4FE2-94BC-858E4FC0C976
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the API execution.
	Result *AddWorkspaceUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Value range:
	//
	// - true: The request was successful. There may be cases where some members are added successfully and others fail. For the reasons of failure, refer to the FailureDetail in the response.
	//
	// - false: The request failed, and no data will be persisted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddWorkspaceUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWorkspaceUsersResponseBody) GetResult() *AddWorkspaceUsersResponseBodyResult {
	return s.Result
}

func (s *AddWorkspaceUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddWorkspaceUsersResponseBody) SetRequestId(v string) *AddWorkspaceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkspaceUsersResponseBody) SetResult(v *AddWorkspaceUsersResponseBodyResult) *AddWorkspaceUsersResponseBody {
	s.Result = v
	return s
}

func (s *AddWorkspaceUsersResponseBody) SetSuccess(v bool) *AddWorkspaceUsersResponseBody {
	s.Success = &v
	return s
}

func (s *AddWorkspaceUsersResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddWorkspaceUsersResponseBodyResult struct {
	// Number of users that failed to be added.
	//
	// example:
	//
	// 2
	Failure *int32 `json:"Failure,omitempty" xml:"Failure,omitempty"`
	// Reasons for the failures.
	//
	// example:
	//
	// {"2046274934845893" : "AE0150010001: This user already exists.", "1213444447906552" : "AE0150010001: This user already exists."}
	FailureDetail map[string]interface{} `json:"FailureDetail,omitempty" xml:"FailureDetail,omitempty"`
	// Number of users that were added successfully.
	//
	// example:
	//
	// 1
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of users being added.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s AddWorkspaceUsersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUsersResponseBodyResult) GetFailure() *int32 {
	return s.Failure
}

func (s *AddWorkspaceUsersResponseBodyResult) GetFailureDetail() map[string]interface{} {
	return s.FailureDetail
}

func (s *AddWorkspaceUsersResponseBodyResult) GetSuccess() *int32 {
	return s.Success
}

func (s *AddWorkspaceUsersResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *AddWorkspaceUsersResponseBodyResult) SetFailure(v int32) *AddWorkspaceUsersResponseBodyResult {
	s.Failure = &v
	return s
}

func (s *AddWorkspaceUsersResponseBodyResult) SetFailureDetail(v map[string]interface{}) *AddWorkspaceUsersResponseBodyResult {
	s.FailureDetail = v
	return s
}

func (s *AddWorkspaceUsersResponseBodyResult) SetSuccess(v int32) *AddWorkspaceUsersResponseBodyResult {
	s.Success = &v
	return s
}

func (s *AddWorkspaceUsersResponseBodyResult) SetTotal(v int32) *AddWorkspaceUsersResponseBodyResult {
	s.Total = &v
	return s
}

func (s *AddWorkspaceUsersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
