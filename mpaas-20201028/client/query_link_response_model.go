// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLinkResponse
	GetStatusCode() *int32
	SetBody(v *QueryLinkResponseBody) *QueryLinkResponse
	GetBody() *QueryLinkResponseBody
}

type QueryLinkResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLinkResponse) GoString() string {
	return s.String()
}

func (s *QueryLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLinkResponse) GetBody() *QueryLinkResponseBody {
	return s.Body
}

func (s *QueryLinkResponse) SetHeaders(v map[string]*string) *QueryLinkResponse {
	s.Headers = v
	return s
}

func (s *QueryLinkResponse) SetStatusCode(v int32) *QueryLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLinkResponse) SetBody(v *QueryLinkResponseBody) *QueryLinkResponse {
	s.Body = v
	return s
}

func (s *QueryLinkResponse) Validate() error {
	return dara.Validate(s)
}
