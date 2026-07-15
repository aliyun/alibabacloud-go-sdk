// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultVersion(v int32) *GetJobTemplateResponseBody
	GetDefaultVersion() *int32
	SetDescription(v string) *GetJobTemplateResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *GetJobTemplateResponseBody
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *GetJobTemplateResponseBody
	GetGmtModifyTime() *string
	SetMetadata(v map[string]interface{}) *GetJobTemplateResponseBody
	GetMetadata() map[string]interface{}
	SetModifiedBy(v string) *GetJobTemplateResponseBody
	GetModifiedBy() *string
	SetRequestId(v string) *GetJobTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *GetJobTemplateResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *GetJobTemplateResponseBody
	GetTemplateName() *string
	SetTenantId(v string) *GetJobTemplateResponseBody
	GetTenantId() *string
	SetTotalCount(v int32) *GetJobTemplateResponseBody
	GetTotalCount() *int32
	SetUserId(v string) *GetJobTemplateResponseBody
	GetUserId() *string
	SetVersions(v []*GetJobTemplateResponseBodyVersions) *GetJobTemplateResponseBody
	GetVersions() []*GetJobTemplateResponseBodyVersions
	SetWorkspaceId(v string) *GetJobTemplateResponseBody
	GetWorkspaceId() *string
}

type GetJobTemplateResponseBody struct {
	// The default version of the job template.
	//
	// example:
	//
	// 2
	DefaultVersion *int32 `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// A description of the job template.
	//
	// example:
	//
	// job template description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time the job template was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-01-08T14:17:55Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time the job template was last modified.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-03-03T05:48:02Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// A collection of user-defined key-value pairs.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The ID of the user who last modified the job template.
	//
	// example:
	//
	// 20**************26
	ModifiedBy *string `json:"ModifiedBy,omitempty" xml:"ModifiedBy,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the job template.
	//
	// example:
	//
	// tplmceolmf2****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the job template.
	//
	// example:
	//
	// job-template-example-1778047****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The ID of the tenant that owns the job template.
	//
	// example:
	//
	// 142388383837****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The total number of versions returned. This value is 1 if a specific version is queried, or the total count if all versions are queried.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The ID of the user who created the job template.
	//
	// example:
	//
	// 20**************26
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// An array of template versions. This array contains only one version if a specific version is requested, or all versions if `all` is specified.
	Versions []*GetJobTemplateResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
	// The ID of the workspace that contains the job template.
	//
	// example:
	//
	// 4***9
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBody) GetDefaultVersion() *int32 {
	return s.DefaultVersion
}

func (s *GetJobTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetJobTemplateResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetJobTemplateResponseBody) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *GetJobTemplateResponseBody) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetJobTemplateResponseBody) GetModifiedBy() *string {
	return s.ModifiedBy
}

func (s *GetJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetJobTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetJobTemplateResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *GetJobTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetJobTemplateResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetJobTemplateResponseBody) GetVersions() []*GetJobTemplateResponseBodyVersions {
	return s.Versions
}

func (s *GetJobTemplateResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetJobTemplateResponseBody) SetDefaultVersion(v int32) *GetJobTemplateResponseBody {
	s.DefaultVersion = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetDescription(v string) *GetJobTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetGmtCreateTime(v string) *GetJobTemplateResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetGmtModifyTime(v string) *GetJobTemplateResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetMetadata(v map[string]interface{}) *GetJobTemplateResponseBody {
	s.Metadata = v
	return s
}

func (s *GetJobTemplateResponseBody) SetModifiedBy(v string) *GetJobTemplateResponseBody {
	s.ModifiedBy = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetRequestId(v string) *GetJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetTemplateId(v string) *GetJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetTemplateName(v string) *GetJobTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetTenantId(v string) *GetJobTemplateResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetTotalCount(v int32) *GetJobTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetUserId(v string) *GetJobTemplateResponseBody {
	s.UserId = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetVersions(v []*GetJobTemplateResponseBodyVersions) *GetJobTemplateResponseBody {
	s.Versions = v
	return s
}

func (s *GetJobTemplateResponseBody) SetWorkspaceId(v string) *GetJobTemplateResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetJobTemplateResponseBody) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobTemplateResponseBodyVersions struct {
	// The field constraint rules. The key is a JSONPath expression and the value is a constraint type.
	//
	// example:
	//
	// {\\"JobSpecs[0].Image\\":\\"locked\\",\\"UserCommand\\":\\"locked\\",\\"JobType\\":\\"locked\\"}
	Constraints map[string]interface{} `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The configuration of the version, in JSON format.
	//
	// example:
	//
	// {\\"WorkspaceId\\":\\"15****05\\",\\"JobType\\":\\"PyTorchJob\\",\\"UserCommand\\":\\"echo hello\\",\\"JobSpecs\\":[{\\"Type\\":\\"Worker\\",\\"PodCount\\":1,\\"Image\\":\\"dsw-registry-vpc.cn-hangzhou.cr.aliyuncs.com/pai/pytorch:2.8.0-gpu-py313-cu129-ubuntu22.04-3995b779-1764361782\\",\\"EcsSpec\\":\\"ecs.gn7i-c8g1.2xlarge\\"}],\\"ResourceType\\":\\"ECS\\",\\"_ResourcePaymentType\\":\\"PostPaid\\",\\"CredentialConfig\\":{\\"EnableCredentialInject\\":false},\\"Accessibility\\":\\"PRIVATE\\",\\"Settings\\":{\\"JobReservedMinutes\\":0,\\"Tags\\":{}}}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the user who created the version.
	//
	// example:
	//
	// 20**************26
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time the version was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-01-08T14:17:55Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetJobTemplateResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyVersions) GetConstraints() map[string]interface{} {
	return s.Constraints
}

func (s *GetJobTemplateResponseBodyVersions) GetContent() interface{} {
	return s.Content
}

func (s *GetJobTemplateResponseBodyVersions) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetJobTemplateResponseBodyVersions) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetJobTemplateResponseBodyVersions) GetVersion() *int32 {
	return s.Version
}

func (s *GetJobTemplateResponseBodyVersions) SetConstraints(v map[string]interface{}) *GetJobTemplateResponseBodyVersions {
	s.Constraints = v
	return s
}

func (s *GetJobTemplateResponseBodyVersions) SetContent(v interface{}) *GetJobTemplateResponseBodyVersions {
	s.Content = v
	return s
}

func (s *GetJobTemplateResponseBodyVersions) SetCreatedBy(v string) *GetJobTemplateResponseBodyVersions {
	s.CreatedBy = &v
	return s
}

func (s *GetJobTemplateResponseBodyVersions) SetGmtCreateTime(v string) *GetJobTemplateResponseBodyVersions {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobTemplateResponseBodyVersions) SetVersion(v int32) *GetJobTemplateResponseBodyVersions {
	s.Version = &v
	return s
}

func (s *GetJobTemplateResponseBodyVersions) Validate() error {
	return dara.Validate(s)
}
