// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateDirectoryResponseBody) *CreateDirectoryResponse
	GetBody() *CreateDirectoryResponseBody
}

type CreateDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDirectoryResponse) GetBody() *CreateDirectoryResponseBody {
	return s.Body
}

func (s *CreateDirectoryResponse) SetHeaders(v map[string]*string) *CreateDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateDirectoryResponse) SetStatusCode(v int32) *CreateDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDirectoryResponse) SetBody(v *CreateDirectoryResponseBody) *CreateDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreateDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
