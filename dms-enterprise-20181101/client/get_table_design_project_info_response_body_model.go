// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDesignProjectInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTableDesignProjectInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTableDesignProjectInfoResponseBody
	GetErrorMessage() *string
	SetProjectInfo(v *GetTableDesignProjectInfoResponseBodyProjectInfo) *GetTableDesignProjectInfoResponseBody
	GetProjectInfo() *GetTableDesignProjectInfoResponseBodyProjectInfo
	SetRequestId(v string) *GetTableDesignProjectInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableDesignProjectInfoResponseBody
	GetSuccess() *bool
}

type GetTableDesignProjectInfoResponseBody struct {
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
	// The information about the schema design project.
	ProjectInfo *GetTableDesignProjectInfoResponseBodyProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 48778434-5796-571A-8455-A59146588401
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
}

func (s GetTableDesignProjectInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTableDesignProjectInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTableDesignProjectInfoResponseBody) GetProjectInfo() *GetTableDesignProjectInfoResponseBodyProjectInfo {
	return s.ProjectInfo
}

func (s *GetTableDesignProjectInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableDesignProjectInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableDesignProjectInfoResponseBody) SetErrorCode(v string) *GetTableDesignProjectInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBody) SetErrorMessage(v string) *GetTableDesignProjectInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBody) SetProjectInfo(v *GetTableDesignProjectInfoResponseBodyProjectInfo) *GetTableDesignProjectInfoResponseBody {
	s.ProjectInfo = v
	return s
}

func (s *GetTableDesignProjectInfoResponseBody) SetRequestId(v string) *GetTableDesignProjectInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBody) SetSuccess(v bool) *GetTableDesignProjectInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTableDesignProjectInfoResponseBodyProjectInfo struct {
	// The information about the change base database of the schema design ticket.
	BaseDatabase *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase `json:"BaseDatabase,omitempty" xml:"BaseDatabase,omitempty" type:"Struct"`
	// The ID of the user who created the ticket.
	//
	// example:
	//
	// 71****
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the schema design project.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the ticket was created.
	//
	// example:
	//
	// 2024-04-23 02:57:01
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the ticket was last modified.
	//
	// example:
	//
	// 2024-04-23 02:57:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ticket ID.
	//
	// example:
	//
	// 95****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 12****
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The state of the schema design project. Valid values:
	//
	// 	- **DESIGN**: The schema is being designed.
	//
	// 	- **PUBLISHED**: The schema is published.
	//
	// 	- **CLOSE**: The ticket is closed.
	//
	// example:
	//
	// DESIGN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the schema design project.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTableDesignProjectInfoResponseBodyProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectInfoResponseBodyProjectInfo) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetBaseDatabase() *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase {
	return s.BaseDatabase
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetDescription() *string {
	return s.Description
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetStatus() *string {
	return s.Status
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) GetTitle() *string {
	return s.Title
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetBaseDatabase(v *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.BaseDatabase = v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetCreatorId(v int64) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.CreatorId = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetDescription(v string) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.Description = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetGmtCreate(v string) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetGmtModified(v string) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.GmtModified = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetOrderId(v int64) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.OrderId = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetProjectId(v int64) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.ProjectId = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetStatus(v string) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.Status = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) SetTitle(v string) *GetTableDesignProjectInfoResponseBodyProjectInfo {
	s.Title = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfo) Validate() error {
	return dara.Validate(s)
}

type GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase struct {
	// The alias of the database instance.
	//
	// example:
	//
	// poc_test
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The database ID.
	//
	// example:
	//
	// 11****
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// POLARDB
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment in which the database instance is deployed. Valid values:
	//
	// 	- **product**: production environment.
	//
	// 	- **dev**: development environment.
	//
	// 	- **pre**: pre-release environment.
	//
	// 	- **test**: test environment.
	//
	// 	- **sit**: system integration testing (SIT) environment.
	//
	// 	- **uat**: user acceptance testing (UAT) environment.
	//
	// 	- **pet**: stress testing environment.
	//
	// 	- **stag**: staging environment.
	//
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The database name.
	//
	// example:
	//
	// bk_atc020
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// schema_name@127.0.XX.XX
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) GetAlias() *string {
	return s.Alias
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) GetDbId() *int32 {
	return s.DbId
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) GetDbType() *string {
	return s.DbType
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) GetEnvType() *string {
	return s.EnvType
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) GetLogic() *bool {
	return s.Logic
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) SetAlias(v string) *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase {
	s.Alias = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) SetDbId(v int32) *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase {
	s.DbId = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) SetDbType(v string) *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase {
	s.DbType = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) SetEnvType(v string) *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase {
	s.EnvType = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) SetLogic(v bool) *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase {
	s.Logic = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) SetSchemaName(v string) *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase {
	s.SchemaName = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) SetSearchName(v string) *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase {
	s.SearchName = &v
	return s
}

func (s *GetTableDesignProjectInfoResponseBodyProjectInfoBaseDatabase) Validate() error {
	return dara.Validate(s)
}
