// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnsSaleControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnsSaleControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnsSaleControlResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnsSaleControlResponseBody) *UpdateEnsSaleControlResponse
	GetBody() *UpdateEnsSaleControlResponseBody
}

type UpdateEnsSaleControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnsSaleControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnsSaleControlResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnsSaleControlResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnsSaleControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnsSaleControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnsSaleControlResponse) GetBody() *UpdateEnsSaleControlResponseBody {
	return s.Body
}

func (s *UpdateEnsSaleControlResponse) SetHeaders(v map[string]*string) *UpdateEnsSaleControlResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnsSaleControlResponse) SetStatusCode(v int32) *UpdateEnsSaleControlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnsSaleControlResponse) SetBody(v *UpdateEnsSaleControlResponseBody) *UpdateEnsSaleControlResponse {
	s.Body = v
	return s
}

func (s *UpdateEnsSaleControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
