// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNotifyConfigResponseBody) *DescribeNotifyConfigResponse
	GetBody() *DescribeNotifyConfigResponseBody
}

type DescribeNotifyConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNotifyConfigResponse) GetBody() *DescribeNotifyConfigResponseBody {
	return s.Body
}

func (s *DescribeNotifyConfigResponse) SetHeaders(v map[string]*string) *DescribeNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeNotifyConfigResponse) SetStatusCode(v int32) *DescribeNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNotifyConfigResponse) SetBody(v *DescribeNotifyConfigResponseBody) *DescribeNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeNotifyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
