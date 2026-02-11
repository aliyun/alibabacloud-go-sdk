// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportedLogCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImportedLogCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImportedLogCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImportedLogCountResponseBody) *DescribeImportedLogCountResponse
	GetBody() *DescribeImportedLogCountResponseBody
}

type DescribeImportedLogCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImportedLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImportedLogCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImportedLogCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImportedLogCountResponse) GetBody() *DescribeImportedLogCountResponseBody {
	return s.Body
}

func (s *DescribeImportedLogCountResponse) SetHeaders(v map[string]*string) *DescribeImportedLogCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeImportedLogCountResponse) SetStatusCode(v int32) *DescribeImportedLogCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImportedLogCountResponse) SetBody(v *DescribeImportedLogCountResponseBody) *DescribeImportedLogCountResponse {
	s.Body = v
	return s
}

func (s *DescribeImportedLogCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
