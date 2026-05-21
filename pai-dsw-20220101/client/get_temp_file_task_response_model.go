// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempFileTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTempFileTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTempFileTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetTempFileTaskResponseBody) *GetTempFileTaskResponse
	GetBody() *GetTempFileTaskResponseBody
}

type GetTempFileTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTempFileTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTempFileTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTempFileTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTempFileTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTempFileTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTempFileTaskResponse) GetBody() *GetTempFileTaskResponseBody {
	return s.Body
}

func (s *GetTempFileTaskResponse) SetHeaders(v map[string]*string) *GetTempFileTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTempFileTaskResponse) SetStatusCode(v int32) *GetTempFileTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTempFileTaskResponse) SetBody(v *GetTempFileTaskResponseBody) *GetTempFileTaskResponse {
	s.Body = v
	return s
}

func (s *GetTempFileTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
