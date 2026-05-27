// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLargeModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLargeModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLargeModelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLargeModelResponseBody) *UpdateLargeModelResponse
	GetBody() *UpdateLargeModelResponseBody
}

type UpdateLargeModelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLargeModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLargeModelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLargeModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateLargeModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLargeModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLargeModelResponse) GetBody() *UpdateLargeModelResponseBody {
	return s.Body
}

func (s *UpdateLargeModelResponse) SetHeaders(v map[string]*string) *UpdateLargeModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateLargeModelResponse) SetStatusCode(v int32) *UpdateLargeModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLargeModelResponse) SetBody(v *UpdateLargeModelResponseBody) *UpdateLargeModelResponse {
	s.Body = v
	return s
}

func (s *UpdateLargeModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
