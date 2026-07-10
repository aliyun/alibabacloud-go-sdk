// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MealApplyModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MealApplyModifyResponse
	GetStatusCode() *int32
	SetBody(v *MealApplyModifyResponseBody) *MealApplyModifyResponse
	GetBody() *MealApplyModifyResponseBody
}

type MealApplyModifyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MealApplyModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MealApplyModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyResponse) GoString() string {
	return s.String()
}

func (s *MealApplyModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MealApplyModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MealApplyModifyResponse) GetBody() *MealApplyModifyResponseBody {
	return s.Body
}

func (s *MealApplyModifyResponse) SetHeaders(v map[string]*string) *MealApplyModifyResponse {
	s.Headers = v
	return s
}

func (s *MealApplyModifyResponse) SetStatusCode(v int32) *MealApplyModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *MealApplyModifyResponse) SetBody(v *MealApplyModifyResponseBody) *MealApplyModifyResponse {
	s.Body = v
	return s
}

func (s *MealApplyModifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
