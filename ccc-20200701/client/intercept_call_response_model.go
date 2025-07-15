// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterceptCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InterceptCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InterceptCallResponse
	GetStatusCode() *int32
	SetBody(v *InterceptCallResponseBody) *InterceptCallResponse
	GetBody() *InterceptCallResponseBody
}

type InterceptCallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InterceptCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InterceptCallResponse) String() string {
	return dara.Prettify(s)
}

func (s InterceptCallResponse) GoString() string {
	return s.String()
}

func (s *InterceptCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InterceptCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InterceptCallResponse) GetBody() *InterceptCallResponseBody {
	return s.Body
}

func (s *InterceptCallResponse) SetHeaders(v map[string]*string) *InterceptCallResponse {
	s.Headers = v
	return s
}

func (s *InterceptCallResponse) SetStatusCode(v int32) *InterceptCallResponse {
	s.StatusCode = &v
	return s
}

func (s *InterceptCallResponse) SetBody(v *InterceptCallResponseBody) *InterceptCallResponse {
	s.Body = v
	return s
}

func (s *InterceptCallResponse) Validate() error {
	return dara.Validate(s)
}
