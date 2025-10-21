// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelInstanceResponseBody) *CreateModelInstanceResponse
	GetBody() *CreateModelInstanceResponseBody
}

type CreateModelInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateModelInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelInstanceResponse) GetBody() *CreateModelInstanceResponseBody {
	return s.Body
}

func (s *CreateModelInstanceResponse) SetHeaders(v map[string]*string) *CreateModelInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateModelInstanceResponse) SetStatusCode(v int32) *CreateModelInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelInstanceResponse) SetBody(v *CreateModelInstanceResponseBody) *CreateModelInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateModelInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
