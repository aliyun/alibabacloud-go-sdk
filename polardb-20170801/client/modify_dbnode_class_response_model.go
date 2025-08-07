// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodeClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodeClassResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodeClassResponseBody) *ModifyDBNodeClassResponse
	GetBody() *ModifyDBNodeClassResponseBody
}

type ModifyDBNodeClassResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodeClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodeClassResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodeClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodeClassResponse) GetBody() *ModifyDBNodeClassResponseBody {
	return s.Body
}

func (s *ModifyDBNodeClassResponse) SetHeaders(v map[string]*string) *ModifyDBNodeClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodeClassResponse) SetStatusCode(v int32) *ModifyDBNodeClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodeClassResponse) SetBody(v *ModifyDBNodeClassResponseBody) *ModifyDBNodeClassResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodeClassResponse) Validate() error {
	return dara.Validate(s)
}
