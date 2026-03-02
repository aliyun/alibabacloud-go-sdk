// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoForkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoForkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoForkResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoForkResponseBody) *CreateRepoForkResponse
	GetBody() *CreateRepoForkResponseBody
}

type CreateRepoForkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoForkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoForkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoForkResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoForkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoForkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoForkResponse) GetBody() *CreateRepoForkResponseBody {
	return s.Body
}

func (s *CreateRepoForkResponse) SetHeaders(v map[string]*string) *CreateRepoForkResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoForkResponse) SetStatusCode(v int32) *CreateRepoForkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoForkResponse) SetBody(v *CreateRepoForkResponseBody) *CreateRepoForkResponse {
	s.Body = v
	return s
}

func (s *CreateRepoForkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
