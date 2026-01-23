// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStandardSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStandardSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateStandardSetResponseBody) *CreateStandardSetResponse
	GetBody() *CreateStandardSetResponseBody
}

type CreateStandardSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStandardSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStandardSetResponse) GetBody() *CreateStandardSetResponseBody {
	return s.Body
}

func (s *CreateStandardSetResponse) SetHeaders(v map[string]*string) *CreateStandardSetResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardSetResponse) SetStatusCode(v int32) *CreateStandardSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardSetResponse) SetBody(v *CreateStandardSetResponseBody) *CreateStandardSetResponse {
	s.Body = v
	return s
}

func (s *CreateStandardSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
