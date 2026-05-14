// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBudgetNameExistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckBudgetNameExistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckBudgetNameExistsResponse
	GetStatusCode() *int32
	SetBody(v *CheckBudgetNameExistsResponseBody) *CheckBudgetNameExistsResponse
	GetBody() *CheckBudgetNameExistsResponseBody
}

type CheckBudgetNameExistsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckBudgetNameExistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckBudgetNameExistsResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckBudgetNameExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckBudgetNameExistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckBudgetNameExistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckBudgetNameExistsResponse) GetBody() *CheckBudgetNameExistsResponseBody {
	return s.Body
}

func (s *CheckBudgetNameExistsResponse) SetHeaders(v map[string]*string) *CheckBudgetNameExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckBudgetNameExistsResponse) SetStatusCode(v int32) *CheckBudgetNameExistsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckBudgetNameExistsResponse) SetBody(v *CheckBudgetNameExistsResponseBody) *CheckBudgetNameExistsResponse {
	s.Body = v
	return s
}

func (s *CheckBudgetNameExistsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
