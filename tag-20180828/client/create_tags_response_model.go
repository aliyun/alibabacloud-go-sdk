// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTagsResponse
	GetStatusCode() *int32
	SetBody(v *CreateTagsResponseBody) *CreateTagsResponse
	GetBody() *CreateTagsResponseBody
}

type CreateTagsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTagsResponse) GoString() string {
	return s.String()
}

func (s *CreateTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTagsResponse) GetBody() *CreateTagsResponseBody {
	return s.Body
}

func (s *CreateTagsResponse) SetHeaders(v map[string]*string) *CreateTagsResponse {
	s.Headers = v
	return s
}

func (s *CreateTagsResponse) SetStatusCode(v int32) *CreateTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagsResponse) SetBody(v *CreateTagsResponseBody) *CreateTagsResponse {
	s.Body = v
	return s
}

func (s *CreateTagsResponse) Validate() error {
	return dara.Validate(s)
}
