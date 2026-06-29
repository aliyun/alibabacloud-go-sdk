// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskWorkforceStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskWorkforceStatisticResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTaskWorkforceStatisticResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTaskWorkforceStatisticResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTaskWorkforceStatisticResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetTaskWorkforceStatisticResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetTaskWorkforceStatisticResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetTaskWorkforceStatisticResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskWorkforceStatisticResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetTaskWorkforceStatisticResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *GetTaskWorkforceStatisticResponseBody
	GetTotalPage() *int32
	SetUsersStatistic(v []*UserStatistic) *GetTaskWorkforceStatisticResponseBody
	GetUsersStatistic() []*UserStatistic
}

type GetTaskWorkforceStatisticResponseBody struct {
	// Return encoding. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// - When Success is false, returns a business error code.
	//
	// - When Success is true, returns an empty value.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Return message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number of the returned list of queried job members.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of job members displayed per page in the response.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of job members.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// List of user statistics.
	UsersStatistic []*UserStatistic `json:"UsersStatistic,omitempty" xml:"UsersStatistic,omitempty" type:"Repeated"`
}

func (s GetTaskWorkforceStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskWorkforceStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceStatisticResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskWorkforceStatisticResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTaskWorkforceStatisticResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskWorkforceStatisticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskWorkforceStatisticResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetTaskWorkforceStatisticResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTaskWorkforceStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskWorkforceStatisticResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskWorkforceStatisticResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetTaskWorkforceStatisticResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *GetTaskWorkforceStatisticResponseBody) GetUsersStatistic() []*UserStatistic {
	return s.UsersStatistic
}

func (s *GetTaskWorkforceStatisticResponseBody) SetCode(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetDetails(v string) *GetTaskWorkforceStatisticResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetErrorCode(v string) *GetTaskWorkforceStatisticResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetMessage(v string) *GetTaskWorkforceStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetPageNumber(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetPageSize(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetRequestId(v string) *GetTaskWorkforceStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetSuccess(v bool) *GetTaskWorkforceStatisticResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetTotalCount(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetTotalPage(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.TotalPage = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetUsersStatistic(v []*UserStatistic) *GetTaskWorkforceStatisticResponseBody {
	s.UsersStatistic = v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) Validate() error {
	if s.UsersStatistic != nil {
		for _, item := range s.UsersStatistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
