// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateStoryResponseBody) *CreateStoryResponse
	GetBody() *CreateStoryResponseBody
}

type CreateStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStoryResponse) GoString() string {
	return s.String()
}

func (s *CreateStoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStoryResponse) GetBody() *CreateStoryResponseBody {
	return s.Body
}

func (s *CreateStoryResponse) SetHeaders(v map[string]*string) *CreateStoryResponse {
	s.Headers = v
	return s
}

func (s *CreateStoryResponse) SetStatusCode(v int32) *CreateStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStoryResponse) SetBody(v *CreateStoryResponseBody) *CreateStoryResponse {
	s.Body = v
	return s
}

func (s *CreateStoryResponse) Validate() error {
	return dara.Validate(s)
}
