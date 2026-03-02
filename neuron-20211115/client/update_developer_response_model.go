// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeveloperResponse
	GetStatusCode() *int32
	SetBody(v *Developer) *UpdateDeveloperResponse
	GetBody() *Developer
}

type UpdateDeveloperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Developer         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeveloperResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeveloperResponse) GetBody() *Developer {
	return s.Body
}

func (s *UpdateDeveloperResponse) SetHeaders(v map[string]*string) *UpdateDeveloperResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeveloperResponse) SetStatusCode(v int32) *UpdateDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeveloperResponse) SetBody(v *Developer) *UpdateDeveloperResponse {
	s.Body = v
	return s
}

func (s *UpdateDeveloperResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
