// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBudgetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBudgetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBudgetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBudgetResponseBody) *UpdateBudgetResponse
	GetBody() *UpdateBudgetResponseBody
}

type UpdateBudgetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBudgetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBudgetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBudgetResponse) GoString() string {
	return s.String()
}

func (s *UpdateBudgetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBudgetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBudgetResponse) GetBody() *UpdateBudgetResponseBody {
	return s.Body
}

func (s *UpdateBudgetResponse) SetHeaders(v map[string]*string) *UpdateBudgetResponse {
	s.Headers = v
	return s
}

func (s *UpdateBudgetResponse) SetStatusCode(v int32) *UpdateBudgetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBudgetResponse) SetBody(v *UpdateBudgetResponseBody) *UpdateBudgetResponse {
	s.Body = v
	return s
}

func (s *UpdateBudgetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
