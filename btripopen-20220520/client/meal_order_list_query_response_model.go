// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealOrderListQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MealOrderListQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MealOrderListQueryResponse
	GetStatusCode() *int32
	SetBody(v *MealOrderListQueryResponseBody) *MealOrderListQueryResponse
	GetBody() *MealOrderListQueryResponseBody
}

type MealOrderListQueryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MealOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MealOrderListQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s MealOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *MealOrderListQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MealOrderListQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MealOrderListQueryResponse) GetBody() *MealOrderListQueryResponseBody {
	return s.Body
}

func (s *MealOrderListQueryResponse) SetHeaders(v map[string]*string) *MealOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *MealOrderListQueryResponse) SetStatusCode(v int32) *MealOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MealOrderListQueryResponse) SetBody(v *MealOrderListQueryResponseBody) *MealOrderListQueryResponse {
	s.Body = v
	return s
}

func (s *MealOrderListQueryResponse) Validate() error {
	return dara.Validate(s)
}
