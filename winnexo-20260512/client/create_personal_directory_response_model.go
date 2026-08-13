// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalDirectoryResponseBody) *CreatePersonalDirectoryResponse
	GetBody() *CreatePersonalDirectoryResponseBody
}

type CreatePersonalDirectoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalDirectoryResponse) GetBody() *CreatePersonalDirectoryResponseBody {
	return s.Body
}

func (s *CreatePersonalDirectoryResponse) SetHeaders(v map[string]*string) *CreatePersonalDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalDirectoryResponse) SetStatusCode(v int32) *CreatePersonalDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalDirectoryResponse) SetBody(v *CreatePersonalDirectoryResponseBody) *CreatePersonalDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
