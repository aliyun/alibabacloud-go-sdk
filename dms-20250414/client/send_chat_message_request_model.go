// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *SendChatMessageRequest
	GetAgentId() *string
	SetDMSUnit(v string) *SendChatMessageRequest
	GetDMSUnit() *string
	SetDataSource(v *SendChatMessageRequestDataSource) *SendChatMessageRequest
	GetDataSource() *SendChatMessageRequestDataSource
	SetDataSources(v []*SendChatMessageRequestDataSources) *SendChatMessageRequest
	GetDataSources() []*SendChatMessageRequestDataSources
	SetMessage(v string) *SendChatMessageRequest
	GetMessage() *string
	SetMessageType(v string) *SendChatMessageRequest
	GetMessageType() *string
	SetParentSessionId(v string) *SendChatMessageRequest
	GetParentSessionId() *string
	SetQuestion(v string) *SendChatMessageRequest
	GetQuestion() *string
	SetQuotedMessage(v string) *SendChatMessageRequest
	GetQuotedMessage() *string
	SetReplyTo(v string) *SendChatMessageRequest
	GetReplyTo() *string
	SetSessionConfig(v *SendChatMessageRequestSessionConfig) *SendChatMessageRequest
	GetSessionConfig() *SendChatMessageRequestSessionConfig
	SetSessionId(v string) *SendChatMessageRequest
	GetSessionId() *string
	SetTaskConfig(v *SendChatMessageRequestTaskConfig) *SendChatMessageRequest
	GetTaskConfig() *SendChatMessageRequestTaskConfig
	SetUserOssBucket(v string) *SendChatMessageRequest
	GetUserOssBucket() *string
	SetWorkspaceId(v string) *SendChatMessageRequest
	GetWorkspaceId() *string
}

