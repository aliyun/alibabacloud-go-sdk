// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePushProxyLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLivePushProxyLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLivePushProxyLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLivePushProxyLogResponseBody) *DescribeLivePushProxyLogResponse
	GetBody() *DescribeLivePushProxyLogResponseBody
}

type DescribeLivePushProxyLogResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLivePushProxyLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLivePushProxyLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePushProxyLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePushProxyLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLivePushProxyLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLivePushProxyLogResponse) GetBody() *DescribeLivePushProxyLogResponseBody {
	return s.Body
}

func (s *DescribeLivePushProxyLogResponse) SetHeaders(v map[string]*string) *DescribeLivePushProxyLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeLivePushProxyLogResponse) SetStatusCode(v int32) *DescribeLivePushProxyLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLivePushProxyLogResponse) SetBody(v *DescribeLivePushProxyLogResponseBody) *DescribeLivePushProxyLogResponse {
	s.Body = v
	return s
}

func (s *DescribeLivePushProxyLogResponse) Validate() error {
	return dara.Validate(s)
}
