// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserOwnedResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListUserOwnedResourcesResponseBodyData) *ListUserOwnedResourcesResponseBody
	GetData() *ListUserOwnedResourcesResponseBodyData
	SetErrorCode(v string) *ListUserOwnedResourcesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListUserOwnedResourcesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListUserOwnedResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserOwnedResourcesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListUserOwnedResourcesResponseBody
	GetTotalCount() *int64
}

type ListUserOwnedResourcesResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	Data *ListUserOwnedResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 8E88933E-E3D4-5BA8-8CBF-0A1CAE666690
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned. By default, this parameter is not returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserOwnedResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserOwnedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserOwnedResourcesResponseBody) GetData() *ListUserOwnedResourcesResponseBodyData {
	return s.Data
}

func (s *ListUserOwnedResourcesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUserOwnedResourcesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUserOwnedResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserOwnedResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserOwnedResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserOwnedResourcesResponseBody) SetData(v *ListUserOwnedResourcesResponseBodyData) *ListUserOwnedResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserOwnedResourcesResponseBody) SetErrorCode(v string) *ListUserOwnedResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBody) SetErrorMessage(v string) *ListUserOwnedResourcesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBody) SetRequestId(v string) *ListUserOwnedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBody) SetSuccess(v bool) *ListUserOwnedResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBody) SetTotalCount(v int64) *ListUserOwnedResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserOwnedResourcesResponseBodyData struct {
	ResourceList []*ListUserOwnedResourcesResponseBodyDataResourceList `json:"resourceList,omitempty" xml:"resourceList,omitempty" type:"Repeated"`
}

func (s ListUserOwnedResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserOwnedResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserOwnedResourcesResponseBodyData) GetResourceList() []*ListUserOwnedResourcesResponseBodyDataResourceList {
	return s.ResourceList
}

func (s *ListUserOwnedResourcesResponseBodyData) SetResourceList(v []*ListUserOwnedResourcesResponseBodyDataResourceList) *ListUserOwnedResourcesResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyData) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserOwnedResourcesResponseBodyDataResourceList struct {
	// The alias of the instance.
	//
	// example:
	//
	// DMS_GYX_TESTupdata
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The ID of the database in DMS.
	//
	// example:
	//
	// 29697059
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The ID of the instance to which the database belongs.
	//
	// example:
	//
	// 876XXX
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The database engine type. For more information about the valid values of the DbType parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database instance belongs.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The endpoint of the instance to which the database belongs.
	//
	// example:
	//
	// rm-wz98bw60x1i1303c5.mysql.rds.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 291594
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database
	//
	// 	- **false**: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The connection port of the instance to which the database belongs.
	//
	// example:
	//
	// 6379
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// as_task_engine
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The query name of the database.
	//
	// example:
	//
	// dtstest [dtstest_dms]
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The table ID.
	//
	// example:
	//
	// 1760934555
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The table name.
	//
	// example:
	//
	// addepmap
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListUserOwnedResourcesResponseBodyDataResourceList) String() string {
	return dara.Prettify(s)
}

func (s ListUserOwnedResourcesResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetAlias() *string {
	return s.Alias
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetDbId() *string {
	return s.DbId
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetDbType() *string {
	return s.DbType
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetEnvType() *string {
	return s.EnvType
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetHost() *string {
	return s.Host
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetLogic() *bool {
	return s.Logic
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetPort() *int64 {
	return s.Port
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetSearchName() *string {
	return s.SearchName
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetTableId() *string {
	return s.TableId
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) GetTableName() *string {
	return s.TableName
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetAlias(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.Alias = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetDbId(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.DbId = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetDbInstanceId(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.DbInstanceId = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetDbType(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.DbType = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetEnvType(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.EnvType = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetHost(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.Host = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetInstanceId(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.InstanceId = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetLogic(v bool) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.Logic = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetPort(v int64) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.Port = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetSchemaName(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.SchemaName = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetSearchName(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.SearchName = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetTableId(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.TableId = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) SetTableName(v string) *ListUserOwnedResourcesResponseBodyDataResourceList {
	s.TableName = &v
	return s
}

func (s *ListUserOwnedResourcesResponseBodyDataResourceList) Validate() error {
	return dara.Validate(s)
}
