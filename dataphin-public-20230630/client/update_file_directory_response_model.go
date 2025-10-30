// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileDirectoryResponseBody) *UpdateFileDirectoryResponse
	GetBody() *UpdateFileDirectoryResponseBody
}

type UpdateFileDirectoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileDirectoryResponse) GetBody() *UpdateFileDirectoryResponseBody {
	return s.Body
}

func (s *UpdateFileDirectoryResponse) SetHeaders(v map[string]*string) *UpdateFileDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileDirectoryResponse) SetStatusCode(v int32) *UpdateFileDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileDirectoryResponse) SetBody(v *UpdateFileDirectoryResponseBody) *UpdateFileDirectoryResponse {
	s.Body = v
	return s
}

func (s *UpdateFileDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
