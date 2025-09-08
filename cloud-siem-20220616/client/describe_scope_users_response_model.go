// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScopeUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScopeUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScopeUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScopeUsersResponseBody) *DescribeScopeUsersResponse
	GetBody() *DescribeScopeUsersResponseBody
}

type DescribeScopeUsersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScopeUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScopeUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScopeUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScopeUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScopeUsersResponse) GetBody() *DescribeScopeUsersResponseBody {
	return s.Body
}

func (s *DescribeScopeUsersResponse) SetHeaders(v map[string]*string) *DescribeScopeUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeScopeUsersResponse) SetStatusCode(v int32) *DescribeScopeUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScopeUsersResponse) SetBody(v *DescribeScopeUsersResponseBody) *DescribeScopeUsersResponse {
	s.Body = v
	return s
}

func (s *DescribeScopeUsersResponse) Validate() error {
	return dara.Validate(s)
}
