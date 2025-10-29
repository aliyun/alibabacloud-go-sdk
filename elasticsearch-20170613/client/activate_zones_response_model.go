// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateZonesResponse
	GetStatusCode() *int32
	SetBody(v *ActivateZonesResponseBody) *ActivateZonesResponse
	GetBody() *ActivateZonesResponseBody
}

type ActivateZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateZonesResponse) GoString() string {
	return s.String()
}

func (s *ActivateZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateZonesResponse) GetBody() *ActivateZonesResponseBody {
	return s.Body
}

func (s *ActivateZonesResponse) SetHeaders(v map[string]*string) *ActivateZonesResponse {
	s.Headers = v
	return s
}

func (s *ActivateZonesResponse) SetStatusCode(v int32) *ActivateZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateZonesResponse) SetBody(v *ActivateZonesResponseBody) *ActivateZonesResponse {
	s.Body = v
	return s
}

func (s *ActivateZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
