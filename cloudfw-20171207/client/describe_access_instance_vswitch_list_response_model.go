// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceVSwitchListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessInstanceVSwitchListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessInstanceVSwitchListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessInstanceVSwitchListResponseBody) *DescribeAccessInstanceVSwitchListResponse
	GetBody() *DescribeAccessInstanceVSwitchListResponseBody
}

type DescribeAccessInstanceVSwitchListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessInstanceVSwitchListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessInstanceVSwitchListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVSwitchListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVSwitchListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessInstanceVSwitchListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessInstanceVSwitchListResponse) GetBody() *DescribeAccessInstanceVSwitchListResponseBody {
	return s.Body
}

func (s *DescribeAccessInstanceVSwitchListResponse) SetHeaders(v map[string]*string) *DescribeAccessInstanceVSwitchListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponse) SetStatusCode(v int32) *DescribeAccessInstanceVSwitchListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponse) SetBody(v *DescribeAccessInstanceVSwitchListResponseBody) *DescribeAccessInstanceVSwitchListResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
