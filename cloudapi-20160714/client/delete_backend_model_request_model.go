// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackendModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendId(v string) *DeleteBackendModelRequest
	GetBackendId() *string
	SetBackendModelId(v string) *DeleteBackendModelRequest
	GetBackendModelId() *string
	SetSecurityToken(v string) *DeleteBackendModelRequest
	GetSecurityToken() *string
	SetStageName(v string) *DeleteBackendModelRequest
	GetStageName() *string
}

type DeleteBackendModelRequest struct {
	// The ID of the backend service.
	//
	// example:
	//
	// 20bcdc9453524b78a8beb1f6de21edb7
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// The ID of the backend model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4be6b110b7aa40b0bf0c83cc00b3bd86
	BackendModelId *string `json:"BackendModelId,omitempty" xml:"BackendModelId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **PRE**
	//
	// 	- **TEST**
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DeleteBackendModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackendModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackendModelRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *DeleteBackendModelRequest) GetBackendModelId() *string {
	return s.BackendModelId
}

func (s *DeleteBackendModelRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteBackendModelRequest) GetStageName() *string {
	return s.StageName
}

func (s *DeleteBackendModelRequest) SetBackendId(v string) *DeleteBackendModelRequest {
	s.BackendId = &v
	return s
}

func (s *DeleteBackendModelRequest) SetBackendModelId(v string) *DeleteBackendModelRequest {
	s.BackendModelId = &v
	return s
}

func (s *DeleteBackendModelRequest) SetSecurityToken(v string) *DeleteBackendModelRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteBackendModelRequest) SetStageName(v string) *DeleteBackendModelRequest {
	s.StageName = &v
	return s
}

func (s *DeleteBackendModelRequest) Validate() error {
	return dara.Validate(s)
}
