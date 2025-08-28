// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCallProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCallProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetCallProgressResponseBody) *GetCallProgressResponse
	GetBody() *GetCallProgressResponseBody
}

type GetCallProgressResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCallProgressResponse) GoString() string {
	return s.String()
}

func (s *GetCallProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCallProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCallProgressResponse) GetBody() *GetCallProgressResponseBody {
	return s.Body
}

func (s *GetCallProgressResponse) SetHeaders(v map[string]*string) *GetCallProgressResponse {
	s.Headers = v
	return s
}

func (s *GetCallProgressResponse) SetStatusCode(v int32) *GetCallProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallProgressResponse) SetBody(v *GetCallProgressResponseBody) *GetCallProgressResponse {
	s.Body = v
	return s
}

func (s *GetCallProgressResponse) Validate() error {
	return dara.Validate(s)
}
