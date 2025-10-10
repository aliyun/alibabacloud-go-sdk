// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppSecurityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAppSecurityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAppSecurityInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryAppSecurityInfoResponseBody) *QueryAppSecurityInfoResponse
	GetBody() *QueryAppSecurityInfoResponseBody
}

type QueryAppSecurityInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppSecurityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppSecurityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAppSecurityInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAppSecurityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAppSecurityInfoResponse) GetBody() *QueryAppSecurityInfoResponseBody {
	return s.Body
}

func (s *QueryAppSecurityInfoResponse) SetHeaders(v map[string]*string) *QueryAppSecurityInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAppSecurityInfoResponse) SetStatusCode(v int32) *QueryAppSecurityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppSecurityInfoResponse) SetBody(v *QueryAppSecurityInfoResponseBody) *QueryAppSecurityInfoResponse {
	s.Body = v
	return s
}

func (s *QueryAppSecurityInfoResponse) Validate() error {
	return dara.Validate(s)
}
