// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyApproveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MealApplyApproveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MealApplyApproveResponse
	GetStatusCode() *int32
	SetBody(v *MealApplyApproveResponseBody) *MealApplyApproveResponse
	GetBody() *MealApplyApproveResponseBody
}

type MealApplyApproveResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MealApplyApproveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MealApplyApproveResponse) String() string {
	return dara.Prettify(s)
}

func (s MealApplyApproveResponse) GoString() string {
	return s.String()
}

func (s *MealApplyApproveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MealApplyApproveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MealApplyApproveResponse) GetBody() *MealApplyApproveResponseBody {
	return s.Body
}

func (s *MealApplyApproveResponse) SetHeaders(v map[string]*string) *MealApplyApproveResponse {
	s.Headers = v
	return s
}

func (s *MealApplyApproveResponse) SetStatusCode(v int32) *MealApplyApproveResponse {
	s.StatusCode = &v
	return s
}

func (s *MealApplyApproveResponse) SetBody(v *MealApplyApproveResponseBody) *MealApplyApproveResponse {
	s.Body = v
	return s
}

func (s *MealApplyApproveResponse) Validate() error {
	return dara.Validate(s)
}
