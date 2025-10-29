// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeAutoShowListTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeAutoShowListTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeAutoShowListTaskResponse
	GetStatusCode() *int32
	SetBody(v *InitializeAutoShowListTaskResponseBody) *InitializeAutoShowListTaskResponse
	GetBody() *InitializeAutoShowListTaskResponseBody
}

type InitializeAutoShowListTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeAutoShowListTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeAutoShowListTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeAutoShowListTaskResponse) GoString() string {
	return s.String()
}

func (s *InitializeAutoShowListTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeAutoShowListTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeAutoShowListTaskResponse) GetBody() *InitializeAutoShowListTaskResponseBody {
	return s.Body
}

func (s *InitializeAutoShowListTaskResponse) SetHeaders(v map[string]*string) *InitializeAutoShowListTaskResponse {
	s.Headers = v
	return s
}

func (s *InitializeAutoShowListTaskResponse) SetStatusCode(v int32) *InitializeAutoShowListTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeAutoShowListTaskResponse) SetBody(v *InitializeAutoShowListTaskResponseBody) *InitializeAutoShowListTaskResponse {
	s.Body = v
	return s
}

func (s *InitializeAutoShowListTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
