// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePushProxyUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLivePushProxyUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLivePushProxyUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLivePushProxyUsageDataResponseBody) *DescribeLivePushProxyUsageDataResponse
	GetBody() *DescribeLivePushProxyUsageDataResponseBody
}

type DescribeLivePushProxyUsageDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLivePushProxyUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLivePushProxyUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePushProxyUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePushProxyUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLivePushProxyUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLivePushProxyUsageDataResponse) GetBody() *DescribeLivePushProxyUsageDataResponseBody {
	return s.Body
}

func (s *DescribeLivePushProxyUsageDataResponse) SetHeaders(v map[string]*string) *DescribeLivePushProxyUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponse) SetStatusCode(v int32) *DescribeLivePushProxyUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponse) SetBody(v *DescribeLivePushProxyUsageDataResponseBody) *DescribeLivePushProxyUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLivePushProxyUsageDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
