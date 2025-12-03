// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubscriptionPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubscriptionPermissionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubscriptionPermissionResponseBody) *DescribeSubscriptionPermissionResponse
	GetBody() *DescribeSubscriptionPermissionResponseBody
}

type DescribeSubscriptionPermissionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubscriptionPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubscriptionPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubscriptionPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubscriptionPermissionResponse) GetBody() *DescribeSubscriptionPermissionResponseBody {
	return s.Body
}

func (s *DescribeSubscriptionPermissionResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionPermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionPermissionResponse) SetStatusCode(v int32) *DescribeSubscriptionPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionPermissionResponse) SetBody(v *DescribeSubscriptionPermissionResponseBody) *DescribeSubscriptionPermissionResponse {
	s.Body = v
	return s
}

func (s *DescribeSubscriptionPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
