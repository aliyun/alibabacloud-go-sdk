// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetOpLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetOpLogResponseBody
	GetErrorMessage() *string
	SetOpLogDetails(v *GetOpLogResponseBodyOpLogDetails) *GetOpLogResponseBody
	GetOpLogDetails() *GetOpLogResponseBodyOpLogDetails
	SetRequestId(v string) *GetOpLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOpLogResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetOpLogResponseBody
	GetTotalCount() *int64
}

type GetOpLogResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details of the operation log.
	OpLogDetails *GetOpLogResponseBodyOpLogDetails `json:"OpLogDetails,omitempty" xml:"OpLogDetails,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 47D56208-DB1D-4FD3-BE32-300E43185488
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
	// The total number of operation logs that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetOpLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOpLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOpLogResponseBody) GetOpLogDetails() *GetOpLogResponseBodyOpLogDetails {
	return s.OpLogDetails
}

func (s *GetOpLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetOpLogResponseBody) SetErrorCode(v string) *GetOpLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOpLogResponseBody) SetErrorMessage(v string) *GetOpLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOpLogResponseBody) SetOpLogDetails(v *GetOpLogResponseBodyOpLogDetails) *GetOpLogResponseBody {
	s.OpLogDetails = v
	return s
}

func (s *GetOpLogResponseBody) SetRequestId(v string) *GetOpLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpLogResponseBody) SetSuccess(v bool) *GetOpLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpLogResponseBody) SetTotalCount(v int64) *GetOpLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetOpLogResponseBody) Validate() error {
	if s.OpLogDetails != nil {
		if err := s.OpLogDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOpLogResponseBodyOpLogDetails struct {
	OpLogDetail []*GetOpLogResponseBodyOpLogDetailsOpLogDetail `json:"OpLogDetail,omitempty" xml:"OpLogDetail,omitempty" type:"Repeated"`
}

func (s GetOpLogResponseBodyOpLogDetails) String() string {
	return dara.Prettify(s)
}

func (s GetOpLogResponseBodyOpLogDetails) GoString() string {
	return s.String()
}

func (s *GetOpLogResponseBodyOpLogDetails) GetOpLogDetail() []*GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	return s.OpLogDetail
}

func (s *GetOpLogResponseBodyOpLogDetails) SetOpLogDetail(v []*GetOpLogResponseBodyOpLogDetailsOpLogDetail) *GetOpLogResponseBodyOpLogDetails {
	s.OpLogDetail = v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetails) Validate() error {
	if s.OpLogDetail != nil {
		for _, item := range s.OpLogDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOpLogResponseBodyOpLogDetailsOpLogDetail struct {
	// The endpoint of the database instance.
	//
	// >
	//
	// 	- This parameter is valid only for database instances of the LocalInstance type.
	//
	// 	- This parameter is valid only for operations on the functional modules related to tasks.
	//
	// example:
	//
	// dmstest_prod_database@dmstest.rds... [Test instance]
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The functional module for which the operation log is queried.
	//
	// example:
	//
	// SECURITY_RULE
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The details of the operation.
	//
	// example:
	//
	// SELECT 	- FROM `orders` \\n LIMIT 20
	OpContent *string `json:"OpContent,omitempty" xml:"OpContent,omitempty"`
	// The time when the operation was performed.
	//
	// example:
	//
	// 2022-03-28 16:45:19
	OpTime *string `json:"OpTime,omitempty" xml:"OpTime,omitempty"`
	// The ID of the user who performed the operation.
	//
	// example:
	//
	// 51****
	OpUserId *int64 `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID of the ticket or task.
	//
	// >  This parameter is valid only for operations on the functional modules related to tasks and the task management module in system management.
	//
	// example:
	//
	// 509****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 22275482072787****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// test_name
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s GetOpLogResponseBodyOpLogDetailsOpLogDetail) String() string {
	return dara.Prettify(s)
}

func (s GetOpLogResponseBodyOpLogDetailsOpLogDetail) GoString() string {
	return s.String()
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) GetDatabase() *string {
	return s.Database
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) GetModule() *string {
	return s.Module
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) GetOpContent() *string {
	return s.OpContent
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) GetOpTime() *string {
	return s.OpTime
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) GetOpUserId() *int64 {
	return s.OpUserId
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) GetUserId() *string {
	return s.UserId
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) GetUserNick() *string {
	return s.UserNick
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetDatabase(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.Database = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetModule(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.Module = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetOpContent(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.OpContent = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetOpTime(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.OpTime = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetOpUserId(v int64) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.OpUserId = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetOrderId(v int64) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.OrderId = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetUserId(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.UserId = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetUserNick(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.UserNick = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) Validate() error {
	return dara.Validate(s)
}
