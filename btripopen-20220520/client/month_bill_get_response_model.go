// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillGetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MonthBillGetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MonthBillGetResponse
	GetStatusCode() *int32
	SetBody(v *MonthBillGetResponseBody) *MonthBillGetResponse
	GetBody() *MonthBillGetResponseBody
}

type MonthBillGetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonthBillGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MonthBillGetResponse) String() string {
	return dara.Prettify(s)
}

func (s MonthBillGetResponse) GoString() string {
	return s.String()
}

func (s *MonthBillGetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MonthBillGetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MonthBillGetResponse) GetBody() *MonthBillGetResponseBody {
	return s.Body
}

func (s *MonthBillGetResponse) SetHeaders(v map[string]*string) *MonthBillGetResponse {
	s.Headers = v
	return s
}

func (s *MonthBillGetResponse) SetStatusCode(v int32) *MonthBillGetResponse {
	s.StatusCode = &v
	return s
}

func (s *MonthBillGetResponse) SetBody(v *MonthBillGetResponseBody) *MonthBillGetResponse {
	s.Body = v
	return s
}

func (s *MonthBillGetResponse) Validate() error {
	return dara.Validate(s)
}
