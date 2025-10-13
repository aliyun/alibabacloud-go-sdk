// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebInstanceLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebInstanceLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebInstanceLogsResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationInstanceLogsBody) *DescribeWebInstanceLogsResponse
	GetBody() *WebApplicationInstanceLogsBody
}

type DescribeWebInstanceLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationInstanceLogsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebInstanceLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebInstanceLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebInstanceLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebInstanceLogsResponse) GetBody() *WebApplicationInstanceLogsBody {
	return s.Body
}

func (s *DescribeWebInstanceLogsResponse) SetHeaders(v map[string]*string) *DescribeWebInstanceLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebInstanceLogsResponse) SetStatusCode(v int32) *DescribeWebInstanceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebInstanceLogsResponse) SetBody(v *WebApplicationInstanceLogsBody) *DescribeWebInstanceLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebInstanceLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
