// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNumberMccMncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNumberMccMncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNumberMccMncResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNumberMccMncResponseBody) *DescribeNumberMccMncResponse
	GetBody() *DescribeNumberMccMncResponseBody
}

type DescribeNumberMccMncResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNumberMccMncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNumberMccMncResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberMccMncResponse) GoString() string {
	return s.String()
}

func (s *DescribeNumberMccMncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNumberMccMncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNumberMccMncResponse) GetBody() *DescribeNumberMccMncResponseBody {
	return s.Body
}

func (s *DescribeNumberMccMncResponse) SetHeaders(v map[string]*string) *DescribeNumberMccMncResponse {
	s.Headers = v
	return s
}

func (s *DescribeNumberMccMncResponse) SetStatusCode(v int32) *DescribeNumberMccMncResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNumberMccMncResponse) SetBody(v *DescribeNumberMccMncResponseBody) *DescribeNumberMccMncResponse {
	s.Body = v
	return s
}

func (s *DescribeNumberMccMncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
