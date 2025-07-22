// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePubUserListBySubUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePubUserListBySubUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePubUserListBySubUserResponse
	GetStatusCode() *int32
	SetBody(v *DescribePubUserListBySubUserResponseBody) *DescribePubUserListBySubUserResponse
	GetBody() *DescribePubUserListBySubUserResponseBody
}

type DescribePubUserListBySubUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePubUserListBySubUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePubUserListBySubUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePubUserListBySubUserResponse) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePubUserListBySubUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePubUserListBySubUserResponse) GetBody() *DescribePubUserListBySubUserResponseBody {
	return s.Body
}

func (s *DescribePubUserListBySubUserResponse) SetHeaders(v map[string]*string) *DescribePubUserListBySubUserResponse {
	s.Headers = v
	return s
}

func (s *DescribePubUserListBySubUserResponse) SetStatusCode(v int32) *DescribePubUserListBySubUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePubUserListBySubUserResponse) SetBody(v *DescribePubUserListBySubUserResponseBody) *DescribePubUserListBySubUserResponse {
	s.Body = v
	return s
}

func (s *DescribePubUserListBySubUserResponse) Validate() error {
	return dara.Validate(s)
}
