// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCasterComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCasterComponentResponse
	GetStatusCode() *int32
	SetBody(v *AddCasterComponentResponseBody) *AddCasterComponentResponse
	GetBody() *AddCasterComponentResponseBody
}

type AddCasterComponentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCasterComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCasterComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCasterComponentResponse) GoString() string {
	return s.String()
}

func (s *AddCasterComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCasterComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCasterComponentResponse) GetBody() *AddCasterComponentResponseBody {
	return s.Body
}

func (s *AddCasterComponentResponse) SetHeaders(v map[string]*string) *AddCasterComponentResponse {
	s.Headers = v
	return s
}

func (s *AddCasterComponentResponse) SetStatusCode(v int32) *AddCasterComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCasterComponentResponse) SetBody(v *AddCasterComponentResponseBody) *AddCasterComponentResponse {
	s.Body = v
	return s
}

func (s *AddCasterComponentResponse) Validate() error {
	return dara.Validate(s)
}
