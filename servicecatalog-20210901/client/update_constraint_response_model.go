// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConstraintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConstraintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConstraintResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConstraintResponseBody) *UpdateConstraintResponse
	GetBody() *UpdateConstraintResponseBody
}

type UpdateConstraintResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConstraintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConstraintResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConstraintResponse) GoString() string {
	return s.String()
}

func (s *UpdateConstraintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConstraintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConstraintResponse) GetBody() *UpdateConstraintResponseBody {
	return s.Body
}

func (s *UpdateConstraintResponse) SetHeaders(v map[string]*string) *UpdateConstraintResponse {
	s.Headers = v
	return s
}

func (s *UpdateConstraintResponse) SetStatusCode(v int32) *UpdateConstraintResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConstraintResponse) SetBody(v *UpdateConstraintResponseBody) *UpdateConstraintResponse {
	s.Body = v
	return s
}

func (s *UpdateConstraintResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
