// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrecisionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrecisionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrecisionTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetPrecisionTaskResponseBody) *GetPrecisionTaskResponse
	GetBody() *GetPrecisionTaskResponseBody
}

type GetPrecisionTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrecisionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrecisionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrecisionTaskResponse) GetBody() *GetPrecisionTaskResponseBody {
	return s.Body
}

func (s *GetPrecisionTaskResponse) SetHeaders(v map[string]*string) *GetPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetPrecisionTaskResponse) SetStatusCode(v int32) *GetPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrecisionTaskResponse) SetBody(v *GetPrecisionTaskResponseBody) *GetPrecisionTaskResponse {
	s.Body = v
	return s
}

func (s *GetPrecisionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
