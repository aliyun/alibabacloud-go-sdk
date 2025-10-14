// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNCInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNCInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNCInformationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNCInformationResponseBody) *DescribeNCInformationResponse
	GetBody() *DescribeNCInformationResponseBody
}

type DescribeNCInformationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNCInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNCInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponse) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNCInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNCInformationResponse) GetBody() *DescribeNCInformationResponseBody {
	return s.Body
}

func (s *DescribeNCInformationResponse) SetHeaders(v map[string]*string) *DescribeNCInformationResponse {
	s.Headers = v
	return s
}

func (s *DescribeNCInformationResponse) SetStatusCode(v int32) *DescribeNCInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNCInformationResponse) SetBody(v *DescribeNCInformationResponseBody) *DescribeNCInformationResponse {
	s.Body = v
	return s
}

func (s *DescribeNCInformationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
