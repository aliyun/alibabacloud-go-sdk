// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportTrainingSetTableScript(v string) *GetModelFeatureResponseBody
	GetExportTrainingSetTableScript() *string
	SetFeatures(v []*GetModelFeatureResponseBodyFeatures) *GetModelFeatureResponseBody
	GetFeatures() []*GetModelFeatureResponseBodyFeatures
	SetGmtCreateTime(v string) *GetModelFeatureResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetModelFeatureResponseBody
	GetGmtModifiedTime() *string
	SetLabelPriorityLevel(v int64) *GetModelFeatureResponseBody
	GetLabelPriorityLevel() *int64
	SetLabelTableId(v string) *GetModelFeatureResponseBody
	GetLabelTableId() *string
	SetLabelTableName(v string) *GetModelFeatureResponseBody
	GetLabelTableName() *string
	SetName(v string) *GetModelFeatureResponseBody
	GetName() *string
	SetOwner(v string) *GetModelFeatureResponseBody
	GetOwner() *string
	SetProjectId(v string) *GetModelFeatureResponseBody
	GetProjectId() *string
	SetProjectName(v string) *GetModelFeatureResponseBody
	GetProjectName() *string
	SetRelations(v *GetModelFeatureResponseBodyRelations) *GetModelFeatureResponseBody
	GetRelations() *GetModelFeatureResponseBodyRelations
	SetRequestId(v string) *GetModelFeatureResponseBody
	GetRequestId() *string
	SetTrainingSetFGTable(v string) *GetModelFeatureResponseBody
	GetTrainingSetFGTable() *string
	SetTrainingSetTable(v string) *GetModelFeatureResponseBody
	GetTrainingSetTable() *string
}

type GetModelFeatureResponseBody struct {
	// example:
	//
	// from feature_store_py.fs_client import FeatureStoreClient\nfrom feature_store_py.fs_project import FeatureStoreProject\nfrom feature_store_py.fs_datasource import LabelInput, MaxComputeDataSource, TrainingSetOutput\nfrom feature_store_py.fs_features import FeatureSelector\nfrom feature_store_py.fs_config import LabelInputConfig, PartitionConfig, FeatureViewConfig\nfrom feature_store_py.fs_config import TrainSetOutputConfig, EASDeployConfig\nimport datetime\nimport sys\n\ncur_day = args[\"dt\"]\nprint(\"cur_day = \", cur_day)\noffset = datetime.timedelta(days=-1)\npre_day = (datetime.datetime.strptime(cur_day, \"%Y%m%d\") + offset).strftime(\"%Y%m%d\")\nprint(\"pre_day = \", pre_day)\n\n\naccess_key_id = o.account.access_id\naccess_key_secret = o.account.secret_access_key\nfs = FeatureStoreClient(access_key_id=access_key_id, access_key_secret=access_key_secret, region=\"cn-beijing\")\ncur_project_name = \"p1\"\nproject = fs.get_project(cur_project_name)\n\nlabel_partitions = PartitionConfig(name = \"ds\", value = cur_day)\nlabel_input_config = LabelInputConfig(partition_config=label_partitions)\n\nfeature_view_1_partitions = PartitionConfig(name = \"ds\", value = pre_day)\nfeature_view_1_config = FeatureViewConfig(name = \"user_fea\",\npartition_config=feature_view_1_partitions)\n\nfeature_view_2_partitions = PartitionConfig(name = \"ds\", value = pre_day)\nfeature_view_2_config = FeatureViewConfig(name = \"seq_fea\",\npartition_config=feature_view_2_partitions)\n\nfeature_view_3_partitions = PartitionConfig(name = \"ds\", value = pre_day)\nfeature_view_3_config = FeatureViewConfig(name = \"item_fea\",\npartition_config=feature_view_3_partitions)\n\nfeature_view_config_list = [feature_view_1_config,feature_view_2_config,feature_view_3_config]\ntrain_set_partitions = PartitionConfig(name = \"ds\", value = cur_day)\ntrain_set_output_config = TrainSetOutputConfig(partition_config=train_set_partitions)\n\n\nmodel_name = \"rank_v1\"\ncur_model = project.get_model(model_name)\ntask = cur_model.export_train_set(label_input_config, feature_view_config_list, train_set_output_config)\ntask.wait()\nprint(\"task_summary = \", task.task_summary)\n
	ExportTrainingSetTableScript *string                                `json:"ExportTrainingSetTableScript,omitempty" xml:"ExportTrainingSetTableScript,omitempty"`
	Features                     []*GetModelFeatureResponseBodyFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-07-04T14:46:22.227+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T14:46:22.227+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 0
	LabelPriorityLevel *int64 `json:"LabelPriorityLevel,omitempty" xml:"LabelPriorityLevel,omitempty"`
	// example:
	//
	// 3
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// example:
	//
	// label_table1
	LabelTableName *string `json:"LabelTableName,omitempty" xml:"LabelTableName,omitempty"`
	// example:
	//
	// model_feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1231243253****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 5
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string                               `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Relations   *GetModelFeatureResponseBodyRelations `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Struct"`
	// example:
	//
	// 0C89F5E1-7F24-5EEC-9F05-508A39278CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// table2
	TrainingSetFGTable *string `json:"TrainingSetFGTable,omitempty" xml:"TrainingSetFGTable,omitempty"`
	// example:
	//
	// table1
	TrainingSetTable *string `json:"TrainingSetTable,omitempty" xml:"TrainingSetTable,omitempty"`
}

