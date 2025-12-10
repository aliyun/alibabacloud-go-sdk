// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetExperimentResponseBody
	GetAccessibility() *string
	SetContent(v string) *GetExperimentResponseBody
	GetContent() *string
	SetCreator(v string) *GetExperimentResponseBody
	GetCreator() *string
	SetDescription(v string) *GetExperimentResponseBody
	GetDescription() *string
	SetExperimentId(v string) *GetExperimentResponseBody
	GetExperimentId() *string
	SetGmtCreateTime(v string) *GetExperimentResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetExperimentResponseBody
	GetGmtModifiedTime() *string
	SetName(v string) *GetExperimentResponseBody
	GetName() *string
	SetOptions(v string) *GetExperimentResponseBody
	GetOptions() *string
	SetRequestId(v string) *GetExperimentResponseBody
	GetRequestId() *string
	SetSource(v string) *GetExperimentResponseBody
	GetSource() *string
	SetVersion(v int64) *GetExperimentResponseBody
	GetVersion() *int64
	SetWorkspaceId(v string) *GetExperimentResponseBody
	GetWorkspaceId() *string
}

type GetExperimentResponseBody struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// {     "nodes": [         {             "id": "id-4c50-1609236241393-76174",             "name": "读数据表",             "metadata": {                 "identifier": "data_source",                 "version": "v1",                 "provider": "pai"             },             "properties": [                 {                     "name": "hasPartition",                     "value": false                 },                 {                     "name": "inputTableName",                     "value": "pai_online_project.e_commerce_test_data"                 }             ],             "position": {                 "x": 427,                 "y": 123             }         },         {             "id": "id-23ce-1609236252041-30106",             "name": "读数据表",             "metadata": {                 "identifier": "data_source",                 "version": "v1",                 "provider": "pai"             },             "properties": [                 {                     "name": "inputTableName",                     "value": "pai_online_project.e_commerce_train_data"                 },                 {                     "name": "hasPartition",                     "value": false                 }             ],             "position": {                 "x": 226,                 "y": 127             }         },         {             "id": "id-97a7-1609236275421-84245",             "name": "DeepFM",             "metadata": {                 "identifier": "deepfm",                 "version": "v1",                 "provider": "pai"             },             "properties": [                 {                     "name": "trainModel",                     "value": "train"                 },                 {                     "name": "arn",                     "value": true                 },                 {                     "name": "cluster",                     "value": "\n{\n    \"ps\": {\n        \"count\": 2,\n        \"cpu\": 1000,\n        \"memory\": 40000\n    },\n    \"worker\": {\n        \"count\": 8,\n        \"cpu\": 1000,\n        \"memory\": 40000\n    }\n}"                 }             ],             "position": {                 "x": 323,                 "y": 345             }         }     ],     "edges": [         {             "source": "id-23ce-1609236252041-30106",             "sourceAnchor": "outputTable",             "targetAnchor": "input1TableName",             "target": "id-97a7-1609236275421-84245"         },         {             "source": "id-4c50-1609236241393-76174",             "sourceAnchor": "outputTable",             "targetAnchor": "input2TableName",             "target": "id-97a7-1609236275421-84245"         }     ],     "globalParams": [],     "globalSettings": []  }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1326689413376250
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// Pipeline Draft Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// draft-rbvg5wzljzjhc9ks92
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// Pipeline Draft Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"mlflow":{"experimentId":"exp-1"}}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// DD405810-73C9-5721-996A-EA04BCC4BBF2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 12
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 23487
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetExperimentResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetExperimentResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetExperimentResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetExperimentResponseBody) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *GetExperimentResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetExperimentResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetExperimentResponseBody) GetName() *string {
	return s.Name
}

func (s *GetExperimentResponseBody) GetOptions() *string {
	return s.Options
}

func (s *GetExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentResponseBody) GetSource() *string {
	return s.Source
}

func (s *GetExperimentResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *GetExperimentResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetExperimentResponseBody) SetAccessibility(v string) *GetExperimentResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetExperimentResponseBody) SetContent(v string) *GetExperimentResponseBody {
	s.Content = &v
	return s
}

func (s *GetExperimentResponseBody) SetCreator(v string) *GetExperimentResponseBody {
	s.Creator = &v
	return s
}

func (s *GetExperimentResponseBody) SetDescription(v string) *GetExperimentResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentResponseBody) SetExperimentId(v string) *GetExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResponseBody) SetGmtCreateTime(v string) *GetExperimentResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetGmtModifiedTime(v string) *GetExperimentResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetName(v string) *GetExperimentResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentResponseBody) SetOptions(v string) *GetExperimentResponseBody {
	s.Options = &v
	return s
}

func (s *GetExperimentResponseBody) SetRequestId(v string) *GetExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentResponseBody) SetSource(v string) *GetExperimentResponseBody {
	s.Source = &v
	return s
}

func (s *GetExperimentResponseBody) SetVersion(v int64) *GetExperimentResponseBody {
	s.Version = &v
	return s
}

func (s *GetExperimentResponseBody) SetWorkspaceId(v string) *GetExperimentResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
