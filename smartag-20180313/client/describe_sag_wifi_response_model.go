// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWifiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagWifiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagWifiResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagWifiResponseBody) *DescribeSagWifiResponse
	GetBody() *DescribeSagWifiResponseBody
}

type DescribeSagWifiResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagWifiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagWifiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWifiResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagWifiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagWifiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagWifiResponse) GetBody() *DescribeSagWifiResponseBody {
	return s.Body
}

func (s *DescribeSagWifiResponse) SetHeaders(v map[string]*string) *DescribeSagWifiResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagWifiResponse) SetStatusCode(v int32) *DescribeSagWifiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagWifiResponse) SetBody(v *DescribeSagWifiResponseBody) *DescribeSagWifiResponse {
	s.Body = v
	return s
}

func (s *DescribeSagWifiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
