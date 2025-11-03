// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGrantRulesToCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGrantRulesToCenResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGrantRulesToCenResponseBody) *DescribeGrantRulesToCenResponse
	GetBody() *DescribeGrantRulesToCenResponseBody
}

type DescribeGrantRulesToCenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGrantRulesToCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGrantRulesToCenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenResponse) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGrantRulesToCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGrantRulesToCenResponse) GetBody() *DescribeGrantRulesToCenResponseBody {
	return s.Body
}

func (s *DescribeGrantRulesToCenResponse) SetHeaders(v map[string]*string) *DescribeGrantRulesToCenResponse {
	s.Headers = v
	return s
}

func (s *DescribeGrantRulesToCenResponse) SetStatusCode(v int32) *DescribeGrantRulesToCenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGrantRulesToCenResponse) SetBody(v *DescribeGrantRulesToCenResponseBody) *DescribeGrantRulesToCenResponse {
	s.Body = v
	return s
}

func (s *DescribeGrantRulesToCenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
