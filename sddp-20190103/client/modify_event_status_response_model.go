// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEventStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEventStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEventStatusResponseBody) *ModifyEventStatusResponse
	GetBody() *ModifyEventStatusResponseBody
}

type ModifyEventStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEventStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEventStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEventStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEventStatusResponse) GetBody() *ModifyEventStatusResponseBody {
	return s.Body
}

func (s *ModifyEventStatusResponse) SetHeaders(v map[string]*string) *ModifyEventStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyEventStatusResponse) SetStatusCode(v int32) *ModifyEventStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEventStatusResponse) SetBody(v *ModifyEventStatusResponseBody) *ModifyEventStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyEventStatusResponse) Validate() error {
	return dara.Validate(s)
}
