// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonthlyBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMonthlyBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMonthlyBillResponse
	GetStatusCode() *int32
	SetBody(v *GetMonthlyBillResponseBody) *GetMonthlyBillResponse
	GetBody() *GetMonthlyBillResponseBody
}

type GetMonthlyBillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMonthlyBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonthlyBillResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMonthlyBillResponse) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMonthlyBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMonthlyBillResponse) GetBody() *GetMonthlyBillResponseBody {
	return s.Body
}

func (s *GetMonthlyBillResponse) SetHeaders(v map[string]*string) *GetMonthlyBillResponse {
	s.Headers = v
	return s
}

func (s *GetMonthlyBillResponse) SetStatusCode(v int32) *GetMonthlyBillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonthlyBillResponse) SetBody(v *GetMonthlyBillResponseBody) *GetMonthlyBillResponse {
	s.Body = v
	return s
}

func (s *GetMonthlyBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
