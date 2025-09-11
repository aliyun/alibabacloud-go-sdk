// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryExtCodeSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryExtCodeSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryExtCodeSignResponse
	GetStatusCode() *int32
	SetBody(v *QueryExtCodeSignResponseBody) *QueryExtCodeSignResponse
	GetBody() *QueryExtCodeSignResponseBody
}

type QueryExtCodeSignResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryExtCodeSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryExtCodeSignResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryExtCodeSignResponse) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryExtCodeSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryExtCodeSignResponse) GetBody() *QueryExtCodeSignResponseBody {
	return s.Body
}

func (s *QueryExtCodeSignResponse) SetHeaders(v map[string]*string) *QueryExtCodeSignResponse {
	s.Headers = v
	return s
}

func (s *QueryExtCodeSignResponse) SetStatusCode(v int32) *QueryExtCodeSignResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryExtCodeSignResponse) SetBody(v *QueryExtCodeSignResponseBody) *QueryExtCodeSignResponse {
	s.Body = v
	return s
}

func (s *QueryExtCodeSignResponse) Validate() error {
	return dara.Validate(s)
}
