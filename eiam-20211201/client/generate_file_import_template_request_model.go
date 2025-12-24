// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileImportTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GenerateFileImportTemplateRequest
	GetInstanceId() *string
	SetTargetType(v string) *GenerateFileImportTemplateRequest
	GetTargetType() *string
}

type GenerateFileImportTemplateRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 同步目标类型
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

func (s *GenerateFileImportTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateFileImportTemplateRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *GenerateFileImportTemplateRequest) SetInstanceId(v string) *GenerateFileImportTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateFileImportTemplateRequest) SetTargetType(v string) *GenerateFileImportTemplateRequest {
	s.TargetType = &v
	return s
}

func (s *GenerateFileImportTemplateRequest) Validate() error {
	return dara.Validate(s)
}
