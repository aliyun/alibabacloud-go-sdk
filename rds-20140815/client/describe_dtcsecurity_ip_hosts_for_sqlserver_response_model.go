// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDTCSecurityIpHostsForSQLServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDTCSecurityIpHostsForSQLServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDTCSecurityIpHostsForSQLServerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDTCSecurityIpHostsForSQLServerResponseBody) *DescribeDTCSecurityIpHostsForSQLServerResponse
	GetBody() *DescribeDTCSecurityIpHostsForSQLServerResponseBody
}

type DescribeDTCSecurityIpHostsForSQLServerResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDTCSecurityIpHostsForSQLServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDTCSecurityIpHostsForSQLServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDTCSecurityIpHostsForSQLServerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponse) GetBody() *DescribeDTCSecurityIpHostsForSQLServerResponseBody {
	return s.Body
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponse) SetHeaders(v map[string]*string) *DescribeDTCSecurityIpHostsForSQLServerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponse) SetStatusCode(v int32) *DescribeDTCSecurityIpHostsForSQLServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponse) SetBody(v *DescribeDTCSecurityIpHostsForSQLServerResponseBody) *DescribeDTCSecurityIpHostsForSQLServerResponse {
	s.Body = v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
