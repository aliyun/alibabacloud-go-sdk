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
	SetOutputName(v string) *GetPhysicalNodeByOutputNameRequest
	GetOutputName() *string
}

type GetPhysicalNodeByOutputNameRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *GetPhysicalNodeByOutputNameRequest) SetOutputName(v string) *GetPhysicalNodeByOutputNameRequest {
	s.OutputName = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameRequest) Validate() error {
	return dara.Validate(s)
}
