// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetConfigShrink(v string) *UpdateProjectShrinkRequest
	GetDatasetConfigShrink() *string
	SetDatasetMaxBindCount(v int64) *UpdateProjectShrinkRequest
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *UpdateProjectShrinkRequest
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *UpdateProjectShrinkRequest
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *UpdateProjectShrinkRequest
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *UpdateProjectShrinkRequest
	GetDatasetMaxTotalFileSize() *int64
	SetDescription(v string) *UpdateProjectShrinkRequest
	GetDescription() *string
	SetProjectMaxDatasetCount(v int64) *UpdateProjectShrinkRequest
	GetProjectMaxDatasetCount() *int64
	SetProjectName(v string) *UpdateProjectShrinkRequest
	GetProjectName() *string
	SetServiceRole(v string) *UpdateProjectShrinkRequest
	GetServiceRole() *string
	SetTagShrink(v string) *UpdateProjectShrinkRequest
	GetTagShrink() *string
	SetTemplateId(v string) *UpdateProjectShrinkRequest
	GetTemplateId() *string
}

type UpdateProjectShrinkRequest struct {
	DatasetConfigShrink *string `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty"`
	// The maximum number of bindings for each dataset. Valid values: 1 to 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities in each dataset.
	//
	// >This is a reserved parameter and is not enforced during use.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files in each dataset. Valid values: 1 to 100000000.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships in each dataset.
	//
	// >This is a reserved parameter and is not enforced during use.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum total file size in each dataset. After the limit is exceeded, no more indexes can be added. Unit: bytes.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// The project description. The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// immtest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of datasets in the project. Valid values: 1 to 1000000000.
	//
	// example:
	//
	// 1000000000
	ProjectMaxDatasetCount *int64 `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	// The project name. For information about how to obtain the project name, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The service role that grants Intelligent Media Management (IMM) permissions to access other cloud resources such as Object Storage Service (OSS).
	//
	// To customize a service role, create a regular service role in the Resource Access Management (RAM) console and grant permissions to the role. For more information, see [Create a regular service role](https://help.aliyun.com/document_detail/116800.html) and [Grant permissions to a role](https://help.aliyun.com/document_detail/116147.html).
	//
	// example:
	//
	// AliyunIMMDefaultRole
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// The list of tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The workflow template ID. For more information, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
	//
	// example:
	//
	// Official:ImageManagement
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectShrinkRequest) GetDatasetConfigShrink() *string {
	return s.DatasetConfigShrink
}

func (s *UpdateProjectShrinkRequest) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *UpdateProjectShrinkRequest) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *UpdateProjectShrinkRequest) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *UpdateProjectShrinkRequest) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *UpdateProjectShrinkRequest) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *UpdateProjectShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProjectShrinkRequest) GetProjectMaxDatasetCount() *int64 {
	return s.ProjectMaxDatasetCount
}

func (s *UpdateProjectShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateProjectShrinkRequest) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *UpdateProjectShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *UpdateProjectShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateProjectShrinkRequest) SetDatasetConfigShrink(v string) *UpdateProjectShrinkRequest {
	s.DatasetConfigShrink = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxBindCount(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxEntityCount(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxFileCount(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxRelationCount(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxTotalFileSize(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDescription(v string) *UpdateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetProjectMaxDatasetCount(v int64) *UpdateProjectShrinkRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetProjectName(v string) *UpdateProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetServiceRole(v string) *UpdateProjectShrinkRequest {
	s.ServiceRole = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetTagShrink(v string) *UpdateProjectShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetTemplateId(v string) *UpdateProjectShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
