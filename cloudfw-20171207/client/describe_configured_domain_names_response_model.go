// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfiguredDomainNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConfiguredDomainNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConfiguredDomainNamesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConfiguredDomainNamesResponseBody) *DescribeConfiguredDomainNamesResponse
	GetBody() *DescribeConfiguredDomainNamesResponseBody
}

type DescribeConfiguredDomainNamesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfiguredDomainNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfiguredDomainNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfiguredDomainNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfiguredDomainNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConfiguredDomainNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConfiguredDomainNamesResponse) GetBody() *DescribeConfiguredDomainNamesResponseBody {
	return s.Body
}

func (s *DescribeConfiguredDomainNamesResponse) SetHeaders(v map[string]*string) *DescribeConfiguredDomainNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfiguredDomainNamesResponse) SetStatusCode(v int32) *DescribeConfiguredDomainNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfiguredDomainNamesResponse) SetBody(v *DescribeConfiguredDomainNamesResponseBody) *DescribeConfiguredDomainNamesResponse {
	s.Body = v
	return s
}

func (s *DescribeConfiguredDomainNamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
