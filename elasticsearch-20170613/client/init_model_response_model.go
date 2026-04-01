// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitModelResponse
	GetStatusCode() *int32
	SetBody(v *InitModelResponseBody) *InitModelResponse
	GetBody() *InitModelResponseBody
}

type InitModelResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitModelResponse) String() string {
	return dara.Prettify(s)
}

func (s InitModelResponse) GoString() string {
	return s.String()
}

func (s *InitModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitModelResponse) GetBody() *InitModelResponseBody {
	return s.Body
}

func (s *InitModelResponse) SetHeaders(v map[string]*string) *InitModelResponse {
	s.Headers = v
	return s
}

func (s *InitModelResponse) SetStatusCode(v int32) *InitModelResponse {
	s.StatusCode = &v
	return s
}

func (s *InitModelResponse) SetBody(v *InitModelResponseBody) *InitModelResponse {
	s.Body = v
	return s
}

func (s *InitModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
