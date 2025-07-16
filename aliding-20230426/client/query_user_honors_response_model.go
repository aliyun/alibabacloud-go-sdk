// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserHonorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserHonorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserHonorsResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserHonorsResponseBody) *QueryUserHonorsResponse
	GetBody() *QueryUserHonorsResponseBody
}

type QueryUserHonorsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserHonorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserHonorsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserHonorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserHonorsResponse) GetBody() *QueryUserHonorsResponseBody {
	return s.Body
}

func (s *QueryUserHonorsResponse) SetHeaders(v map[string]*string) *QueryUserHonorsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserHonorsResponse) SetStatusCode(v int32) *QueryUserHonorsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserHonorsResponse) SetBody(v *QueryUserHonorsResponseBody) *QueryUserHonorsResponse {
	s.Body = v
	return s
}

func (s *QueryUserHonorsResponse) Validate() error {
	return dara.Validate(s)
}
