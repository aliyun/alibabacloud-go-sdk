// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMessageAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMessageAppResponse
	GetStatusCode() *int32
	SetBody(v *QueryMessageAppResponseBody) *QueryMessageAppResponse
	GetBody() *QueryMessageAppResponseBody
}

type QueryMessageAppResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMessageAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMessageAppResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageAppResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMessageAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMessageAppResponse) GetBody() *QueryMessageAppResponseBody {
	return s.Body
}

func (s *QueryMessageAppResponse) SetHeaders(v map[string]*string) *QueryMessageAppResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageAppResponse) SetStatusCode(v int32) *QueryMessageAppResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageAppResponse) SetBody(v *QueryMessageAppResponseBody) *QueryMessageAppResponse {
	s.Body = v
	return s
}

func (s *QueryMessageAppResponse) Validate() error {
	return dara.Validate(s)
}
