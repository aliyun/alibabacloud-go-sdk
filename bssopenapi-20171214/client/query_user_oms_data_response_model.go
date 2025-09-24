// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserOmsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserOmsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserOmsDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserOmsDataResponseBody) *QueryUserOmsDataResponse
	GetBody() *QueryUserOmsDataResponseBody
}

type QueryUserOmsDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserOmsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserOmsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserOmsDataResponse) GoString() string {
	return s.String()
}

func (s *QueryUserOmsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserOmsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserOmsDataResponse) GetBody() *QueryUserOmsDataResponseBody {
	return s.Body
}

func (s *QueryUserOmsDataResponse) SetHeaders(v map[string]*string) *QueryUserOmsDataResponse {
	s.Headers = v
	return s
}

func (s *QueryUserOmsDataResponse) SetStatusCode(v int32) *QueryUserOmsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserOmsDataResponse) SetBody(v *QueryUserOmsDataResponseBody) *QueryUserOmsDataResponse {
	s.Body = v
	return s
}

func (s *QueryUserOmsDataResponse) Validate() error {
	return dara.Validate(s)
}
