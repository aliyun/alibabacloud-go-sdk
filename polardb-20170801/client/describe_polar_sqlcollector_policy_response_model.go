// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarSQLCollectorPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarSQLCollectorPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarSQLCollectorPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarSQLCollectorPolicyResponseBody) *DescribePolarSQLCollectorPolicyResponse
	GetBody() *DescribePolarSQLCollectorPolicyResponseBody
}

type DescribePolarSQLCollectorPolicyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarSQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarSQLCollectorPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarSQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarSQLCollectorPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarSQLCollectorPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarSQLCollectorPolicyResponse) GetBody() *DescribePolarSQLCollectorPolicyResponseBody {
	return s.Body
}

func (s *DescribePolarSQLCollectorPolicyResponse) SetHeaders(v map[string]*string) *DescribePolarSQLCollectorPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarSQLCollectorPolicyResponse) SetStatusCode(v int32) *DescribePolarSQLCollectorPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyResponse) SetBody(v *DescribePolarSQLCollectorPolicyResponseBody) *DescribePolarSQLCollectorPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribePolarSQLCollectorPolicyResponse) Validate() error {
	return dara.Validate(s)
}
