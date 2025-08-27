// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MealApplyAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MealApplyAddResponse
	GetStatusCode() *int32
	SetBody(v *MealApplyAddResponseBody) *MealApplyAddResponse
	GetBody() *MealApplyAddResponseBody
}

type MealApplyAddResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MealApplyAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MealApplyAddResponse) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddResponse) GoString() string {
	return s.String()
}

func (s *MealApplyAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MealApplyAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MealApplyAddResponse) GetBody() *MealApplyAddResponseBody {
	return s.Body
}

func (s *MealApplyAddResponse) SetHeaders(v map[string]*string) *MealApplyAddResponse {
	s.Headers = v
	return s
}

func (s *MealApplyAddResponse) SetStatusCode(v int32) *MealApplyAddResponse {
	s.StatusCode = &v
	return s
}

func (s *MealApplyAddResponse) SetBody(v *MealApplyAddResponseBody) *MealApplyAddResponse {
	s.Body = v
	return s
}

func (s *MealApplyAddResponse) Validate() error {
	return dara.Validate(s)
}
