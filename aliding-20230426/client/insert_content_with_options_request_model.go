// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertContentWithOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *InsertContentWithOptionsRequest
	GetContent() map[string]interface{}
	SetDocumentId(v string) *InsertContentWithOptionsRequest
	GetDocumentId() *string
	SetIndex(v int32) *InsertContentWithOptionsRequest
	GetIndex() *int32
	SetPath(v []*int32) *InsertContentWithOptionsRequest
	GetPath() []*int32
	SetTenantContext(v *InsertContentWithOptionsRequestTenantContext) *InsertContentWithOptionsRequest
	GetTenantContext() *InsertContentWithOptionsRequestTenantContext
}

type InsertContentWithOptionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// content
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// documentId
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// [0,0]
	Path          []*int32                                      `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	TenantContext *InsertContentWithOptionsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s InsertContentWithOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertContentWithOptionsRequest) GoString() string {
	return s.String()
}

func (s *InsertContentWithOptionsRequest) GetContent() map[string]interface{} {
	return s.Content
}

func (s *InsertContentWithOptionsRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *InsertContentWithOptionsRequest) GetIndex() *int32 {
	return s.Index
}

func (s *InsertContentWithOptionsRequest) GetPath() []*int32 {
	return s.Path
}

func (s *InsertContentWithOptionsRequest) GetTenantContext() *InsertContentWithOptionsRequestTenantContext {
	return s.TenantContext
}

func (s *InsertContentWithOptionsRequest) SetContent(v map[string]interface{}) *InsertContentWithOptionsRequest {
	s.Content = v
	return s
}

func (s *InsertContentWithOptionsRequest) SetDocumentId(v string) *InsertContentWithOptionsRequest {
	s.DocumentId = &v
	return s
}

func (s *InsertContentWithOptionsRequest) SetIndex(v int32) *InsertContentWithOptionsRequest {
	s.Index = &v
	return s
}

func (s *InsertContentWithOptionsRequest) SetPath(v []*int32) *InsertContentWithOptionsRequest {
	s.Path = v
	return s
}

func (s *InsertContentWithOptionsRequest) SetTenantContext(v *InsertContentWithOptionsRequestTenantContext) *InsertContentWithOptionsRequest {
	s.TenantContext = v
	return s
}

func (s *InsertContentWithOptionsRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertContentWithOptionsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InsertContentWithOptionsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s InsertContentWithOptionsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InsertContentWithOptionsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *InsertContentWithOptionsRequestTenantContext) SetTenantId(v string) *InsertContentWithOptionsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *InsertContentWithOptionsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
