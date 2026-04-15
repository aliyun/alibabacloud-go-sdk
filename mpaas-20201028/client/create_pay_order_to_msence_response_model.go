// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePayOrderToMsenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePayOrderToMsenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePayOrderToMsenceResponse
	GetStatusCode() *int32
	SetBody(v *CreatePayOrderToMsenceResponseBody) *CreatePayOrderToMsenceResponse
	GetBody() *CreatePayOrderToMsenceResponseBody
}

type CreatePayOrderToMsenceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePayOrderToMsenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePayOrderToMsenceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePayOrderToMsenceResponse) GoString() string {
	return s.String()
}

func (s *CreatePayOrderToMsenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePayOrderToMsenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePayOrderToMsenceResponse) GetBody() *CreatePayOrderToMsenceResponseBody {
	return s.Body
}

func (s *CreatePayOrderToMsenceResponse) SetHeaders(v map[string]*string) *CreatePayOrderToMsenceResponse {
	s.Headers = v
	return s
}

func (s *CreatePayOrderToMsenceResponse) SetStatusCode(v int32) *CreatePayOrderToMsenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePayOrderToMsenceResponse) SetBody(v *CreatePayOrderToMsenceResponseBody) *CreatePayOrderToMsenceResponse {
	s.Body = v
	return s
}

func (s *CreatePayOrderToMsenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
