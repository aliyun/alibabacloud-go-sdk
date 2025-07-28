// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomLineResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomLineResponseBody) *AddCustomLineResponse
	GetBody() *AddCustomLineResponseBody
}

type AddCustomLineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomLineResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLineResponse) GoString() string {
	return s.String()
}

func (s *AddCustomLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomLineResponse) GetBody() *AddCustomLineResponseBody {
	return s.Body
}

func (s *AddCustomLineResponse) SetHeaders(v map[string]*string) *AddCustomLineResponse {
	s.Headers = v
	return s
}

func (s *AddCustomLineResponse) SetStatusCode(v int32) *AddCustomLineResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomLineResponse) SetBody(v *AddCustomLineResponseBody) *AddCustomLineResponse {
	s.Body = v
	return s
}

func (s *AddCustomLineResponse) Validate() error {
	return dara.Validate(s)
}
