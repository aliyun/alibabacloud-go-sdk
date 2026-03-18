// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInventoryResponse
	GetStatusCode() *int32
	SetBody(v *CheckInventoryResponseBody) *CheckInventoryResponse
	GetBody() *CheckInventoryResponseBody
}

type CheckInventoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInventoryResponse) GoString() string {
	return s.String()
}

func (s *CheckInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInventoryResponse) GetBody() *CheckInventoryResponseBody {
	return s.Body
}

func (s *CheckInventoryResponse) SetHeaders(v map[string]*string) *CheckInventoryResponse {
	s.Headers = v
	return s
}

func (s *CheckInventoryResponse) SetStatusCode(v int32) *CheckInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInventoryResponse) SetBody(v *CheckInventoryResponseBody) *CheckInventoryResponse {
	s.Body = v
	return s
}

func (s *CheckInventoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
