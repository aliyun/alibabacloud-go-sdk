// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsLifecycleStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApsLifecycleStrategyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListApsLifecycleStrategyResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListApsLifecycleStrategyResponseBodyItems) *ListApsLifecycleStrategyResponseBody
	GetItems() []*ListApsLifecycleStrategyResponseBodyItems
	SetMessage(v string) *ListApsLifecycleStrategyResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListApsLifecycleStrategyResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApsLifecycleStrategyResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApsLifecycleStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApsLifecycleStrategyResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListApsLifecycleStrategyResponseBody
	GetTotalCount() *int64
}

type ListApsLifecycleStrategyResponseBody struct {
	// The HTTP status code or the error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The queried lifecycle management policies.
	Items []*ListApsLifecycleStrategyResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-******-9F06-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApsLifecycleStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApsLifecycleStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListApsLifecycleStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApsLifecycleStrategyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApsLifecycleStrategyResponseBody) GetItems() []*ListApsLifecycleStrategyResponseBodyItems {
	return s.Items
}

func (s *ListApsLifecycleStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApsLifecycleStrategyResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApsLifecycleStrategyResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApsLifecycleStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApsLifecycleStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApsLifecycleStrategyResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApsLifecycleStrategyResponseBody) SetCode(v string) *ListApsLifecycleStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) SetHttpStatusCode(v int32) *ListApsLifecycleStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) SetItems(v []*ListApsLifecycleStrategyResponseBodyItems) *ListApsLifecycleStrategyResponseBody {
	s.Items = v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) SetMessage(v string) *ListApsLifecycleStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) SetPageNumber(v int32) *ListApsLifecycleStrategyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) SetPageSize(v int32) *ListApsLifecycleStrategyResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) SetRequestId(v string) *ListApsLifecycleStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) SetSuccess(v bool) *ListApsLifecycleStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) SetTotalCount(v int64) *ListApsLifecycleStrategyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApsLifecycleStrategyResponseBodyItems struct {
	// The job ID.
	//
	// example:
	//
	// aps-******
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The time when the policy was created.
	//
	// example:
	//
	// 2021-06-30T02:44:27Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The time when the policy was modified.
	//
	// example:
	//
	// 2021-07-03T06:33:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The operation tables.
	OperationTables []*ListApsLifecycleStrategyResponseBodyItemsOperationTables `json:"OperationTables,omitempty" xml:"OperationTables,omitempty" type:"Repeated"`
	// The status of the lifecycle management policy. Valid values:
	//
	// 1.  on: enables the current policy.
	//
	// 2.  off: disables the current policy.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of databases that are managed during the lifecycle management.
	//
	// example:
	//
	// 5
	StrategyDatabases *int64 `json:"StrategyDatabases,omitempty" xml:"StrategyDatabases,omitempty"`
	// The description of the lifecycle management policy.
	//
	// example:
	//
	// test
	StrategyDesc *string `json:"StrategyDesc,omitempty" xml:"StrategyDesc,omitempty"`
	// The name of the lifecycle management policy.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The number of tables that are managed during the lifecycle management.
	//
	// example:
	//
	// 5
	StrategyTables *int64 `json:"StrategyTables,omitempty" xml:"StrategyTables,omitempty"`
	// The type of the lifecycle management policy.
	//
	// example:
	//
	// KEEP_BY_TIME
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// The value of the lifecycle management policy.
	//
	// example:
	//
	// 10
	StrategyValue *string `json:"StrategyValue,omitempty" xml:"StrategyValue,omitempty"`
}

func (s ListApsLifecycleStrategyResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListApsLifecycleStrategyResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetOperationTables() []*ListApsLifecycleStrategyResponseBodyItemsOperationTables {
	return s.OperationTables
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetStrategyDatabases() *int64 {
	return s.StrategyDatabases
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetStrategyDesc() *string {
	return s.StrategyDesc
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetStrategyTables() *int64 {
	return s.StrategyTables
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListApsLifecycleStrategyResponseBodyItems) GetStrategyValue() *string {
	return s.StrategyValue
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetApsJobId(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.ApsJobId = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetCreatedTime(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetDBClusterId(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetModifiedTime(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetOperationTables(v []*ListApsLifecycleStrategyResponseBodyItemsOperationTables) *ListApsLifecycleStrategyResponseBodyItems {
	s.OperationTables = v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetStatus(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetStrategyDatabases(v int64) *ListApsLifecycleStrategyResponseBodyItems {
	s.StrategyDatabases = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetStrategyDesc(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.StrategyDesc = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetStrategyName(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.StrategyName = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetStrategyTables(v int64) *ListApsLifecycleStrategyResponseBodyItems {
	s.StrategyTables = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetStrategyType(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.StrategyType = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) SetStrategyValue(v string) *ListApsLifecycleStrategyResponseBodyItems {
	s.StrategyValue = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItems) Validate() error {
	if s.OperationTables != nil {
		for _, item := range s.OperationTables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApsLifecycleStrategyResponseBodyItemsOperationTables struct {
	// The name of the database.
	//
	// example:
	//
	// test
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Indicates whether all tables in the database are selected.
	//
	// example:
	//
	// true
	ProcessAll *string `json:"ProcessAll,omitempty" xml:"ProcessAll,omitempty"`
	// The names of the tables.
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s ListApsLifecycleStrategyResponseBodyItemsOperationTables) String() string {
	return dara.Prettify(s)
}

func (s ListApsLifecycleStrategyResponseBodyItemsOperationTables) GoString() string {
	return s.String()
}

func (s *ListApsLifecycleStrategyResponseBodyItemsOperationTables) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListApsLifecycleStrategyResponseBodyItemsOperationTables) GetProcessAll() *string {
	return s.ProcessAll
}

func (s *ListApsLifecycleStrategyResponseBodyItemsOperationTables) GetTableNames() []*string {
	return s.TableNames
}

func (s *ListApsLifecycleStrategyResponseBodyItemsOperationTables) SetDatabaseName(v string) *ListApsLifecycleStrategyResponseBodyItemsOperationTables {
	s.DatabaseName = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItemsOperationTables) SetProcessAll(v string) *ListApsLifecycleStrategyResponseBodyItemsOperationTables {
	s.ProcessAll = &v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItemsOperationTables) SetTableNames(v []*string) *ListApsLifecycleStrategyResponseBodyItemsOperationTables {
	s.TableNames = v
	return s
}

func (s *ListApsLifecycleStrategyResponseBodyItemsOperationTables) Validate() error {
	return dara.Validate(s)
}
