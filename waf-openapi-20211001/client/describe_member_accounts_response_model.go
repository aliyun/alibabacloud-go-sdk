// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMemberAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMemberAccountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMemberAccountsResponseBody) *DescribeMemberAccountsResponse
	GetBody() *DescribeMemberAccountsResponseBody
}

type DescribeMemberAccountsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMemberAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMemberAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMemberAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMemberAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMemberAccountsResponse) GetBody() *DescribeMemberAccountsResponseBody {
	return s.Body
}

func (s *DescribeMemberAccountsResponse) SetHeaders(v map[string]*string) *DescribeMemberAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMemberAccountsResponse) SetStatusCode(v int32) *DescribeMemberAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMemberAccountsResponse) SetBody(v *DescribeMemberAccountsResponseBody) *DescribeMemberAccountsResponse {
	s.Body = v
	return s
}

func (s *DescribeMemberAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
