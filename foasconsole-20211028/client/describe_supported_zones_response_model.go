// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportedZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupportedZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupportedZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupportedZonesResponseBody) *DescribeSupportedZonesResponse
	GetBody() *DescribeSupportedZonesResponseBody
}

type DescribeSupportedZonesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportedZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportedZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportedZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupportedZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupportedZonesResponse) GetBody() *DescribeSupportedZonesResponseBody {
	return s.Body
}

func (s *DescribeSupportedZonesResponse) SetHeaders(v map[string]*string) *DescribeSupportedZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportedZonesResponse) SetStatusCode(v int32) *DescribeSupportedZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportedZonesResponse) SetBody(v *DescribeSupportedZonesResponseBody) *DescribeSupportedZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeSupportedZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
