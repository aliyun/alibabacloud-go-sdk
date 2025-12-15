// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse
	GetBody() *DescribeBackupsResponseBody
}

type DescribeBackupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupsResponse) GetBody() *DescribeBackupsResponseBody {
	return s.Body
}

func (s *DescribeBackupsResponse) SetHeaders(v map[string]*string) *DescribeBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupsResponse) SetStatusCode(v int32) *DescribeBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupsResponse) SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
