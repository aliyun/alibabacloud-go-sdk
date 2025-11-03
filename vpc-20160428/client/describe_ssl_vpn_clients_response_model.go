// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSslVpnClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSslVpnClientsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSslVpnClientsResponseBody) *DescribeSslVpnClientsResponse
	GetBody() *DescribeSslVpnClientsResponseBody
}

type DescribeSslVpnClientsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSslVpnClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSslVpnClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSslVpnClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSslVpnClientsResponse) GetBody() *DescribeSslVpnClientsResponseBody {
	return s.Body
}

func (s *DescribeSslVpnClientsResponse) SetHeaders(v map[string]*string) *DescribeSslVpnClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSslVpnClientsResponse) SetStatusCode(v int32) *DescribeSslVpnClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSslVpnClientsResponse) SetBody(v *DescribeSslVpnClientsResponseBody) *DescribeSslVpnClientsResponse {
	s.Body = v
	return s
}

func (s *DescribeSslVpnClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
