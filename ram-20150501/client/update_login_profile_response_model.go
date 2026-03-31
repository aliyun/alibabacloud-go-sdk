// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoginProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLoginProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLoginProfileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLoginProfileResponseBody) *UpdateLoginProfileResponse
	GetBody() *UpdateLoginProfileResponseBody
}

type UpdateLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLoginProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLoginProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLoginProfileResponse) GetBody() *UpdateLoginProfileResponseBody {
	return s.Body
}

func (s *UpdateLoginProfileResponse) SetHeaders(v map[string]*string) *UpdateLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoginProfileResponse) SetStatusCode(v int32) *UpdateLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoginProfileResponse) SetBody(v *UpdateLoginProfileResponseBody) *UpdateLoginProfileResponse {
	s.Body = v
	return s
}

func (s *UpdateLoginProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
