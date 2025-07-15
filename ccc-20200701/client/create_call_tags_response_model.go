// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCallTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCallTagsResponse
	GetStatusCode() *int32
	SetBody(v *CreateCallTagsResponseBody) *CreateCallTagsResponse
	GetBody() *CreateCallTagsResponseBody
}

type CreateCallTagsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCallTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCallTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCallTagsResponse) GoString() string {
	return s.String()
}

func (s *CreateCallTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCallTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCallTagsResponse) GetBody() *CreateCallTagsResponseBody {
	return s.Body
}

func (s *CreateCallTagsResponse) SetHeaders(v map[string]*string) *CreateCallTagsResponse {
	s.Headers = v
	return s
}

func (s *CreateCallTagsResponse) SetStatusCode(v int32) *CreateCallTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCallTagsResponse) SetBody(v *CreateCallTagsResponseBody) *CreateCallTagsResponse {
	s.Body = v
	return s
}

func (s *CreateCallTagsResponse) Validate() error {
	return dara.Validate(s)
}
