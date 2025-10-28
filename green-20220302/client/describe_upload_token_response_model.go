// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUploadTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUploadTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUploadTokenResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUploadTokenResponseBody) *DescribeUploadTokenResponse
	GetBody() *DescribeUploadTokenResponseBody
}

type DescribeUploadTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUploadTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeUploadTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUploadTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUploadTokenResponse) GetBody() *DescribeUploadTokenResponseBody {
	return s.Body
}

func (s *DescribeUploadTokenResponse) SetHeaders(v map[string]*string) *DescribeUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeUploadTokenResponse) SetStatusCode(v int32) *DescribeUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUploadTokenResponse) SetBody(v *DescribeUploadTokenResponseBody) *DescribeUploadTokenResponse {
	s.Body = v
	return s
}

func (s *DescribeUploadTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
