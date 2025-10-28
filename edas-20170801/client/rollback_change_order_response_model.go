// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackChangeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackChangeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackChangeOrderResponse
	GetStatusCode() *int32
	SetBody(v *RollbackChangeOrderResponseBody) *RollbackChangeOrderResponse
	GetBody() *RollbackChangeOrderResponseBody
}

type RollbackChangeOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackChangeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackChangeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackChangeOrderResponse) GetBody() *RollbackChangeOrderResponseBody {
	return s.Body
}

func (s *RollbackChangeOrderResponse) SetHeaders(v map[string]*string) *RollbackChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *RollbackChangeOrderResponse) SetStatusCode(v int32) *RollbackChangeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackChangeOrderResponse) SetBody(v *RollbackChangeOrderResponseBody) *RollbackChangeOrderResponse {
	s.Body = v
	return s
}

func (s *RollbackChangeOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
