// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeptNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetDeptNoRequestTenantContext) *GetDeptNoRequest
	GetTenantContext() *GetDeptNoRequestTenantContext
	SetDeptId(v string) *GetDeptNoRequest
	GetDeptId() *string
}

type GetDeptNoRequest struct {
	TenantContext *GetDeptNoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s GetDeptNoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeptNoRequest) GoString() string {
	return s.String()
}

func (s *GetDeptNoRequest) GetTenantContext() *GetDeptNoRequestTenantContext {
	return s.TenantContext
}

func (s *GetDeptNoRequest) GetDeptId() *string {
	return s.DeptId
}

func (s *GetDeptNoRequest) SetTenantContext(v *GetDeptNoRequestTenantContext) *GetDeptNoRequest {
	s.TenantContext = v
	return s
}

func (s *GetDeptNoRequest) SetDeptId(v string) *GetDeptNoRequest {
	s.DeptId = &v
	return s
}

func (s *GetDeptNoRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeptNoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDeptNoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDeptNoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDeptNoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDeptNoRequestTenantContext) SetTenantId(v string) *GetDeptNoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDeptNoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
