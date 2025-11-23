// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDifyInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDifyInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDifyInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDifyInstanceResponseBody) *CreateDifyInstanceResponse
	GetBody() *CreateDifyInstanceResponseBody
}

type CreateDifyInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDifyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDifyInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDifyInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDifyInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDifyInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDifyInstanceResponse) GetBody() *CreateDifyInstanceResponseBody {
	return s.Body
}

func (s *CreateDifyInstanceResponse) SetHeaders(v map[string]*string) *CreateDifyInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDifyInstanceResponse) SetStatusCode(v int32) *CreateDifyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDifyInstanceResponse) SetBody(v *CreateDifyInstanceResponseBody) *CreateDifyInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateDifyInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
