// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateDatasetJobConfigRequest
	GetConfig() *string
	SetConfigType(v string) *CreateDatasetJobConfigRequest
	GetConfigType() *string
	SetDatasetVersion(v string) *CreateDatasetJobConfigRequest
	GetDatasetVersion() *string
	SetWorkspaceId(v string) *CreateDatasetJobConfigRequest
	GetWorkspaceId() *string
}

type CreateDatasetJobConfigRequest struct {
	// The configuration content. Format:
	//
	// 	- MultimodalIntelligentTag
	//
	// { "apiKey":"sk-xxxxxxxxxxxxxxxxxxxxx" }
	//
	// 	- MultimodalSemanticIndex
	//
	// { "defaultModelId": "xxx" "defaultModelVersion":"1.0.0" }
	//
	// This parameter is required.
	//
	// example:
	//
	// { "apiKey":"sk-xxxxxxxxxxxxxxxxxxxxx" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The configuration type.
	//
	// Valid values:
	//
	// 	- MultimodalIntelligentTag
	//
	// 	- MultimodalSemanticIndex
	//
	// This parameter is required.
	//
	// example:
	//
	// MultimodalIntelligentTag
	ConfigType     *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 454716
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetJobConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetJobConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateDatasetJobConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *CreateDatasetJobConfigRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *CreateDatasetJobConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDatasetJobConfigRequest) SetConfig(v string) *CreateDatasetJobConfigRequest {
	s.Config = &v
	return s
}

func (s *CreateDatasetJobConfigRequest) SetConfigType(v string) *CreateDatasetJobConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *CreateDatasetJobConfigRequest) SetDatasetVersion(v string) *CreateDatasetJobConfigRequest {
	s.DatasetVersion = &v
	return s
}

func (s *CreateDatasetJobConfigRequest) SetWorkspaceId(v string) *CreateDatasetJobConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDatasetJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
