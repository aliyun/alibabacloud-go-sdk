// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHDMInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddHDMInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddHDMInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AddHDMInstanceResponseBody) *AddHDMInstanceResponse
	GetBody() *AddHDMInstanceResponseBody
}

type AddHDMInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddHDMInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddHDMInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddHDMInstanceResponse) GoString() string {
	return s.String()
}

func (s *AddHDMInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddHDMInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddHDMInstanceResponse) GetBody() *AddHDMInstanceResponseBody {
	return s.Body
}

func (s *AddHDMInstanceResponse) SetHeaders(v map[string]*string) *AddHDMInstanceResponse {
	s.Headers = v
	return s
}

func (s *AddHDMInstanceResponse) SetStatusCode(v int32) *AddHDMInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHDMInstanceResponse) SetBody(v *AddHDMInstanceResponseBody) *AddHDMInstanceResponse {
	s.Body = v
	return s
}

func (s *AddHDMInstanceResponse) Validate() error {
	return dara.Validate(s)
}
