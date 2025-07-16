// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocUpdateContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DocUpdateContentShrinkRequest
	GetContent() *string
	SetDataType(v string) *DocUpdateContentShrinkRequest
	GetDataType() *string
	SetDocKey(v string) *DocUpdateContentShrinkRequest
	GetDocKey() *string
	SetTenantContextShrink(v string) *DocUpdateContentShrinkRequest
	GetTenantContextShrink() *string
}

type DocUpdateContentShrinkRequest struct {
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
	DocKey              *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DocUpdateContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DocUpdateContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *DocUpdateContentShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *DocUpdateContentShrinkRequest) GetDataType() *string {
	return s.DataType
}

func (s *DocUpdateContentShrinkRequest) GetDocKey() *string {
	return s.DocKey
}

func (s *DocUpdateContentShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DocUpdateContentShrinkRequest) SetContent(v string) *DocUpdateContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *DocUpdateContentShrinkRequest) SetDataType(v string) *DocUpdateContentShrinkRequest {
	s.DataType = &v
	return s
}

func (s *DocUpdateContentShrinkRequest) SetDocKey(v string) *DocUpdateContentShrinkRequest {
	s.DocKey = &v
	return s
}

func (s *DocUpdateContentShrinkRequest) SetTenantContextShrink(v string) *DocUpdateContentShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DocUpdateContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
