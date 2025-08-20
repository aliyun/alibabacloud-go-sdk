// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRepositoryResponseBody) *UpdateRepositoryResponse
	GetBody() *UpdateRepositoryResponseBody
}

type UpdateRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRepositoryResponse) GetBody() *UpdateRepositoryResponseBody {
	return s.Body
}

func (s *UpdateRepositoryResponse) SetHeaders(v map[string]*string) *UpdateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepositoryResponse) SetStatusCode(v int32) *UpdateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepositoryResponse) SetBody(v *UpdateRepositoryResponseBody) *UpdateRepositoryResponse {
	s.Body = v
	return s
}

func (s *UpdateRepositoryResponse) Validate() error {
	return dara.Validate(s)
}
