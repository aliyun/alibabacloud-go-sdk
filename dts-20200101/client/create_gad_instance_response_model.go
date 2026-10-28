// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGadInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGadInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateGadInstanceResponseBody) *CreateGadInstanceResponse
	GetBody() *CreateGadInstanceResponseBody
}

type CreateGadInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGadInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGadInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGadInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateGadInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGadInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGadInstanceResponse) GetBody() *CreateGadInstanceResponseBody {
	return s.Body
}

func (s *CreateGadInstanceResponse) SetHeaders(v map[string]*string) *CreateGadInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateGadInstanceResponse) SetStatusCode(v int32) *CreateGadInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGadInstanceResponse) SetBody(v *CreateGadInstanceResponseBody) *CreateGadInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateGadInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
