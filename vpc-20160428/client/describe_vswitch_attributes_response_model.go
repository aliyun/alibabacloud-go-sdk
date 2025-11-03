// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVSwitchAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVSwitchAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVSwitchAttributesResponseBody) *DescribeVSwitchAttributesResponse
	GetBody() *DescribeVSwitchAttributesResponseBody
}

type DescribeVSwitchAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVSwitchAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVSwitchAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVSwitchAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVSwitchAttributesResponse) GetBody() *DescribeVSwitchAttributesResponseBody {
	return s.Body
}

func (s *DescribeVSwitchAttributesResponse) SetHeaders(v map[string]*string) *DescribeVSwitchAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVSwitchAttributesResponse) SetStatusCode(v int32) *DescribeVSwitchAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVSwitchAttributesResponse) SetBody(v *DescribeVSwitchAttributesResponseBody) *DescribeVSwitchAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeVSwitchAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
