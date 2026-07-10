// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLangfuseUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLangfuseUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLangfuseUsersResponseBody) *DescribeLangfuseUsersResponse
	GetBody() *DescribeLangfuseUsersResponseBody
}

type DescribeLangfuseUsersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLangfuseUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLangfuseUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLangfuseUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLangfuseUsersResponse) GetBody() *DescribeLangfuseUsersResponseBody {
	return s.Body
}

func (s *DescribeLangfuseUsersResponse) SetHeaders(v map[string]*string) *DescribeLangfuseUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeLangfuseUsersResponse) SetStatusCode(v int32) *DescribeLangfuseUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLangfuseUsersResponse) SetBody(v *DescribeLangfuseUsersResponseBody) *DescribeLangfuseUsersResponse {
	s.Body = v
	return s
}

func (s *DescribeLangfuseUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
