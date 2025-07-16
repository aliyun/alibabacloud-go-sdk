// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocUpdateContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DocUpdateContentRequest
	GetContent() *string
	SetDataType(v string) *DocUpdateContentRequest
	GetDataType() *string
	SetDocKey(v string) *DocUpdateContentRequest
	GetDocKey() *string
	SetTenantContext(v *DocUpdateContentRequestTenantContext) *DocUpdateContentRequest
	GetTenantContext() *DocUpdateContentRequestTenantContext
}

type DocUpdateContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// []
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// markdown
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4j6OJzVEG4jJO3p8
	DocKey        *string                               `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	TenantContext *DocUpdateContentRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DocUpdateContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DocUpdateContentRequest) GoString() string {
	return s.String()
}

func (s *DocUpdateContentRequest) GetContent() *string {
	return s.Content
}

func (s *DocUpdateContentRequest) GetDataType() *string {
	return s.DataType
}

func (s *DocUpdateContentRequest) GetDocKey() *string {
	return s.DocKey
}

func (s *DocUpdateContentRequest) GetTenantContext() *DocUpdateContentRequestTenantContext {
	return s.TenantContext
}

func (s *DocUpdateContentRequest) SetContent(v string) *DocUpdateContentRequest {
	s.Content = &v
	return s
}

func (s *DocUpdateContentRequest) SetDataType(v string) *DocUpdateContentRequest {
	s.DataType = &v
	return s
}

func (s *DocUpdateContentRequest) SetDocKey(v string) *DocUpdateContentRequest {
	s.DocKey = &v
	return s
}

func (s *DocUpdateContentRequest) SetTenantContext(v *DocUpdateContentRequestTenantContext) *DocUpdateContentRequest {
	s.TenantContext = v
	return s
}

func (s *DocUpdateContentRequest) Validate() error {
	return dara.Validate(s)
}

type DocUpdateContentRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DocUpdateContentRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DocUpdateContentRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DocUpdateContentRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DocUpdateContentRequestTenantContext) SetTenantId(v string) *DocUpdateContentRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DocUpdateContentRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
