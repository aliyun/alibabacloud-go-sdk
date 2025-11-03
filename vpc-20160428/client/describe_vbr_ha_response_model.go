// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVbrHaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVbrHaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVbrHaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVbrHaResponseBody) *DescribeVbrHaResponse
	GetBody() *DescribeVbrHaResponseBody
}

type DescribeVbrHaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVbrHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVbrHaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVbrHaResponse) GoString() string {
	return s.String()
}

func (s *DescribeVbrHaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVbrHaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVbrHaResponse) GetBody() *DescribeVbrHaResponseBody {
	return s.Body
}

func (s *DescribeVbrHaResponse) SetHeaders(v map[string]*string) *DescribeVbrHaResponse {
	s.Headers = v
	return s
}

func (s *DescribeVbrHaResponse) SetStatusCode(v int32) *DescribeVbrHaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVbrHaResponse) SetBody(v *DescribeVbrHaResponseBody) *DescribeVbrHaResponse {
	s.Body = v
	return s
}

func (s *DescribeVbrHaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
