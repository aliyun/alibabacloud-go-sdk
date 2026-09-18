// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGroupDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGroupDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGroupDirectoryResponseBody) *UpdateGroupDirectoryResponse
	GetBody() *UpdateGroupDirectoryResponseBody
}

type UpdateGroupDirectoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGroupDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGroupDirectoryResponse) GetBody() *UpdateGroupDirectoryResponseBody {
	return s.Body
}

func (s *UpdateGroupDirectoryResponse) SetHeaders(v map[string]*string) *UpdateGroupDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupDirectoryResponse) SetStatusCode(v int32) *UpdateGroupDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupDirectoryResponse) SetBody(v *UpdateGroupDirectoryResponseBody) *UpdateGroupDirectoryResponse {
	s.Body = v
	return s
}

func (s *UpdateGroupDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
