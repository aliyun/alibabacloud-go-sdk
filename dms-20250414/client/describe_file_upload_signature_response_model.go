// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileUploadSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFileUploadSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFileUploadSignatureResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFileUploadSignatureResponseBody) *DescribeFileUploadSignatureResponse
	GetBody() *DescribeFileUploadSignatureResponseBody
}

type DescribeFileUploadSignatureResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileUploadSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileUploadSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileUploadSignatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFileUploadSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFileUploadSignatureResponse) GetBody() *DescribeFileUploadSignatureResponseBody {
	return s.Body
}

func (s *DescribeFileUploadSignatureResponse) SetHeaders(v map[string]*string) *DescribeFileUploadSignatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileUploadSignatureResponse) SetStatusCode(v int32) *DescribeFileUploadSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileUploadSignatureResponse) SetBody(v *DescribeFileUploadSignatureResponseBody) *DescribeFileUploadSignatureResponse {
	s.Body = v
	return s
}

func (s *DescribeFileUploadSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
