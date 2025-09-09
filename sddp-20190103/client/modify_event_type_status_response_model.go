// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventTypeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEventTypeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEventTypeStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEventTypeStatusResponseBody) *ModifyEventTypeStatusResponse
	GetBody() *ModifyEventTypeStatusResponseBody
}

type ModifyEventTypeStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEventTypeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEventTypeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventTypeStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyEventTypeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEventTypeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEventTypeStatusResponse) GetBody() *ModifyEventTypeStatusResponseBody {
	return s.Body
}

func (s *ModifyEventTypeStatusResponse) SetHeaders(v map[string]*string) *ModifyEventTypeStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyEventTypeStatusResponse) SetStatusCode(v int32) *ModifyEventTypeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEventTypeStatusResponse) SetBody(v *ModifyEventTypeStatusResponseBody) *ModifyEventTypeStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyEventTypeStatusResponse) Validate() error {
	return dara.Validate(s)
}
