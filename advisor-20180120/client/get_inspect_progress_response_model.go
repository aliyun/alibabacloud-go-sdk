// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInspectProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInspectProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetInspectProgressResponseBody) *GetInspectProgressResponse
	GetBody() *GetInspectProgressResponseBody
}

type GetInspectProgressResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInspectProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInspectProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInspectProgressResponse) GoString() string {
	return s.String()
}

func (s *GetInspectProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInspectProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInspectProgressResponse) GetBody() *GetInspectProgressResponseBody {
	return s.Body
}

func (s *GetInspectProgressResponse) SetHeaders(v map[string]*string) *GetInspectProgressResponse {
	s.Headers = v
	return s
}

func (s *GetInspectProgressResponse) SetStatusCode(v int32) *GetInspectProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInspectProgressResponse) SetBody(v *GetInspectProgressResponseBody) *GetInspectProgressResponse {
	s.Body = v
	return s
}

func (s *GetInspectProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
