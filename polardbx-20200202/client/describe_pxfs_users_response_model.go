// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePxfsUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePxfsUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribePxfsUsersResponseBody) *DescribePxfsUsersResponse
	GetBody() *DescribePxfsUsersResponseBody
}

type DescribePxfsUsersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePxfsUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePxfsUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribePxfsUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePxfsUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePxfsUsersResponse) GetBody() *DescribePxfsUsersResponseBody {
	return s.Body
}

func (s *DescribePxfsUsersResponse) SetHeaders(v map[string]*string) *DescribePxfsUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribePxfsUsersResponse) SetStatusCode(v int32) *DescribePxfsUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePxfsUsersResponse) SetBody(v *DescribePxfsUsersResponseBody) *DescribePxfsUsersResponse {
	s.Body = v
	return s
}

func (s *DescribePxfsUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
