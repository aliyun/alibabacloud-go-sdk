// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEditingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchEditingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchEditingProjectResponse
	GetStatusCode() *int32
	SetBody(v *SearchEditingProjectResponseBody) *SearchEditingProjectResponse
	GetBody() *SearchEditingProjectResponseBody
}

type SearchEditingProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchEditingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchEditingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchEditingProjectResponse) GetBody() *SearchEditingProjectResponseBody {
	return s.Body
}

func (s *SearchEditingProjectResponse) SetHeaders(v map[string]*string) *SearchEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *SearchEditingProjectResponse) SetStatusCode(v int32) *SearchEditingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchEditingProjectResponse) SetBody(v *SearchEditingProjectResponseBody) *SearchEditingProjectResponse {
	s.Body = v
	return s
}

func (s *SearchEditingProjectResponse) Validate() error {
	return dara.Validate(s)
}
