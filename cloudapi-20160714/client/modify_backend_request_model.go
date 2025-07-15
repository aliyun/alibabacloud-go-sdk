// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendId(v string) *ModifyBackendRequest
	GetBackendId() *string
	SetBackendName(v string) *ModifyBackendRequest
	GetBackendName() *string
	SetBackendType(v string) *ModifyBackendRequest
	GetBackendType() *string
	SetDescription(v string) *ModifyBackendRequest
	GetDescription() *string
	SetSecurityToken(v string) *ModifyBackendRequest
	GetSecurityToken() *string
}

type ModifyBackendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20bcdc9453524b78a8beb1f6de21edb7
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testHttpModify
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
	// example:
	//
	// test
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyBackendRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackendRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackendRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *ModifyBackendRequest) GetBackendName() *string {
	return s.BackendName
}

func (s *ModifyBackendRequest) GetBackendType() *string {
	return s.BackendType
}

func (s *ModifyBackendRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyBackendRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyBackendRequest) SetBackendId(v string) *ModifyBackendRequest {
	s.BackendId = &v
	return s
}

func (s *ModifyBackendRequest) SetBackendName(v string) *ModifyBackendRequest {
	s.BackendName = &v
	return s
}

func (s *ModifyBackendRequest) SetBackendType(v string) *ModifyBackendRequest {
	s.BackendType = &v
	return s
}

func (s *ModifyBackendRequest) SetDescription(v string) *ModifyBackendRequest {
	s.Description = &v
	return s
}

func (s *ModifyBackendRequest) SetSecurityToken(v string) *ModifyBackendRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyBackendRequest) Validate() error {
	return dara.Validate(s)
}
