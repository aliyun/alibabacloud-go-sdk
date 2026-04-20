// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropUserResponse
	GetStatusCode() *int32
	SetBody(v *DropUserResponseBody) *DropUserResponse
	GetBody() *DropUserResponseBody
}

type DropUserResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DropUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DropUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DropUserResponse) GoString() string {
	return s.String()
}

func (s *DropUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropUserResponse) GetBody() *DropUserResponseBody {
	return s.Body
}

func (s *DropUserResponse) SetHeaders(v map[string]*string) *DropUserResponse {
	s.Headers = v
	return s
}

func (s *DropUserResponse) SetStatusCode(v int32) *DropUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DropUserResponse) SetBody(v *DropUserResponseBody) *DropUserResponse {
	s.Body = v
	return s
}

func (s *DropUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
