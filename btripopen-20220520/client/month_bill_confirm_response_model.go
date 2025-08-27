// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillConfirmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MonthBillConfirmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MonthBillConfirmResponse
	GetStatusCode() *int32
	SetBody(v *MonthBillConfirmResponseBody) *MonthBillConfirmResponse
	GetBody() *MonthBillConfirmResponseBody
}

type MonthBillConfirmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonthBillConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MonthBillConfirmResponse) String() string {
	return dara.Prettify(s)
}

func (s MonthBillConfirmResponse) GoString() string {
	return s.String()
}

func (s *MonthBillConfirmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MonthBillConfirmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MonthBillConfirmResponse) GetBody() *MonthBillConfirmResponseBody {
	return s.Body
}

func (s *MonthBillConfirmResponse) SetHeaders(v map[string]*string) *MonthBillConfirmResponse {
	s.Headers = v
	return s
}

func (s *MonthBillConfirmResponse) SetStatusCode(v int32) *MonthBillConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *MonthBillConfirmResponse) SetBody(v *MonthBillConfirmResponseBody) *MonthBillConfirmResponse {
	s.Body = v
	return s
}

func (s *MonthBillConfirmResponse) Validate() error {
	return dara.Validate(s)
}
