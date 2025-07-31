// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddImageLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddImageLibResponse
	GetStatusCode() *int32
	SetBody(v *AddImageLibResponseBody) *AddImageLibResponse
	GetBody() *AddImageLibResponseBody
}

type AddImageLibResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageLibResponse) String() string {
	return dara.Prettify(s)
}

func (s AddImageLibResponse) GoString() string {
	return s.String()
}

func (s *AddImageLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddImageLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddImageLibResponse) GetBody() *AddImageLibResponseBody {
	return s.Body
}

func (s *AddImageLibResponse) SetHeaders(v map[string]*string) *AddImageLibResponse {
	s.Headers = v
	return s
}

func (s *AddImageLibResponse) SetStatusCode(v int32) *AddImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageLibResponse) SetBody(v *AddImageLibResponseBody) *AddImageLibResponse {
	s.Body = v
	return s
}

func (s *AddImageLibResponse) Validate() error {
	return dara.Validate(s)
}
