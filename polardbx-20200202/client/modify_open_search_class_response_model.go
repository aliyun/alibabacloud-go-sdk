// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOpenSearchClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOpenSearchClassResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOpenSearchClassResponseBody) *ModifyOpenSearchClassResponse
	GetBody() *ModifyOpenSearchClassResponseBody
}

type ModifyOpenSearchClassResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOpenSearchClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOpenSearchClassResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOpenSearchClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOpenSearchClassResponse) GetBody() *ModifyOpenSearchClassResponseBody {
	return s.Body
}

func (s *ModifyOpenSearchClassResponse) SetHeaders(v map[string]*string) *ModifyOpenSearchClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyOpenSearchClassResponse) SetStatusCode(v int32) *ModifyOpenSearchClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOpenSearchClassResponse) SetBody(v *ModifyOpenSearchClassResponseBody) *ModifyOpenSearchClassResponse {
	s.Body = v
	return s
}

func (s *ModifyOpenSearchClassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
