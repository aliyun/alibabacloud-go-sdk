// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyColumnarClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyColumnarClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyColumnarClassResponse
	GetStatusCode() *int32
	SetBody(v *ModifyColumnarClassResponseBody) *ModifyColumnarClassResponse
	GetBody() *ModifyColumnarClassResponseBody
}

type ModifyColumnarClassResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyColumnarClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyColumnarClassResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyColumnarClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyColumnarClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyColumnarClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyColumnarClassResponse) GetBody() *ModifyColumnarClassResponseBody {
	return s.Body
}

func (s *ModifyColumnarClassResponse) SetHeaders(v map[string]*string) *ModifyColumnarClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyColumnarClassResponse) SetStatusCode(v int32) *ModifyColumnarClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyColumnarClassResponse) SetBody(v *ModifyColumnarClassResponseBody) *ModifyColumnarClassResponse {
	s.Body = v
	return s
}

func (s *ModifyColumnarClassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
