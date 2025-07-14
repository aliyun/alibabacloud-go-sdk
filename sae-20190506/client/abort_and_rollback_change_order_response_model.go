// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortAndRollbackChangeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbortAndRollbackChangeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbortAndRollbackChangeOrderResponse
	GetStatusCode() *int32
	SetBody(v *AbortAndRollbackChangeOrderResponseBody) *AbortAndRollbackChangeOrderResponse
	GetBody() *AbortAndRollbackChangeOrderResponseBody
}

type AbortAndRollbackChangeOrderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbortAndRollbackChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbortAndRollbackChangeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbortAndRollbackChangeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbortAndRollbackChangeOrderResponse) GetBody() *AbortAndRollbackChangeOrderResponseBody {
	return s.Body
}

func (s *AbortAndRollbackChangeOrderResponse) SetHeaders(v map[string]*string) *AbortAndRollbackChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *AbortAndRollbackChangeOrderResponse) SetStatusCode(v int32) *AbortAndRollbackChangeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponse) SetBody(v *AbortAndRollbackChangeOrderResponseBody) *AbortAndRollbackChangeOrderResponse {
	s.Body = v
	return s
}

func (s *AbortAndRollbackChangeOrderResponse) Validate() error {
	return dara.Validate(s)
}
