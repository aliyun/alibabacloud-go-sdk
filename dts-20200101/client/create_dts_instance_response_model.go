// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDtsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDtsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDtsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDtsInstanceResponseBody) *CreateDtsInstanceResponse
	GetBody() *CreateDtsInstanceResponseBody
}

type CreateDtsInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDtsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDtsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDtsInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDtsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDtsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDtsInstanceResponse) GetBody() *CreateDtsInstanceResponseBody {
	return s.Body
}

func (s *CreateDtsInstanceResponse) SetHeaders(v map[string]*string) *CreateDtsInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDtsInstanceResponse) SetStatusCode(v int32) *CreateDtsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDtsInstanceResponse) SetBody(v *CreateDtsInstanceResponseBody) *CreateDtsInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateDtsInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
