// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUsersRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkspaceUsersRoleResponseBody
	GetRequestId() *string
	SetResult(v *UpdateWorkspaceUsersRoleResponseBodyResult) *UpdateWorkspaceUsersRoleResponseBody
	GetResult() *UpdateWorkspaceUsersRoleResponseBodyResult
	SetSuccess(v bool) *UpdateWorkspaceUsersRoleResponseBody
	GetSuccess() *bool
}

type UpdateWorkspaceUsersRoleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 7AAB95D7-2E11-4FE2-94BC-858E4FC0C976
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the interface execution.
	Result *UpdateWorkspaceUsersRoleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Value range:
	//
	// - true: The request was successful, some members may have been updated successfully while others failed, refer to FailureDetail in the response for reasons of failure
	//
	// - false: The request failed, no data will be persisted
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkspaceUsersRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUsersRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUsersRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceUsersRoleResponseBody) GetResult() *UpdateWorkspaceUsersRoleResponseBodyResult {
	return s.Result
}

func (s *UpdateWorkspaceUsersRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkspaceUsersRoleResponseBody) SetRequestId(v string) *UpdateWorkspaceUsersRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBody) SetResult(v *UpdateWorkspaceUsersRoleResponseBodyResult) *UpdateWorkspaceUsersRoleResponseBody {
	s.Result = v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBody) SetSuccess(v bool) *UpdateWorkspaceUsersRoleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceUsersRoleResponseBodyResult struct {
	// Number of users that failed to update.
	//
	// example:
	//
	// 0
	Failure *int32 `json:"Failure,omitempty" xml:"Failure,omitempty"`
	// Reasons for the update failures.
	FailureDetail map[string]interface{} `json:"FailureDetail,omitempty" xml:"FailureDetail,omitempty"`
	// Number of users that were updated successfully.
	//
	// example:
	//
	// 2
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// Modify the total number of users.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s UpdateWorkspaceUsersRoleResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUsersRoleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) GetFailure() *int32 {
	return s.Failure
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) GetFailureDetail() map[string]interface{} {
	return s.FailureDetail
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) GetSuccess() *int32 {
	return s.Success
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) SetFailure(v int32) *UpdateWorkspaceUsersRoleResponseBodyResult {
	s.Failure = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) SetFailureDetail(v map[string]interface{}) *UpdateWorkspaceUsersRoleResponseBodyResult {
	s.FailureDetail = v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) SetSuccess(v int32) *UpdateWorkspaceUsersRoleResponseBodyResult {
	s.Success = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) SetTotal(v int32) *UpdateWorkspaceUsersRoleResponseBodyResult {
	s.Total = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
