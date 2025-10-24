// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrustedProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrustedProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrustedProjectsResponse
	GetStatusCode() *int32
	SetBody(v *GetTrustedProjectsResponseBody) *GetTrustedProjectsResponse
	GetBody() *GetTrustedProjectsResponseBody
}

type GetTrustedProjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrustedProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrustedProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrustedProjectsResponse) GoString() string {
	return s.String()
}

func (s *GetTrustedProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrustedProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrustedProjectsResponse) GetBody() *GetTrustedProjectsResponseBody {
	return s.Body
}

func (s *GetTrustedProjectsResponse) SetHeaders(v map[string]*string) *GetTrustedProjectsResponse {
	s.Headers = v
	return s
}

func (s *GetTrustedProjectsResponse) SetStatusCode(v int32) *GetTrustedProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrustedProjectsResponse) SetBody(v *GetTrustedProjectsResponseBody) *GetTrustedProjectsResponse {
	s.Body = v
	return s
}

func (s *GetTrustedProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
