// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectManagerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddProjectManagerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddProjectManagerResponse
	GetStatusCode() *int32
	SetBody(v *AddProjectManagerResponseBody) *AddProjectManagerResponse
	GetBody() *AddProjectManagerResponseBody
}

type AddProjectManagerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddProjectManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddProjectManagerResponse) String() string {
	return dara.Prettify(s)
}

func (s AddProjectManagerResponse) GoString() string {
	return s.String()
}

func (s *AddProjectManagerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddProjectManagerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddProjectManagerResponse) GetBody() *AddProjectManagerResponseBody {
	return s.Body
}

func (s *AddProjectManagerResponse) SetHeaders(v map[string]*string) *AddProjectManagerResponse {
	s.Headers = v
	return s
}

func (s *AddProjectManagerResponse) SetStatusCode(v int32) *AddProjectManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProjectManagerResponse) SetBody(v *AddProjectManagerResponseBody) *AddProjectManagerResponse {
	s.Body = v
	return s
}

func (s *AddProjectManagerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
