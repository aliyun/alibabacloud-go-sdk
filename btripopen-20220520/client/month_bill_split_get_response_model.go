// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillSplitGetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MonthBillSplitGetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MonthBillSplitGetResponse
	GetStatusCode() *int32
	SetBody(v *MonthBillSplitGetResponseBody) *MonthBillSplitGetResponse
	GetBody() *MonthBillSplitGetResponseBody
}

type MonthBillSplitGetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonthBillSplitGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MonthBillSplitGetResponse) String() string {
	return dara.Prettify(s)
}

func (s MonthBillSplitGetResponse) GoString() string {
	return s.String()
}

func (s *MonthBillSplitGetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MonthBillSplitGetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MonthBillSplitGetResponse) GetBody() *MonthBillSplitGetResponseBody {
	return s.Body
}

func (s *MonthBillSplitGetResponse) SetHeaders(v map[string]*string) *MonthBillSplitGetResponse {
	s.Headers = v
	return s
}

func (s *MonthBillSplitGetResponse) SetStatusCode(v int32) *MonthBillSplitGetResponse {
	s.StatusCode = &v
	return s
}

func (s *MonthBillSplitGetResponse) SetBody(v *MonthBillSplitGetResponseBody) *MonthBillSplitGetResponse {
	s.Body = v
	return s
}

func (s *MonthBillSplitGetResponse) Validate() error {
	return dara.Validate(s)
}
