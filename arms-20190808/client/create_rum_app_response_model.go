// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRumAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRumAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRumAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateRumAppResponseBody) *CreateRumAppResponse
	GetBody() *CreateRumAppResponseBody
}

type CreateRumAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRumAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRumAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRumAppResponse) GoString() string {
	return s.String()
}

func (s *CreateRumAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRumAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRumAppResponse) GetBody() *CreateRumAppResponseBody {
	return s.Body
}

func (s *CreateRumAppResponse) SetHeaders(v map[string]*string) *CreateRumAppResponse {
	s.Headers = v
	return s
}

func (s *CreateRumAppResponse) SetStatusCode(v int32) *CreateRumAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRumAppResponse) SetBody(v *CreateRumAppResponseBody) *CreateRumAppResponse {
	s.Body = v
	return s
}

func (s *CreateRumAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
