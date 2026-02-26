// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectionGroupAddProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectionGroupAddProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectionGroupAddProductResponse
	GetStatusCode() *int32
	SetBody(v *SelectionGroupAddProductResponseBody) *SelectionGroupAddProductResponse
	GetBody() *SelectionGroupAddProductResponseBody
}

type SelectionGroupAddProductResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectionGroupAddProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectionGroupAddProductResponse) String() string {
	return dara.Prettify(s)
}

func (s SelectionGroupAddProductResponse) GoString() string {
	return s.String()
}

func (s *SelectionGroupAddProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectionGroupAddProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectionGroupAddProductResponse) GetBody() *SelectionGroupAddProductResponseBody {
	return s.Body
}

func (s *SelectionGroupAddProductResponse) SetHeaders(v map[string]*string) *SelectionGroupAddProductResponse {
	s.Headers = v
	return s
}

func (s *SelectionGroupAddProductResponse) SetStatusCode(v int32) *SelectionGroupAddProductResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectionGroupAddProductResponse) SetBody(v *SelectionGroupAddProductResponseBody) *SelectionGroupAddProductResponse {
	s.Body = v
	return s
}

func (s *SelectionGroupAddProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
