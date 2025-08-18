// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdatePageRequest
	GetContent() *string
	SetContentType(v string) *UpdatePageRequest
	GetContentType() *string
	SetDescription(v string) *UpdatePageRequest
	GetDescription() *string
	SetId(v int64) *UpdatePageRequest
	GetId() *int64
	SetName(v string) *UpdatePageRequest
	GetName() *string
}

type UpdatePageRequest struct {
	// The Base64-encoded content of the error page. The content type is specified by the Content-Type field.
	//
	// This parameter is required.
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The Content-Type field in the HTTP header. Valid values:
	//
	// 	- text/html
	//
	// 	- application/json
	//
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the custom error page.
	//
	// This parameter is required.
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom error page, which can be obtained by calling the [ListPages](https://help.aliyun.com/document_detail/2850223.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the custom error page.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdatePageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePageRequest) GoString() string {
	return s.String()
}

func (s *UpdatePageRequest) GetContent() *string {
	return s.Content
}

func (s *UpdatePageRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UpdatePageRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePageRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdatePageRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePageRequest) SetContent(v string) *UpdatePageRequest {
	s.Content = &v
	return s
}

func (s *UpdatePageRequest) SetContentType(v string) *UpdatePageRequest {
	s.ContentType = &v
	return s
}

func (s *UpdatePageRequest) SetDescription(v string) *UpdatePageRequest {
	s.Description = &v
	return s
}

func (s *UpdatePageRequest) SetId(v int64) *UpdatePageRequest {
	s.Id = &v
	return s
}

func (s *UpdatePageRequest) SetName(v string) *UpdatePageRequest {
	s.Name = &v
	return s
}

func (s *UpdatePageRequest) Validate() error {
	return dara.Validate(s)
}
