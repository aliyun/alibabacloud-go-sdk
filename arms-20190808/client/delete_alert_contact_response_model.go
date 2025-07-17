// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertContactResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertContactResponseBody) *DeleteAlertContactResponse
	GetBody() *DeleteAlertContactResponseBody
}

type DeleteAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertContactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertContactResponse) GetBody() *DeleteAlertContactResponseBody {
	return s.Body
}

func (s *DeleteAlertContactResponse) SetHeaders(v map[string]*string) *DeleteAlertContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertContactResponse) SetStatusCode(v int32) *DeleteAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertContactResponse) SetBody(v *DeleteAlertContactResponseBody) *DeleteAlertContactResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertContactResponse) Validate() error {
	return dara.Validate(s)
}
