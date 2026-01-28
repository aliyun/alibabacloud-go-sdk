// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationVersionResponseBody) *UpdateApplicationVersionResponse
	GetBody() *UpdateApplicationVersionResponseBody
}

type UpdateApplicationVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationVersionResponse) GetBody() *UpdateApplicationVersionResponseBody {
	return s.Body
}

func (s *UpdateApplicationVersionResponse) SetHeaders(v map[string]*string) *UpdateApplicationVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationVersionResponse) SetStatusCode(v int32) *UpdateApplicationVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationVersionResponse) SetBody(v *UpdateApplicationVersionResponseBody) *UpdateApplicationVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
