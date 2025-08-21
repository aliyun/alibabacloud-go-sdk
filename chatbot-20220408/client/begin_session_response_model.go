// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BeginSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BeginSessionResponse
	GetStatusCode() *int32
	SetBody(v *BeginSessionResponseBody) *BeginSessionResponse
	GetBody() *BeginSessionResponseBody
}

type BeginSessionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BeginSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BeginSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s BeginSessionResponse) GoString() string {
	return s.String()
}

func (s *BeginSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BeginSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BeginSessionResponse) GetBody() *BeginSessionResponseBody {
	return s.Body
}

func (s *BeginSessionResponse) SetHeaders(v map[string]*string) *BeginSessionResponse {
	s.Headers = v
	return s
}

func (s *BeginSessionResponse) SetStatusCode(v int32) *BeginSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *BeginSessionResponse) SetBody(v *BeginSessionResponseBody) *BeginSessionResponse {
	s.Body = v
	return s
}

func (s *BeginSessionResponse) Validate() error {
	return dara.Validate(s)
}
