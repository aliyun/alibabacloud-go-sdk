// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRumAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRumAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRumAppResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRumAppResponseBody) *UpdateRumAppResponse
	GetBody() *UpdateRumAppResponseBody
}

type UpdateRumAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRumAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRumAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRumAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateRumAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRumAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRumAppResponse) GetBody() *UpdateRumAppResponseBody {
	return s.Body
}

func (s *UpdateRumAppResponse) SetHeaders(v map[string]*string) *UpdateRumAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateRumAppResponse) SetStatusCode(v int32) *UpdateRumAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRumAppResponse) SetBody(v *UpdateRumAppResponseBody) *UpdateRumAppResponse {
	s.Body = v
	return s
}

func (s *UpdateRumAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
