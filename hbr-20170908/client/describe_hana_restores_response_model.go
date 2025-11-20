// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaRestoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHanaRestoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHanaRestoresResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHanaRestoresResponseBody) *DescribeHanaRestoresResponse
	GetBody() *DescribeHanaRestoresResponseBody
}

type DescribeHanaRestoresResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHanaRestoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHanaRestoresResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaRestoresResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHanaRestoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHanaRestoresResponse) GetBody() *DescribeHanaRestoresResponseBody {
	return s.Body
}

func (s *DescribeHanaRestoresResponse) SetHeaders(v map[string]*string) *DescribeHanaRestoresResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaRestoresResponse) SetStatusCode(v int32) *DescribeHanaRestoresResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaRestoresResponse) SetBody(v *DescribeHanaRestoresResponseBody) *DescribeHanaRestoresResponse {
	s.Body = v
	return s
}

func (s *DescribeHanaRestoresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