type SendChatMessageRequest struct {
	// **[Optimized]*	- This field is now automatically retrieved in the background. You do not need to specify it.
	//
	// example:
	//
	// agent_***
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// **[Optimized]*	- This field is now automatically retrieved in the background. You do not need to specify it when calling the API.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The data source information. This parameter can be left empty. You can specify only one data source for this parameter. Use the DataSources parameter instead.
	//
	// example:
	//
	// null
	DataSource *SendChatMessageRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The detailed data source information. This parameter can be left empty.
	DataSources []*SendChatMessageRequestDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The content of the message sent to the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// what can you do?
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The message type. Default value: `primary`.
	//
	// - In normal interactions with the agent, the message type is `primary`.
	//
	// - When the message is a response to a human-in-the-loop question from the agent, the type must be `additional`.
	//
	// - When the message is intended to trigger report generation, the type must be `report`.
	//
	// - When the message is intended to cancel the current session, the type must be `cancel`.
	//
	// example:
	//
	// primary
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The ID of the parent session.
	//
	// example:
	//
	// 20qrliuoo7p2vlsfg*****
	ParentSessionId *string `json:"ParentSessionId,omitempty" xml:"ParentSessionId,omitempty"`
	// Required when the message type is `additional`. Specifies the specific question asked by the agent to the user through human-in-the-loop.
	//
	// example:
	//
	// Provide the calculation criteria for GMV
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// The quoted content. This is typically used when interacting with the agent.
	//
	// example:
	//
	// {"version":"v0"}
	QuotedMessage *string `json:"QuotedMessage,omitempty" xml:"QuotedMessage,omitempty"`
	// **Important**
	//
	// When this message is a reply to an agent message, such as an agent asking for clarification through ASK_HUMAN, `reply_to` must exactly match the checkpoint sequence number carried by that agent message. If this message is not a targeted reply, such as requesting the agent to perform further in-depth analysis after the initial analysis is complete, you can leave `reply_to` empty or set it to "0".
	//
	// This field affects how the agent decides to process the message. Passing an incorrect value may lead to suboptimal analysis results.
	//
	// example:
	//
	// 0
	ReplyTo *string `json:"ReplyTo,omitempty" xml:"ReplyTo,omitempty"`
	// The special configurations for the current session. For the same session, only the configurations specified in the first SendMessage call take effect.
	//
	// if can be null:
	// true
	SessionConfig *SendChatMessageRequestSessionConfig `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty" type:"Struct"`
	// The session ID. This optional field is used for multi-turn sessions.
	//
	// - You can start a session without specifying this field. The response will include the session ID for the current session.
	//
	// - You can also manually create a session ID by calling the CreateDataAgentSession operation and include the ID when initiating a session.
	//
	// - If you need multi-turn conversations, such as follow-up questions or confirming execution plans, specify the session ID returned by the previous SendChatMessage operation in this operation.
	//
	// example:
	//
	// sess_***
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The configuration items that affect only the current task.
	TaskConfig *SendChatMessageRequestTaskConfig `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty" type:"Struct"`
	// The Object Storage Service (OSS) bucket of the user. If this parameter is left empty, the analysis process is securely stored in the built-in storage.
	//
	// example:
	//
	// my-bucket
	UserOssBucket *string `json:"UserOssBucket,omitempty" xml:"UserOssBucket,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// c1p71ne***baexrt3o
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SendChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequest) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *SendChatMessageRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *SendChatMessageRequest) GetDataSource() *SendChatMessageRequestDataSource {
	return s.DataSource
}

func (s *SendChatMessageRequest) GetDataSources() []*SendChatMessageRequestDataSources {
	return s.DataSources
}

func (s *SendChatMessageRequest) GetMessage() *string {
	return s.Message
}

func (s *SendChatMessageRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *SendChatMessageRequest) GetParentSessionId() *string {
	return s.ParentSessionId
}

func (s *SendChatMessageRequest) GetQuestion() *string {
	return s.Question
}

func (s *SendChatMessageRequest) GetQuotedMessage() *string {
	return s.QuotedMessage
}

func (s *SendChatMessageRequest) GetReplyTo() *string {
	return s.ReplyTo
}

func (s *SendChatMessageRequest) GetSessionConfig() *SendChatMessageRequestSessionConfig {
	return s.SessionConfig
}

func (s *SendChatMessageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendChatMessageRequest) GetTaskConfig() *SendChatMessageRequestTaskConfig {
	return s.TaskConfig
}

func (s *SendChatMessageRequest) GetUserOssBucket() *string {
	return s.UserOssBucket
}

func (s *SendChatMessageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SendChatMessageRequest) SetAgentId(v string) *SendChatMessageRequest {
	s.AgentId = &v
	return s
}

func (s *SendChatMessageRequest) SetDMSUnit(v string) *SendChatMessageRequest {
	s.DMSUnit = &v
	return s
}

func (s *SendChatMessageRequest) SetDataSource(v *SendChatMessageRequestDataSource) *SendChatMessageRequest {
	s.DataSource = v
	return s
}

func (s *SendChatMessageRequest) SetDataSources(v []*SendChatMessageRequestDataSources) *SendChatMessageRequest {
	s.DataSources = v
	return s
}

func (s *SendChatMessageRequest) SetMessage(v string) *SendChatMessageRequest {
	s.Message = &v
	return s
}

func (s *SendChatMessageRequest) SetMessageType(v string) *SendChatMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendChatMessageRequest) SetParentSessionId(v string) *SendChatMessageRequest {
	s.ParentSessionId = &v
	return s
}

func (s *SendChatMessageRequest) SetQuestion(v string) *SendChatMessageRequest {
	s.Question = &v
	return s
}

func (s *SendChatMessageRequest) SetQuotedMessage(v string) *SendChatMessageRequest {
	s.QuotedMessage = &v
	return s
}

func (s *SendChatMessageRequest) SetReplyTo(v string) *SendChatMessageRequest {
	s.ReplyTo = &v
	return s
}

func (s *SendChatMessageRequest) SetSessionConfig(v *SendChatMessageRequestSessionConfig) *SendChatMessageRequest {
	s.SessionConfig = v
	return s
}

func (s *SendChatMessageRequest) SetSessionId(v string) *SendChatMessageRequest {
	s.SessionId = &v
	return s
}

func (s *SendChatMessageRequest) SetTaskConfig(v *SendChatMessageRequestTaskConfig) *SendChatMessageRequest {
	s.TaskConfig = v
	return s
}

func (s *SendChatMessageRequest) SetUserOssBucket(v string) *SendChatMessageRequest {
	s.UserOssBucket = &v
	return s
}

func (s *SendChatMessageRequest) SetWorkspaceId(v string) *SendChatMessageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SendChatMessageRequest) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SessionConfig != nil {
		if err := s.SessionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TaskConfig != nil {
		if err := s.TaskConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendChatMessageRequestDataSource struct {
	// **[Deprecated]*	- You do not need to specify this field.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source type. Valid values: `remote_data_center` and `database`. `remote_data_center` indicates a file, and `database` indicates a database.
	//
	// example:
	//
	// remote_data_center
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// **[Deprecated]*	- You do not need to specify this field.
	//
	// example:
	//
	// test_db
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// ******
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the database in Data Management (DMS).
	//
	// example:
	//
	// 23******
	DmsDatabaseId *string `json:"DmsDatabaseId,omitempty" xml:"DmsDatabaseId,omitempty"`
	// The ID of the instance in DMS.
	//
	// example:
	//
	// 12******
	DmsInstanceId *string `json:"DmsInstanceId,omitempty" xml:"DmsInstanceId,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 35****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// **[Deprecated]*	- You do not need to specify this field.
	//
	// example:
	//
	// localhost
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The permission constraints for querying the current data source. The permission constraint feature is in canary release. This field does not take effect for users outside the canary release scope.
	Permission *SendChatMessageRequestDataSourcePermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of names of the tables to be analyzed.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The data source resource identifier.
	//
	// example:
	//
	// datalake://abc/def
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s SendChatMessageRequestDataSource) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestDataSource) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestDataSource) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *SendChatMessageRequestDataSource) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *SendChatMessageRequestDataSource) GetDatabase() *string {
	return s.Database
}

