// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTimingSyntheticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTimingSyntheticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTimingSyntheticTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTimingSyntheticTaskResponseBody) *UpdateTimingSyntheticTaskResponse
	GetBody() *UpdateTimingSyntheticTaskResponseBody
}

type UpdateTimingSyntheticTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTimingSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTimingSyntheticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTimingSyntheticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTimingSyntheticTaskResponse) GetBody() *UpdateTimingSyntheticTaskResponseBody {
	return s.Body
}

func (s *UpdateTimingSyntheticTaskResponse) SetHeaders(v map[string]*string) *UpdateTimingSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTimingSyntheticTaskResponse) SetStatusCode(v int32) *UpdateTimingSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTimingSyntheticTaskResponse) SetBody(v *UpdateTimingSyntheticTaskResponseBody) *UpdateTimingSyntheticTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateTimingSyntheticTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
