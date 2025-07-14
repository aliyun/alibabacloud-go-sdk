// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceConnectionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDataSourceConnectionInfoResponseBody
	GetRequestId() *string
	SetResult(v *GetDataSourceConnectionInfoResponseBodyResult) *GetDataSourceConnectionInfoResponseBody
	GetResult() *GetDataSourceConnectionInfoResponseBodyResult
	SetSuccess(v bool) *GetDataSourceConnectionInfoResponseBody
	GetSuccess() *bool
}

type GetDataSourceConnectionInfoResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 7AAB95D-*****-****-*4FC0C976
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Data source information.
	Result *GetDataSourceConnectionInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataSourceConnectionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceConnectionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceConnectionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataSourceConnectionInfoResponseBody) GetResult() *GetDataSourceConnectionInfoResponseBodyResult {
	return s.Result
}

func (s *GetDataSourceConnectionInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataSourceConnectionInfoResponseBody) SetRequestId(v string) *GetDataSourceConnectionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBody) SetResult(v *GetDataSourceConnectionInfoResponseBodyResult) *GetDataSourceConnectionInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBody) SetSuccess(v bool) *GetDataSourceConnectionInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceConnectionInfoResponseBodyResult struct {
	// Database connection string address (domain or IP).
	//
	// example:
	//
	// 172.**.**.48
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Permission level:
	//
	// - 0 -- Private
	//
	// - 1 -- Collaborative Editing (old)
	//
	// - 11 -- Collaborative Editing - Space Members
	//
	// - 12 -- Collaborative Editing - Specified to Individuals
	//
	// example:
	//
	// 0
	AuthLevel *string `json:"AuthLevel,omitempty" xml:"AuthLevel,omitempty"`
	// Quick BI user ID of the creator.
	//
	// example:
	//
	// U240****0880C6095
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// Data source ID.
	//
	// example:
	//
	// a201c85c-******
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// Data source type.
	//
	// example:
	//
	// mysql
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// Version of the data source.
	//
	// example:
	//
	// 5.7
	DsVersion *string `json:"DsVersion,omitempty" xml:"DsVersion,omitempty"`
	// Database instance, corresponding to the database name, and for ODPS, it is the project.
	//
	// example:
	//
	// rm*********t44ju1
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// rm*********t44ju1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Quick BI user ID of the modifier.
	//
	// example:
	//
	// U240****0880C6095
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// Whether the impala data source requires authentication to log in:
	//
	// - true - Requires account and password login
	//
	// - false - No authentication required (default)
	//
	// example:
	//
	// true
	NoSasl *bool `json:"NoSasl,omitempty" xml:"NoSasl,omitempty"`
	// Primary data source type for multi-engine data sources.
	//
	// example:
	//
	// dataphin
	ParentDsType *string `json:"ParentDsType,omitempty" xml:"ParentDsType,omitempty"`
	// Port.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Used for front-end display when obtaining connection details for ODPS.
	//
	// example:
	//
	// prod-ossdoc
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// Database schema, only needs to be set for databases that support schemas.
	//
	// example:
	//
	// Analysis
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Display name of the data source on the front end.
	//
	// example:
	//
	// 0327
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// Workspace ID to which the data source belongs.
	//
	// example:
	//
	// 0de6**2-d**-4720-8836-0cc****1394c
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDataSourceConnectionInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceConnectionInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetAddress() *string {
	return s.Address
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetAuthLevel() *string {
	return s.AuthLevel
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetDsId() *string {
	return s.DsId
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetDsType() *string {
	return s.DsType
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetDsVersion() *string {
	return s.DsVersion
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetInstance() *string {
	return s.Instance
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetNoSasl() *bool {
	return s.NoSasl
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetParentDsType() *string {
	return s.ParentDsType
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetPort() *string {
	return s.Port
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetProject() *string {
	return s.Project
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetSchema() *string {
	return s.Schema
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetShowName() *string {
	return s.ShowName
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetAddress(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.Address = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetAuthLevel(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.AuthLevel = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetCreatorId(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetDsId(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.DsId = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetDsType(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.DsType = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetDsVersion(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.DsVersion = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetInstance(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.Instance = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetInstanceId(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetModifyUser(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetNoSasl(v bool) *GetDataSourceConnectionInfoResponseBodyResult {
	s.NoSasl = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetParentDsType(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.ParentDsType = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetPort(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.Port = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetProject(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.Project = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetSchema(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.Schema = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetShowName(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.ShowName = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) SetWorkspaceId(v string) *GetDataSourceConnectionInfoResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataSourceConnectionInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
