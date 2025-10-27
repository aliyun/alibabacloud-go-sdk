// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspiciousUUIDConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSuspiciousUUIDConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSuspiciousUUIDConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSuspiciousUUIDConfigResponseBody) *DescribeSuspiciousUUIDConfigResponse
	GetBody() *DescribeSuspiciousUUIDConfigResponseBody
}

type DescribeSuspiciousUUIDConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSuspiciousUUIDConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSuspiciousUUIDConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspiciousUUIDConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspiciousUUIDConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSuspiciousUUIDConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSuspiciousUUIDConfigResponse) GetBody() *DescribeSuspiciousUUIDConfigResponseBody {
	return s.Body
}

func (s *DescribeSuspiciousUUIDConfigResponse) SetHeaders(v map[string]*string) *DescribeSuspiciousUUIDConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspiciousUUIDConfigResponse) SetStatusCode(v int32) *DescribeSuspiciousUUIDConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSuspiciousUUIDConfigResponse) SetBody(v *DescribeSuspiciousUUIDConfigResponseBody) *DescribeSuspiciousUUIDConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSuspiciousUUIDConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
