// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreatePageRequest
	GetContent() *string
	SetContentType(v string) *CreatePageRequest
	GetContentType() *string
	SetDescription(v string) *CreatePageRequest
	GetDescription() *string
	SetName(v string) *CreatePageRequest
	GetName() *string
}

type CreatePageRequest struct {
	// The Base64-encoded page content. Example: "PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=", which indicates "hello page".
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
	// The description of the page.
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom error page.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePageRequest) GoString() string {
	return s.String()
}

func (s *CreatePageRequest) GetContent() *string {
	return s.Content
}

func (s *CreatePageRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreatePageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePageRequest) GetName() *string {
	return s.Name
}

func (s *CreatePageRequest) SetContent(v string) *CreatePageRequest {
	s.Content = &v
	return s
}

func (s *CreatePageRequest) SetContentType(v string) *CreatePageRequest {
	s.ContentType = &v
	return s
}

func (s *CreatePageRequest) SetDescription(v string) *CreatePageRequest {
	s.Description = &v
	return s
}

func (s *CreatePageRequest) SetName(v string) *CreatePageRequest {
	s.Name = &v
	return s
}

func (s *CreatePageRequest) Validate() error {
	return dara.Validate(s)
}
