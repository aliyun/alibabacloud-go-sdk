// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetPageResponseBody
	GetContent() *string
	SetContentType(v string) *GetPageResponseBody
	GetContentType() *string
	SetDescription(v string) *GetPageResponseBody
	GetDescription() *string
	SetId(v int64) *GetPageResponseBody
	GetId() *int64
	SetKind(v string) *GetPageResponseBody
	GetKind() *string
	SetName(v string) *GetPageResponseBody
	GetName() *string
	SetRequestId(v string) *GetPageResponseBody
	GetRequestId() *string
	SetUpdateTime(v string) *GetPageResponseBody
	GetUpdateTime() *string
}

type GetPageResponseBody struct {
	// The Base64-encoded content of the error page. The content type is specified by the Content-Type field.
	//
	// This parameter is required.
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The Content-Type field in the HTTP header.
	//
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the custom error page.
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom error page.[](~~2850223~~)
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the custom response page.
	//
	// example:
	//
	// custom
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The name of the custom response page.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the custom error page was last modified.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetPageResponseBody) GetContentType() *string {
	return s.ContentType
}

func (s *GetPageResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetPageResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetPageResponseBody) GetKind() *string {
	return s.Kind
}

func (s *GetPageResponseBody) GetName() *string {
	return s.Name
}

func (s *GetPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPageResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPageResponseBody) SetContent(v string) *GetPageResponseBody {
	s.Content = &v
	return s
}

func (s *GetPageResponseBody) SetContentType(v string) *GetPageResponseBody {
	s.ContentType = &v
	return s
}

func (s *GetPageResponseBody) SetDescription(v string) *GetPageResponseBody {
	s.Description = &v
	return s
}

func (s *GetPageResponseBody) SetId(v int64) *GetPageResponseBody {
	s.Id = &v
	return s
}

func (s *GetPageResponseBody) SetKind(v string) *GetPageResponseBody {
	s.Kind = &v
	return s
}

func (s *GetPageResponseBody) SetName(v string) *GetPageResponseBody {
	s.Name = &v
	return s
}

func (s *GetPageResponseBody) SetRequestId(v string) *GetPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageResponseBody) SetUpdateTime(v string) *GetPageResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetPageResponseBody) Validate() error {
	return dara.Validate(s)
}
