// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileImportTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFieldSelectionEnabled(v bool) *GenerateFileImportTemplateRequest
	GetFieldSelectionEnabled() *bool
	SetInstanceId(v string) *GenerateFileImportTemplateRequest
	GetInstanceId() *string
	SetSelectedFieldNames(v []*string) *GenerateFileImportTemplateRequest
	GetSelectedFieldNames() []*string
	SetTargetType(v string) *GenerateFileImportTemplateRequest
	GetTargetType() *string
}

type GenerateFileImportTemplateRequest struct {
	// Specifies whether field selection mode is enabled.
	FieldSelectionEnabled *bool `json:"FieldSelectionEnabled,omitempty" xml:"FieldSelectionEnabled,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The selected fields.
	SelectedFieldNames []*string `json:"SelectedFieldNames,omitempty" xml:"SelectedFieldNames,omitempty" type:"Repeated"`
	// The synchronization target type.
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam_v2_user_import,
	//
	// application, identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GenerateFileImportTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileImportTemplateRequest) GoString() string {
	return s.String()
}

func (s *GenerateFileImportTemplateRequest) GetFieldSelectionEnabled() *bool {
	return s.FieldSelectionEnabled
}

func (s *GenerateFileImportTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateFileImportTemplateRequest) GetSelectedFieldNames() []*string {
	return s.SelectedFieldNames
}

func (s *GenerateFileImportTemplateRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *GenerateFileImportTemplateRequest) SetFieldSelectionEnabled(v bool) *GenerateFileImportTemplateRequest {
	s.FieldSelectionEnabled = &v
	return s
}

func (s *GenerateFileImportTemplateRequest) SetInstanceId(v string) *GenerateFileImportTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateFileImportTemplateRequest) SetSelectedFieldNames(v []*string) *GenerateFileImportTemplateRequest {
	s.SelectedFieldNames = v
	return s
}

func (s *GenerateFileImportTemplateRequest) SetTargetType(v string) *GenerateFileImportTemplateRequest {
	s.TargetType = &v
	return s
}

func (s *GenerateFileImportTemplateRequest) Validate() error {
	return dara.Validate(s)
}
