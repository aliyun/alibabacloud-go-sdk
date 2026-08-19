// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDDoSSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDDoSSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDDoSSpecResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDDoSSpecResponseBody) *UpdateDDoSSpecResponse
	GetBody() *UpdateDDoSSpecResponseBody
}

type UpdateDDoSSpecResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDDoSSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDDoSSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDDoSSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateDDoSSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDDoSSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDDoSSpecResponse) GetBody() *UpdateDDoSSpecResponseBody {
	return s.Body
}

func (s *UpdateDDoSSpecResponse) SetHeaders(v map[string]*string) *UpdateDDoSSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateDDoSSpecResponse) SetStatusCode(v int32) *UpdateDDoSSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDDoSSpecResponse) SetBody(v *UpdateDDoSSpecResponseBody) *UpdateDDoSSpecResponse {
	s.Body = v
	return s
}

func (s *UpdateDDoSSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
