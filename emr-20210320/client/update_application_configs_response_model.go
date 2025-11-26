// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationConfigsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationConfigsResponseBody) *UpdateApplicationConfigsResponse
	GetBody() *UpdateApplicationConfigsResponseBody
}

type UpdateApplicationConfigsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationConfigsResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationConfigsResponse) GetBody() *UpdateApplicationConfigsResponseBody {
	return s.Body
}

func (s *UpdateApplicationConfigsResponse) SetHeaders(v map[string]*string) *UpdateApplicationConfigsResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationConfigsResponse) SetStatusCode(v int32) *UpdateApplicationConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationConfigsResponse) SetBody(v *UpdateApplicationConfigsResponseBody) *UpdateApplicationConfigsResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