func (s *SendChatMessageRequestDataSource) GetDbName() *string {
	return s.DbName
}

func (s *SendChatMessageRequestDataSource) GetDmsDatabaseId() *string {
	return s.DmsDatabaseId
}

func (s *SendChatMessageRequestDataSource) GetDmsInstanceId() *string {
	return s.DmsInstanceId
}

func (s *SendChatMessageRequestDataSource) GetEngine() *string {
	return s.Engine
}

func (s *SendChatMessageRequestDataSource) GetFileId() *string {
	return s.FileId
}

func (s *SendChatMessageRequestDataSource) GetLocation() *string {
	return s.Location
}

func (s *SendChatMessageRequestDataSource) GetPermission() *SendChatMessageRequestDataSourcePermission {
	return s.Permission
}

func (s *SendChatMessageRequestDataSource) GetRegionId() *string {
	return s.RegionId
}

func (s *SendChatMessageRequestDataSource) GetTables() []*string {
	return s.Tables
}

func (s *SendChatMessageRequestDataSource) GetUri() *string {
	return s.Uri
}

func (s *SendChatMessageRequestDataSource) SetDataSourceId(v string) *SendChatMessageRequestDataSource {
	s.DataSourceId = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetDataSourceType(v string) *SendChatMessageRequestDataSource {
	s.DataSourceType = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetDatabase(v string) *SendChatMessageRequestDataSource {
	s.Database = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetDbName(v string) *SendChatMessageRequestDataSource {
	s.DbName = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetDmsDatabaseId(v string) *SendChatMessageRequestDataSource {
	s.DmsDatabaseId = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetDmsInstanceId(v string) *SendChatMessageRequestDataSource {
	s.DmsInstanceId = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetEngine(v string) *SendChatMessageRequestDataSource {
	s.Engine = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetFileId(v string) *SendChatMessageRequestDataSource {
	s.FileId = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetLocation(v string) *SendChatMessageRequestDataSource {
	s.Location = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetPermission(v *SendChatMessageRequestDataSourcePermission) *SendChatMessageRequestDataSource {
	s.Permission = v
	return s
}

func (s *SendChatMessageRequestDataSource) SetRegionId(v string) *SendChatMessageRequestDataSource {
	s.RegionId = &v
	return s
}

func (s *SendChatMessageRequestDataSource) SetTables(v []*string) *SendChatMessageRequestDataSource {
	s.Tables = v
	return s
}

func (s *SendChatMessageRequestDataSource) SetUri(v string) *SendChatMessageRequestDataSource {
	s.Uri = &v
	return s
}

func (s *SendChatMessageRequestDataSource) Validate() error {
	if s.Permission != nil {
		if err := s.Permission.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendChatMessageRequestDataSourcePermission struct {
	// The table-level permission constraints for querying the current data source. Each item in the list represents the permission constraints for a specific table.
	Tables []*SendChatMessageRequestDataSourcePermissionTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s SendChatMessageRequestDataSourcePermission) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestDataSourcePermission) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestDataSourcePermission) GetTables() []*SendChatMessageRequestDataSourcePermissionTables {
	return s.Tables
}

func (s *SendChatMessageRequestDataSourcePermission) SetTables(v []*SendChatMessageRequestDataSourcePermissionTables) *SendChatMessageRequestDataSourcePermission {
	s.Tables = v
	return s
}

func (s *SendChatMessageRequestDataSourcePermission) Validate() error {
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendChatMessageRequestDataSourcePermissionTables struct {
	// The list of columns that are allowed to be queried in the current table. If you leave this field empty, all columns can be queried. If you specify this field, SQL statements that query columns beyond the allowed scope, such as `SELECT *`, will be blocked. To ensure the analysis performance of DataAgent, make sure that the prompts, knowledge, and instructions modules of DataAgent do not include columns beyond the allowed scope. Otherwise, unauthorized SQL statements may be generated and blocked, which degrades the analysis speed and performance of DataAgent.
	AllowedColumns []*string `json:"AllowedColumns,omitempty" xml:"AllowedColumns,omitempty" type:"Repeated"`
	// The list of columns that are not allowed to be queried in the current table. If you leave this field empty, all columns **can be queried**. If you specify this field, SQL statements that query columns beyond the allowed scope, such as `SELECT *`, will be blocked.
	//
	// If both the disallowed list and the allowed list are configured, the disallowed list takes precedence. For example, if Table 1 has columns A, B, C, and D, and you configure columns A and B as not allowed, while configuring columns B and C as allowed, only column C is ultimately allowed to be queried.
	//
	// To ensure the analysis performance of DataAgent, make sure that the prompts, knowledge, and instructions modules of DataAgent do not include columns beyond the allowed scope. Otherwise, unauthorized SQL statements may be generated and blocked, which degrades the analysis speed and performance of DataAgent.
	DisallowedColumns []*string `json:"DisallowedColumns,omitempty" xml:"DisallowedColumns,omitempty" type:"Repeated"`
	// The row filter condition that must be included for the current table. If you leave this field empty, it is ignored. If you specify this field, the system checks whether all SQL statements involving this table include the filter column and whether the WHERE clause meets the constraints. SQL statements that do not meet the constraints will be rejected. Pay attention to the format of the validation condition.
	//
	// example:
	//
	// region = \\"east\\"
	RequiredRowFilter *string `json:"RequiredRowFilter,omitempty" xml:"RequiredRowFilter,omitempty"`
	// The name of the table to which the permission constraint rules apply.
	//
	// example:
	//
	// sample_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s SendChatMessageRequestDataSourcePermissionTables) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestDataSourcePermissionTables) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestDataSourcePermissionTables) GetAllowedColumns() []*string {
	return s.AllowedColumns
}

func (s *SendChatMessageRequestDataSourcePermissionTables) GetDisallowedColumns() []*string {
	return s.DisallowedColumns
}

func (s *SendChatMessageRequestDataSourcePermissionTables) GetRequiredRowFilter() *string {
	return s.RequiredRowFilter
}

func (s *SendChatMessageRequestDataSourcePermissionTables) GetTableName() *string {
	return s.TableName
}

func (s *SendChatMessageRequestDataSourcePermissionTables) SetAllowedColumns(v []*string) *SendChatMessageRequestDataSourcePermissionTables {
	s.AllowedColumns = v
	return s
}

func (s *SendChatMessageRequestDataSourcePermissionTables) SetDisallowedColumns(v []*string) *SendChatMessageRequestDataSourcePermissionTables {
	s.DisallowedColumns = v
	return s
}

func (s *SendChatMessageRequestDataSourcePermissionTables) SetRequiredRowFilter(v string) *SendChatMessageRequestDataSourcePermissionTables {
	s.RequiredRowFilter = &v
	return s
}

func (s *SendChatMessageRequestDataSourcePermissionTables) SetTableName(v string) *SendChatMessageRequestDataSourcePermissionTables {
	s.TableName = &v
	return s
}

func (s *SendChatMessageRequestDataSourcePermissionTables) Validate() error {
	return dara.Validate(s)
}

type SendChatMessageRequestDataSources struct {
	// **[Deprecated]*	- You do not need to specify this field.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source type. Valid values: remote_data_center and database, which indicate that the analysis is based on a file or a database, respectively.
	//
	// example:
	//
	// remote_data_center
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// **[Deprecated]*	- You do not need to specify this field.
	//
	// example:
	//
	// test_db
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// mydatabase
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the database in Data Management (DMS).
	//
	// example:
	//
	// 123****
	DmsDatabaseId *string `json:"DmsDatabaseId,omitempty" xml:"DmsDatabaseId,omitempty"`
	// The ID of the instance in DMS.
	//
	// example:
	//
	// 248*****
	DmsInstanceId *string `json:"DmsInstanceId,omitempty" xml:"DmsInstanceId,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The file ID.
	//
	// example:
	//
	// f-4w*******
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// **[Deprecated]*	- You do not need to specify this field.
	//
	// example:
	//
	// localhost
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The permission constraints for querying the current data source. The permission constraint feature is in canary release. This field does not take effect for users outside the canary release scope.
	Permission *SendChatMessageRequestDataSourcesPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of names of the tables to be analyzed.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The data source resource identifier.
	//
	// example:
	//
	// datalake://abc/def
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s SendChatMessageRequestDataSources) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestDataSources) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestDataSources) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *SendChatMessageRequestDataSources) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *SendChatMessageRequestDataSources) GetDatabase() *string {
	return s.Database
}

func (s *SendChatMessageRequestDataSources) GetDbName() *string {
	return s.DbName
}

func (s *SendChatMessageRequestDataSources) GetDmsDatabaseId() *string {
	return s.DmsDatabaseId
}

func (s *SendChatMessageRequestDataSources) GetDmsInstanceId() *string {
	return s.DmsInstanceId
}

func (s *SendChatMessageRequestDataSources) GetEngine() *string {
	return s.Engine
}

func (s *SendChatMessageRequestDataSources) GetFileId() *string {
	return s.FileId
}

func (s *SendChatMessageRequestDataSources) GetLocation() *string {
	return s.Location
}

func (s *SendChatMessageRequestDataSources) GetPermission() *SendChatMessageRequestDataSourcesPermission {
	return s.Permission
}

func (s *SendChatMessageRequestDataSources) GetRegionId() *string {
	return s.RegionId
}

func (s *SendChatMessageRequestDataSources) GetTables() []*string {
	return s.Tables
}

func (s *SendChatMessageRequestDataSources) GetUri() *string {
	return s.Uri
}

func (s *SendChatMessageRequestDataSources) SetDataSourceId(v string) *SendChatMessageRequestDataSources {
	s.DataSourceId = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetDataSourceType(v string) *SendChatMessageRequestDataSources {
	s.DataSourceType = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetDatabase(v string) *SendChatMessageRequestDataSources {
	s.Database = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetDbName(v string) *SendChatMessageRequestDataSources {
	s.DbName = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetDmsDatabaseId(v string) *SendChatMessageRequestDataSources {
	s.DmsDatabaseId = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetDmsInstanceId(v string) *SendChatMessageRequestDataSources {
	s.DmsInstanceId = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetEngine(v string) *SendChatMessageRequestDataSources {
	s.Engine = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetFileId(v string) *SendChatMessageRequestDataSources {
	s.FileId = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetLocation(v string) *SendChatMessageRequestDataSources {
	s.Location = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetPermission(v *SendChatMessageRequestDataSourcesPermission) *SendChatMessageRequestDataSources {
	s.Permission = v
	return s
}

func (s *SendChatMessageRequestDataSources) SetRegionId(v string) *SendChatMessageRequestDataSources {
	s.RegionId = &v
	return s
}

func (s *SendChatMessageRequestDataSources) SetTables(v []*string) *SendChatMessageRequestDataSources {
	s.Tables = v
	return s
}

func (s *SendChatMessageRequestDataSources) SetUri(v string) *SendChatMessageRequestDataSources {
	s.Uri = &v
	return s
}

func (s *SendChatMessageRequestDataSources) Validate() error {
	if s.Permission != nil {
		if err := s.Permission.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendChatMessageRequestDataSourcesPermission struct {
	// The table-level permission constraints for querying the current data source. Each item in the list represents the permission constraints for a specific table.
	Tables []*SendChatMessageRequestDataSourcesPermissionTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s SendChatMessageRequestDataSourcesPermission) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestDataSourcesPermission) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestDataSourcesPermission) GetTables() []*SendChatMessageRequestDataSourcesPermissionTables {
	return s.Tables
}

func (s *SendChatMessageRequestDataSourcesPermission) SetTables(v []*SendChatMessageRequestDataSourcesPermissionTables) *SendChatMessageRequestDataSourcesPermission {
	s.Tables = v
	return s
}

func (s *SendChatMessageRequestDataSourcesPermission) Validate() error {
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendChatMessageRequestDataSourcesPermissionTables struct {
	// The list of columns that are allowed to be queried in the current table. If this parameter is left empty, all columns can be queried. If this parameter is specified, SQL statements that query columns beyond the allowed scope, such as SELECT *, are prohibited. To ensure the analysis performance of DataAgent, make sure that the prompts, knowledge, and instructions of DataAgent do not include columns beyond the allowed scope. Otherwise, unauthorized SQL statements may be generated and blocked, which degrades the analysis speed and performance of DataAgent.
	AllowedColumns []*string `json:"AllowedColumns,omitempty" xml:"AllowedColumns,omitempty" type:"Repeated"`
	// The list of columns that are not allowed to be queried in the current table. If you leave this field empty, all columns **can be queried**. If you specify this field, SQL statements that query columns beyond the allowed scope, such as `SELECT *`, will be blocked.
	//
	// If both the disallowed list and the allowed list are configured, the disallowed list takes precedence. For example, if Table 1 has columns A, B, C, and D, and you configure columns A and B as not allowed, while configuring columns B and C as allowed, only column C is ultimately allowed to be queried.
	//
	// To ensure the analysis performance of DataAgent, make sure that the prompts, knowledge, and instructions modules of DataAgent do not include columns beyond the allowed scope. Otherwise, unauthorized SQL statements may be generated and blocked, which degrades the analysis speed and performance of DataAgent.
	DisallowedColumns []*string `json:"DisallowedColumns,omitempty" xml:"DisallowedColumns,omitempty" type:"Repeated"`
	// The row filter condition that must be included for the current table. If you leave this field empty, it is ignored. If you specify this field, the system checks whether all SQL statements involving this table include the filter column and whether the WHERE clause meets the constraints. SQL statements that do not meet the constraints will be rejected. Pay attention to the format of the validation condition.
	//
	// example:
	//
	// region = \\"east\\"
	RequiredRowFilter *string `json:"RequiredRowFilter,omitempty" xml:"RequiredRowFilter,omitempty"`
	// The name of the table to which the permission constraint rules apply.
	//
	// example:
	//
	// sample_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s SendChatMessageRequestDataSourcesPermissionTables) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestDataSourcesPermissionTables) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) GetAllowedColumns() []*string {
	return s.AllowedColumns
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) GetDisallowedColumns() []*string {
	return s.DisallowedColumns
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) GetRequiredRowFilter() *string {
	return s.RequiredRowFilter
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) GetTableName() *string {
	return s.TableName
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) SetAllowedColumns(v []*string) *SendChatMessageRequestDataSourcesPermissionTables {
	s.AllowedColumns = v
	return s
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) SetDisallowedColumns(v []*string) *SendChatMessageRequestDataSourcesPermissionTables {
	s.DisallowedColumns = v
	return s
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) SetRequiredRowFilter(v string) *SendChatMessageRequestDataSourcesPermissionTables {
	s.RequiredRowFilter = &v
	return s
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) SetTableName(v string) *SendChatMessageRequestDataSourcesPermissionTables {
	s.TableName = &v
	return s
}

func (s *SendChatMessageRequestDataSourcesPermissionTables) Validate() error {
	return dara.Validate(s)
}

type SendChatMessageRequestSessionConfig struct {
	// The ID of the custom agent. A custom agent is an entity used to customize the analysis process. You can create a custom agent in the DataAgent console or by calling the CreateCustomAgent operation. The ID of a custom agent must start with `ca-`.
	//
	// example:
	//
	// ca-asfkjhqwe***aesrfqa
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The stage of the custom agent.
	//
	// example:
	//
	// null
	CustomAgentStage *string `json:"CustomAgentStage,omitempty" xml:"CustomAgentStage,omitempty"`
	// Specifies whether to enable web search.
	//
	// example:
	//
	// false
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// The list of knowledge base IDs.
	//
	// example:
	//
	// dasd***cc211
	KbUuidList []*string `json:"KbUuidList,omitempty" xml:"KbUuidList,omitempty" type:"Repeated"`
	// The language. Only Chinese and English are supported. The default value is Chinese. Only uppercase letters are supported.
	//
	// example:
	//
	// ENGLISH
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The MCP server IDs in the session configuration.
	//
	// example:
	//
	// 2q1tu90**********6uttu2nw
	McpServerIds []*string `json:"McpServerIds,omitempty" xml:"McpServerIds,omitempty" type:"Repeated"`
	// The mode. Valid values:
	//
	//  - **ASK_DATA**: the data query mode.
	//
	//  - **ANALYSIS**: the analysis mode.
	//
	//  - **INSIGHT**: the insight mode.
	//
	// example:
	//
	// ANALYSIS
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The data permission settings at the session level. If you configure this parameter multiple times in multi-turn conversations within the same session, the last configuration prevails.
	PermissionConfig *SendChatMessageRequestSessionConfigPermissionConfig `json:"PermissionConfig,omitempty" xml:"PermissionConfig,omitempty" type:"Struct"`
	// Specifies whether to enable the plan feature. Valid values: `disable`, `enable`, and `force`. Default value: `enable`.
	//
	// example:
	//
	// disable
	PlanMode *string `json:"PlanMode,omitempty" xml:"PlanMode,omitempty"`
	// The watermark text. You can enter up to 64 characters. The text is used to add a watermark to the generated PDF report.
	//
	// example:
	//
	// sample-watermark
	ReportWaterMark *string `json:"ReportWaterMark,omitempty" xml:"ReportWaterMark,omitempty"`
	// Specifies whether to disable user inquiries during the process.
	//
	// example:
	//
	// True
	SkipAskHuman *bool `json:"SkipAskHuman,omitempty" xml:"SkipAskHuman,omitempty"`
	// Specifies whether to skip the plan confirmation step.
	//
	// example:
	//
	// True
	SkipPlan *bool `json:"SkipPlan,omitempty" xml:"SkipPlan,omitempty"`
	// Specifies whether to skip all SQL confirmations.
	//
	// example:
	//
	// False
	SkipSqlConfirm *bool `json:"SkipSqlConfirm,omitempty" xml:"SkipSqlConfirm,omitempty"`
	// Specifies whether to skip the web report generation confirmation.
	//
	// example:
	//
	// True
	SkipWebReportConfirm *bool `json:"SkipWebReportConfirm,omitempty" xml:"SkipWebReportConfirm,omitempty"`
	// The list of user-specified skills.
	UserSpecifiedSkillList []*string `json:"UserSpecifiedSkillList,omitempty" xml:"UserSpecifiedSkillList,omitempty" type:"Repeated"`
}

func (s SendChatMessageRequestSessionConfig) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestSessionConfig) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestSessionConfig) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *SendChatMessageRequestSessionConfig) GetCustomAgentStage() *string {
	return s.CustomAgentStage
}

func (s *SendChatMessageRequestSessionConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *SendChatMessageRequestSessionConfig) GetKbUuidList() []*string {
	return s.KbUuidList
}

func (s *SendChatMessageRequestSessionConfig) GetLanguage() *string {
	return s.Language
}

func (s *SendChatMessageRequestSessionConfig) GetMcpServerIds() []*string {
	return s.McpServerIds
}

func (s *SendChatMessageRequestSessionConfig) GetMode() *string {
	return s.Mode
}

func (s *SendChatMessageRequestSessionConfig) GetPermissionConfig() *SendChatMessageRequestSessionConfigPermissionConfig {
	return s.PermissionConfig
}

func (s *SendChatMessageRequestSessionConfig) GetPlanMode() *string {
	return s.PlanMode
}

func (s *SendChatMessageRequestSessionConfig) GetReportWaterMark() *string {
	return s.ReportWaterMark
}

func (s *SendChatMessageRequestSessionConfig) GetSkipAskHuman() *bool {
	return s.SkipAskHuman
}

func (s *SendChatMessageRequestSessionConfig) GetSkipPlan() *bool {
	return s.SkipPlan
}

func (s *SendChatMessageRequestSessionConfig) GetSkipSqlConfirm() *bool {
	return s.SkipSqlConfirm
}

func (s *SendChatMessageRequestSessionConfig) GetSkipWebReportConfirm() *bool {
	return s.SkipWebReportConfirm
}

func (s *SendChatMessageRequestSessionConfig) GetUserSpecifiedSkillList() []*string {
	return s.UserSpecifiedSkillList
}

func (s *SendChatMessageRequestSessionConfig) SetCustomAgentId(v string) *SendChatMessageRequestSessionConfig {
	s.CustomAgentId = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetCustomAgentStage(v string) *SendChatMessageRequestSessionConfig {
	s.CustomAgentStage = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetEnableSearch(v bool) *SendChatMessageRequestSessionConfig {
	s.EnableSearch = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetKbUuidList(v []*string) *SendChatMessageRequestSessionConfig {
	s.KbUuidList = v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetLanguage(v string) *SendChatMessageRequestSessionConfig {
	s.Language = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetMcpServerIds(v []*string) *SendChatMessageRequestSessionConfig {
	s.McpServerIds = v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetMode(v string) *SendChatMessageRequestSessionConfig {
	s.Mode = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetPermissionConfig(v *SendChatMessageRequestSessionConfigPermissionConfig) *SendChatMessageRequestSessionConfig {
	s.PermissionConfig = v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetPlanMode(v string) *SendChatMessageRequestSessionConfig {
	s.PlanMode = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetReportWaterMark(v string) *SendChatMessageRequestSessionConfig {
	s.ReportWaterMark = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetSkipAskHuman(v bool) *SendChatMessageRequestSessionConfig {
	s.SkipAskHuman = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetSkipPlan(v bool) *SendChatMessageRequestSessionConfig {
	s.SkipPlan = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetSkipSqlConfirm(v bool) *SendChatMessageRequestSessionConfig {
	s.SkipSqlConfirm = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetSkipWebReportConfirm(v bool) *SendChatMessageRequestSessionConfig {
	s.SkipWebReportConfirm = &v
	return s
}

func (s *SendChatMessageRequestSessionConfig) SetUserSpecifiedSkillList(v []*string) *SendChatMessageRequestSessionConfig {
	s.UserSpecifiedSkillList = v
	return s
}

func (s *SendChatMessageRequestSessionConfig) Validate() error {
	if s.PermissionConfig != nil {
		if err := s.PermissionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendChatMessageRequestSessionConfigPermissionConfig struct {
	// The default behavior for table-level permissions. This parameter controls whether tables that are not configured in DataSources.[item].Permission can be queried. Valid values: `allow` and `deny`.
	//
	// - `allow`: the default value. This value indicates that the blacklist mode is used. By default, all tables can be queried, and the actual permissions are subject to the permissions configured in DMS. Only when you configure **row-level and column-level*	- permissions in DataSources.[item].Permission, the row-level and column-level permissions are used for access control.
	//
	// - `deny`: This value indicates that the whitelist mode is used. You must manually specify this value. By default, no tables can be queried. A table can be queried only when you configure permissions for the table in DataSources.[item].Permission. You can grant full access to a table by not configuring any row-level or column-level restrictions.
	//
	// example:
	//
	// allow
	DefaultAction *string `json:"DefaultAction,omitempty" xml:"DefaultAction,omitempty"`
}

func (s SendChatMessageRequestSessionConfigPermissionConfig) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestSessionConfigPermissionConfig) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestSessionConfigPermissionConfig) GetDefaultAction() *string {
	return s.DefaultAction
}

func (s *SendChatMessageRequestSessionConfigPermissionConfig) SetDefaultAction(v string) *SendChatMessageRequestSessionConfigPermissionConfig {
	s.DefaultAction = &v
	return s
}

func (s *SendChatMessageRequestSessionConfigPermissionConfig) Validate() error {
	return dara.Validate(s)
}

type SendChatMessageRequestTaskConfig struct {
	// The configuration of the report rules. A report task is executed based on this configuration only when MessageType is set to REPORT.
	ReportConfig *SendChatMessageRequestTaskConfigReportConfig `json:"ReportConfig,omitempty" xml:"ReportConfig,omitempty" type:"Struct"`
}

func (s SendChatMessageRequestTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestTaskConfig) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestTaskConfig) GetReportConfig() *SendChatMessageRequestTaskConfigReportConfig {
	return s.ReportConfig
}

func (s *SendChatMessageRequestTaskConfig) SetReportConfig(v *SendChatMessageRequestTaskConfigReportConfig) *SendChatMessageRequestTaskConfig {
	s.ReportConfig = v
	return s
}

func (s *SendChatMessageRequestTaskConfig) Validate() error {
	if s.ReportConfig != nil {
		if err := s.ReportConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendChatMessageRequestTaskConfigReportConfig struct {
	// The prompt that must be followed for the current report.
	//
	// example:
	//
	// generate a report
	ReportPrompt *string `json:"ReportPrompt,omitempty" xml:"ReportPrompt,omitempty"`
	// The report theme. Valid values: default, journal, legacy, and neobrutalism.
	//
	// example:
	//
	// default
	ReportTheme *string `json:"ReportTheme,omitempty" xml:"ReportTheme,omitempty"`
	// The service type. Valid values: TextReport and WebReport, which indicate that the current task generates a text report or a web report, respectively. Only WebReport is supported.
	//
	// example:
	//
	// WebReport
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s SendChatMessageRequestTaskConfigReportConfig) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestTaskConfigReportConfig) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestTaskConfigReportConfig) GetReportPrompt() *string {
	return s.ReportPrompt
}

func (s *SendChatMessageRequestTaskConfigReportConfig) GetReportTheme() *string {
	return s.ReportTheme
}

func (s *SendChatMessageRequestTaskConfigReportConfig) GetReportType() *string {
	return s.ReportType
}

func (s *SendChatMessageRequestTaskConfigReportConfig) SetReportPrompt(v string) *SendChatMessageRequestTaskConfigReportConfig {
	s.ReportPrompt = &v
	return s
}

func (s *SendChatMessageRequestTaskConfigReportConfig) SetReportTheme(v string) *SendChatMessageRequestTaskConfigReportConfig {
	s.ReportTheme = &v
	return s
}

func (s *SendChatMessageRequestTaskConfigReportConfig) SetReportType(v string) *SendChatMessageRequestTaskConfigReportConfig {
	s.ReportType = &v
	return s
}

func (s *SendChatMessageRequestTaskConfigReportConfig) Validate() error {
	return dara.Validate(s)
}
