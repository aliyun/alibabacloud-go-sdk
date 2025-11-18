// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppInstanceResponseBody) *CreateAppInstanceResponse
	GetBody() *CreateAppInstanceResponseBody
}

type CreateAppInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppInstanceResponse) GetBody() *CreateAppInstanceResponseBody {
	return s.Body
}

func (s *CreateAppInstanceResponse) SetHeaders(v map[string]*string) *CreateAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateAppInstanceResponse) SetStatusCode(v int32) *CreateAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppInstanceResponse) SetBody(v *CreateAppInstanceResponseBody) *CreateAppInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
