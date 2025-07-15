// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportOASTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImportOASTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImportOASTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImportOASTaskResponseBody) *DescribeImportOASTaskResponse
	GetBody() *DescribeImportOASTaskResponseBody
}

type DescribeImportOASTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImportOASTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImportOASTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportOASTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeImportOASTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImportOASTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImportOASTaskResponse) GetBody() *DescribeImportOASTaskResponseBody {
	return s.Body
}

func (s *DescribeImportOASTaskResponse) SetHeaders(v map[string]*string) *DescribeImportOASTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeImportOASTaskResponse) SetStatusCode(v int32) *DescribeImportOASTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImportOASTaskResponse) SetBody(v *DescribeImportOASTaskResponseBody) *DescribeImportOASTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeImportOASTaskResponse) Validate() error {
	return dara.Validate(s)
}
