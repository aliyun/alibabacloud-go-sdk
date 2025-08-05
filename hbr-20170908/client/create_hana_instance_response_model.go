// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHanaInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHanaInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateHanaInstanceResponseBody) *CreateHanaInstanceResponse
	GetBody() *CreateHanaInstanceResponseBody
}

type CreateHanaInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHanaInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHanaInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateHanaInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHanaInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHanaInstanceResponse) GetBody() *CreateHanaInstanceResponseBody {
	return s.Body
}

func (s *CreateHanaInstanceResponse) SetHeaders(v map[string]*string) *CreateHanaInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateHanaInstanceResponse) SetStatusCode(v int32) *CreateHanaInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHanaInstanceResponse) SetBody(v *CreateHanaInstanceResponseBody) *CreateHanaInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateHanaInstanceResponse) Validate() error {
	return dara.Validate(s)
}
