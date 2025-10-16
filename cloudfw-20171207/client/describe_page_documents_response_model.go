// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePageDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePageDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePageDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePageDocumentsResponseBody) *DescribePageDocumentsResponse
	GetBody() *DescribePageDocumentsResponseBody
}

type DescribePageDocumentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePageDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePageDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePageDocumentsResponse) GoString() string {
	return s.String()
}

func (s *DescribePageDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePageDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePageDocumentsResponse) GetBody() *DescribePageDocumentsResponseBody {
	return s.Body
}

func (s *DescribePageDocumentsResponse) SetHeaders(v map[string]*string) *DescribePageDocumentsResponse {
	s.Headers = v
	return s
}

func (s *DescribePageDocumentsResponse) SetStatusCode(v int32) *DescribePageDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePageDocumentsResponse) SetBody(v *DescribePageDocumentsResponseBody) *DescribePageDocumentsResponse {
	s.Body = v
	return s
}

func (s *DescribePageDocumentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
