// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRevisionResponse
	GetStatusCode() *int32
	SetBody(v *Revision) *UpdateRevisionResponse
	GetBody() *Revision
}

type UpdateRevisionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRevisionResponse) GoString() string {
	return s.String()
}

func (s *UpdateRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRevisionResponse) GetBody() *Revision {
	return s.Body
}

func (s *UpdateRevisionResponse) SetHeaders(v map[string]*string) *UpdateRevisionResponse {
	s.Headers = v
	return s
}

func (s *UpdateRevisionResponse) SetStatusCode(v int32) *UpdateRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRevisionResponse) SetBody(v *Revision) *UpdateRevisionResponse {
	s.Body = v
	return s
}

func (s *UpdateRevisionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
