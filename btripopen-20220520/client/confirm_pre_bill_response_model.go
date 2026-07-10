// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPreBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmPreBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmPreBillResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmPreBillResponseBody) *ConfirmPreBillResponse
	GetBody() *ConfirmPreBillResponseBody
}

type ConfirmPreBillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmPreBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmPreBillResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPreBillResponse) GoString() string {
	return s.String()
}

func (s *ConfirmPreBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmPreBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmPreBillResponse) GetBody() *ConfirmPreBillResponseBody {
	return s.Body
}

func (s *ConfirmPreBillResponse) SetHeaders(v map[string]*string) *ConfirmPreBillResponse {
	s.Headers = v
	return s
}

func (s *ConfirmPreBillResponse) SetStatusCode(v int32) *ConfirmPreBillResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmPreBillResponse) SetBody(v *ConfirmPreBillResponseBody) *ConfirmPreBillResponse {
	s.Body = v
	return s
}

func (s *ConfirmPreBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
