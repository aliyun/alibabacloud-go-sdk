// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationBackupsResponseBody) *DescribeApplicationBackupsResponse
	GetBody() *DescribeApplicationBackupsResponseBody
}

type DescribeApplicationBackupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationBackupsResponse) GetBody() *DescribeApplicationBackupsResponseBody {
	return s.Body
}

func (s *DescribeApplicationBackupsResponse) SetHeaders(v map[string]*string) *DescribeApplicationBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationBackupsResponse) SetStatusCode(v int32) *DescribeApplicationBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationBackupsResponse) SetBody(v *DescribeApplicationBackupsResponseBody) *DescribeApplicationBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationBackupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
