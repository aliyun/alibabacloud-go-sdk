// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescibeImportsFromDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescibeImportsFromDatabaseResponseBodyItems) *DescibeImportsFromDatabaseResponseBody
	GetItems() *DescibeImportsFromDatabaseResponseBodyItems
	SetPageNumber(v int32) *DescibeImportsFromDatabaseResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescibeImportsFromDatabaseResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescibeImportsFromDatabaseResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescibeImportsFromDatabaseResponseBody
	GetTotalRecordCount() *int32
}

type DescibeImportsFromDatabaseResponseBody struct {
	// The migration tasks.
	Items *DescibeImportsFromDatabaseResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B000AA91-393D-46F9-8D9B-098E28931A3A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescibeImportsFromDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescibeImportsFromDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DescibeImportsFromDatabaseResponseBody) GetItems() *DescibeImportsFromDatabaseResponseBodyItems {
	return s.Items
}

func (s *DescibeImportsFromDatabaseResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescibeImportsFromDatabaseResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescibeImportsFromDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescibeImportsFromDatabaseResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescibeImportsFromDatabaseResponseBody) SetItems(v *DescibeImportsFromDatabaseResponseBodyItems) *DescibeImportsFromDatabaseResponseBody {
	s.Items = v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBody) SetPageNumber(v int32) *DescibeImportsFromDatabaseResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBody) SetPageRecordCount(v int32) *DescibeImportsFromDatabaseResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBody) SetRequestId(v string) *DescibeImportsFromDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBody) SetTotalRecordCount(v int32) *DescibeImportsFromDatabaseResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescibeImportsFromDatabaseResponseBodyItems struct {
	ImportResultFromDB []*DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB `json:"ImportResultFromDB,omitempty" xml:"ImportResultFromDB,omitempty" type:"Repeated"`
}

func (s DescibeImportsFromDatabaseResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescibeImportsFromDatabaseResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescibeImportsFromDatabaseResponseBodyItems) GetImportResultFromDB() []*DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB {
	return s.ImportResultFromDB
}

func (s *DescibeImportsFromDatabaseResponseBodyItems) SetImportResultFromDB(v []*DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) *DescibeImportsFromDatabaseResponseBodyItems {
	s.ImportResultFromDB = v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBodyItems) Validate() error {
	if s.ImportResultFromDB != nil {
		for _, item := range s.ImportResultFromDB {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB struct {
	// The status of the migration task. Valid values:
	//
	// 	- **NotStart**: The migration task has not started.
	//
	// 	- **FullExporting**: The migration task is exporting full data.
	//
	// 	- **FullImporting**: The migration task is importing full data.
	//
	// 	- **Success**: The migration task is successful.
	//
	// 	- **Failed**: The migration task failed.
	//
	// 	- **Canceled**: The migration task is canceled.
	//
	// 	- **Canceling**: The migration task is being canceled.
	//
	// 	- **IncrementalWaiting**: The migration task is waiting to synchronize incremental data.
	//
	// 	- **IncrementalImporting**: The migration task is synchronizing incremental data.
	//
	// 	- **StopSyncing**: The migration task stops synchronizing data.
	//
	// example:
	//
	// NotStart
	ImportDataStatus *string `json:"ImportDataStatus,omitempty" xml:"ImportDataStatus,omitempty"`
	// The description of the migration task.
	//
	// example:
	//
	// Description
	ImportDataStatusDescription *string `json:"ImportDataStatusDescription,omitempty" xml:"ImportDataStatusDescription,omitempty"`
	// The type of the migration task. Valid values:
	//
	// 	- **Full**: full migration
	//
	// 	- **Incremental:**: incremental migration
	//
	// example:
	//
	// Full
	ImportDataType *string `json:"ImportDataType,omitempty" xml:"ImportDataType,omitempty"`
	// The ID of the migration task.
	//
	// example:
	//
	// 123
	ImportId *int32 `json:"ImportId,omitempty" xml:"ImportId,omitempty"`
	// The time when the migration task synchronized incremental data. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2011-06-11T15:00Z
	IncrementalImportingTime *string `json:"IncrementalImportingTime,omitempty" xml:"IncrementalImportingTime,omitempty"`
}

func (s DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) String() string {
	return dara.Prettify(s)
}

func (s DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) GoString() string {
	return s.String()
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) GetImportDataStatus() *string {
	return s.ImportDataStatus
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) GetImportDataStatusDescription() *string {
	return s.ImportDataStatusDescription
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) GetImportDataType() *string {
	return s.ImportDataType
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) GetImportId() *int32 {
	return s.ImportId
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) GetIncrementalImportingTime() *string {
	return s.IncrementalImportingTime
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) SetImportDataStatus(v string) *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB {
	s.ImportDataStatus = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) SetImportDataStatusDescription(v string) *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB {
	s.ImportDataStatusDescription = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) SetImportDataType(v string) *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB {
	s.ImportDataType = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) SetImportId(v int32) *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB {
	s.ImportId = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) SetIncrementalImportingTime(v string) *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB {
	s.IncrementalImportingTime = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponseBodyItemsImportResultFromDB) Validate() error {
	return dara.Validate(s)
}
