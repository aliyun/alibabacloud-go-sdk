// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSaslUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSaslUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSaslUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSaslUsersResponseBody) *DescribeSaslUsersResponse
	GetBody() *DescribeSaslUsersResponseBody
}

type DescribeSaslUsersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSaslUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSaslUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSaslUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSaslUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSaslUsersResponse) GetBody() *DescribeSaslUsersResponseBody {
	return s.Body
}

func (s *DescribeSaslUsersResponse) SetHeaders(v map[string]*string) *DescribeSaslUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSaslUsersResponse) SetStatusCode(v int32) *DescribeSaslUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSaslUsersResponse) SetBody(v *DescribeSaslUsersResponseBody) *DescribeSaslUsersResponse {
	s.Body = v
	return s
}

func (s *DescribeSaslUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
