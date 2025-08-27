// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealOrderDetailQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MealOrderDetailQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MealOrderDetailQueryResponse
	GetStatusCode() *int32
	SetBody(v *MealOrderDetailQueryResponseBody) *MealOrderDetailQueryResponse
	GetBody() *MealOrderDetailQueryResponseBody
}

type MealOrderDetailQueryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MealOrderDetailQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MealOrderDetailQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s MealOrderDetailQueryResponse) GoString() string {
	return s.String()
}

func (s *MealOrderDetailQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MealOrderDetailQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MealOrderDetailQueryResponse) GetBody() *MealOrderDetailQueryResponseBody {
	return s.Body
}

func (s *MealOrderDetailQueryResponse) SetHeaders(v map[string]*string) *MealOrderDetailQueryResponse {
	s.Headers = v
	return s
}

func (s *MealOrderDetailQueryResponse) SetStatusCode(v int32) *MealOrderDetailQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MealOrderDetailQueryResponse) SetBody(v *MealOrderDetailQueryResponseBody) *MealOrderDetailQueryResponse {
	s.Body = v
	return s
}

func (s *MealOrderDetailQueryResponse) Validate() error {
	return dara.Validate(s)
}
