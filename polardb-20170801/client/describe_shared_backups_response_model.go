// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSharedBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSharedBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSharedBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSharedBackupsResponseBody) *DescribeSharedBackupsResponse
	GetBody() *DescribeSharedBackupsResponseBody
}

type DescribeSharedBackupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSharedBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSharedBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSharedBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSharedBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSharedBackupsResponse) GetBody() *DescribeSharedBackupsResponseBody {
	return s.Body
}

func (s *DescribeSharedBackupsResponse) SetHeaders(v map[string]*string) *DescribeSharedBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSharedBackupsResponse) SetStatusCode(v int32) *DescribeSharedBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSharedBackupsResponse) SetBody(v *DescribeSharedBackupsResponseBody) *DescribeSharedBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeSharedBackupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
