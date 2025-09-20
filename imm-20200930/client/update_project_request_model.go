// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetMaxBindCount(v int64) *UpdateProjectRequest
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *UpdateProjectRequest
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *UpdateProjectRequest
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *UpdateProjectRequest
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *UpdateProjectRequest
	GetDatasetMaxTotalFileSize() *int64
	SetDescription(v string) *UpdateProjectRequest
	GetDescription() *string
	SetProjectMaxDatasetCount(v int64) *UpdateProjectRequest
	GetProjectMaxDatasetCount() *int64
	SetProjectName(v string) *UpdateProjectRequest
	GetProjectName() *string
	SetServiceRole(v string) *UpdateProjectRequest
	GetServiceRole() *string
	SetTag(v []*UpdateProjectRequestTag) *UpdateProjectRequest
	GetTag() []*UpdateProjectRequestTag
	SetTemplateId(v string) *UpdateProjectRequest
	GetTemplateId() *string
}

type UpdateProjectRequest struct {
	// The maximum number of bindings for each dataset. Valid values: 1 to 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities in each dataset.
	//
	// >  This is a precautionary setting that does not impose practical limitations.
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
	// The maximum number of metadata relationships in a dataset.
	//
	// >  This is a precautionary setting that does not impose practical limitations.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum size of files in each dataset. If the maximum size is exceeded, indexes can no longer be added. Unit: bytes.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// The description of the project. The description must be 1 to 256 characters in length.
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
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the Resource Access Management (RAM) role. You must grant the RAM role to Intelligent Media Management (IMM) before IMM can access other cloud resources such as Object Storage Service (OSS).
	//
	// You can also create a custom service role in the RAM console and grant the required permissions to the role based on your business requirements. For more information, see [Create a regular service role](https://help.aliyun.com/document_detail/116800.html) and [Grant permissions to a role](https://help.aliyun.com/document_detail/116147.html).
	//
	// example:
	//
	// AliyunIMMDefaultRole
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// The tags.
	Tag []*UpdateProjectRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the workflow template. For more information, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
	//
	// example:
	//
	// AliyunIMMDefaultRole
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *UpdateProjectRequest) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *UpdateProjectRequest) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *UpdateProjectRequest) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *UpdateProjectRequest) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *UpdateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProjectRequest) GetProjectMaxDatasetCount() *int64 {
	return s.ProjectMaxDatasetCount
}

func (s *UpdateProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateProjectRequest) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *UpdateProjectRequest) GetTag() []*UpdateProjectRequestTag {
	return s.Tag
}

func (s *UpdateProjectRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateProjectRequest) SetDatasetMaxBindCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxEntityCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxFileCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxRelationCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxTotalFileSize(v int64) *UpdateProjectRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectMaxDatasetCount(v int64) *UpdateProjectRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectName(v string) *UpdateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectRequest) SetServiceRole(v string) *UpdateProjectRequest {
	s.ServiceRole = &v
	return s
}

func (s *UpdateProjectRequest) SetTag(v []*UpdateProjectRequestTag) *UpdateProjectRequest {
	s.Tag = v
	return s
}

func (s *UpdateProjectRequest) SetTemplateId(v string) *UpdateProjectRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateProjectRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateProjectRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequestTag) GetKey() *string {
	return s.Key
}

func (s *UpdateProjectRequestTag) GetValue() *string {
	return s.Value
}

func (s *UpdateProjectRequestTag) SetKey(v string) *UpdateProjectRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateProjectRequestTag) SetValue(v string) *UpdateProjectRequestTag {
	s.Value = &v
	return s
}

func (s *UpdateProjectRequestTag) Validate() error {
	return dara.Validate(s)
}
