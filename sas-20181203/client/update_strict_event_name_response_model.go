// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStrictEventNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStrictEventNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStrictEventNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStrictEventNameResponseBody) *UpdateStrictEventNameResponse
	GetBody() *UpdateStrictEventNameResponseBody
}

type UpdateStrictEventNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStrictEventNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStrictEventNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStrictEventNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateStrictEventNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStrictEventNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStrictEventNameResponse) GetBody() *UpdateStrictEventNameResponseBody {
	return s.Body
}

func (s *UpdateStrictEventNameResponse) SetHeaders(v map[string]*string) *UpdateStrictEventNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateStrictEventNameResponse) SetStatusCode(v int32) *UpdateStrictEventNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStrictEventNameResponse) SetBody(v *UpdateStrictEventNameResponseBody) *UpdateStrictEventNameResponse {
	s.Body = v
	return s
}

func (s *UpdateStrictEventNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
