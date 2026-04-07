// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetJobConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DatasetJobConfig
	GetConfig() *string
	SetConfigType(v string) *DatasetJobConfig
	GetConfigType() *string
	SetCreateTime(v string) *DatasetJobConfig
	GetCreateTime() *string
	SetDatasetJobConfigId(v string) *DatasetJobConfig
	GetDatasetJobConfigId() *string
	SetDatasetVersion(v string) *DatasetJobConfig
	GetDatasetVersion() *string
	SetModifyTime(v string) *DatasetJobConfig
	GetModifyTime() *string
	SetWorkspaceId(v string) *DatasetJobConfig
	GetWorkspaceId() *string
}

type DatasetJobConfig struct {
	// The content of the dataset job configuration, in the JSON format.
	//
	// example:
	//
	// { "apiKey":"sk-xxxxxxxxxxxxxxxxxxxxx" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of the dataset job configuration.
	//
	// example:
	//
	// MultimodalIntelligentTag
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The time when the dataset job was created.
	//
	// example:
	//
	// 2025-01-14T01:37:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the dataset job configuration.
	//
	// example:
	//
	// dscfg-xxxxxxxxxxxx
	DatasetJobConfigId *string `json:"DatasetJobConfigId,omitempty" xml:"DatasetJobConfigId,omitempty"`
	DatasetVersion     *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The time when the dataset job was modified.
	//
	// example:
	//
	// 2024-10-11T02:18:54Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 234*34
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DatasetJobConfig) String() string {
	return dara.Prettify(s)
}

func (s DatasetJobConfig) GoString() string {
	return s.String()
}

func (s *DatasetJobConfig) GetConfig() *string {
	return s.Config
}

func (s *DatasetJobConfig) GetConfigType() *string {
	return s.ConfigType
}

func (s *DatasetJobConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DatasetJobConfig) GetDatasetJobConfigId() *string {
	return s.DatasetJobConfigId
}

func (s *DatasetJobConfig) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *DatasetJobConfig) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DatasetJobConfig) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DatasetJobConfig) SetConfig(v string) *DatasetJobConfig {
	s.Config = &v
	return s
}

func (s *DatasetJobConfig) SetConfigType(v string) *DatasetJobConfig {
	s.ConfigType = &v
	return s
}

func (s *DatasetJobConfig) SetCreateTime(v string) *DatasetJobConfig {
	s.CreateTime = &v
	return s
}

func (s *DatasetJobConfig) SetDatasetJobConfigId(v string) *DatasetJobConfig {
	s.DatasetJobConfigId = &v
	return s
}

func (s *DatasetJobConfig) SetDatasetVersion(v string) *DatasetJobConfig {
	s.DatasetVersion = &v
	return s
}

func (s *DatasetJobConfig) SetModifyTime(v string) *DatasetJobConfig {
	s.ModifyTime = &v
	return s
}

func (s *DatasetJobConfig) SetWorkspaceId(v string) *DatasetJobConfig {
	s.WorkspaceId = &v
	return s
}

func (s *DatasetJobConfig) Validate() error {
	return dara.Validate(s)
}
