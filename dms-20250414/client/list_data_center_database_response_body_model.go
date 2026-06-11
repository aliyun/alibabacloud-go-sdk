// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCenterDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataCenterDatabaseResponseBodyData) *ListDataCenterDatabaseResponseBody
	GetData() []*ListDataCenterDatabaseResponseBodyData
	SetErrorCode(v string) *ListDataCenterDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataCenterDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDataCenterDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCenterDatabaseResponseBody
	GetSuccess() *bool
}

type ListDataCenterDatabaseResponseBody struct {
	// The list of databases.
	Data []*ListDataCenterDatabaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request fails.
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
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataCenterDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCenterDatabaseResponseBody) GetData() []*ListDataCenterDatabaseResponseBodyData {
	return s.Data
}

func (s *ListDataCenterDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataCenterDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataCenterDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCenterDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCenterDatabaseResponseBody) SetData(v []*ListDataCenterDatabaseResponseBodyData) *ListDataCenterDatabaseResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCenterDatabaseResponseBody) SetErrorCode(v string) *ListDataCenterDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBody) SetErrorMessage(v string) *ListDataCenterDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBody) SetRequestId(v string) *ListDataCenterDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBody) SetSuccess(v bool) *ListDataCenterDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataCenterDatabaseResponseBodyData struct {
	// The description of the database.
	//
	// example:
	//
	// This is a sample database.
	DatabaseDesc *string `json:"DatabaseDesc,omitempty" xml:"DatabaseDesc,omitempty"`
	// The name of the database.
	//
	// - If `ImportType` is `FILE`, this is the file name.
	//
	// example:
	//
	// diamonds.csv
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 6kv159u9vtpvl**********b8
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// - If `ImportType` is `FILE`:
	//
	//   - The file format, such as `csv`, `xlsx`, or `xls`.
	//
	// example:
	//
	// csv
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The time the database description was last updated.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	DescUpdateTime *string `json:"DescUpdateTime,omitempty" xml:"DescUpdateTime,omitempty"`
	// The ID of the database in DMS.
	//
	// - This parameter is not returned if `ImportType` is `FILE`.
	//
	// example:
	//
	// 73088962
	DmsDbId *int64 `json:"DmsDbId,omitempty" xml:"DmsDbId,omitempty"`
	// The ID of the DMS instance that manages the database.
	//
	// - This parameter is not returned if `ImportType` is `FILE`.
	//
	// example:
	//
	// 2740966
	DmsInstanceId *int64  `json:"DmsInstanceId,omitempty" xml:"DmsInstanceId,omitempty"`
	DownloadLink  *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The time the entry was created.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The import type. Valid values:
	//
	// - FILE
	//
	// - RDS
	//
	// - ADB
	//
	// - PolarDB
	//
	// - Hologres
	//
	// - DMS
	//
	// example:
	//
	// FILE
	ImportType *string `json:"ImportType,omitempty" xml:"ImportType,omitempty"`
	// The name of the instance.
	//
	// - If `ImportType` is `FILE`, this parameter specifies the file ID in the data center.
	//
	// example:
	//
	// f-ean8u5881qk4*********xh5y
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	// Indicates whether the dataset is built-in. Valid values:
	//
	// - Y: The dataset is built-in.
	//
	// - N: The dataset is not built-in.
	//
	// example:
	//
	// N
	IsInternal *string `json:"IsInternal,omitempty" xml:"IsInternal,omitempty"`
	OssBucket  *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The size of the file, in bytes.
	//
	// example:
	//
	// 999
	Size             *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	UseUserOssBucket *bool  `json:"UseUserOssBucket,omitempty" xml:"UseUserOssBucket,omitempty"`
}

func (s ListDataCenterDatabaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterDatabaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCenterDatabaseResponseBodyData) GetDatabaseDesc() *string {
	return s.DatabaseDesc
}

func (s *ListDataCenterDatabaseResponseBodyData) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDataCenterDatabaseResponseBodyData) GetDbId() *string {
	return s.DbId
}

func (s *ListDataCenterDatabaseResponseBodyData) GetDbType() *string {
	return s.DbType
}

func (s *ListDataCenterDatabaseResponseBodyData) GetDescUpdateTime() *string {
	return s.DescUpdateTime
}

func (s *ListDataCenterDatabaseResponseBodyData) GetDmsDbId() *int64 {
	return s.DmsDbId
}

func (s *ListDataCenterDatabaseResponseBodyData) GetDmsInstanceId() *int64 {
	return s.DmsInstanceId
}

func (s *ListDataCenterDatabaseResponseBodyData) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *ListDataCenterDatabaseResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListDataCenterDatabaseResponseBodyData) GetImportType() *string {
	return s.ImportType
}

func (s *ListDataCenterDatabaseResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListDataCenterDatabaseResponseBodyData) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *ListDataCenterDatabaseResponseBodyData) GetIsInternal() *string {
	return s.IsInternal
}

func (s *ListDataCenterDatabaseResponseBodyData) GetOssBucket() *string {
	return s.OssBucket
}

func (s *ListDataCenterDatabaseResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *ListDataCenterDatabaseResponseBodyData) GetUseUserOssBucket() *bool {
	return s.UseUserOssBucket
}

func (s *ListDataCenterDatabaseResponseBodyData) SetDatabaseDesc(v string) *ListDataCenterDatabaseResponseBodyData {
	s.DatabaseDesc = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetDatabaseName(v string) *ListDataCenterDatabaseResponseBodyData {
	s.DatabaseName = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetDbId(v string) *ListDataCenterDatabaseResponseBodyData {
	s.DbId = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetDbType(v string) *ListDataCenterDatabaseResponseBodyData {
	s.DbType = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetDescUpdateTime(v string) *ListDataCenterDatabaseResponseBodyData {
	s.DescUpdateTime = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetDmsDbId(v int64) *ListDataCenterDatabaseResponseBodyData {
	s.DmsDbId = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetDmsInstanceId(v int64) *ListDataCenterDatabaseResponseBodyData {
	s.DmsInstanceId = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetDownloadLink(v string) *ListDataCenterDatabaseResponseBodyData {
	s.DownloadLink = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetGmtCreated(v string) *ListDataCenterDatabaseResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetImportType(v string) *ListDataCenterDatabaseResponseBodyData {
	s.ImportType = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetInstanceName(v string) *ListDataCenterDatabaseResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetIntranetDownloadLink(v string) *ListDataCenterDatabaseResponseBodyData {
	s.IntranetDownloadLink = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetIsInternal(v string) *ListDataCenterDatabaseResponseBodyData {
	s.IsInternal = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetOssBucket(v string) *ListDataCenterDatabaseResponseBodyData {
	s.OssBucket = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetSize(v int64) *ListDataCenterDatabaseResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetUseUserOssBucket(v bool) *ListDataCenterDatabaseResponseBodyData {
	s.UseUserOssBucket = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
