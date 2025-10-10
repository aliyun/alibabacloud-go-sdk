// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyProductResponse
	GetStatusCode() *int32
	SetBody(v *ModifyProductResponseBody) *ModifyProductResponse
	GetBody() *ModifyProductResponseBody
}

type ModifyProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyProductResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyProductResponse) GoString() string {
	return s.String()
}

func (s *ModifyProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyProductResponse) GetBody() *ModifyProductResponseBody {
	return s.Body
}

func (s *ModifyProductResponse) SetHeaders(v map[string]*string) *ModifyProductResponse {
	s.Headers = v
	return s
}

func (s *ModifyProductResponse) SetStatusCode(v int32) *ModifyProductResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProductResponse) SetBody(v *ModifyProductResponseBody) *ModifyProductResponse {
	s.Body = v
	return s
}

func (s *ModifyProductResponse) Validate() error {
	return dara.Validate(s)
}
