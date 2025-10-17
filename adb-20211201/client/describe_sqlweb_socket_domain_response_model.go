// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLWebSocketDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLWebSocketDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLWebSocketDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLWebSocketDomainResponseBody) *DescribeSQLWebSocketDomainResponse
	GetBody() *DescribeSQLWebSocketDomainResponseBody
}

type DescribeSQLWebSocketDomainResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLWebSocketDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLWebSocketDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLWebSocketDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLWebSocketDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLWebSocketDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLWebSocketDomainResponse) GetBody() *DescribeSQLWebSocketDomainResponseBody {
	return s.Body
}

func (s *DescribeSQLWebSocketDomainResponse) SetHeaders(v map[string]*string) *DescribeSQLWebSocketDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLWebSocketDomainResponse) SetStatusCode(v int32) *DescribeSQLWebSocketDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLWebSocketDomainResponse) SetBody(v *DescribeSQLWebSocketDomainResponseBody) *DescribeSQLWebSocketDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLWebSocketDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
