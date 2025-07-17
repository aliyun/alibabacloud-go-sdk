// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveOrderResponse
	GetStatusCode() *int32
	SetBody(v *ApproveOrderResponseBody) *ApproveOrderResponse
	GetBody() *ApproveOrderResponseBody
}

type ApproveOrderResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveOrderResponse) GoString() string {
	return s.String()
}

func (s *ApproveOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveOrderResponse) GetBody() *ApproveOrderResponseBody {
	return s.Body
}

func (s *ApproveOrderResponse) SetHeaders(v map[string]*string) *ApproveOrderResponse {
	s.Headers = v
	return s
}

func (s *ApproveOrderResponse) SetStatusCode(v int32) *ApproveOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveOrderResponse) SetBody(v *ApproveOrderResponseBody) *ApproveOrderResponse {
	s.Body = v
	return s
}

func (s *ApproveOrderResponse) Validate() error {
	return dara.Validate(s)
}
