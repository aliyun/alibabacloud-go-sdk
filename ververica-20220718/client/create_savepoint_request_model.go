// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavepointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *CreateSavepointRequest
	GetDeploymentId() *string
	SetDescription(v string) *CreateSavepointRequest
	GetDescription() *string
	SetNativeFormat(v bool) *CreateSavepointRequest
	GetNativeFormat() *bool
}

type CreateSavepointRequest struct {
	// The deployment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The description of the savepoint.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to use the native format mode. Valid values:
	//
	// 	- true: The native format mode is used.
	//
	// 	- false: The native format mode is not used.
	//
	// example:
	//
	// true
	NativeFormat *bool `json:"nativeFormat,omitempty" xml:"nativeFormat,omitempty"`
}

func (s CreateSavepointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSavepointRequest) GoString() string {
	return s.String()
}

func (s *CreateSavepointRequest) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *CreateSavepointRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSavepointRequest) GetNativeFormat() *bool {
	return s.NativeFormat
}

func (s *CreateSavepointRequest) SetDeploymentId(v string) *CreateSavepointRequest {
	s.DeploymentId = &v
	return s
}

func (s *CreateSavepointRequest) SetDescription(v string) *CreateSavepointRequest {
	s.Description = &v
	return s
}

func (s *CreateSavepointRequest) SetNativeFormat(v bool) *CreateSavepointRequest {
	s.NativeFormat = &v
	return s
}

func (s *CreateSavepointRequest) Validate() error {
	return dara.Validate(s)
}
