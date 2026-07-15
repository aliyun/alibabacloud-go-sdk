// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultVersion(v int32) *CreateJobTemplateResponseBody
	GetDefaultVersion() *int32
	SetDescription(v string) *CreateJobTemplateResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *CreateJobTemplateResponseBody
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *CreateJobTemplateResponseBody
	GetGmtModifyTime() *string
	SetMetadata(v map[string]interface{}) *CreateJobTemplateResponseBody
	GetMetadata() map[string]interface{}
	SetRequestId(v string) *CreateJobTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateJobTemplateResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *CreateJobTemplateResponseBody
	GetTemplateName() *string
	SetTenantId(v string) *CreateJobTemplateResponseBody
	GetTenantId() *string
	SetUserId(v string) *CreateJobTemplateResponseBody
	GetUserId() *string
	SetVersion(v int32) *CreateJobTemplateResponseBody
	GetVersion() *int32
	SetWorkspaceId(v string) *CreateJobTemplateResponseBody
	GetWorkspaceId() *string
}

type CreateJobTemplateResponseBody struct {
	// The default version number of the job template.
	//
	// example:
	//
	// 1
	DefaultVersion *int32 `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the job template.
	//
	// example:
	//
	// Template description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time of the job template.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-12-31T02:18:09Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modification time of the job template.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// User-defined key-value metadata.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The ID of the request, used for troubleshooting.
	//
	// example:
	//
	// 8762921A-911C-515F-A3A4-*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unique ID of the job template.
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
	// The ID of the tenant that contains the job template.
	//
	// example:
	//
	// 10**************14
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the user who created the job template.
	//
	// example:
	//
	// 20**************02
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The version number of the created job template.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the workspace that contains the job template.
	//
	// example:
	//
	// 15****05
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponseBody) GetDefaultVersion() *int32 {
	return s.DefaultVersion
}

func (s *CreateJobTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateJobTemplateResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *CreateJobTemplateResponseBody) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *CreateJobTemplateResponseBody) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *CreateJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateJobTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateJobTemplateResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateJobTemplateResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *CreateJobTemplateResponseBody) GetVersion() *int32 {
	return s.Version
}

func (s *CreateJobTemplateResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateJobTemplateResponseBody) SetDefaultVersion(v int32) *CreateJobTemplateResponseBody {
	s.DefaultVersion = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetDescription(v string) *CreateJobTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetGmtCreateTime(v string) *CreateJobTemplateResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetGmtModifyTime(v string) *CreateJobTemplateResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetMetadata(v map[string]interface{}) *CreateJobTemplateResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateJobTemplateResponseBody) SetRequestId(v string) *CreateJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetTemplateId(v string) *CreateJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetTemplateName(v string) *CreateJobTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetTenantId(v string) *CreateJobTemplateResponseBody {
	s.TenantId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetUserId(v string) *CreateJobTemplateResponseBody {
	s.UserId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetVersion(v int32) *CreateJobTemplateResponseBody {
	s.Version = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetWorkspaceId(v string) *CreateJobTemplateResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
