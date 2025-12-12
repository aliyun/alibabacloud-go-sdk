// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarningConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWarningConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWarningConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWarningConfigResponseBody) *UpdateWarningConfigResponse
	GetBody() *UpdateWarningConfigResponseBody
}

type UpdateWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWarningConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWarningConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWarningConfigResponse) GetBody() *UpdateWarningConfigResponseBody {
	return s.Body
}

func (s *UpdateWarningConfigResponse) SetHeaders(v map[string]*string) *UpdateWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateWarningConfigResponse) SetStatusCode(v int32) *UpdateWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWarningConfigResponse) SetBody(v *UpdateWarningConfigResponseBody) *UpdateWarningConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateWarningConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
