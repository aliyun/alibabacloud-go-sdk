// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReimbursementOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReimbursementOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReimbursementOrderResponse
	GetStatusCode() *int32
	SetBody(v *QueryReimbursementOrderResponseBody) *QueryReimbursementOrderResponse
	GetBody() *QueryReimbursementOrderResponseBody
}

type QueryReimbursementOrderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReimbursementOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReimbursementOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReimbursementOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReimbursementOrderResponse) GetBody() *QueryReimbursementOrderResponseBody {
	return s.Body
}

func (s *QueryReimbursementOrderResponse) SetHeaders(v map[string]*string) *QueryReimbursementOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryReimbursementOrderResponse) SetStatusCode(v int32) *QueryReimbursementOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReimbursementOrderResponse) SetBody(v *QueryReimbursementOrderResponseBody) *QueryReimbursementOrderResponse {
	s.Body = v
	return s
}

func (s *QueryReimbursementOrderResponse) Validate() error {
	return dara.Validate(s)
}
