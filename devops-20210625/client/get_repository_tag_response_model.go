// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepositoryTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepositoryTagResponse
	GetStatusCode() *int32
	SetBody(v *GetRepositoryTagResponseBody) *GetRepositoryTagResponse
	GetBody() *GetRepositoryTagResponseBody
}

type GetRepositoryTagResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepositoryTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepositoryTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryTagResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepositoryTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepositoryTagResponse) GetBody() *GetRepositoryTagResponseBody {
	return s.Body
}

func (s *GetRepositoryTagResponse) SetHeaders(v map[string]*string) *GetRepositoryTagResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryTagResponse) SetStatusCode(v int32) *GetRepositoryTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryTagResponse) SetBody(v *GetRepositoryTagResponseBody) *GetRepositoryTagResponse {
	s.Body = v
	return s
}

func (s *GetRepositoryTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
