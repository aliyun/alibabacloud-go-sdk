// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeveloperResponse
	GetStatusCode() *int32
	SetBody(v *Developer) *CreateDeveloperResponse
	GetBody() *Developer
}

type CreateDeveloperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Developer         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeveloperResponse) GoString() string {
	return s.String()
}

func (s *CreateDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeveloperResponse) GetBody() *Developer {
	return s.Body
}

func (s *CreateDeveloperResponse) SetHeaders(v map[string]*string) *CreateDeveloperResponse {
	s.Headers = v
	return s
}

func (s *CreateDeveloperResponse) SetStatusCode(v int32) *CreateDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeveloperResponse) SetBody(v *Developer) *CreateDeveloperResponse {
	s.Body = v
	return s
}

func (s *CreateDeveloperResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
