// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectionGroupRemoveProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectionGroupRemoveProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectionGroupRemoveProductResponse
	GetStatusCode() *int32
	SetBody(v *SelectionGroupRemoveProductResponseBody) *SelectionGroupRemoveProductResponse
	GetBody() *SelectionGroupRemoveProductResponseBody
}

type SelectionGroupRemoveProductResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectionGroupRemoveProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectionGroupRemoveProductResponse) String() string {
	return dara.Prettify(s)
}

func (s SelectionGroupRemoveProductResponse) GoString() string {
	return s.String()
}

func (s *SelectionGroupRemoveProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectionGroupRemoveProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectionGroupRemoveProductResponse) GetBody() *SelectionGroupRemoveProductResponseBody {
	return s.Body
}

func (s *SelectionGroupRemoveProductResponse) SetHeaders(v map[string]*string) *SelectionGroupRemoveProductResponse {
	s.Headers = v
	return s
}

func (s *SelectionGroupRemoveProductResponse) SetStatusCode(v int32) *SelectionGroupRemoveProductResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectionGroupRemoveProductResponse) SetBody(v *SelectionGroupRemoveProductResponseBody) *SelectionGroupRemoveProductResponse {
	s.Body = v
	return s
}

func (s *SelectionGroupRemoveProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
