// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelOperatorApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelOperatorApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelOperatorApiKeyResponseBody) *DescribeModelOperatorApiKeyResponse
	GetBody() *DescribeModelOperatorApiKeyResponseBody
}

type DescribeModelOperatorApiKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelOperatorApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelOperatorApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorApiKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelOperatorApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelOperatorApiKeyResponse) GetBody() *DescribeModelOperatorApiKeyResponseBody {
	return s.Body
}

func (s *DescribeModelOperatorApiKeyResponse) SetHeaders(v map[string]*string) *DescribeModelOperatorApiKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelOperatorApiKeyResponse) SetStatusCode(v int32) *DescribeModelOperatorApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelOperatorApiKeyResponse) SetBody(v *DescribeModelOperatorApiKeyResponseBody) *DescribeModelOperatorApiKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeModelOperatorApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
