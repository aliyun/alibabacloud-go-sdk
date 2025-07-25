// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDohUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDohUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDohUserInfoResponseBody) *DescribeDohUserInfoResponse
	GetBody() *DescribeDohUserInfoResponseBody
}

type DescribeDohUserInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDohUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDohUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDohUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDohUserInfoResponse) GetBody() *DescribeDohUserInfoResponseBody {
	return s.Body
}

func (s *DescribeDohUserInfoResponse) SetHeaders(v map[string]*string) *DescribeDohUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohUserInfoResponse) SetStatusCode(v int32) *DescribeDohUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDohUserInfoResponse) SetBody(v *DescribeDohUserInfoResponseBody) *DescribeDohUserInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDohUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
