// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRAMDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRAMDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRAMDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateRAMDirectoryResponseBody) *CreateRAMDirectoryResponse
	GetBody() *CreateRAMDirectoryResponseBody
}

type CreateRAMDirectoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRAMDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRAMDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRAMDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRAMDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRAMDirectoryResponse) GetBody() *CreateRAMDirectoryResponseBody {
	return s.Body
}

func (s *CreateRAMDirectoryResponse) SetHeaders(v map[string]*string) *CreateRAMDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateRAMDirectoryResponse) SetStatusCode(v int32) *CreateRAMDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRAMDirectoryResponse) SetBody(v *CreateRAMDirectoryResponseBody) *CreateRAMDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreateRAMDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
