// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenPrivateZoneRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenPrivateZoneRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenPrivateZoneRoutesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenPrivateZoneRoutesResponseBody) *DescribeCenPrivateZoneRoutesResponse
	GetBody() *DescribeCenPrivateZoneRoutesResponseBody
}

type DescribeCenPrivateZoneRoutesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenPrivateZoneRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenPrivateZoneRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenPrivateZoneRoutesResponse) GetBody() *DescribeCenPrivateZoneRoutesResponseBody {
	return s.Body
}

func (s *DescribeCenPrivateZoneRoutesResponse) SetHeaders(v map[string]*string) *DescribeCenPrivateZoneRoutesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponse) SetStatusCode(v int32) *DescribeCenPrivateZoneRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponse) SetBody(v *DescribeCenPrivateZoneRoutesResponseBody) *DescribeCenPrivateZoneRoutesResponse {
	s.Body = v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
