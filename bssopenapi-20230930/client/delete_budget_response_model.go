// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBudgetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBudgetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBudgetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBudgetResponseBody) *DeleteBudgetResponse
	GetBody() *DeleteBudgetResponseBody
}

type DeleteBudgetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBudgetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBudgetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBudgetResponse) GoString() string {
	return s.String()
}

func (s *DeleteBudgetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBudgetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBudgetResponse) GetBody() *DeleteBudgetResponseBody {
	return s.Body
}

func (s *DeleteBudgetResponse) SetHeaders(v map[string]*string) *DeleteBudgetResponse {
	s.Headers = v
	return s
}

func (s *DeleteBudgetResponse) SetStatusCode(v int32) *DeleteBudgetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBudgetResponse) SetBody(v *DeleteBudgetResponseBody) *DeleteBudgetResponse {
	s.Body = v
	return s
}

func (s *DeleteBudgetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
