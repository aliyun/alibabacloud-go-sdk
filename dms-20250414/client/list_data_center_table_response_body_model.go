// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCenterTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataCenterTableResponseBodyData) *ListDataCenterTableResponseBody
	GetData() *ListDataCenterTableResponseBodyData
	SetErrorCode(v string) *ListDataCenterTableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataCenterTableResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDataCenterTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCenterTableResponseBody
	GetSuccess() *bool
}

type ListDataCenterTableResponseBody struct {
	// The data returned in the response.
	Data *ListDataCenterTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The code that indicates the result of the request. If the request fails, an error code is returned.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataCenterTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterTableResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCenterTableResponseBody) GetData() *ListDataCenterTableResponseBodyData {
	return s.Data
}

func (s *ListDataCenterTableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataCenterTableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataCenterTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCenterTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCenterTableResponseBody) SetData(v *ListDataCenterTableResponseBodyData) *ListDataCenterTableResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCenterTableResponseBody) SetErrorCode(v string) *ListDataCenterTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataCenterTableResponseBody) SetErrorMessage(v string) *ListDataCenterTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataCenterTableResponseBody) SetRequestId(v string) *ListDataCenterTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCenterTableResponseBody) SetSuccess(v bool) *ListDataCenterTableResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCenterTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataCenterTableResponseBodyData struct {
	// A list of the data tables.
	Content []*ListDataCenterTableResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDataCenterTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCenterTableResponseBodyData) GetContent() []*ListDataCenterTableResponseBodyDataContent {
	return s.Content
}

func (s *ListDataCenterTableResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataCenterTableResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataCenterTableResponseBodyData) GetTotalElements() *int64 {
	return s.TotalElements
}

func (s *ListDataCenterTableResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *ListDataCenterTableResponseBodyData) SetContent(v []*ListDataCenterTableResponseBodyDataContent) *ListDataCenterTableResponseBodyData {
	s.Content = v
	return s
}

func (s *ListDataCenterTableResponseBodyData) SetPageNumber(v int64) *ListDataCenterTableResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataCenterTableResponseBodyData) SetPageSize(v int64) *ListDataCenterTableResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataCenterTableResponseBodyData) SetTotalElements(v int64) *ListDataCenterTableResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDataCenterTableResponseBodyData) SetTotalPages(v int64) *ListDataCenterTableResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDataCenterTableResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataCenterTableResponseBodyDataContent struct {
	// The database name.
	//
	// - If `ImportType` is `FILE`, this parameter returns the file name.
	//
	// example:
	//
	// diamonds.csv
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The time when the data table description was last updated.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	DescUpdateTime *string `json:"DescUpdateTime,omitempty" xml:"DescUpdateTime,omitempty"`
	// The ID of the DMS database.
	//
	// - This parameter is not returned if `ImportType` is `FILE`.
	//
	// example:
	//
	// 69950353
	DmsDbId *int64 `json:"DmsDbId,omitempty" xml:"DmsDbId,omitempty"`
	// The ID of the DMS instance that hosts the database.
	//
	// - This parameter is not returned if `ImportType` is `FILE`.
	//
	// example:
	//
	// 2310246
	DmsInstanceId *int64 `json:"DmsInstanceId,omitempty" xml:"DmsInstanceId,omitempty"`
	// The time when the data table was created.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The import type. Valid value:
	//
	// - **FILE**: The data is imported from a file.
	//
	// example:
	//
	// FILE
	ImportType *string `json:"ImportType,omitempty" xml:"ImportType,omitempty"`
	// The instance name.
	//
	// - If `ImportType` is `FILE`, this parameter returns the file ID.
	//
	// example:
	//
	// f-ean8u5881qk4*********xh5y
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The description of the data table.
	//
	// example:
	//
	// 这是一份钻石数据。
	TableDesc *string `json:"TableDesc,omitempty" xml:"TableDesc,omitempty"`
	// The ID of the data table.
	//
	// example:
	//
	// xa8wib4ga3a2*********fjbx
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The table name.
	//
	// example:
	//
	// diamonds
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDataCenterTableResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterTableResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListDataCenterTableResponseBodyDataContent) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDataCenterTableResponseBodyDataContent) GetDescUpdateTime() *string {
	return s.DescUpdateTime
}

func (s *ListDataCenterTableResponseBodyDataContent) GetDmsDbId() *int64 {
	return s.DmsDbId
}

func (s *ListDataCenterTableResponseBodyDataContent) GetDmsInstanceId() *int64 {
	return s.DmsInstanceId
}

func (s *ListDataCenterTableResponseBodyDataContent) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListDataCenterTableResponseBodyDataContent) GetImportType() *string {
	return s.ImportType
}

func (s *ListDataCenterTableResponseBodyDataContent) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListDataCenterTableResponseBodyDataContent) GetTableDesc() *string {
	return s.TableDesc
}

func (s *ListDataCenterTableResponseBodyDataContent) GetTableId() *string {
	return s.TableId
}

func (s *ListDataCenterTableResponseBodyDataContent) GetTableName() *string {
	return s.TableName
}

func (s *ListDataCenterTableResponseBodyDataContent) SetDatabaseName(v string) *ListDataCenterTableResponseBodyDataContent {
	s.DatabaseName = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetDescUpdateTime(v string) *ListDataCenterTableResponseBodyDataContent {
	s.DescUpdateTime = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetDmsDbId(v int64) *ListDataCenterTableResponseBodyDataContent {
	s.DmsDbId = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetDmsInstanceId(v int64) *ListDataCenterTableResponseBodyDataContent {
	s.DmsInstanceId = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetGmtCreated(v string) *ListDataCenterTableResponseBodyDataContent {
	s.GmtCreated = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetImportType(v string) *ListDataCenterTableResponseBodyDataContent {
	s.ImportType = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetInstanceName(v string) *ListDataCenterTableResponseBodyDataContent {
	s.InstanceName = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetTableDesc(v string) *ListDataCenterTableResponseBodyDataContent {
	s.TableDesc = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetTableId(v string) *ListDataCenterTableResponseBodyDataContent {
	s.TableId = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) SetTableName(v string) *ListDataCenterTableResponseBodyDataContent {
	s.TableName = &v
	return s
}

func (s *ListDataCenterTableResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
