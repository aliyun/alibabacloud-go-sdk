// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAddonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAddonResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAddonResponseBody) *DescribeAddonResponse
	GetBody() *DescribeAddonResponseBody
}

type DescribeAddonResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddonResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAddonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAddonResponse) GetBody() *DescribeAddonResponseBody {
	return s.Body
}

func (s *DescribeAddonResponse) SetHeaders(v map[string]*string) *DescribeAddonResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonResponse) SetStatusCode(v int32) *DescribeAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddonResponse) SetBody(v *DescribeAddonResponseBody) *DescribeAddonResponse {
	s.Body = v
	return s
}

func (s *DescribeAddonResponse) Validate() error {
	return dara.Validate(s)
}
