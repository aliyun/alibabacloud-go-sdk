// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AddInstanceResponseBody) *AddInstanceResponse
	GetBody() *AddInstanceResponseBody
}

type AddInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceResponse) GoString() string {
	return s.String()
}

func (s *AddInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddInstanceResponse) GetBody() *AddInstanceResponseBody {
	return s.Body
}

func (s *AddInstanceResponse) SetHeaders(v map[string]*string) *AddInstanceResponse {
	s.Headers = v
	return s
}

func (s *AddInstanceResponse) SetStatusCode(v int32) *AddInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInstanceResponse) SetBody(v *AddInstanceResponseBody) *AddInstanceResponse {
	s.Body = v
	return s
}

func (s *AddInstanceResponse) Validate() error {
	return dara.Validate(s)
}
