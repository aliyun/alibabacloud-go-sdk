// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessAssignmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessAssignmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessAssignmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessAssignmentResponseBody) *CreateAccessAssignmentResponse
	GetBody() *CreateAccessAssignmentResponseBody
}

type CreateAccessAssignmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessAssignmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessAssignmentResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessAssignmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessAssignmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessAssignmentResponse) GetBody() *CreateAccessAssignmentResponseBody {
	return s.Body
}

func (s *CreateAccessAssignmentResponse) SetHeaders(v map[string]*string) *CreateAccessAssignmentResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessAssignmentResponse) SetStatusCode(v int32) *CreateAccessAssignmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessAssignmentResponse) SetBody(v *CreateAccessAssignmentResponseBody) *CreateAccessAssignmentResponse {
	s.Body = v
	return s
}

func (s *CreateAccessAssignmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
