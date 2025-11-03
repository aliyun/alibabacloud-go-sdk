// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoSourceCodeRepoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepoSourceCodeRepoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepoSourceCodeRepoResponse
	GetStatusCode() *int32
	SetBody(v *GetRepoSourceCodeRepoResponseBody) *GetRepoSourceCodeRepoResponse
	GetBody() *GetRepoSourceCodeRepoResponseBody
}

type GetRepoSourceCodeRepoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoSourceCodeRepoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoSourceCodeRepoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSourceCodeRepoResponse) GoString() string {
	return s.String()
}

func (s *GetRepoSourceCodeRepoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepoSourceCodeRepoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepoSourceCodeRepoResponse) GetBody() *GetRepoSourceCodeRepoResponseBody {
	return s.Body
}

func (s *GetRepoSourceCodeRepoResponse) SetHeaders(v map[string]*string) *GetRepoSourceCodeRepoResponse {
	s.Headers = v
	return s
}

func (s *GetRepoSourceCodeRepoResponse) SetStatusCode(v int32) *GetRepoSourceCodeRepoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponse) SetBody(v *GetRepoSourceCodeRepoResponseBody) *GetRepoSourceCodeRepoResponse {
	s.Body = v
	return s
}

func (s *GetRepoSourceCodeRepoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
