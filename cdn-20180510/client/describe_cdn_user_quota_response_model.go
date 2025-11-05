// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnUserQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnUserQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnUserQuotaResponseBody) *DescribeCdnUserQuotaResponse
	GetBody() *DescribeCdnUserQuotaResponseBody
}

type DescribeCdnUserQuotaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnUserQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnUserQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnUserQuotaResponse) GetBody() *DescribeCdnUserQuotaResponseBody {
	return s.Body
}

func (s *DescribeCdnUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeCdnUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserQuotaResponse) SetStatusCode(v int32) *DescribeCdnUserQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnUserQuotaResponse) SetBody(v *DescribeCdnUserQuotaResponseBody) *DescribeCdnUserQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnUserQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
