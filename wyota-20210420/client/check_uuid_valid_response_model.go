// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUuidValidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUuidValidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUuidValidResponse
	GetStatusCode() *int32
	SetBody(v *CheckUuidValidResponseBody) *CheckUuidValidResponse
	GetBody() *CheckUuidValidResponseBody
}

type CheckUuidValidResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUuidValidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUuidValidResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUuidValidResponse) GoString() string {
	return s.String()
}

func (s *CheckUuidValidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUuidValidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUuidValidResponse) GetBody() *CheckUuidValidResponseBody {
	return s.Body
}

func (s *CheckUuidValidResponse) SetHeaders(v map[string]*string) *CheckUuidValidResponse {
	s.Headers = v
	return s
}

func (s *CheckUuidValidResponse) SetStatusCode(v int32) *CheckUuidValidResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUuidValidResponse) SetBody(v *CheckUuidValidResponseBody) *CheckUuidValidResponse {
	s.Body = v
	return s
}

func (s *CheckUuidValidResponse) Validate() error {
	return dara.Validate(s)
}
