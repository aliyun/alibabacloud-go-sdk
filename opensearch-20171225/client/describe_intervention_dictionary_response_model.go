// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterventionDictionaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInterventionDictionaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInterventionDictionaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInterventionDictionaryResponseBody) *DescribeInterventionDictionaryResponse
	GetBody() *DescribeInterventionDictionaryResponseBody
}

type DescribeInterventionDictionaryResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInterventionDictionaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInterventionDictionaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInterventionDictionaryResponse) GetBody() *DescribeInterventionDictionaryResponseBody {
	return s.Body
}

func (s *DescribeInterventionDictionaryResponse) SetHeaders(v map[string]*string) *DescribeInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeInterventionDictionaryResponse) SetStatusCode(v int32) *DescribeInterventionDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInterventionDictionaryResponse) SetBody(v *DescribeInterventionDictionaryResponseBody) *DescribeInterventionDictionaryResponse {
	s.Body = v
	return s
}

func (s *DescribeInterventionDictionaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
