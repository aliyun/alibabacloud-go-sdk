// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBudgetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBudgetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBudgetResponse
	GetStatusCode() *int32
	SetBody(v *CreateBudgetResponseBody) *CreateBudgetResponse
	GetBody() *CreateBudgetResponseBody
}

type CreateBudgetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBudgetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBudgetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetResponse) GoString() string {
	return s.String()
}

func (s *CreateBudgetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBudgetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBudgetResponse) GetBody() *CreateBudgetResponseBody {
	return s.Body
}

func (s *CreateBudgetResponse) SetHeaders(v map[string]*string) *CreateBudgetResponse {
	s.Headers = v
	return s
}

func (s *CreateBudgetResponse) SetStatusCode(v int32) *CreateBudgetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBudgetResponse) SetBody(v *CreateBudgetResponseBody) *CreateBudgetResponse {
	s.Body = v
	return s
}

func (s *CreateBudgetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
