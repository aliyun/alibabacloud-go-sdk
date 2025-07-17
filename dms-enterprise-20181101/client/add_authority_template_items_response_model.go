// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthorityTemplateItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAuthorityTemplateItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAuthorityTemplateItemsResponse
	GetStatusCode() *int32
	SetBody(v *AddAuthorityTemplateItemsResponseBody) *AddAuthorityTemplateItemsResponse
	GetBody() *AddAuthorityTemplateItemsResponseBody
}

type AddAuthorityTemplateItemsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAuthorityTemplateItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAuthorityTemplateItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAuthorityTemplateItemsResponse) GoString() string {
	return s.String()
}

func (s *AddAuthorityTemplateItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAuthorityTemplateItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAuthorityTemplateItemsResponse) GetBody() *AddAuthorityTemplateItemsResponseBody {
	return s.Body
}

func (s *AddAuthorityTemplateItemsResponse) SetHeaders(v map[string]*string) *AddAuthorityTemplateItemsResponse {
	s.Headers = v
	return s
}

func (s *AddAuthorityTemplateItemsResponse) SetStatusCode(v int32) *AddAuthorityTemplateItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAuthorityTemplateItemsResponse) SetBody(v *AddAuthorityTemplateItemsResponseBody) *AddAuthorityTemplateItemsResponse {
	s.Body = v
	return s
}

func (s *AddAuthorityTemplateItemsResponse) Validate() error {
	return dara.Validate(s)
}
