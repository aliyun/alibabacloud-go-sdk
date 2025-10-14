// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewARMServerInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewARMServerInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewARMServerInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewARMServerInstanceResponseBody) *RenewARMServerInstanceResponse
	GetBody() *RenewARMServerInstanceResponseBody
}

type RenewARMServerInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewARMServerInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewARMServerInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewARMServerInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewARMServerInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewARMServerInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewARMServerInstanceResponse) GetBody() *RenewARMServerInstanceResponseBody {
	return s.Body
}

func (s *RenewARMServerInstanceResponse) SetHeaders(v map[string]*string) *RenewARMServerInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewARMServerInstanceResponse) SetStatusCode(v int32) *RenewARMServerInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewARMServerInstanceResponse) SetBody(v *RenewARMServerInstanceResponseBody) *RenewARMServerInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewARMServerInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
