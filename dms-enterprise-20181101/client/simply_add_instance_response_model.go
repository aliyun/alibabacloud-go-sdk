// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimplyAddInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SimplyAddInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SimplyAddInstanceResponse
	GetStatusCode() *int32
	SetBody(v *SimplyAddInstanceResponseBody) *SimplyAddInstanceResponse
	GetBody() *SimplyAddInstanceResponseBody
}

type SimplyAddInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SimplyAddInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SimplyAddInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s SimplyAddInstanceResponse) GoString() string {
	return s.String()
}

func (s *SimplyAddInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SimplyAddInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SimplyAddInstanceResponse) GetBody() *SimplyAddInstanceResponseBody {
	return s.Body
}

func (s *SimplyAddInstanceResponse) SetHeaders(v map[string]*string) *SimplyAddInstanceResponse {
	s.Headers = v
	return s
}

func (s *SimplyAddInstanceResponse) SetStatusCode(v int32) *SimplyAddInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SimplyAddInstanceResponse) SetBody(v *SimplyAddInstanceResponseBody) *SimplyAddInstanceResponse {
	s.Body = v
	return s
}

func (s *SimplyAddInstanceResponse) Validate() error {
	return dara.Validate(s)
}
