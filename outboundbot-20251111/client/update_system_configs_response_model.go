// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSystemConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSystemConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSystemConfigsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSystemConfigsResponseBody) *UpdateSystemConfigsResponse
	GetBody() *UpdateSystemConfigsResponseBody
}

type UpdateSystemConfigsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSystemConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSystemConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSystemConfigsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSystemConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSystemConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSystemConfigsResponse) GetBody() *UpdateSystemConfigsResponseBody {
	return s.Body
}

func (s *UpdateSystemConfigsResponse) SetHeaders(v map[string]*string) *UpdateSystemConfigsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSystemConfigsResponse) SetStatusCode(v int32) *UpdateSystemConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSystemConfigsResponse) SetBody(v *UpdateSystemConfigsResponseBody) *UpdateSystemConfigsResponse {
	s.Body = v
	return s
}

func (s *UpdateSystemConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
