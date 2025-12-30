// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateZoneRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateZoneRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateZoneRemarkResponseBody) *UpdateZoneRemarkResponse
	GetBody() *UpdateZoneRemarkResponseBody
}

type UpdateZoneRemarkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateZoneRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateZoneRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateZoneRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateZoneRemarkResponse) GetBody() *UpdateZoneRemarkResponseBody {
	return s.Body
}

func (s *UpdateZoneRemarkResponse) SetHeaders(v map[string]*string) *UpdateZoneRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateZoneRemarkResponse) SetStatusCode(v int32) *UpdateZoneRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateZoneRemarkResponse) SetBody(v *UpdateZoneRemarkResponseBody) *UpdateZoneRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateZoneRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
