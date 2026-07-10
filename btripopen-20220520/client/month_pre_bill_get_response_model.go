// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthPreBillGetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MonthPreBillGetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MonthPreBillGetResponse
	GetStatusCode() *int32
	SetBody(v *MonthPreBillGetResponseBody) *MonthPreBillGetResponse
	GetBody() *MonthPreBillGetResponseBody
}

type MonthPreBillGetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonthPreBillGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MonthPreBillGetResponse) String() string {
	return dara.Prettify(s)
}

func (s MonthPreBillGetResponse) GoString() string {
	return s.String()
}

func (s *MonthPreBillGetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MonthPreBillGetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MonthPreBillGetResponse) GetBody() *MonthPreBillGetResponseBody {
	return s.Body
}

func (s *MonthPreBillGetResponse) SetHeaders(v map[string]*string) *MonthPreBillGetResponse {
	s.Headers = v
	return s
}

func (s *MonthPreBillGetResponse) SetStatusCode(v int32) *MonthPreBillGetResponse {
	s.StatusCode = &v
	return s
}

func (s *MonthPreBillGetResponse) SetBody(v *MonthPreBillGetResponseBody) *MonthPreBillGetResponse {
	s.Body = v
	return s
}

func (s *MonthPreBillGetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
