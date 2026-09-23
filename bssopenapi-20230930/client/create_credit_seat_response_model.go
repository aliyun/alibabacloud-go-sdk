// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCreditSeatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCreditSeatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCreditSeatResponse
	GetStatusCode() *int32
	SetBody(v *CreateCreditSeatResponseBody) *CreateCreditSeatResponse
	GetBody() *CreateCreditSeatResponseBody
}

type CreateCreditSeatResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCreditSeatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCreditSeatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditSeatResponse) GoString() string {
	return s.String()
}

func (s *CreateCreditSeatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCreditSeatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCreditSeatResponse) GetBody() *CreateCreditSeatResponseBody {
	return s.Body
}

func (s *CreateCreditSeatResponse) SetHeaders(v map[string]*string) *CreateCreditSeatResponse {
	s.Headers = v
	return s
}

func (s *CreateCreditSeatResponse) SetStatusCode(v int32) *CreateCreditSeatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCreditSeatResponse) SetBody(v *CreateCreditSeatResponseBody) *CreateCreditSeatResponse {
	s.Body = v
	return s
}

func (s *CreateCreditSeatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
