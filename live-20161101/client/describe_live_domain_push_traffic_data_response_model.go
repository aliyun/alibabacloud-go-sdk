// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPushTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainPushTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainPushTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainPushTrafficDataResponseBody) *DescribeLiveDomainPushTrafficDataResponse
	GetBody() *DescribeLiveDomainPushTrafficDataResponseBody
}

type DescribeLiveDomainPushTrafficDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainPushTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainPushTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainPushTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainPushTrafficDataResponse) GetBody() *DescribeLiveDomainPushTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainPushTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetStatusCode(v int32) *DescribeLiveDomainPushTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetBody(v *DescribeLiveDomainPushTrafficDataResponseBody) *DescribeLiveDomainPushTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
