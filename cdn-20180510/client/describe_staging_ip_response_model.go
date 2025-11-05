// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStagingIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStagingIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStagingIpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStagingIpResponseBody) *DescribeStagingIpResponse
	GetBody() *DescribeStagingIpResponseBody
}

type DescribeStagingIpResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStagingIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStagingIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStagingIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeStagingIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStagingIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStagingIpResponse) GetBody() *DescribeStagingIpResponseBody {
	return s.Body
}

func (s *DescribeStagingIpResponse) SetHeaders(v map[string]*string) *DescribeStagingIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeStagingIpResponse) SetStatusCode(v int32) *DescribeStagingIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStagingIpResponse) SetBody(v *DescribeStagingIpResponseBody) *DescribeStagingIpResponse {
	s.Body = v
	return s
}

func (s *DescribeStagingIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
