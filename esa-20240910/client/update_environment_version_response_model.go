// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnvironmentVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnvironmentVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnvironmentVersionResponseBody) *UpdateEnvironmentVersionResponse
	GetBody() *UpdateEnvironmentVersionResponseBody
}

type UpdateEnvironmentVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnvironmentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvironmentVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnvironmentVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnvironmentVersionResponse) GetBody() *UpdateEnvironmentVersionResponseBody {
	return s.Body
}

func (s *UpdateEnvironmentVersionResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentVersionResponse) SetStatusCode(v int32) *UpdateEnvironmentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvironmentVersionResponse) SetBody(v *UpdateEnvironmentVersionResponseBody) *UpdateEnvironmentVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateEnvironmentVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
