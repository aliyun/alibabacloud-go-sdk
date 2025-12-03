// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreTablesResponseBody) *DescribeRestoreTablesResponse
	GetBody() *DescribeRestoreTablesResponseBody
}

type DescribeRestoreTablesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreTablesResponse) GetBody() *DescribeRestoreTablesResponseBody {
	return s.Body
}

func (s *DescribeRestoreTablesResponse) SetHeaders(v map[string]*string) *DescribeRestoreTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreTablesResponse) SetStatusCode(v int32) *DescribeRestoreTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreTablesResponse) SetBody(v *DescribeRestoreTablesResponseBody) *DescribeRestoreTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
