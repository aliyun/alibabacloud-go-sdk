// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDetectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDetectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDetectConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDetectConfigResponseBody) *UpdateDetectConfigResponse
	GetBody() *UpdateDetectConfigResponseBody
}

type UpdateDetectConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDetectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDetectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDetectConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDetectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDetectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDetectConfigResponse) GetBody() *UpdateDetectConfigResponseBody {
	return s.Body
}

func (s *UpdateDetectConfigResponse) SetHeaders(v map[string]*string) *UpdateDetectConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDetectConfigResponse) SetStatusCode(v int32) *UpdateDetectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDetectConfigResponse) SetBody(v *UpdateDetectConfigResponseBody) *UpdateDetectConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateDetectConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
