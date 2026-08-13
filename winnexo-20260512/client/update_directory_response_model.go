// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDirectoryResponseBody) *UpdateDirectoryResponse
	GetBody() *UpdateDirectoryResponseBody
}

type UpdateDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDirectoryResponse) GetBody() *UpdateDirectoryResponseBody {
	return s.Body
}

func (s *UpdateDirectoryResponse) SetHeaders(v map[string]*string) *UpdateDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateDirectoryResponse) SetStatusCode(v int32) *UpdateDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDirectoryResponse) SetBody(v *UpdateDirectoryResponseBody) *UpdateDirectoryResponse {
	s.Body = v
	return s
}

func (s *UpdateDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
