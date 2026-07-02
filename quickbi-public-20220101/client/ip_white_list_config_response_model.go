// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIpWhiteListConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IpWhiteListConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IpWhiteListConfigResponse
	GetStatusCode() *int32
	SetBody(v *IpWhiteListConfigResponseBody) *IpWhiteListConfigResponse
	GetBody() *IpWhiteListConfigResponseBody
}

type IpWhiteListConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IpWhiteListConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IpWhiteListConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s IpWhiteListConfigResponse) GoString() string {
	return s.String()
}

func (s *IpWhiteListConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IpWhiteListConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IpWhiteListConfigResponse) GetBody() *IpWhiteListConfigResponseBody {
	return s.Body
}

func (s *IpWhiteListConfigResponse) SetHeaders(v map[string]*string) *IpWhiteListConfigResponse {
	s.Headers = v
	return s
}

func (s *IpWhiteListConfigResponse) SetStatusCode(v int32) *IpWhiteListConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *IpWhiteListConfigResponse) SetBody(v *IpWhiteListConfigResponseBody) *IpWhiteListConfigResponse {
	s.Body = v
	return s
}

func (s *IpWhiteListConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
