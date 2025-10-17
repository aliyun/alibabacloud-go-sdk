// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInstanceAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubscriptionInstanceAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubscriptionInstanceAlertResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubscriptionInstanceAlertResponseBody) *DescribeSubscriptionInstanceAlertResponse
	GetBody() *DescribeSubscriptionInstanceAlertResponseBody
}

type DescribeSubscriptionInstanceAlertResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionInstanceAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionInstanceAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstanceAlertResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubscriptionInstanceAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubscriptionInstanceAlertResponse) GetBody() *DescribeSubscriptionInstanceAlertResponseBody {
	return s.Body
}

func (s *DescribeSubscriptionInstanceAlertResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionInstanceAlertResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponse) SetStatusCode(v int32) *DescribeSubscriptionInstanceAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponse) SetBody(v *DescribeSubscriptionInstanceAlertResponseBody) *DescribeSubscriptionInstanceAlertResponse {
	s.Body = v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
