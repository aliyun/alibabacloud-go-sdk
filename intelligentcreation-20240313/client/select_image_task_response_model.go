// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectImageTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectImageTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectImageTaskResponse
	GetStatusCode() *int32
	SetBody(v *SelectImageTaskResponseBody) *SelectImageTaskResponse
	GetBody() *SelectImageTaskResponseBody
}

type SelectImageTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectImageTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SelectImageTaskResponse) GoString() string {
	return s.String()
}

func (s *SelectImageTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectImageTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectImageTaskResponse) GetBody() *SelectImageTaskResponseBody {
	return s.Body
}

func (s *SelectImageTaskResponse) SetHeaders(v map[string]*string) *SelectImageTaskResponse {
	s.Headers = v
	return s
}

func (s *SelectImageTaskResponse) SetStatusCode(v int32) *SelectImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectImageTaskResponse) SetBody(v *SelectImageTaskResponseBody) *SelectImageTaskResponse {
	s.Body = v
	return s
}

func (s *SelectImageTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
