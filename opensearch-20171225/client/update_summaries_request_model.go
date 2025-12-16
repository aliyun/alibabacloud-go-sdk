// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*UpdateSummariesRequestBody) *UpdateSummariesRequest
	GetBody() []*UpdateSummariesRequestBody
	SetDryRun(v bool) *UpdateSummariesRequest
	GetDryRun() *bool
}

type UpdateSummariesRequest struct {
	// The request body.
	Body []*UpdateSummariesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSummariesRequest) GoString() string {
	return s.String()
}

func (s *UpdateSummariesRequest) GetBody() []*UpdateSummariesRequestBody {
	return s.Body
}

func (s *UpdateSummariesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateSummariesRequest) SetBody(v []*UpdateSummariesRequestBody) *UpdateSummariesRequest {
	s.Body = v
	return s
}

func (s *UpdateSummariesRequest) SetDryRun(v bool) *UpdateSummariesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSummariesRequest) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateSummariesRequestBody struct {
	// The HTML tag that is used to highlight terms in red.
	//
	// example:
	//
	// "em"
	Element *string `json:"element,omitempty" xml:"element,omitempty"`
	// The connector that is used to connect segments.
	//
	// example:
	//
	// "..."
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	// The field.
	//
	// example:
	//
	// "title"
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The length of a segment.
	//
	// example:
	//
	// 50
	Len *int32 `json:"len,omitempty" xml:"len,omitempty"`
	// The number of segments.
	//
	// example:
	//
	// 1
	Snippet *int32 `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s UpdateSummariesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSummariesRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateSummariesRequestBody) GetElement() *string {
	return s.Element
}

func (s *UpdateSummariesRequestBody) GetEllipsis() *string {
	return s.Ellipsis
}

func (s *UpdateSummariesRequestBody) GetField() *string {
	return s.Field
}

func (s *UpdateSummariesRequestBody) GetLen() *int32 {
	return s.Len
}

func (s *UpdateSummariesRequestBody) GetSnippet() *int32 {
	return s.Snippet
}

func (s *UpdateSummariesRequestBody) SetElement(v string) *UpdateSummariesRequestBody {
	s.Element = &v
	return s
}

func (s *UpdateSummariesRequestBody) SetEllipsis(v string) *UpdateSummariesRequestBody {
	s.Ellipsis = &v
	return s
}

func (s *UpdateSummariesRequestBody) SetField(v string) *UpdateSummariesRequestBody {
	s.Field = &v
	return s
}

func (s *UpdateSummariesRequestBody) SetLen(v int32) *UpdateSummariesRequestBody {
	s.Len = &v
	return s
}

func (s *UpdateSummariesRequestBody) SetSnippet(v int32) *UpdateSummariesRequestBody {
	s.Snippet = &v
	return s
}

func (s *UpdateSummariesRequestBody) Validate() error {
	return dara.Validate(s)
}
