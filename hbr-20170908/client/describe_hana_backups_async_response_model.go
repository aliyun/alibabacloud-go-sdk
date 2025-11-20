// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupsAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHanaBackupsAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHanaBackupsAsyncResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHanaBackupsAsyncResponseBody) *DescribeHanaBackupsAsyncResponse
	GetBody() *DescribeHanaBackupsAsyncResponseBody
}

type DescribeHanaBackupsAsyncResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHanaBackupsAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHanaBackupsAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupsAsyncResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupsAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHanaBackupsAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHanaBackupsAsyncResponse) GetBody() *DescribeHanaBackupsAsyncResponseBody {
	return s.Body
}

func (s *DescribeHanaBackupsAsyncResponse) SetHeaders(v map[string]*string) *DescribeHanaBackupsAsyncResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaBackupsAsyncResponse) SetStatusCode(v int32) *DescribeHanaBackupsAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponse) SetBody(v *DescribeHanaBackupsAsyncResponseBody) *DescribeHanaBackupsAsyncResponse {
	s.Body = v
	return s
}

func (s *DescribeHanaBackupsAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
