// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsMpsAiUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMeterImsMpsAiUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMeterImsMpsAiUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMeterImsMpsAiUsageResponseBody) *DescribeMeterImsMpsAiUsageResponse
	GetBody() *DescribeMeterImsMpsAiUsageResponseBody
}

type DescribeMeterImsMpsAiUsageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterImsMpsAiUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterImsMpsAiUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMpsAiUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMpsAiUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMeterImsMpsAiUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMeterImsMpsAiUsageResponse) GetBody() *DescribeMeterImsMpsAiUsageResponseBody {
	return s.Body
}

func (s *DescribeMeterImsMpsAiUsageResponse) SetHeaders(v map[string]*string) *DescribeMeterImsMpsAiUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterImsMpsAiUsageResponse) SetStatusCode(v int32) *DescribeMeterImsMpsAiUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageResponse) SetBody(v *DescribeMeterImsMpsAiUsageResponseBody) *DescribeMeterImsMpsAiUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeMeterImsMpsAiUsageResponse) Validate() error {
	return dara.Validate(s)
}
