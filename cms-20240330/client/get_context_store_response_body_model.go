// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *GetContextStoreResponseBodyConfig) *GetContextStoreResponseBody
	GetConfig() *GetContextStoreResponseBodyConfig
	SetContextStoreName(v string) *GetContextStoreResponseBody
	GetContextStoreName() *string
	SetContextType(v string) *GetContextStoreResponseBody
	GetContextType() *string
	SetCreateTime(v string) *GetContextStoreResponseBody
	GetCreateTime() *string
	SetDataset(v *GetContextStoreResponseBodyDataset) *GetContextStoreResponseBody
	GetDataset() *GetContextStoreResponseBodyDataset
	SetDescription(v string) *GetContextStoreResponseBody
	GetDescription() *string
	SetRegionId(v string) *GetContextStoreResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetContextStoreResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetContextStoreResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *GetContextStoreResponseBody
	GetUpdateTime() *string
	SetWorkspace(v string) *GetContextStoreResponseBody
	GetWorkspace() *string
}

type GetContextStoreResponseBody struct {
	Config *GetContextStoreResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// test-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// example:
	//
	// memory
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1695090077
	CreateTime *string                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Dataset    *GetContextStoreResponseBodyDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1695090077
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetContextStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBody) GetConfig() *GetContextStoreResponseBodyConfig {
	return s.Config
}

func (s *GetContextStoreResponseBody) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *GetContextStoreResponseBody) GetContextType() *string {
	return s.ContextType
}

func (s *GetContextStoreResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetContextStoreResponseBody) GetDataset() *GetContextStoreResponseBodyDataset {
	return s.Dataset
}

func (s *GetContextStoreResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetContextStoreResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetContextStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextStoreResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetContextStoreResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetContextStoreResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetContextStoreResponseBody) SetConfig(v *GetContextStoreResponseBodyConfig) *GetContextStoreResponseBody {
	s.Config = v
	return s
}

func (s *GetContextStoreResponseBody) SetContextStoreName(v string) *GetContextStoreResponseBody {
	s.ContextStoreName = &v
	return s
}

func (s *GetContextStoreResponseBody) SetContextType(v string) *GetContextStoreResponseBody {
	s.ContextType = &v
	return s
}

func (s *GetContextStoreResponseBody) SetCreateTime(v string) *GetContextStoreResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetContextStoreResponseBody) SetDataset(v *GetContextStoreResponseBodyDataset) *GetContextStoreResponseBody {
	s.Dataset = v
	return s
}

func (s *GetContextStoreResponseBody) SetDescription(v string) *GetContextStoreResponseBody {
	s.Description = &v
	return s
}

func (s *GetContextStoreResponseBody) SetRegionId(v string) *GetContextStoreResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetContextStoreResponseBody) SetRequestId(v string) *GetContextStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextStoreResponseBody) SetStatus(v string) *GetContextStoreResponseBody {
	s.Status = &v
	return s
}

func (s *GetContextStoreResponseBody) SetUpdateTime(v string) *GetContextStoreResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetContextStoreResponseBody) SetWorkspace(v string) *GetContextStoreResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetContextStoreResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfig struct {
	InnerSource   *GetContextStoreResponseBodyConfigInnerSource `json:"innerSource,omitempty" xml:"innerSource,omitempty" type:"Struct"`
	MetadataField map[string]*string                            `json:"metadataField,omitempty" xml:"metadataField,omitempty"`
	Source        *GetContextStoreResponseBodyConfigSource      `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s GetContextStoreResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfig) GetInnerSource() *GetContextStoreResponseBodyConfigInnerSource {
	return s.InnerSource
}

func (s *GetContextStoreResponseBodyConfig) GetMetadataField() map[string]*string {
	return s.MetadataField
}

func (s *GetContextStoreResponseBodyConfig) GetSource() *GetContextStoreResponseBodyConfigSource {
	return s.Source
}

func (s *GetContextStoreResponseBodyConfig) SetInnerSource(v *GetContextStoreResponseBodyConfigInnerSource) *GetContextStoreResponseBodyConfig {
	s.InnerSource = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetMetadataField(v map[string]*string) *GetContextStoreResponseBodyConfig {
	s.MetadataField = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetSource(v *GetContextStoreResponseBodyConfigSource) *GetContextStoreResponseBodyConfig {
	s.Source = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) Validate() error {
	if s.InnerSource != nil {
		if err := s.InnerSource.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfigInnerSource struct {
	// example:
	//
	// sls-test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// sls-test-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s GetContextStoreResponseBodyConfigInnerSource) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigInnerSource) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigInnerSource) GetLogstore() *string {
	return s.Logstore
}

func (s *GetContextStoreResponseBodyConfigInnerSource) GetProject() *string {
	return s.Project
}

func (s *GetContextStoreResponseBodyConfigInnerSource) SetLogstore(v string) *GetContextStoreResponseBodyConfigInnerSource {
	s.Logstore = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigInnerSource) SetProject(v string) *GetContextStoreResponseBodyConfigInnerSource {
	s.Project = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigInnerSource) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyConfigSource struct {
	// example:
	//
	// sls-test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// sls-test-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1754962200000
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSource) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSource) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSource) GetLogstore() *string {
	return s.Logstore
}

func (s *GetContextStoreResponseBodyConfigSource) GetProject() *string {
	return s.Project
}

func (s *GetContextStoreResponseBodyConfigSource) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextStoreResponseBodyConfigSource) SetLogstore(v string) *GetContextStoreResponseBodyConfigSource {
	s.Logstore = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetProject(v string) *GetContextStoreResponseBodyConfigSource {
	s.Project = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetStartTime(v string) *GetContextStoreResponseBodyConfigSource {
	s.StartTime = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) Validate() error {
	return dara.Validate(s)
}

type GetContextStoreResponseBodyDataset struct {
	// example:
	//
	// test_dataset
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetContextStoreResponseBodyDataset) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyDataset) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyDataset) GetName() *string {
	return s.Name
}

func (s *GetContextStoreResponseBodyDataset) SetName(v string) *GetContextStoreResponseBodyDataset {
	s.Name = &v
	return s
}

func (s *GetContextStoreResponseBodyDataset) Validate() error {
	return dara.Validate(s)
}
