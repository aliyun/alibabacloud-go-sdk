// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetMaxBindCount(v int64) *CreateProjectShrinkRequest
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *CreateProjectShrinkRequest
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *CreateProjectShrinkRequest
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *CreateProjectShrinkRequest
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *CreateProjectShrinkRequest
	GetDatasetMaxTotalFileSize() *int64
	SetDescription(v string) *CreateProjectShrinkRequest
	GetDescription() *string
	SetProjectMaxDatasetCount(v int64) *CreateProjectShrinkRequest
	GetProjectMaxDatasetCount() *int64
	SetProjectName(v string) *CreateProjectShrinkRequest
	GetProjectName() *string
	SetServiceRole(v string) *CreateProjectShrinkRequest
	GetServiceRole() *string
	SetTagShrink(v string) *CreateProjectShrinkRequest
	GetTagShrink() *string
	SetTemplateId(v string) *CreateProjectShrinkRequest
	GetTemplateId() *string
}

type CreateProjectShrinkRequest struct {
	// The maximum number of bindings for each dataset. Valid values: 1 to 10. Default value: 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities in each dataset. Default value: 10000000000.
	//
	// >  This is a precautionary setting that does not impose practical limitations.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files in each dataset. Valid values: 1 to 100000000. Default value: 10000000000.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships in each dataset. Default value: 100000000000.
	//
	// >  This is a precautionary setting that does not impose practical limitations.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum size of files in each dataset. If the maximum size is exceeded, no indexes can be added. Unit: bytes. Default value: 90000000000000000.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// The description of the project. The description must be 1 to 256 characters in length. You can leave this parameter empty.
	//
	// example:
	//
	// immtest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of datasets in the project. Valid values: 1 to 1000000000. Default value: 1000000000.
	//
	// example:
	//
	// 1000000000
	ProjectMaxDatasetCount *int64 `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	// The name of the project. The name must meet the following requirements:
	//
	// 	- The name must be 1 to 128 characters in length
	//
	// 	- and can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must start with a letter or an underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the Resource Access Management (RAM) role. You must attach the RAM role to IMM to allow IMM to access other cloud resources, such as Object Storage Service (OSS). Default value: `AliyunIMMDefaultRole`.
	//
	// You can also create a custom role in the RAM console and grant the required permissions to the role based on your business requirements. For more information, see [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/477257.html).
	//
	// example:
	//
	// AliyunIMMDefaultRole
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the workflow template. You can leave this parameter empty. For more information, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
	//
	// example:
	//
	// Official:AllFunction
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *CreateProjectShrinkRequest) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *CreateProjectShrinkRequest) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *CreateProjectShrinkRequest) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *CreateProjectShrinkRequest) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *CreateProjectShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectShrinkRequest) GetProjectMaxDatasetCount() *int64 {
	return s.ProjectMaxDatasetCount
}

func (s *CreateProjectShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateProjectShrinkRequest) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *CreateProjectShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateProjectShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateProjectShrinkRequest) SetDatasetMaxBindCount(v int64) *CreateProjectShrinkRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDatasetMaxEntityCount(v int64) *CreateProjectShrinkRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDatasetMaxFileCount(v int64) *CreateProjectShrinkRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDatasetMaxRelationCount(v int64) *CreateProjectShrinkRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDatasetMaxTotalFileSize(v int64) *CreateProjectShrinkRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDescription(v string) *CreateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectMaxDatasetCount(v int64) *CreateProjectShrinkRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectName(v string) *CreateProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetServiceRole(v string) *CreateProjectShrinkRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetTagShrink(v string) *CreateProjectShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetTemplateId(v string) *CreateProjectShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
