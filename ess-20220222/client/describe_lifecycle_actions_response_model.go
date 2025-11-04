// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecycleActionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLifecycleActionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLifecycleActionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLifecycleActionsResponseBody) *DescribeLifecycleActionsResponse
	GetBody() *DescribeLifecycleActionsResponseBody
}

type DescribeLifecycleActionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLifecycleActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLifecycleActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecycleActionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLifecycleActionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLifecycleActionsResponse) GetBody() *DescribeLifecycleActionsResponseBody {
	return s.Body
}

func (s *DescribeLifecycleActionsResponse) SetHeaders(v map[string]*string) *DescribeLifecycleActionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLifecycleActionsResponse) SetStatusCode(v int32) *DescribeLifecycleActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLifecycleActionsResponse) SetBody(v *DescribeLifecycleActionsResponseBody) *DescribeLifecycleActionsResponse {
	s.Body = v
	return s
}

func (s *DescribeLifecycleActionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
