// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateByPassShuntEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateByPassShuntEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateByPassShuntEventResponse
	GetStatusCode() *int32
	SetBody(v *UpdateByPassShuntEventResponseBody) *UpdateByPassShuntEventResponse
	GetBody() *UpdateByPassShuntEventResponseBody
}

type UpdateByPassShuntEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateByPassShuntEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateByPassShuntEventResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateByPassShuntEventResponse) GoString() string {
	return s.String()
}

func (s *UpdateByPassShuntEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateByPassShuntEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateByPassShuntEventResponse) GetBody() *UpdateByPassShuntEventResponseBody {
	return s.Body
}

func (s *UpdateByPassShuntEventResponse) SetHeaders(v map[string]*string) *UpdateByPassShuntEventResponse {
	s.Headers = v
	return s
}

func (s *UpdateByPassShuntEventResponse) SetStatusCode(v int32) *UpdateByPassShuntEventResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateByPassShuntEventResponse) SetBody(v *UpdateByPassShuntEventResponseBody) *UpdateByPassShuntEventResponse {
	s.Body = v
	return s
}

func (s *UpdateByPassShuntEventResponse) Validate() error {
	return dara.Validate(s)
}
