// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSourceCodeRepoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoSourceCodeRepoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoSourceCodeRepoResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoSourceCodeRepoResponseBody) *CreateRepoSourceCodeRepoResponse
	GetBody() *CreateRepoSourceCodeRepoResponseBody
}

type CreateRepoSourceCodeRepoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoSourceCodeRepoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoSourceCodeRepoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSourceCodeRepoResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoSourceCodeRepoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoSourceCodeRepoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoSourceCodeRepoResponse) GetBody() *CreateRepoSourceCodeRepoResponseBody {
	return s.Body
}

func (s *CreateRepoSourceCodeRepoResponse) SetHeaders(v map[string]*string) *CreateRepoSourceCodeRepoResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoSourceCodeRepoResponse) SetStatusCode(v int32) *CreateRepoSourceCodeRepoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoSourceCodeRepoResponse) SetBody(v *CreateRepoSourceCodeRepoResponseBody) *CreateRepoSourceCodeRepoResponse {
	s.Body = v
	return s
}

func (s *CreateRepoSourceCodeRepoResponse) Validate() error {
	return dara.Validate(s)
}
