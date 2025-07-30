// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsEtlJobVersionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDtsEtlJobVersionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDtsEtlJobVersionInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDtsEtlJobVersionInfoResponseBody) *DescribeDtsEtlJobVersionInfoResponse
	GetBody() *DescribeDtsEtlJobVersionInfoResponseBody
}

type DescribeDtsEtlJobVersionInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDtsEtlJobVersionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDtsEtlJobVersionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsEtlJobVersionInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDtsEtlJobVersionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDtsEtlJobVersionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDtsEtlJobVersionInfoResponse) GetBody() *DescribeDtsEtlJobVersionInfoResponseBody {
	return s.Body
}

func (s *DescribeDtsEtlJobVersionInfoResponse) SetHeaders(v map[string]*string) *DescribeDtsEtlJobVersionInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponse) SetStatusCode(v int32) *DescribeDtsEtlJobVersionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponse) SetBody(v *DescribeDtsEtlJobVersionInfoResponseBody) *DescribeDtsEtlJobVersionInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponse) Validate() error {
	return dara.Validate(s)
}
