// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnStagingIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnStagingIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnStagingIpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnStagingIpResponseBody) *DescribeDcdnStagingIpResponse
	GetBody() *DescribeDcdnStagingIpResponseBody
}

type DescribeDcdnStagingIpResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnStagingIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnStagingIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnStagingIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnStagingIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnStagingIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnStagingIpResponse) GetBody() *DescribeDcdnStagingIpResponseBody {
	return s.Body
}

func (s *DescribeDcdnStagingIpResponse) SetHeaders(v map[string]*string) *DescribeDcdnStagingIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnStagingIpResponse) SetStatusCode(v int32) *DescribeDcdnStagingIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnStagingIpResponse) SetBody(v *DescribeDcdnStagingIpResponseBody) *DescribeDcdnStagingIpResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnStagingIpResponse) Validate() error {
	return dara.Validate(s)
}
