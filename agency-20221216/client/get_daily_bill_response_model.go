// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDailyBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDailyBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDailyBillResponse
	GetStatusCode() *int32
	SetBody(v *GetDailyBillResponseBody) *GetDailyBillResponse
	GetBody() *GetDailyBillResponseBody
}

type GetDailyBillResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDailyBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDailyBillResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDailyBillResponse) GoString() string {
	return s.String()
}

func (s *GetDailyBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDailyBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDailyBillResponse) GetBody() *GetDailyBillResponseBody {
	return s.Body
}

func (s *GetDailyBillResponse) SetHeaders(v map[string]*string) *GetDailyBillResponse {
	s.Headers = v
	return s
}

func (s *GetDailyBillResponse) SetStatusCode(v int32) *GetDailyBillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDailyBillResponse) SetBody(v *GetDailyBillResponseBody) *GetDailyBillResponse {
	s.Body = v
	return s
}

func (s *GetDailyBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
