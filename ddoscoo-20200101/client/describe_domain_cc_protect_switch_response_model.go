// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCcProtectSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainCcProtectSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainCcProtectSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainCcProtectSwitchResponseBody) *DescribeDomainCcProtectSwitchResponse
	GetBody() *DescribeDomainCcProtectSwitchResponseBody
}

type DescribeDomainCcProtectSwitchResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainCcProtectSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainCcProtectSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCcProtectSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcProtectSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainCcProtectSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainCcProtectSwitchResponse) GetBody() *DescribeDomainCcProtectSwitchResponseBody {
	return s.Body
}

func (s *DescribeDomainCcProtectSwitchResponse) SetHeaders(v map[string]*string) *DescribeDomainCcProtectSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponse) SetStatusCode(v int32) *DescribeDomainCcProtectSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponse) SetBody(v *DescribeDomainCcProtectSwitchResponseBody) *DescribeDomainCcProtectSwitchResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainCcProtectSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
