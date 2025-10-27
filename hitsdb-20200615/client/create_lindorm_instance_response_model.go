// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLindormInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLindormInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLindormInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateLindormInstanceResponseBody) *CreateLindormInstanceResponse
	GetBody() *CreateLindormInstanceResponseBody
}

type CreateLindormInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLindormInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLindormInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLindormInstanceResponse) GetBody() *CreateLindormInstanceResponseBody {
	return s.Body
}

func (s *CreateLindormInstanceResponse) SetHeaders(v map[string]*string) *CreateLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateLindormInstanceResponse) SetStatusCode(v int32) *CreateLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLindormInstanceResponse) SetBody(v *CreateLindormInstanceResponseBody) *CreateLindormInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateLindormInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
