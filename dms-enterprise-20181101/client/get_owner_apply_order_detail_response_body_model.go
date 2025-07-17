// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOwnerApplyOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetOwnerApplyOrderDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetOwnerApplyOrderDetailResponseBody
	GetErrorMessage() *string
	SetOwnerApplyOrderDetail(v *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) *GetOwnerApplyOrderDetailResponseBody
	GetOwnerApplyOrderDetail() *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail
	SetRequestId(v string) *GetOwnerApplyOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOwnerApplyOrderDetailResponseBody
	GetSuccess() *bool
}

type GetOwnerApplyOrderDetailResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details of the ticket.
	OwnerApplyOrderDetail *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail `json:"OwnerApplyOrderDetail,omitempty" xml:"OwnerApplyOrderDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CB784055-E8CB-4461-AB0B-483A1DA32BB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOwnerApplyOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOwnerApplyOrderDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOwnerApplyOrderDetailResponseBody) GetOwnerApplyOrderDetail() *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail {
	return s.OwnerApplyOrderDetail
}

func (s *GetOwnerApplyOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOwnerApplyOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetErrorCode(v string) *GetOwnerApplyOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetErrorMessage(v string) *GetOwnerApplyOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetOwnerApplyOrderDetail(v *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) *GetOwnerApplyOrderDetailResponseBody {
	s.OwnerApplyOrderDetail = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetRequestId(v string) *GetOwnerApplyOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetSuccess(v bool) *GetOwnerApplyOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail struct {
	// The type of the submitted ticket. Valid values:
	//
	// 	- **INSTANCE**: the ticket that applies for the permissions to be an instance owner
	//
	// 	- **DB**: the ticket that applies for the permissions to be a database owner
	//
	// 	- **TABLE**: the ticket that applies for the permissions to be a table owner
	//
	// example:
	//
	// DB
	ApplyType *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	// The details of the requested resource.
	Resources []*GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) GetApplyType() *string {
	return s.ApplyType
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) GetResources() []*GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources {
	return s.Resources
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) SetApplyType(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail {
	s.ApplyType = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) SetResources(v []*GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail {
	s.Resources = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) Validate() error {
	return dara.Validate(s)
}

type GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources struct {
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The instance is a logical database.
	//
	// 	- **false**: The instance is not a logical database.
	//
	// example:
	//
	// true
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The details of the resource.
	ResourceDetail *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail `json:"ResourceDetail,omitempty" xml:"ResourceDetail,omitempty" type:"Struct"`
	// The ID of the resource.
	//
	// example:
	//
	// 12345
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) String() string {
	return dara.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) GetLogic() *bool {
	return s.Logic
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) GetResourceDetail() *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	return s.ResourceDetail
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) GetTargetId() *string {
	return s.TargetId
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) SetLogic(v bool) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources {
	s.Logic = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) SetResourceDetail(v *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources {
	s.ResourceDetail = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) SetTargetId(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources {
	s.TargetId = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) Validate() error {
	return dara.Validate(s)
}

type GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail struct {
	// The type of the database engine.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the instance belongs. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// PRODUCT
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The IDs of the original owners.
	OwnerIds []*int64 `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	// The nicknames of the owners.
	OwnerNickNames []*string `json:"OwnerNickNames,omitempty" xml:"OwnerNickNames,omitempty" type:"Repeated"`
	// The search name of the resource.
	//
	// example:
	//
	// yuyang【test】
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The name of the table.
	//
	// > : This parameter is returned when you submit a Database-OWNER ticket.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) String() string {
	return dara.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) GetDbType() *string {
	return s.DbType
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) GetEnvType() *string {
	return s.EnvType
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) GetOwnerIds() []*int64 {
	return s.OwnerIds
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) GetOwnerNickNames() []*string {
	return s.OwnerNickNames
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) GetSearchName() *string {
	return s.SearchName
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) GetTableName() *string {
	return s.TableName
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetDbType(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.DbType = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetEnvType(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.EnvType = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetOwnerIds(v []*int64) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.OwnerIds = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetOwnerNickNames(v []*string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.OwnerNickNames = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetSearchName(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.SearchName = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetTableName(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.TableName = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) Validate() error {
	return dara.Validate(s)
}
