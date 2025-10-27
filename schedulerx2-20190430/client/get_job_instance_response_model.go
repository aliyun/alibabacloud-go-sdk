// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetJobInstanceResponseBody) *GetJobInstanceResponse
	GetBody() *GetJobInstanceResponseBody
}

type GetJobInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetJobInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobInstanceResponse) GetBody() *GetJobInstanceResponseBody {
	return s.Body
}

func (s *GetJobInstanceResponse) SetHeaders(v map[string]*string) *GetJobInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetJobInstanceResponse) SetStatusCode(v int32) *GetJobInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobInstanceResponse) SetBody(v *GetJobInstanceResponseBody) *GetJobInstanceResponse {
	s.Body = v
	return s
}

func (s *GetJobInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
