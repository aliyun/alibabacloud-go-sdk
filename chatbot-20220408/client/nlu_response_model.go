// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNluResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *NluResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *NluResponse
	GetStatusCode() *int32
	SetBody(v *NluResponseBody) *NluResponse
	GetBody() *NluResponseBody
}

type NluResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NluResponseBody   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NluResponse) String() string {
	return dara.Prettify(s)
}

func (s NluResponse) GoString() string {
	return s.String()
}

func (s *NluResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *NluResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *NluResponse) GetBody() *NluResponseBody {
	return s.Body
}

func (s *NluResponse) SetHeaders(v map[string]*string) *NluResponse {
	s.Headers = v
	return s
}

func (s *NluResponse) SetStatusCode(v int32) *NluResponse {
	s.StatusCode = &v
	return s
}

func (s *NluResponse) SetBody(v *NluResponseBody) *NluResponse {
	s.Body = v
	return s
}

func (s *NluResponse) Validate() error {
	return dara.Validate(s)
}
