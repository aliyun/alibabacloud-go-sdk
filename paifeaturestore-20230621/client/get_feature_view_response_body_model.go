// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetFeatureViewResponseBody
	GetConfig() *string
	SetFeatureEntityId(v string) *GetFeatureViewResponseBody
	GetFeatureEntityId() *string
	SetFeatureEntityName(v string) *GetFeatureViewResponseBody
	GetFeatureEntityName() *string
	SetFields(v []*GetFeatureViewResponseBodyFields) *GetFeatureViewResponseBody
	GetFields() []*GetFeatureViewResponseBodyFields
	SetGmtCreateTime(v string) *GetFeatureViewResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetFeatureViewResponseBody
	GetGmtModifiedTime() *string
	SetGmtSyncTime(v string) *GetFeatureViewResponseBody
	GetGmtSyncTime() *string
	SetJoinId(v string) *GetFeatureViewResponseBody
	GetJoinId() *string
	SetLastSyncConfig(v string) *GetFeatureViewResponseBody
	GetLastSyncConfig() *string
	SetMockTableName(v string) *GetFeatureViewResponseBody
	GetMockTableName() *string
	SetName(v string) *GetFeatureViewResponseBody
	GetName() *string
	SetOwner(v string) *GetFeatureViewResponseBody
	GetOwner() *string
	SetProjectId(v string) *GetFeatureViewResponseBody
	GetProjectId() *string
	SetProjectName(v string) *GetFeatureViewResponseBody
	GetProjectName() *string
	SetPublishTableScript(v string) *GetFeatureViewResponseBody
	GetPublishTableScript() *string
	SetRegisterDatasourceId(v string) *GetFeatureViewResponseBody
	GetRegisterDatasourceId() *string
	SetRegisterDatasourceName(v string) *GetFeatureViewResponseBody
	GetRegisterDatasourceName() *string
	SetRegisterTable(v string) *GetFeatureViewResponseBody
	GetRegisterTable() *string
	SetRequestId(v string) *GetFeatureViewResponseBody
	GetRequestId() *string
	SetSyncOnlineTable(v bool) *GetFeatureViewResponseBody
	GetSyncOnlineTable() *bool
	SetTTL(v int32) *GetFeatureViewResponseBody
	GetTTL() *int32
	SetTags(v []*string) *GetFeatureViewResponseBody
	GetTags() []*string
	SetType(v string) *GetFeatureViewResponseBody
	GetType() *string
	SetWriteMethod(v string) *GetFeatureViewResponseBody
	GetWriteMethod() *string
	SetWriteToFeatureDB(v bool) *GetFeatureViewResponseBody
	GetWriteToFeatureDB() *bool
}

