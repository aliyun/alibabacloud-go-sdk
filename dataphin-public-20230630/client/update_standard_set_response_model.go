// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStandardSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStandardSetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStandardSetResponseBody) *UpdateStandardSetResponse
	GetBody() *UpdateStandardSetResponseBody
}

type UpdateStandardSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStandardSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStandardSetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateStandardSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStandardSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStandardSetResponse) GetBody() *UpdateStandardSetResponseBody {
	return s.Body
}

func (s *UpdateStandardSetResponse) SetHeaders(v map[string]*string) *UpdateStandardSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateStandardSetResponse) SetStatusCode(v int32) *UpdateStandardSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStandardSetResponse) SetBody(v *UpdateStandardSetResponseBody) *UpdateStandardSetResponse {
	s.Body = v
	return s
}

func (s *UpdateStandardSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
