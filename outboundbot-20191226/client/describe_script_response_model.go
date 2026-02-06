// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScriptResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScriptResponseBody) *DescribeScriptResponse
	GetBody() *DescribeScriptResponseBody
}

type DescribeScriptResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptResponse) GoString() string {
	return s.String()
}

func (s *DescribeScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScriptResponse) GetBody() *DescribeScriptResponseBody {
	return s.Body
}

func (s *DescribeScriptResponse) SetHeaders(v map[string]*string) *DescribeScriptResponse {
	s.Headers = v
	return s
}

func (s *DescribeScriptResponse) SetStatusCode(v int32) *DescribeScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScriptResponse) SetBody(v *DescribeScriptResponseBody) *DescribeScriptResponse {
	s.Body = v
	return s
}

func (s *DescribeScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