type GetFeatureViewResponseBody struct {
	// example:
	//
	// {"save_original_field":true}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// example:
	//
	// featureEntity1
	FeatureEntityName *string                             `json:"FeatureEntityName,omitempty" xml:"FeatureEntityName,omitempty"`
	Fields            []*GetFeatureViewResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtSyncTime *string `json:"GmtSyncTime,omitempty" xml:"GmtSyncTime,omitempty"`
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// example:
	//
	// {
	//
	// 	"mode": "overwrite",
	//
	// 	"partitions": {
	//
	// 		"ds": {
	//
	// 			"value": "20230820"
	//
	// 		}
	//
	// 	},
	//
	// 	"event_time": "",
	//
	// 	"config": {},
	//
	// 	"offline_to_online": true
	//
	// }
	LastSyncConfig *string `json:"LastSyncConfig,omitempty" xml:"LastSyncConfig,omitempty"`
	// example:
	//
	// item_table_mock_1
	MockTableName *string `json:"MockTableName,omitempty" xml:"MockTableName,omitempty"`
	// example:
	//
	// featureView1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12321421412****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// from feature_store_py.fs_client import FeatureStoreClient\nimport datetime\nfrom feature_store_py.fs_datasource import MaxComputeDataSource\nimport sys\n\ncur_day = args[\"dt\"]\nprint(\"cur_day = \", cur_day)\n\naccess_key_id = o.account.access_id\naccess_key_secret = o.account.secret_access_key\nfs = FeatureStoreClient(access_key_id=access_key_id, access_key_secret=access_key_secret, region=\"cn-beijing\")\ncur_project_name = \"p1\"\nproject = fs.get_project(cur_project_name)\n\nfeature_view_name = \"user_fea\"\nbatch_feature_view = project.get_feature_view(feature_view_name)\ntask = batch_feature_view.publish_table(partitions={\"ds\":cur_day}, mode=\"Overwrite\")\ntask.wait()\ntask.print_summary()\n
	PublishTableScript *string `json:"PublishTableScript,omitempty" xml:"PublishTableScript,omitempty"`
	// example:
	//
	// 4
	RegisterDatasourceId *string `json:"RegisterDatasourceId,omitempty" xml:"RegisterDatasourceId,omitempty"`
	// example:
	//
	// datasource1
	RegisterDatasourceName *string `json:"RegisterDatasourceName,omitempty" xml:"RegisterDatasourceName,omitempty"`
	// example:
	//
	// table1
	RegisterTable *string `json:"RegisterTable,omitempty" xml:"RegisterTable,omitempty"`
	// example:
	//
	// 72F15A8A-5A28-5B18-A0DE-0EABD7D3245A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SyncOnlineTable *bool `json:"SyncOnlineTable,omitempty" xml:"SyncOnlineTable,omitempty"`
	// example:
	//
	// 90
	TTL  *int32    `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// Batch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// custom
	WriteMethod      *string `json:"WriteMethod,omitempty" xml:"WriteMethod,omitempty"`
	WriteToFeatureDB *bool   `json:"WriteToFeatureDB,omitempty" xml:"WriteToFeatureDB,omitempty"`
}

func (s GetFeatureViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureViewResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetFeatureViewResponseBody) GetFeatureEntityId() *string {
	return s.FeatureEntityId
}

func (s *GetFeatureViewResponseBody) GetFeatureEntityName() *string {
	return s.FeatureEntityName
}

func (s *GetFeatureViewResponseBody) GetFields() []*GetFeatureViewResponseBodyFields {
	return s.Fields
}

func (s *GetFeatureViewResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetFeatureViewResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetFeatureViewResponseBody) GetGmtSyncTime() *string {
	return s.GmtSyncTime
}

func (s *GetFeatureViewResponseBody) GetJoinId() *string {
	return s.JoinId
}

func (s *GetFeatureViewResponseBody) GetLastSyncConfig() *string {
	return s.LastSyncConfig
}

func (s *GetFeatureViewResponseBody) GetMockTableName() *string {
	return s.MockTableName
}

func (s *GetFeatureViewResponseBody) GetName() *string {
	return s.Name
}

func (s *GetFeatureViewResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *GetFeatureViewResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetFeatureViewResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetFeatureViewResponseBody) GetPublishTableScript() *string {
	return s.PublishTableScript
}

func (s *GetFeatureViewResponseBody) GetRegisterDatasourceId() *string {
	return s.RegisterDatasourceId
}

func (s *GetFeatureViewResponseBody) GetRegisterDatasourceName() *string {
	return s.RegisterDatasourceName
}

func (s *GetFeatureViewResponseBody) GetRegisterTable() *string {
	return s.RegisterTable
}

func (s *GetFeatureViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFeatureViewResponseBody) GetSyncOnlineTable() *bool {
	return s.SyncOnlineTable
}

func (s *GetFeatureViewResponseBody) GetTTL() *int32 {
	return s.TTL
}

func (s *GetFeatureViewResponseBody) GetTags() []*string {
	return s.Tags
}

func (s *GetFeatureViewResponseBody) GetType() *string {
	return s.Type
}

func (s *GetFeatureViewResponseBody) GetWriteMethod() *string {
	return s.WriteMethod
}

func (s *GetFeatureViewResponseBody) GetWriteToFeatureDB() *bool {
	return s.WriteToFeatureDB
}

func (s *GetFeatureViewResponseBody) SetConfig(v string) *GetFeatureViewResponseBody {
	s.Config = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetFeatureEntityId(v string) *GetFeatureViewResponseBody {
	s.FeatureEntityId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetFeatureEntityName(v string) *GetFeatureViewResponseBody {
	s.FeatureEntityName = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetFields(v []*GetFeatureViewResponseBodyFields) *GetFeatureViewResponseBody {
	s.Fields = v
	return s
}

func (s *GetFeatureViewResponseBody) SetGmtCreateTime(v string) *GetFeatureViewResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetGmtModifiedTime(v string) *GetFeatureViewResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetGmtSyncTime(v string) *GetFeatureViewResponseBody {
	s.GmtSyncTime = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetJoinId(v string) *GetFeatureViewResponseBody {
	s.JoinId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetLastSyncConfig(v string) *GetFeatureViewResponseBody {
	s.LastSyncConfig = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetMockTableName(v string) *GetFeatureViewResponseBody {
	s.MockTableName = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetName(v string) *GetFeatureViewResponseBody {
	s.Name = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetOwner(v string) *GetFeatureViewResponseBody {
	s.Owner = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetProjectId(v string) *GetFeatureViewResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetProjectName(v string) *GetFeatureViewResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetPublishTableScript(v string) *GetFeatureViewResponseBody {
	s.PublishTableScript = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetRegisterDatasourceId(v string) *GetFeatureViewResponseBody {
	s.RegisterDatasourceId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetRegisterDatasourceName(v string) *GetFeatureViewResponseBody {
	s.RegisterDatasourceName = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetRegisterTable(v string) *GetFeatureViewResponseBody {
	s.RegisterTable = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetRequestId(v string) *GetFeatureViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetSyncOnlineTable(v bool) *GetFeatureViewResponseBody {
	s.SyncOnlineTable = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetTTL(v int32) *GetFeatureViewResponseBody {
	s.TTL = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetTags(v []*string) *GetFeatureViewResponseBody {
	s.Tags = v
	return s
}

func (s *GetFeatureViewResponseBody) SetType(v string) *GetFeatureViewResponseBody {
	s.Type = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetWriteMethod(v string) *GetFeatureViewResponseBody {
	s.WriteMethod = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetWriteToFeatureDB(v bool) *GetFeatureViewResponseBody {
	s.WriteToFeatureDB = &v
	return s
}

func (s *GetFeatureViewResponseBody) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFeatureViewResponseBodyFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// user
	Name      *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Transform []*GetFeatureViewResponseBodyFieldsTransform `json:"Transform,omitempty" xml:"Transform,omitempty" type:"Repeated"`
	// example:
	//
	// int
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFeatureViewResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureViewResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetFeatureViewResponseBodyFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *GetFeatureViewResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *GetFeatureViewResponseBodyFields) GetTransform() []*GetFeatureViewResponseBodyFieldsTransform {
	return s.Transform
}

func (s *GetFeatureViewResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *GetFeatureViewResponseBodyFields) SetAttributes(v []*string) *GetFeatureViewResponseBodyFields {
	s.Attributes = v
	return s
}

func (s *GetFeatureViewResponseBodyFields) SetName(v string) *GetFeatureViewResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetFeatureViewResponseBodyFields) SetTransform(v []*GetFeatureViewResponseBodyFieldsTransform) *GetFeatureViewResponseBodyFields {
	s.Transform = v
	return s
}

func (s *GetFeatureViewResponseBodyFields) SetType(v string) *GetFeatureViewResponseBodyFields {
	s.Type = &v
	return s
}

func (s *GetFeatureViewResponseBodyFields) Validate() error {
	if s.Transform != nil {
		for _, item := range s.Transform {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFeatureViewResponseBodyFieldsTransform struct {
	Input []*GetFeatureViewResponseBodyFieldsTransformInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	LLMConfigId *int32 `json:"LLMConfigId,omitempty" xml:"LLMConfigId,omitempty"`
	// example:
	//
	// LLMEmbedding
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFeatureViewResponseBodyFieldsTransform) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureViewResponseBodyFieldsTransform) GoString() string {
	return s.String()
}

func (s *GetFeatureViewResponseBodyFieldsTransform) GetInput() []*GetFeatureViewResponseBodyFieldsTransformInput {
	return s.Input
}

func (s *GetFeatureViewResponseBodyFieldsTransform) GetLLMConfigId() *int32 {
	return s.LLMConfigId
}

func (s *GetFeatureViewResponseBodyFieldsTransform) GetType() *string {
	return s.Type
}

func (s *GetFeatureViewResponseBodyFieldsTransform) SetInput(v []*GetFeatureViewResponseBodyFieldsTransformInput) *GetFeatureViewResponseBodyFieldsTransform {
	s.Input = v
	return s
}

func (s *GetFeatureViewResponseBodyFieldsTransform) SetLLMConfigId(v int32) *GetFeatureViewResponseBodyFieldsTransform {
	s.LLMConfigId = &v
	return s
}

func (s *GetFeatureViewResponseBodyFieldsTransform) SetType(v string) *GetFeatureViewResponseBodyFieldsTransform {
	s.Type = &v
	return s
}

func (s *GetFeatureViewResponseBodyFieldsTransform) Validate() error {
	if s.Input != nil {
		for _, item := range s.Input {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFeatureViewResponseBodyFieldsTransformInput struct {
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFeatureViewResponseBodyFieldsTransformInput) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureViewResponseBodyFieldsTransformInput) GoString() string {
	return s.String()
}

func (s *GetFeatureViewResponseBodyFieldsTransformInput) GetName() *string {
	return s.Name
}

func (s *GetFeatureViewResponseBodyFieldsTransformInput) GetType() *string {
	return s.Type
}

func (s *GetFeatureViewResponseBodyFieldsTransformInput) SetName(v string) *GetFeatureViewResponseBodyFieldsTransformInput {
	s.Name = &v
	return s
}

func (s *GetFeatureViewResponseBodyFieldsTransformInput) SetType(v string) *GetFeatureViewResponseBodyFieldsTransformInput {
	s.Type = &v
	return s
}

func (s *GetFeatureViewResponseBodyFieldsTransformInput) Validate() error {
	return dara.Validate(s)
}
