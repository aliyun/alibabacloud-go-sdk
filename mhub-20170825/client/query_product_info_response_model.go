// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryProductInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryProductInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryProductInfoResponseBody) *QueryProductInfoResponse
	GetBody() *QueryProductInfoResponseBody
}

type QueryProductInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProductInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProductInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryProductInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryProductInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryProductInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryProductInfoResponse) GetBody() *QueryProductInfoResponseBody {
	return s.Body
}

func (s *QueryProductInfoResponse) SetHeaders(v map[string]*string) *QueryProductInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryProductInfoResponse) SetStatusCode(v int32) *QueryProductInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProductInfoResponse) SetBody(v *QueryProductInfoResponseBody) *QueryProductInfoResponse {
	s.Body = v
	return s
}

func (s *QueryProductInfoResponse) Validate() error {
	return dara.Validate(s)
}
