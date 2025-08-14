// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySearchLibResponse
	GetStatusCode() *int32
	SetBody(v *QuerySearchLibResponseBody) *QuerySearchLibResponse
	GetBody() *QuerySearchLibResponseBody
}

type QuerySearchLibResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySearchLibResponse) GoString() string {
	return s.String()
}

func (s *QuerySearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySearchLibResponse) GetBody() *QuerySearchLibResponseBody {
	return s.Body
}

func (s *QuerySearchLibResponse) SetHeaders(v map[string]*string) *QuerySearchLibResponse {
	s.Headers = v
	return s
}

func (s *QuerySearchLibResponse) SetStatusCode(v int32) *QuerySearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySearchLibResponse) SetBody(v *QuerySearchLibResponseBody) *QuerySearchLibResponse {
	s.Body = v
	return s
}

func (s *QuerySearchLibResponse) Validate() error {
	return dara.Validate(s)
}
