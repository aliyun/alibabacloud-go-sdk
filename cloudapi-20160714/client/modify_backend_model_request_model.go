// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackendModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendId(v string) *ModifyBackendModelRequest
	GetBackendId() *string
	SetBackendModelData(v string) *ModifyBackendModelRequest
	GetBackendModelData() *string
	SetBackendModelId(v string) *ModifyBackendModelRequest
	GetBackendModelId() *string
	SetBackendType(v string) *ModifyBackendModelRequest
	GetBackendType() *string
	SetDescription(v string) *ModifyBackendModelRequest
	GetDescription() *string
	SetSecurityToken(v string) *ModifyBackendModelRequest
	GetSecurityToken() *string
	SetStageName(v string) *ModifyBackendModelRequest
	GetStageName() *string
}

type ModifyBackendModelRequest struct {
	// example:
	//
	// 20bcdc9453524b78a8beb1f6de21edb7
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// example:
	//
	// {\\"ServiceAddress\\":\\"http://121.40.XX.XX\\"}
	BackendModelData *string `json:"BackendModelData,omitempty" xml:"BackendModelData,omitempty"`
	// example:
	//
	// 3bb6375bc71c4e4c95ce05b4e7a55a9d
	BackendModelId *string `json:"BackendModelId,omitempty" xml:"BackendModelId,omitempty"`
	// example:
	//
	// OSS
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
	// example:
	//
	// modify plugin first
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s ModifyBackendModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackendModelRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackendModelRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *ModifyBackendModelRequest) GetBackendModelData() *string {
	return s.BackendModelData
}

func (s *ModifyBackendModelRequest) GetBackendModelId() *string {
	return s.BackendModelId
}

func (s *ModifyBackendModelRequest) GetBackendType() *string {
	return s.BackendType
}

func (s *ModifyBackendModelRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyBackendModelRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyBackendModelRequest) GetStageName() *string {
	return s.StageName
}

func (s *ModifyBackendModelRequest) SetBackendId(v string) *ModifyBackendModelRequest {
	s.BackendId = &v
	return s
}

func (s *ModifyBackendModelRequest) SetBackendModelData(v string) *ModifyBackendModelRequest {
	s.BackendModelData = &v
	return s
}

func (s *ModifyBackendModelRequest) SetBackendModelId(v string) *ModifyBackendModelRequest {
	s.BackendModelId = &v
	return s
}

func (s *ModifyBackendModelRequest) SetBackendType(v string) *ModifyBackendModelRequest {
	s.BackendType = &v
	return s
}

func (s *ModifyBackendModelRequest) SetDescription(v string) *ModifyBackendModelRequest {
	s.Description = &v
	return s
}

func (s *ModifyBackendModelRequest) SetSecurityToken(v string) *ModifyBackendModelRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyBackendModelRequest) SetStageName(v string) *ModifyBackendModelRequest {
	s.StageName = &v
	return s
}

func (s *ModifyBackendModelRequest) Validate() error {
	return dara.Validate(s)
}
