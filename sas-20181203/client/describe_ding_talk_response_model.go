// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDingTalkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDingTalkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDingTalkResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDingTalkResponseBody) *DescribeDingTalkResponse
	GetBody() *DescribeDingTalkResponseBody
}

type DescribeDingTalkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDingTalkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDingTalkResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDingTalkResponse) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDingTalkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDingTalkResponse) GetBody() *DescribeDingTalkResponseBody {
	return s.Body
}

func (s *DescribeDingTalkResponse) SetHeaders(v map[string]*string) *DescribeDingTalkResponse {
	s.Headers = v
	return s
}

func (s *DescribeDingTalkResponse) SetStatusCode(v int32) *DescribeDingTalkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDingTalkResponse) SetBody(v *DescribeDingTalkResponseBody) *DescribeDingTalkResponse {
	s.Body = v
	return s
}

func (s *DescribeDingTalkResponse) Validate() error {
	return dara.Validate(s)
}
