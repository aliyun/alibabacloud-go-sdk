// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOuterAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOuterAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOuterAccountResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOuterAccountResponseBody) *UpdateOuterAccountResponse
	GetBody() *UpdateOuterAccountResponseBody
}

type UpdateOuterAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOuterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOuterAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOuterAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateOuterAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOuterAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOuterAccountResponse) GetBody() *UpdateOuterAccountResponseBody {
	return s.Body
}

func (s *UpdateOuterAccountResponse) SetHeaders(v map[string]*string) *UpdateOuterAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateOuterAccountResponse) SetStatusCode(v int32) *UpdateOuterAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOuterAccountResponse) SetBody(v *UpdateOuterAccountResponseBody) *UpdateOuterAccountResponse {
	s.Body = v
	return s
}

func (s *UpdateOuterAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
