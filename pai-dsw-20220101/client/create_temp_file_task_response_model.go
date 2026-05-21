// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTempFileTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTempFileTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTempFileTaskResponseBody) *CreateTempFileTaskResponse
	GetBody() *CreateTempFileTaskResponseBody
}

type CreateTempFileTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTempFileTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTempFileTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTempFileTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTempFileTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTempFileTaskResponse) GetBody() *CreateTempFileTaskResponseBody {
	return s.Body
}

func (s *CreateTempFileTaskResponse) SetHeaders(v map[string]*string) *CreateTempFileTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTempFileTaskResponse) SetStatusCode(v int32) *CreateTempFileTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTempFileTaskResponse) SetBody(v *CreateTempFileTaskResponseBody) *CreateTempFileTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTempFileTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
