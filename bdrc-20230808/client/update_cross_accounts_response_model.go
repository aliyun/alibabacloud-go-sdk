// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrossAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCrossAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCrossAccountsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCrossAccountsResponseBody) *UpdateCrossAccountsResponse
	GetBody() *UpdateCrossAccountsResponseBody
}

type UpdateCrossAccountsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCrossAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCrossAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossAccountsResponse) GoString() string {
	return s.String()
}

func (s *UpdateCrossAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCrossAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCrossAccountsResponse) GetBody() *UpdateCrossAccountsResponseBody {
	return s.Body
}

func (s *UpdateCrossAccountsResponse) SetHeaders(v map[string]*string) *UpdateCrossAccountsResponse {
	s.Headers = v
	return s
}

func (s *UpdateCrossAccountsResponse) SetStatusCode(v int32) *UpdateCrossAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCrossAccountsResponse) SetBody(v *UpdateCrossAccountsResponseBody) *UpdateCrossAccountsResponse {
	s.Body = v
	return s
}

func (s *UpdateCrossAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
