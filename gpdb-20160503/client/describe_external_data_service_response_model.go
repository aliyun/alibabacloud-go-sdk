// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExternalDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExternalDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExternalDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExternalDataServiceResponseBody) *DescribeExternalDataServiceResponse
	GetBody() *DescribeExternalDataServiceResponseBody
}

type DescribeExternalDataServiceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExternalDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExternalDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExternalDataServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeExternalDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExternalDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExternalDataServiceResponse) GetBody() *DescribeExternalDataServiceResponseBody {
	return s.Body
}

func (s *DescribeExternalDataServiceResponse) SetHeaders(v map[string]*string) *DescribeExternalDataServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeExternalDataServiceResponse) SetStatusCode(v int32) *DescribeExternalDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExternalDataServiceResponse) SetBody(v *DescribeExternalDataServiceResponseBody) *DescribeExternalDataServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeExternalDataServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
