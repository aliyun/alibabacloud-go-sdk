// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIPv6TranslatorEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIPv6TranslatorEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIPv6TranslatorEntriesResponseBody) *DescribeIPv6TranslatorEntriesResponse
	GetBody() *DescribeIPv6TranslatorEntriesResponseBody
}

type DescribeIPv6TranslatorEntriesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIPv6TranslatorEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIPv6TranslatorEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIPv6TranslatorEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIPv6TranslatorEntriesResponse) GetBody() *DescribeIPv6TranslatorEntriesResponseBody {
	return s.Body
}

func (s *DescribeIPv6TranslatorEntriesResponse) SetHeaders(v map[string]*string) *DescribeIPv6TranslatorEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponse) SetStatusCode(v int32) *DescribeIPv6TranslatorEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponse) SetBody(v *DescribeIPv6TranslatorEntriesResponseBody) *DescribeIPv6TranslatorEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
