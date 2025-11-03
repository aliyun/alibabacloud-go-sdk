// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoSourceCodeRepoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRepoSourceCodeRepoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRepoSourceCodeRepoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRepoSourceCodeRepoResponseBody) *UpdateRepoSourceCodeRepoResponse
	GetBody() *UpdateRepoSourceCodeRepoResponseBody
}

type UpdateRepoSourceCodeRepoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepoSourceCodeRepoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepoSourceCodeRepoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoSourceCodeRepoResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepoSourceCodeRepoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRepoSourceCodeRepoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRepoSourceCodeRepoResponse) GetBody() *UpdateRepoSourceCodeRepoResponseBody {
	return s.Body
}

func (s *UpdateRepoSourceCodeRepoResponse) SetHeaders(v map[string]*string) *UpdateRepoSourceCodeRepoResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponse) SetStatusCode(v int32) *UpdateRepoSourceCodeRepoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponse) SetBody(v *UpdateRepoSourceCodeRepoResponseBody) *UpdateRepoSourceCodeRepoResponse {
	s.Body = v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
