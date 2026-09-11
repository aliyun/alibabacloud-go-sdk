// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckTableDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiffTableCount(v int64) *DescribeDataCheckTableDetailsResponseBody
	GetDiffTableCount() *int64
	SetDynamicCode(v string) *DescribeDataCheckTableDetailsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDataCheckTableDetailsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDataCheckTableDetailsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDataCheckTableDetailsResponseBody
	GetErrMessage() *string
	SetFailedTableCount(v int64) *DescribeDataCheckTableDetailsResponseBody
	GetFailedTableCount() *int64
	SetFinishedCount(v int64) *DescribeDataCheckTableDetailsResponseBody
	GetFinishedCount() *int64
	SetHttpStatusCode(v int32) *DescribeDataCheckTableDetailsResponseBody
	GetHttpStatusCode() *int32
	SetPageNumber(v int32) *DescribeDataCheckTableDetailsResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeDataCheckTableDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDataCheckTableDetailsResponseBody
	GetSuccess() *bool
	SetTableDetails(v []*DescribeDataCheckTableDetailsResponseBodyTableDetails) *DescribeDataCheckTableDetailsResponseBody
	GetTableDetails() []*DescribeDataCheckTableDetailsResponseBodyTableDetails
	SetTotalCount(v int64) *DescribeDataCheckTableDetailsResponseBody
	GetTotalCount() *int64
}

type DescribeDataCheckTableDetailsResponseBody struct {
	// The number of tables with data inconsistency.
	//
	// example:
	//
	// 1
	DiffTableCount *int64 `json:"DiffTableCount,omitempty" xml:"DiffTableCount,omitempty"`
	// The dynamic error code. This parameter will be deprecated.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message used to replace the **%s*	- variable in the **ErrMessage*	- response parameter.
	//
	// > For example, if **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **Type**, the request parameter **Type*	- is invalid.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The number of tables that failed the verification.
	//
	// example:
	//
	// 1
	FailedTableCount *int64 `json:"FailedTableCount,omitempty" xml:"FailedTableCount,omitempty"`
	// The total number of rows that have been verified.
	//
	// example:
	//
	// 7
	FinishedCount *int64 `json:"FinishedCount,omitempty" xml:"FinishedCount,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the data consistency verification results for tables.
	TableDetails []*DescribeDataCheckTableDetailsResponseBodyTableDetails `json:"TableDetails,omitempty" xml:"TableDetails,omitempty" type:"Repeated"`
	// The total number of tables to be verified.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataCheckTableDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckTableDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetDiffTableCount() *int64 {
	return s.DiffTableCount
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetFailedTableCount() *int64 {
	return s.FailedTableCount
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetFinishedCount() *int64 {
	return s.FinishedCount
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetTableDetails() []*DescribeDataCheckTableDetailsResponseBodyTableDetails {
	return s.TableDetails
}

func (s *DescribeDataCheckTableDetailsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetDiffTableCount(v int64) *DescribeDataCheckTableDetailsResponseBody {
	s.DiffTableCount = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetDynamicCode(v string) *DescribeDataCheckTableDetailsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetDynamicMessage(v string) *DescribeDataCheckTableDetailsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetErrCode(v string) *DescribeDataCheckTableDetailsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetErrMessage(v string) *DescribeDataCheckTableDetailsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetFailedTableCount(v int64) *DescribeDataCheckTableDetailsResponseBody {
	s.FailedTableCount = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetFinishedCount(v int64) *DescribeDataCheckTableDetailsResponseBody {
	s.FinishedCount = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetHttpStatusCode(v int32) *DescribeDataCheckTableDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetPageNumber(v int32) *DescribeDataCheckTableDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetRequestId(v string) *DescribeDataCheckTableDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetSuccess(v bool) *DescribeDataCheckTableDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetTableDetails(v []*DescribeDataCheckTableDetailsResponseBodyTableDetails) *DescribeDataCheckTableDetailsResponseBody {
	s.TableDetails = v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) SetTotalCount(v int64) *DescribeDataCheckTableDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBody) Validate() error {
	if s.TableDetails != nil {
		for _, item := range s.TableDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataCheckTableDetailsResponseBodyTableDetails struct {
	// The time when the verification was performed.
	//
	// example:
	//
	// 2023-01-18 11:26:59
	BootTime *string `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	// The number of rows with data inconsistency.
	//
	// example:
	//
	// 1
	DiffCount *int64 `json:"DiffCount,omitempty" xml:"DiffCount,omitempty"`
	// The error code returned when the task fails. Valid values:
	//
	// - **1**: the number of tables without primary key exceeds the limit.
	//
	// - **2**: the number of rows with data inconsistency exceeds 300.
	//
	// - **3**: the table to be queried does not exist.
	//
	// - **4**: the SQL statement used to query data contains a syntax error.
	//
	// example:
	//
	// 1
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The number of rows that have been verified in the table.
	//
	// example:
	//
	// 7
	FinishCount *int64 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The auto-increment primary key that identifies a verification result record.
	//
	// example:
	//
	// 167401241974****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the source database.
	//
	// example:
	//
	// testdb
	SourceDbName *string `json:"SourceDbName,omitempty" xml:"SourceDbName,omitempty"`
	// The name of the source table.
	//
	// example:
	//
	// student
	SourceTbName *string `json:"SourceTbName,omitempty" xml:"SourceTbName,omitempty"`
	// The status of the verification result. Valid values:
	//
	// - **0**: completed.
	//
	// - **2**: initializing.
	//
	// - **3**: running.
	//
	// - **5**: failed.
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the destination database.
	//
	// example:
	//
	// testdb
	TargetDbName *string `json:"TargetDbName,omitempty" xml:"TargetDbName,omitempty"`
	// The name of the destination table.
	//
	// example:
	//
	// person
	TargetTbName *string `json:"TargetTbName,omitempty" xml:"TargetTbName,omitempty"`
	// The total number of rows to be verified.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataCheckTableDetailsResponseBodyTableDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckTableDetailsResponseBodyTableDetails) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetBootTime() *string {
	return s.BootTime
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetDiffCount() *int64 {
	return s.DiffCount
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetFinishCount() *int64 {
	return s.FinishCount
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetSourceDbName() *string {
	return s.SourceDbName
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetSourceTbName() *string {
	return s.SourceTbName
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetTargetDbName() *string {
	return s.TargetDbName
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetTargetTbName() *string {
	return s.TargetTbName
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetBootTime(v string) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.BootTime = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetDiffCount(v int64) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.DiffCount = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetErrorCode(v int32) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetFinishCount(v int64) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.FinishCount = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetId(v int64) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.Id = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetSourceDbName(v string) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.SourceDbName = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetSourceTbName(v string) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.SourceTbName = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetStatus(v string) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.Status = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetTargetDbName(v string) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.TargetDbName = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetTargetTbName(v string) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.TargetTbName = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) SetTotalCount(v int64) *DescribeDataCheckTableDetailsResponseBodyTableDetails {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponseBodyTableDetails) Validate() error {
	return dara.Validate(s)
}
