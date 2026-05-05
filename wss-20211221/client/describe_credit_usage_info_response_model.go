// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditUsageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCreditUsageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCreditUsageInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCreditUsageInfoResponseBody) *DescribeCreditUsageInfoResponse
	GetBody() *DescribeCreditUsageInfoResponseBody
}

type DescribeCreditUsageInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCreditUsageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCreditUsageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCreditUsageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCreditUsageInfoResponse) GetBody() *DescribeCreditUsageInfoResponseBody {
	return s.Body
}

func (s *DescribeCreditUsageInfoResponse) SetHeaders(v map[string]*string) *DescribeCreditUsageInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreditUsageInfoResponse) SetStatusCode(v int32) *DescribeCreditUsageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCreditUsageInfoResponse) SetBody(v *DescribeCreditUsageInfoResponseBody) *DescribeCreditUsageInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCreditUsageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