func (s GetModelFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBody) GetExportTrainingSetTableScript() *string {
	return s.ExportTrainingSetTableScript
}

func (s *GetModelFeatureResponseBody) GetFeatures() []*GetModelFeatureResponseBodyFeatures {
	return s.Features
}

func (s *GetModelFeatureResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetModelFeatureResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetModelFeatureResponseBody) GetLabelPriorityLevel() *int64 {
	return s.LabelPriorityLevel
}

func (s *GetModelFeatureResponseBody) GetLabelTableId() *string {
	return s.LabelTableId
}

func (s *GetModelFeatureResponseBody) GetLabelTableName() *string {
	return s.LabelTableName
}

func (s *GetModelFeatureResponseBody) GetName() *string {
	return s.Name
}

func (s *GetModelFeatureResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *GetModelFeatureResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetModelFeatureResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetModelFeatureResponseBody) GetRelations() *GetModelFeatureResponseBodyRelations {
	return s.Relations
}

func (s *GetModelFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelFeatureResponseBody) GetTrainingSetFGTable() *string {
	return s.TrainingSetFGTable
}

func (s *GetModelFeatureResponseBody) GetTrainingSetTable() *string {
	return s.TrainingSetTable
}

func (s *GetModelFeatureResponseBody) SetExportTrainingSetTableScript(v string) *GetModelFeatureResponseBody {
	s.ExportTrainingSetTableScript = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetFeatures(v []*GetModelFeatureResponseBodyFeatures) *GetModelFeatureResponseBody {
	s.Features = v
	return s
}

func (s *GetModelFeatureResponseBody) SetGmtCreateTime(v string) *GetModelFeatureResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetGmtModifiedTime(v string) *GetModelFeatureResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetLabelPriorityLevel(v int64) *GetModelFeatureResponseBody {
	s.LabelPriorityLevel = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetLabelTableId(v string) *GetModelFeatureResponseBody {
	s.LabelTableId = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetLabelTableName(v string) *GetModelFeatureResponseBody {
	s.LabelTableName = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetName(v string) *GetModelFeatureResponseBody {
	s.Name = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetOwner(v string) *GetModelFeatureResponseBody {
	s.Owner = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetProjectId(v string) *GetModelFeatureResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetProjectName(v string) *GetModelFeatureResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetRelations(v *GetModelFeatureResponseBodyRelations) *GetModelFeatureResponseBody {
	s.Relations = v
	return s
}

func (s *GetModelFeatureResponseBody) SetRequestId(v string) *GetModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetTrainingSetFGTable(v string) *GetModelFeatureResponseBody {
	s.TrainingSetFGTable = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetTrainingSetTable(v string) *GetModelFeatureResponseBody {
	s.TrainingSetTable = &v
	return s
}

func (s *GetModelFeatureResponseBody) Validate() error {
	if s.Features != nil {
		for _, item := range s.Features {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Relations != nil {
		if err := s.Relations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelFeatureResponseBodyFeatures struct {
	// example:
	//
	// feature2
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 3
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// example:
	//
	// feature_view_1
	FeatureViewName *string `json:"FeatureViewName,omitempty" xml:"FeatureViewName,omitempty"`
	// example:
	//
	// feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetModelFeatureResponseBodyFeatures) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureResponseBodyFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBodyFeatures) GetAliasName() *string {
	return s.AliasName
}

func (s *GetModelFeatureResponseBodyFeatures) GetFeatureViewId() *string {
	return s.FeatureViewId
}

func (s *GetModelFeatureResponseBodyFeatures) GetFeatureViewName() *string {
	return s.FeatureViewName
}

func (s *GetModelFeatureResponseBodyFeatures) GetName() *string {
	return s.Name
}

func (s *GetModelFeatureResponseBodyFeatures) GetType() *string {
	return s.Type
}

func (s *GetModelFeatureResponseBodyFeatures) SetAliasName(v string) *GetModelFeatureResponseBodyFeatures {
	s.AliasName = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) SetFeatureViewId(v string) *GetModelFeatureResponseBodyFeatures {
	s.FeatureViewId = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) SetFeatureViewName(v string) *GetModelFeatureResponseBodyFeatures {
	s.FeatureViewName = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) SetName(v string) *GetModelFeatureResponseBodyFeatures {
	s.Name = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) SetType(v string) *GetModelFeatureResponseBodyFeatures {
	s.Type = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) Validate() error {
	return dara.Validate(s)
}

type GetModelFeatureResponseBodyRelations struct {
	Domains []*GetModelFeatureResponseBodyRelationsDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	Links   []*GetModelFeatureResponseBodyRelationsLinks   `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
}

func (s GetModelFeatureResponseBodyRelations) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureResponseBodyRelations) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBodyRelations) GetDomains() []*GetModelFeatureResponseBodyRelationsDomains {
	return s.Domains
}

func (s *GetModelFeatureResponseBodyRelations) GetLinks() []*GetModelFeatureResponseBodyRelationsLinks {
	return s.Links
}

func (s *GetModelFeatureResponseBodyRelations) SetDomains(v []*GetModelFeatureResponseBodyRelationsDomains) *GetModelFeatureResponseBodyRelations {
	s.Domains = v
	return s
}

func (s *GetModelFeatureResponseBodyRelations) SetLinks(v []*GetModelFeatureResponseBodyRelationsLinks) *GetModelFeatureResponseBodyRelations {
	s.Links = v
	return s
}

func (s *GetModelFeatureResponseBodyRelations) Validate() error {
	if s.Domains != nil {
		for _, item := range s.Domains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Links != nil {
		for _, item := range s.Links {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelFeatureResponseBodyRelationsDomains struct {
	// example:
	//
	// FeatureEntity
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// Domain ID。
	//
	// example:
	//
	// 3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetModelFeatureResponseBodyRelationsDomains) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureResponseBodyRelationsDomains) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBodyRelationsDomains) GetDomainType() *string {
	return s.DomainType
}

func (s *GetModelFeatureResponseBodyRelationsDomains) GetId() *string {
	return s.Id
}

func (s *GetModelFeatureResponseBodyRelationsDomains) GetName() *string {
	return s.Name
}

func (s *GetModelFeatureResponseBodyRelationsDomains) SetDomainType(v string) *GetModelFeatureResponseBodyRelationsDomains {
	s.DomainType = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsDomains) SetId(v string) *GetModelFeatureResponseBodyRelationsDomains {
	s.Id = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsDomains) SetName(v string) *GetModelFeatureResponseBodyRelationsDomains {
	s.Name = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsDomains) Validate() error {
	return dara.Validate(s)
}

type GetModelFeatureResponseBodyRelationsLinks struct {
	// example:
	//
	// model_feature_2
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// user_id
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// example:
	//
	// feature_entity_3
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetModelFeatureResponseBodyRelationsLinks) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureResponseBodyRelationsLinks) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBodyRelationsLinks) GetFrom() *string {
	return s.From
}

func (s *GetModelFeatureResponseBodyRelationsLinks) GetLink() *string {
	return s.Link
}

func (s *GetModelFeatureResponseBodyRelationsLinks) GetTo() *string {
	return s.To
}

func (s *GetModelFeatureResponseBodyRelationsLinks) SetFrom(v string) *GetModelFeatureResponseBodyRelationsLinks {
	s.From = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsLinks) SetLink(v string) *GetModelFeatureResponseBodyRelationsLinks {
	s.Link = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsLinks) SetTo(v string) *GetModelFeatureResponseBodyRelationsLinks {
	s.To = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsLinks) Validate() error {
	return dara.Validate(s)
}
