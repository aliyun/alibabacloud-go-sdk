// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourceStockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckResourceStockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckResourceStockResponse
	GetStatusCode() *int32
	SetBody(v *CheckResourceStockResponseBody) *CheckResourceStockResponse
	GetBody() *CheckResourceStockResponseBody
}

type CheckResourceStockResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResourceStockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResourceStockResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceStockResponse) GoString() string {
	return s.String()
}

func (s *CheckResourceStockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckResourceStockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckResourceStockResponse) GetBody() *CheckResourceStockResponseBody {
	return s.Body
}

func (s *CheckResourceStockResponse) SetHeaders(v map[string]*string) *CheckResourceStockResponse {
	s.Headers = v
	return s
}

func (s *CheckResourceStockResponse) SetStatusCode(v int32) *CheckResourceStockResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResourceStockResponse) SetBody(v *CheckResourceStockResponseBody) *CheckResourceStockResponse {
	s.Body = v
	return s
}

func (s *CheckResourceStockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
