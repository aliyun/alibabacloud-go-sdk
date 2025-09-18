// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryArmsEnableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryArmsEnableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryArmsEnableResponse
	GetStatusCode() *int32
	SetBody(v *QueryArmsEnableResponseBody) *QueryArmsEnableResponse
	GetBody() *QueryArmsEnableResponseBody
}

type QueryArmsEnableResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryArmsEnableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryArmsEnableResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryArmsEnableResponse) GoString() string {
	return s.String()
}

func (s *QueryArmsEnableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryArmsEnableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryArmsEnableResponse) GetBody() *QueryArmsEnableResponseBody {
	return s.Body
}

func (s *QueryArmsEnableResponse) SetHeaders(v map[string]*string) *QueryArmsEnableResponse {
	s.Headers = v
	return s
}

func (s *QueryArmsEnableResponse) SetStatusCode(v int32) *QueryArmsEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryArmsEnableResponse) SetBody(v *QueryArmsEnableResponseBody) *QueryArmsEnableResponse {
	s.Body = v
	return s
}

func (s *QueryArmsEnableResponse) Validate() error {
	return dara.Validate(s)
}
