// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelVersionResponseBody) *UpdateModelVersionResponse
	GetBody() *UpdateModelVersionResponseBody
}

type UpdateModelVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelVersionResponse) GetBody() *UpdateModelVersionResponseBody {
	return s.Body
}

func (s *UpdateModelVersionResponse) SetHeaders(v map[string]*string) *UpdateModelVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelVersionResponse) SetStatusCode(v int32) *UpdateModelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelVersionResponse) SetBody(v *UpdateModelVersionResponseBody) *UpdateModelVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateModelVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
