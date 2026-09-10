// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataCheckTaskResponseBody) *UpdateDataCheckTaskResponse
	GetBody() *UpdateDataCheckTaskResponseBody
}

type UpdateDataCheckTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataCheckTaskResponse) GetBody() *UpdateDataCheckTaskResponseBody {
	return s.Body
}

func (s *UpdateDataCheckTaskResponse) SetHeaders(v map[string]*string) *UpdateDataCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataCheckTaskResponse) SetStatusCode(v int32) *UpdateDataCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataCheckTaskResponse) SetBody(v *UpdateDataCheckTaskResponseBody) *UpdateDataCheckTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateDataCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
