// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContents(v []*CreateReportRequestContents) *CreateReportRequest
	GetContents() []*CreateReportRequestContents
	SetDdFrom(v string) *CreateReportRequest
	GetDdFrom() *string
	SetTemplateId(v string) *CreateReportRequest
	GetTemplateId() *string
	SetTenantContext(v *CreateReportRequestTenantContext) *CreateReportRequest
	GetTenantContext() *CreateReportRequestTenantContext
	SetToChat(v bool) *CreateReportRequest
	GetToChat() *bool
	SetToCids(v []*string) *CreateReportRequest
	GetToCids() []*string
	SetToUserids(v []*string) *CreateReportRequest
	GetToUserids() []*string
}

type CreateReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// []
	Contents []*CreateReportRequestContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
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
	TemplateId    *string                           `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantContext *CreateReportRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	ToChat *bool `json:"ToChat,omitempty" xml:"ToChat,omitempty"`
	// example:
	//
	// []
	ToCids []*string `json:"ToCids,omitempty" xml:"ToCids,omitempty" type:"Repeated"`
	// example:
	//
	// [123,456]
	ToUserids []*string `json:"ToUserids,omitempty" xml:"ToUserids,omitempty" type:"Repeated"`
}

func (s CreateReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReportRequest) GoString() string {
	return s.String()
}

func (s *CreateReportRequest) GetContents() []*CreateReportRequestContents {
	return s.Contents
}

func (s *CreateReportRequest) GetDdFrom() *string {
	return s.DdFrom
}

func (s *CreateReportRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateReportRequest) GetTenantContext() *CreateReportRequestTenantContext {
	return s.TenantContext
}

func (s *CreateReportRequest) GetToChat() *bool {
	return s.ToChat
}

func (s *CreateReportRequest) GetToCids() []*string {
	return s.ToCids
}

func (s *CreateReportRequest) GetToUserids() []*string {
	return s.ToUserids
}

func (s *CreateReportRequest) SetContents(v []*CreateReportRequestContents) *CreateReportRequest {
	s.Contents = v
	return s
}

func (s *CreateReportRequest) SetDdFrom(v string) *CreateReportRequest {
	s.DdFrom = &v
	return s
}

func (s *CreateReportRequest) SetTemplateId(v string) *CreateReportRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateReportRequest) SetTenantContext(v *CreateReportRequestTenantContext) *CreateReportRequest {
	s.TenantContext = v
	return s
}

func (s *CreateReportRequest) SetToChat(v bool) *CreateReportRequest {
	s.ToChat = &v
	return s
}

func (s *CreateReportRequest) SetToCids(v []*string) *CreateReportRequest {
	s.ToCids = v
	return s
}

func (s *CreateReportRequest) SetToUserids(v []*string) *CreateReportRequest {
	s.ToUserids = v
	return s
}

func (s *CreateReportRequest) Validate() error {
	return dara.Validate(s)
}

type CreateReportRequestContents struct {
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
	// 0
	Sort *int64 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateReportRequestContents) String() string {
	return dara.Prettify(s)
}

func (s CreateReportRequestContents) GoString() string {
	return s.String()
}

func (s *CreateReportRequestContents) GetContent() *string {
	return s.Content
}

func (s *CreateReportRequestContents) GetContentType() *string {
	return s.ContentType
}

func (s *CreateReportRequestContents) GetKey() *string {
	return s.Key
}

func (s *CreateReportRequestContents) GetSort() *int64 {
	return s.Sort
}

func (s *CreateReportRequestContents) GetType() *int64 {
	return s.Type
}

func (s *CreateReportRequestContents) SetContent(v string) *CreateReportRequestContents {
	s.Content = &v
	return s
}

func (s *CreateReportRequestContents) SetContentType(v string) *CreateReportRequestContents {
	s.ContentType = &v
	return s
}

func (s *CreateReportRequestContents) SetKey(v string) *CreateReportRequestContents {
	s.Key = &v
	return s
}

func (s *CreateReportRequestContents) SetSort(v int64) *CreateReportRequestContents {
	s.Sort = &v
	return s
}

func (s *CreateReportRequestContents) SetType(v int64) *CreateReportRequestContents {
	s.Type = &v
	return s
}

func (s *CreateReportRequestContents) Validate() error {
	return dara.Validate(s)
}

type CreateReportRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateReportRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateReportRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateReportRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateReportRequestTenantContext) SetTenantId(v string) *CreateReportRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateReportRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
