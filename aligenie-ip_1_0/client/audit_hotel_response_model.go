// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditHotelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuditHotelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuditHotelResponse
	GetStatusCode() *int32
	SetBody(v *AuditHotelResponseBody) *AuditHotelResponse
	GetBody() *AuditHotelResponseBody
}

type AuditHotelResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditHotelResponse) String() string {
	return dara.Prettify(s)
}

func (s AuditHotelResponse) GoString() string {
	return s.String()
}

func (s *AuditHotelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuditHotelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuditHotelResponse) GetBody() *AuditHotelResponseBody {
	return s.Body
}

func (s *AuditHotelResponse) SetHeaders(v map[string]*string) *AuditHotelResponse {
	s.Headers = v
	return s
}

func (s *AuditHotelResponse) SetStatusCode(v int32) *AuditHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditHotelResponse) SetBody(v *AuditHotelResponseBody) *AuditHotelResponse {
	s.Body = v
	return s
}

func (s *AuditHotelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
