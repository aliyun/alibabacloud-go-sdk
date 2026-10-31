// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeByOutputNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetPhysicalNodeByOutputNameRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetPhysicalNodeByOutputNameRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetPhysicalNodeByOutputNameRequest
	GetOpUserId() *string
	SetOutputName(v string) *GetPhysicalNodeByOutputNameRequest
	GetOutputName() *string
}

type GetPhysicalNodeByOutputNameRequest struct {
	// The environment identifier. Valid values:
	//
	// - DEV: development environment
	//
	// - PROD: production environment. This is the default value.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The output name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// t_test
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
}

func (s GetPhysicalNodeByOutputNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameRequest) GetEnv() *string {
	return s.Env
}

func (s *GetPhysicalNodeByOutputNameRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPhysicalNodeByOutputNameRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetPhysicalNodeByOutputNameRequest) GetOutputName() *string {
	return s.OutputName
}

func (s *GetPhysicalNodeByOutputNameRequest) SetEnv(v string) *GetPhysicalNodeByOutputNameRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameRequest) SetOpTenantId(v int64) *GetPhysicalNodeByOutputNameRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameRequest) SetOpUserId(v string) *GetPhysicalNodeByOutputNameRequest {
	s.OpUserId = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameRequest) SetOutputName(v string) *GetPhysicalNodeByOutputNameRequest {
	s.OutputName = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameRequest) Validate() error {
	return dara.Validate(s)
}
