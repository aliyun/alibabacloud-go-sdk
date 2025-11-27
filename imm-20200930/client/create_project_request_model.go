// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetMaxBindCount(v int64) *CreateProjectRequest
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *CreateProjectRequest
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *CreateProjectRequest
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *CreateProjectRequest
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *CreateProjectRequest
	GetDatasetMaxTotalFileSize() *int64
	SetDescription(v string) *CreateProjectRequest
	GetDescription() *string
	SetProjectMaxDatasetCount(v int64) *CreateProjectRequest
	GetProjectMaxDatasetCount() *int64
	SetProjectName(v string) *CreateProjectRequest
	GetProjectName() *string
	SetServiceRole(v string) *CreateProjectRequest
	GetServiceRole() *string
	SetTag(v []*CreateProjectRequestTag) *CreateProjectRequest
	GetTag() []*CreateProjectRequestTag
	SetTemplateId(v string) *CreateProjectRequest
	GetTemplateId() *string
}

type CreateProjectRequest struct {
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
	Tag []*CreateProjectRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the workflow template. You can leave this parameter empty. For more information, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
	//
	// example:
	//
	// Official:AllFunction
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *CreateProjectRequest) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *CreateProjectRequest) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *CreateProjectRequest) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *CreateProjectRequest) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *CreateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectRequest) GetProjectMaxDatasetCount() *int64 {
	return s.ProjectMaxDatasetCount
}

func (s *CreateProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateProjectRequest) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *CreateProjectRequest) GetTag() []*CreateProjectRequestTag {
	return s.Tag
}

func (s *CreateProjectRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateProjectRequest) SetDatasetMaxBindCount(v int64) *CreateProjectRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxEntityCount(v int64) *CreateProjectRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxFileCount(v int64) *CreateProjectRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxRelationCount(v int64) *CreateProjectRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxTotalFileSize(v int64) *CreateProjectRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetProjectMaxDatasetCount(v int64) *CreateProjectRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetServiceRole(v string) *CreateProjectRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateProjectRequest) SetTag(v []*CreateProjectRequestTag) *CreateProjectRequest {
	s.Tag = v
	return s
}

func (s *CreateProjectRequest) SetTemplateId(v string) *CreateProjectRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProjectRequestTag struct {
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

func (s CreateProjectRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestTag) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateProjectRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateProjectRequestTag) SetKey(v string) *CreateProjectRequestTag {
	s.Key = &v
	return s
}

func (s *CreateProjectRequestTag) SetValue(v string) *CreateProjectRequestTag {
	s.Value = &v
	return s
}

func (s *CreateProjectRequestTag) Validate() error {
	return dara.Validate(s)
}
