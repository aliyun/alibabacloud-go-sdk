// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRcsSignMenuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRcsSignMenuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRcsSignMenuResponse
	GetStatusCode() *int32
	SetBody(v *AddRcsSignMenuResponseBody) *AddRcsSignMenuResponse
	GetBody() *AddRcsSignMenuResponseBody
}

type AddRcsSignMenuResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRcsSignMenuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRcsSignMenuResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRcsSignMenuResponse) GoString() string {
	return s.String()
}

func (s *AddRcsSignMenuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRcsSignMenuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRcsSignMenuResponse) GetBody() *AddRcsSignMenuResponseBody {
	return s.Body
}

func (s *AddRcsSignMenuResponse) SetHeaders(v map[string]*string) *AddRcsSignMenuResponse {
	s.Headers = v
	return s
}

func (s *AddRcsSignMenuResponse) SetStatusCode(v int32) *AddRcsSignMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRcsSignMenuResponse) SetBody(v *AddRcsSignMenuResponseBody) *AddRcsSignMenuResponse {
	s.Body = v
	return s
}

func (s *AddRcsSignMenuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
