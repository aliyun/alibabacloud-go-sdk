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
	Data []*ListDataCenterDatabaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DatabaseDesc *string `json:"DatabaseDesc,omitempty" xml:"DatabaseDesc,omitempty"`
	// example:
	//
	// diamonds.csv
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// 6kv159u9vtpvl**********b8
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// csv
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	DescUpdateTime *string `json:"DescUpdateTime,omitempty" xml:"DescUpdateTime,omitempty"`
	// example:
	//
	// 73088962
	DmsDbId *int64 `json:"DmsDbId,omitempty" xml:"DmsDbId,omitempty"`
	// example:
	//
	// 2740966
	DmsInstanceId *int64 `json:"DmsInstanceId,omitempty" xml:"DmsInstanceId,omitempty"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// FILE
	ImportType *string `json:"ImportType,omitempty" xml:"ImportType,omitempty"`
	// example:
	//
	// f-ean8u5881qk4*********xh5y
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// N
	IsInternal *string `json:"IsInternal,omitempty" xml:"IsInternal,omitempty"`
	// example:
	//
	// 999
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
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

func (s *ListDataCenterDatabaseResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListDataCenterDatabaseResponseBodyData) GetImportType() *string {
	return s.ImportType
}

func (s *ListDataCenterDatabaseResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListDataCenterDatabaseResponseBodyData) GetIsInternal() *string {
	return s.IsInternal
}

func (s *ListDataCenterDatabaseResponseBodyData) GetSize() *int64 {
	return s.Size
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

func (s *ListDataCenterDatabaseResponseBodyData) SetIsInternal(v string) *ListDataCenterDatabaseResponseBodyData {
	s.IsInternal = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) SetSize(v int64) *ListDataCenterDatabaseResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListDataCenterDatabaseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
