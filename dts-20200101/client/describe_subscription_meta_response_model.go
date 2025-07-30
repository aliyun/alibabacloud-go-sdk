// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubscriptionMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubscriptionMetaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubscriptionMetaResponseBody) *DescribeSubscriptionMetaResponse
	GetBody() *DescribeSubscriptionMetaResponseBody
}

type DescribeSubscriptionMetaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubscriptionMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubscriptionMetaResponse) GetBody() *DescribeSubscriptionMetaResponseBody {
	return s.Body
}

func (s *DescribeSubscriptionMetaResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionMetaResponse) SetStatusCode(v int32) *DescribeSubscriptionMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionMetaResponse) SetBody(v *DescribeSubscriptionMetaResponseBody) *DescribeSubscriptionMetaResponse {
	s.Body = v
	return s
}

func (s *DescribeSubscriptionMetaResponse) Validate() error {
	return dara.Validate(s)
}
