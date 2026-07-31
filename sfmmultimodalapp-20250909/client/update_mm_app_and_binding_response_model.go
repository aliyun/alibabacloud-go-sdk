// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppAndBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmAppAndBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmAppAndBindingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmAppAndBindingResponseBody) *UpdateMmAppAndBindingResponse
	GetBody() *UpdateMmAppAndBindingResponseBody
}

type UpdateMmAppAndBindingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmAppAndBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmAppAndBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmAppAndBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmAppAndBindingResponse) GetBody() *UpdateMmAppAndBindingResponseBody {
	return s.Body
}

func (s *UpdateMmAppAndBindingResponse) SetHeaders(v map[string]*string) *UpdateMmAppAndBindingResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmAppAndBindingResponse) SetStatusCode(v int32) *UpdateMmAppAndBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmAppAndBindingResponse) SetBody(v *UpdateMmAppAndBindingResponseBody) *UpdateMmAppAndBindingResponse {
	s.Body = v
	return s
}

func (s *UpdateMmAppAndBindingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
