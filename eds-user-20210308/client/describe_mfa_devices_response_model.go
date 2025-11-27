// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMfaDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMfaDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMfaDevicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMfaDevicesResponseBody) *DescribeMfaDevicesResponse
	GetBody() *DescribeMfaDevicesResponseBody
}

type DescribeMfaDevicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMfaDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMfaDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMfaDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMfaDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMfaDevicesResponse) GetBody() *DescribeMfaDevicesResponseBody {
	return s.Body
}

func (s *DescribeMfaDevicesResponse) SetHeaders(v map[string]*string) *DescribeMfaDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMfaDevicesResponse) SetStatusCode(v int32) *DescribeMfaDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMfaDevicesResponse) SetBody(v *DescribeMfaDevicesResponseBody) *DescribeMfaDevicesResponse {
	s.Body = v
	return s
}

func (s *DescribeMfaDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
