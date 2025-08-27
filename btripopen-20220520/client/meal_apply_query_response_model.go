// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MealApplyQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MealApplyQueryResponse
	GetStatusCode() *int32
	SetBody(v *MealApplyQueryResponseBody) *MealApplyQueryResponse
	GetBody() *MealApplyQueryResponseBody
}

type MealApplyQueryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MealApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MealApplyQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s MealApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *MealApplyQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MealApplyQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MealApplyQueryResponse) GetBody() *MealApplyQueryResponseBody {
	return s.Body
}

func (s *MealApplyQueryResponse) SetHeaders(v map[string]*string) *MealApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *MealApplyQueryResponse) SetStatusCode(v int32) *MealApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MealApplyQueryResponse) SetBody(v *MealApplyQueryResponseBody) *MealApplyQueryResponse {
	s.Body = v
	return s
}

func (s *MealApplyQueryResponse) Validate() error {
	return dara.Validate(s)
}
