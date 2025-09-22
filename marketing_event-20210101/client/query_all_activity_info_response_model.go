// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllActivityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAllActivityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAllActivityInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryAllActivityInfoResponseBody) *QueryAllActivityInfoResponse
	GetBody() *QueryAllActivityInfoResponseBody
}

type QueryAllActivityInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllActivityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAllActivityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAllActivityInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAllActivityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAllActivityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAllActivityInfoResponse) GetBody() *QueryAllActivityInfoResponseBody {
	return s.Body
}

func (s *QueryAllActivityInfoResponse) SetHeaders(v map[string]*string) *QueryAllActivityInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAllActivityInfoResponse) SetStatusCode(v int32) *QueryAllActivityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllActivityInfoResponse) SetBody(v *QueryAllActivityInfoResponseBody) *QueryAllActivityInfoResponse {
	s.Body = v
	return s
}

func (s *QueryAllActivityInfoResponse) Validate() error {
	return dara.Validate(s)
}
