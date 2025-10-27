// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanTaskProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScanTaskProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScanTaskProgressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScanTaskProgressResponseBody) *DescribeScanTaskProgressResponse
	GetBody() *DescribeScanTaskProgressResponseBody
}

type DescribeScanTaskProgressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScanTaskProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScanTaskProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScanTaskProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScanTaskProgressResponse) GetBody() *DescribeScanTaskProgressResponseBody {
	return s.Body
}

func (s *DescribeScanTaskProgressResponse) SetHeaders(v map[string]*string) *DescribeScanTaskProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeScanTaskProgressResponse) SetStatusCode(v int32) *DescribeScanTaskProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScanTaskProgressResponse) SetBody(v *DescribeScanTaskProgressResponseBody) *DescribeScanTaskProgressResponse {
	s.Body = v
	return s
}

func (s *DescribeScanTaskProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
