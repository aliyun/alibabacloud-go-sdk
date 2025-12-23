// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopMetadataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopMetadataResponseBody) *DescribeDesktopMetadataResponse
	GetBody() *DescribeDesktopMetadataResponseBody
}

type DescribeDesktopMetadataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopMetadataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopMetadataResponse) GetBody() *DescribeDesktopMetadataResponseBody {
	return s.Body
}

func (s *DescribeDesktopMetadataResponse) SetHeaders(v map[string]*string) *DescribeDesktopMetadataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopMetadataResponse) SetStatusCode(v int32) *DescribeDesktopMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopMetadataResponse) SetBody(v *DescribeDesktopMetadataResponseBody) *DescribeDesktopMetadataResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
