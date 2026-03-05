// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonthAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMonthAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMonthAmountResponse
	GetStatusCode() *int32
	SetBody(v *GetMonthAmountResponseBody) *GetMonthAmountResponse
	GetBody() *GetMonthAmountResponseBody
}

type GetMonthAmountResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMonthAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonthAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMonthAmountResponse) GoString() string {
	return s.String()
}

func (s *GetMonthAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMonthAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMonthAmountResponse) GetBody() *GetMonthAmountResponseBody {
	return s.Body
}

func (s *GetMonthAmountResponse) SetHeaders(v map[string]*string) *GetMonthAmountResponse {
	s.Headers = v
	return s
}

func (s *GetMonthAmountResponse) SetStatusCode(v int32) *GetMonthAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonthAmountResponse) SetBody(v *GetMonthAmountResponseBody) *GetMonthAmountResponse {
	s.Body = v
	return s
}

func (s *GetMonthAmountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
