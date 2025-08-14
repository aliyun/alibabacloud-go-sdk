// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEventResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEventResponseBody) *ModifyEventResponse
	GetBody() *ModifyEventResponseBody
}

type ModifyEventResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventResponse) GoString() string {
	return s.String()
}

func (s *ModifyEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEventResponse) GetBody() *ModifyEventResponseBody {
	return s.Body
}

func (s *ModifyEventResponse) SetHeaders(v map[string]*string) *ModifyEventResponse {
	s.Headers = v
	return s
}

func (s *ModifyEventResponse) SetStatusCode(v int32) *ModifyEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEventResponse) SetBody(v *ModifyEventResponseBody) *ModifyEventResponse {
	s.Body = v
	return s
}

func (s *ModifyEventResponse) Validate() error {
	return dara.Validate(s)
}
