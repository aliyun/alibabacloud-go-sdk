// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebAccessModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebAccessModeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebAccessModeResponseBody) *DescribeWebAccessModeResponse
	GetBody() *DescribeWebAccessModeResponseBody
}

type DescribeWebAccessModeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebAccessModeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebAccessModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebAccessModeResponse) GetBody() *DescribeWebAccessModeResponseBody {
	return s.Body
}

func (s *DescribeWebAccessModeResponse) SetHeaders(v map[string]*string) *DescribeWebAccessModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAccessModeResponse) SetStatusCode(v int32) *DescribeWebAccessModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebAccessModeResponse) SetBody(v *DescribeWebAccessModeResponseBody) *DescribeWebAccessModeResponse {
	s.Body = v
	return s
}

func (s *DescribeWebAccessModeResponse) Validate() error {
	return dara.Validate(s)
}
