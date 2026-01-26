// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebhookContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebhookContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebhookContactsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebhookContactsResponseBody) *DescribeWebhookContactsResponse
	GetBody() *DescribeWebhookContactsResponseBody
}

type DescribeWebhookContactsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebhookContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebhookContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebhookContactsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebhookContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebhookContactsResponse) GetBody() *DescribeWebhookContactsResponseBody {
	return s.Body
}

func (s *DescribeWebhookContactsResponse) SetHeaders(v map[string]*string) *DescribeWebhookContactsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebhookContactsResponse) SetStatusCode(v int32) *DescribeWebhookContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebhookContactsResponse) SetBody(v *DescribeWebhookContactsResponseBody) *DescribeWebhookContactsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebhookContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
