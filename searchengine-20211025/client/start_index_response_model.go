// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartIndexResponse
	GetStatusCode() *int32
	SetBody(v *StartIndexResponseBody) *StartIndexResponse
	GetBody() *StartIndexResponseBody
}

type StartIndexResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s StartIndexResponse) GoString() string {
	return s.String()
}

func (s *StartIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartIndexResponse) GetBody() *StartIndexResponseBody {
	return s.Body
}

func (s *StartIndexResponse) SetHeaders(v map[string]*string) *StartIndexResponse {
	s.Headers = v
	return s
}

func (s *StartIndexResponse) SetStatusCode(v int32) *StartIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *StartIndexResponse) SetBody(v *StartIndexResponseBody) *StartIndexResponse {
	s.Body = v
	return s
}

func (s *StartIndexResponse) Validate() error {
	return dara.Validate(s)
}
