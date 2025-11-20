// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContents(v []*SaveContentRequestContents) *SaveContentRequest
	GetContents() []*SaveContentRequestContents
	SetDdFrom(v string) *SaveContentRequest
	GetDdFrom() *string
	SetTemplateId(v string) *SaveContentRequest
	GetTemplateId() *string
	SetTenantContext(v *SaveContentRequestTenantContext) *SaveContentRequest
	GetTenantContext() *SaveContentRequestTenantContext
}

type SaveContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// []
	Contents []*SaveContentRequestContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// client
	DdFrom *string `json:"DdFrom,omitempty" xml:"DdFrom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sdfafdsfsafdfsaf
	TemplateId    *string                          `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantContext *SaveContentRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s SaveContentRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveContentRequest) GoString() string {
	return s.String()
}

func (s *SaveContentRequest) GetContents() []*SaveContentRequestContents {
	return s.Contents
}

func (s *SaveContentRequest) GetDdFrom() *string {
	return s.DdFrom
}

func (s *SaveContentRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SaveContentRequest) GetTenantContext() *SaveContentRequestTenantContext {
	return s.TenantContext
}

func (s *SaveContentRequest) SetContents(v []*SaveContentRequestContents) *SaveContentRequest {
	s.Contents = v
	return s
}

func (s *SaveContentRequest) SetDdFrom(v string) *SaveContentRequest {
	s.DdFrom = &v
	return s
}

func (s *SaveContentRequest) SetTemplateId(v string) *SaveContentRequest {
	s.TemplateId = &v
	return s
}

func (s *SaveContentRequest) SetTenantContext(v *SaveContentRequestTenantContext) *SaveContentRequest {
	s.TenantContext = v
	return s
}

func (s *SaveContentRequest) Validate() error {
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveContentRequestContents struct {
	// This parameter is required.
	//
	// example:
	//
	// ### 序号1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// markdown
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 今日完成工作
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Sort *int64 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SaveContentRequestContents) String() string {
	return dara.Prettify(s)
}

func (s SaveContentRequestContents) GoString() string {
	return s.String()
}

func (s *SaveContentRequestContents) GetContent() *string {
	return s.Content
}

func (s *SaveContentRequestContents) GetContentType() *string {
	return s.ContentType
}

func (s *SaveContentRequestContents) GetKey() *string {
	return s.Key
}

func (s *SaveContentRequestContents) GetSort() *int64 {
	return s.Sort
}

func (s *SaveContentRequestContents) GetType() *int64 {
	return s.Type
}

func (s *SaveContentRequestContents) SetContent(v string) *SaveContentRequestContents {
	s.Content = &v
	return s
}

func (s *SaveContentRequestContents) SetContentType(v string) *SaveContentRequestContents {
	s.ContentType = &v
	return s
}

func (s *SaveContentRequestContents) SetKey(v string) *SaveContentRequestContents {
	s.Key = &v
	return s
}

func (s *SaveContentRequestContents) SetSort(v int64) *SaveContentRequestContents {
	s.Sort = &v
	return s
}

func (s *SaveContentRequestContents) SetType(v int64) *SaveContentRequestContents {
	s.Type = &v
	return s
}

func (s *SaveContentRequestContents) Validate() error {
	return dara.Validate(s)
}

type SaveContentRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SaveContentRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SaveContentRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SaveContentRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveContentRequestTenantContext) SetTenantId(v string) *SaveContentRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SaveContentRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
