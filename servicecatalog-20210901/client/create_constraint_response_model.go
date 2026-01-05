// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConstraintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConstraintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConstraintResponse
	GetStatusCode() *int32
	SetBody(v *CreateConstraintResponseBody) *CreateConstraintResponse
	GetBody() *CreateConstraintResponseBody
}

type CreateConstraintResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConstraintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConstraintResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConstraintResponse) GoString() string {
	return s.String()
}

func (s *CreateConstraintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConstraintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConstraintResponse) GetBody() *CreateConstraintResponseBody {
	return s.Body
}

func (s *CreateConstraintResponse) SetHeaders(v map[string]*string) *CreateConstraintResponse {
	s.Headers = v
	return s
}

func (s *CreateConstraintResponse) SetStatusCode(v int32) *CreateConstraintResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConstraintResponse) SetBody(v *CreateConstraintResponseBody) *CreateConstraintResponse {
	s.Body = v
	return s
}

func (s *CreateConstraintResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
