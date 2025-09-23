// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAccessModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainAccessModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainAccessModeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainAccessModeResponseBody) *DescribeDomainAccessModeResponse
	GetBody() *DescribeDomainAccessModeResponseBody
}

type DescribeDomainAccessModeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainAccessModeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAccessModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainAccessModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainAccessModeResponse) GetBody() *DescribeDomainAccessModeResponseBody {
	return s.Body
}

func (s *DescribeDomainAccessModeResponse) SetHeaders(v map[string]*string) *DescribeDomainAccessModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAccessModeResponse) SetStatusCode(v int32) *DescribeDomainAccessModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAccessModeResponse) SetBody(v *DescribeDomainAccessModeResponseBody) *DescribeDomainAccessModeResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainAccessModeResponse) Validate() error {
	return dara.Validate(s)
}
