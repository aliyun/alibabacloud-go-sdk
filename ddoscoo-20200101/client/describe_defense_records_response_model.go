// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseRecordsResponseBody) *DescribeDefenseRecordsResponse
	GetBody() *DescribeDefenseRecordsResponseBody
}

type DescribeDefenseRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseRecordsResponse) GetBody() *DescribeDefenseRecordsResponseBody {
	return s.Body
}

func (s *DescribeDefenseRecordsResponse) SetHeaders(v map[string]*string) *DescribeDefenseRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRecordsResponse) SetStatusCode(v int32) *DescribeDefenseRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseRecordsResponse) SetBody(v *DescribeDefenseRecordsResponseBody) *DescribeDefenseRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
