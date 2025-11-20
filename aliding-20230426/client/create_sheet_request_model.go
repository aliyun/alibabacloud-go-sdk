// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSheetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSheetRequest
	GetName() *string
	SetTenantContext(v *CreateSheetRequestTenantContext) *CreateSheetRequest
	GetTenantContext() *CreateSheetRequestTenantContext
	SetWorkbookId(v string) *CreateSheetRequest
	GetWorkbookId() *string
}

type CreateSheetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	Name          *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContext *CreateSheetRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s CreateSheetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSheetRequest) GoString() string {
	return s.String()
}

func (s *CreateSheetRequest) GetName() *string {
	return s.Name
}

func (s *CreateSheetRequest) GetTenantContext() *CreateSheetRequestTenantContext {
	return s.TenantContext
}

func (s *CreateSheetRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *CreateSheetRequest) SetName(v string) *CreateSheetRequest {
	s.Name = &v
	return s
}

func (s *CreateSheetRequest) SetTenantContext(v *CreateSheetRequestTenantContext) *CreateSheetRequest {
	s.TenantContext = v
	return s
}

func (s *CreateSheetRequest) SetWorkbookId(v string) *CreateSheetRequest {
	s.WorkbookId = &v
	return s
}

func (s *CreateSheetRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSheetRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateSheetRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateSheetRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateSheetRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateSheetRequestTenantContext) SetTenantId(v string) *CreateSheetRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateSheetRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
