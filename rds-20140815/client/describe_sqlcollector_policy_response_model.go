// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLCollectorPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLCollectorPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLCollectorPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLCollectorPolicyResponseBody) *DescribeSQLCollectorPolicyResponse
	GetBody() *DescribeSQLCollectorPolicyResponseBody
}

type DescribeSQLCollectorPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLCollectorPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLCollectorPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLCollectorPolicyResponse) GetBody() *DescribeSQLCollectorPolicyResponseBody {
	return s.Body
}

func (s *DescribeSQLCollectorPolicyResponse) SetHeaders(v map[string]*string) *DescribeSQLCollectorPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLCollectorPolicyResponse) SetStatusCode(v int32) *DescribeSQLCollectorPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLCollectorPolicyResponse) SetBody(v *DescribeSQLCollectorPolicyResponseBody) *DescribeSQLCollectorPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLCollectorPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
