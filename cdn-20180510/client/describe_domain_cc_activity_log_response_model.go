// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCcActivityLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainCcActivityLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainCcActivityLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainCcActivityLogResponseBody) *DescribeDomainCcActivityLogResponse
	GetBody() *DescribeDomainCcActivityLogResponseBody
}

type DescribeDomainCcActivityLogResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainCcActivityLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainCcActivityLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCcActivityLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcActivityLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainCcActivityLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainCcActivityLogResponse) GetBody() *DescribeDomainCcActivityLogResponseBody {
	return s.Body
}

func (s *DescribeDomainCcActivityLogResponse) SetHeaders(v map[string]*string) *DescribeDomainCcActivityLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCcActivityLogResponse) SetStatusCode(v int32) *DescribeDomainCcActivityLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponse) SetBody(v *DescribeDomainCcActivityLogResponseBody) *DescribeDomainCcActivityLogResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainCcActivityLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
