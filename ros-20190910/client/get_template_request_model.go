// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSetId(v string) *GetTemplateRequest
	GetChangeSetId() *string
	SetIncludePermission(v string) *GetTemplateRequest
	GetIncludePermission() *string
	SetIncludeTags(v string) *GetTemplateRequest
	GetIncludeTags() *string
	SetRegionId(v string) *GetTemplateRequest
	GetRegionId() *string
	SetStackGroupName(v string) *GetTemplateRequest
	GetStackGroupName() *string
	SetStackId(v string) *GetTemplateRequest
	GetStackId() *string
	SetTemplateId(v string) *GetTemplateRequest
	GetTemplateId() *string
	SetTemplateStage(v string) *GetTemplateRequest
	GetTemplateStage() *string
	SetTemplateVersion(v string) *GetTemplateRequest
	GetTemplateVersion() *string
}

type GetTemplateRequest struct {
	// The ID of the change set.
	//
	// > You must specify one of the following parameters: StackId, ChangeSetId, StackGroupName, and TemplateId.
	//
	// example:
	//
	// 1f6521a4-05af-4975-afe9-bc4b45ad****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// Specifies whether to query the shared information about the template. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled (default)
	//
	// > Only the template owner can query the shared information of a template.
	//
	// example:
	//
	// Enabled
	IncludePermission *string `json:"IncludePermission,omitempty" xml:"IncludePermission,omitempty"`
	// Specifies whether to query the information about tags. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled (default)
	//
	// > This parameter takes effect only if you specify TemplateId.
	//
	// example:
	//
	// Enabled
	IncludeTags *string `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The region ID of the stack or stack group that uses the template. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group.
	//
	// > You must specify one of the following parameters: StackId, ChangeSetId, StackGroupName, and TemplateId.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The ID of the stack.
	//
	// > You must specify one of the following parameters: StackId, ChangeSetId, StackGroupName, and TemplateId.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The ID of the template.
	//
	// This parameter applies to shared and private templates. If the template is a shared template, the value of TemplateId is the same as the value of TemplateARN. You can use the template ID to query a shared template.
	//
	// > You must specify one of the following parameters: StackId, ChangeSetId, StackGroupName, and TemplateId.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The stage of the template. This parameter takes effect only if you specify StackId, ChangeSetId, or StackGroupName.
	//
	// Valid values:
	//
	// 	- Processed (default): returns the processed template.
	//
	// 	- Original: returns the original template.
	//
	// example:
	//
	// Processed
	TemplateStage *string `json:"TemplateStage,omitempty" xml:"TemplateStage,omitempty"`
	// The version of the template. This parameter takes effect only if you specify TemplateId.\\
	//
	// If the template is a shared template, you can specify this parameter only if VersionOption is set to AllVersions. For more information, see [SetTemplatePermission](https://help.aliyun.com/document_detail/194768.html).
	//
	// Valid values: v1 to v100.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *GetTemplateRequest) GetIncludePermission() *string {
	return s.IncludePermission
}

func (s *GetTemplateRequest) GetIncludeTags() *string {
	return s.IncludeTags
}

func (s *GetTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *GetTemplateRequest) GetStackId() *string {
	return s.StackId
}

func (s *GetTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateRequest) GetTemplateStage() *string {
	return s.TemplateStage
}

func (s *GetTemplateRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateRequest) SetChangeSetId(v string) *GetTemplateRequest {
	s.ChangeSetId = &v
	return s
}

func (s *GetTemplateRequest) SetIncludePermission(v string) *GetTemplateRequest {
	s.IncludePermission = &v
	return s
}

func (s *GetTemplateRequest) SetIncludeTags(v string) *GetTemplateRequest {
	s.IncludeTags = &v
	return s
}

func (s *GetTemplateRequest) SetRegionId(v string) *GetTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRequest) SetStackGroupName(v string) *GetTemplateRequest {
	s.StackGroupName = &v
	return s
}

func (s *GetTemplateRequest) SetStackId(v string) *GetTemplateRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateStage(v string) *GetTemplateRequest {
	s.TemplateStage = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateVersion(v string) *GetTemplateRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
