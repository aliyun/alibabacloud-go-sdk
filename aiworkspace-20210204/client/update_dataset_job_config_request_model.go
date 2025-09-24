// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateDatasetJobConfigRequest
	GetConfig() *string
	SetConfigType(v string) *UpdateDatasetJobConfigRequest
	GetConfigType() *string
	SetWorkspaceId(v string) *UpdateDatasetJobConfigRequest
	GetWorkspaceId() *string
}

type UpdateDatasetJobConfigRequest struct {
	// The configuration content. Formats:
	//
	// 	- MultimodalIntelligentTag
	//
	// { "apiKey":"sk-xxxxxxxxxxxxxxxxxxxxx" }
	//
	// 	- MultimodalSemanticIndex
	//
	// { "defaultModelId": "xxx" "defaultModelVersion":"1.0.0" }
	//
	// example:
	//
	// { "apiKey":"sk-xxxxxxxxxxxxxxxxxxxxx" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The configuration type.
	//
	// 	- MultimodalIntelligentTag
	//
	// 	- MultimodalSemanticIndex
	//
	// example:
	//
	// MultimodalSemanticIndex
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 167497
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDatasetJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetJobConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetJobConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateDatasetJobConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateDatasetJobConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDatasetJobConfigRequest) SetConfig(v string) *UpdateDatasetJobConfigRequest {
	s.Config = &v
	return s
}

func (s *UpdateDatasetJobConfigRequest) SetConfigType(v string) *UpdateDatasetJobConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *UpdateDatasetJobConfigRequest) SetWorkspaceId(v string) *UpdateDatasetJobConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDatasetJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
