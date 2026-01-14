// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateListenerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateListenerResponseBody) *UpdateListenerResponse
	GetBody() *UpdateListenerResponseBody
}

type UpdateListenerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateListenerResponse) GetBody() *UpdateListenerResponseBody {
	return s.Body
}

func (s *UpdateListenerResponse) SetHeaders(v map[string]*string) *UpdateListenerResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerResponse) SetStatusCode(v int32) *UpdateListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateListenerResponse) SetBody(v *UpdateListenerResponseBody) *UpdateListenerResponse {
	s.Body = v
	return s
}

func (s *UpdateListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
