// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpEventConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpEventConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpEventConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpEventConfigResponseBody) *UpdateHttpEventConfigResponse
	GetBody() *UpdateHttpEventConfigResponseBody
}

type UpdateHttpEventConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpEventConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpEventConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpEventConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpEventConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpEventConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpEventConfigResponse) GetBody() *UpdateHttpEventConfigResponseBody {
	return s.Body
}

func (s *UpdateHttpEventConfigResponse) SetHeaders(v map[string]*string) *UpdateHttpEventConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpEventConfigResponse) SetStatusCode(v int32) *UpdateHttpEventConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpEventConfigResponse) SetBody(v *UpdateHttpEventConfigResponseBody) *UpdateHttpEventConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpEventConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
