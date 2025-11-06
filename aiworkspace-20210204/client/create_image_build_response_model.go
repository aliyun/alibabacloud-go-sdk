// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageBuildResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageBuildResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageBuildResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageBuildResponseBody) *CreateImageBuildResponse
	GetBody() *CreateImageBuildResponseBody
}

type CreateImageBuildResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageBuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageBuildResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildResponse) GoString() string {
	return s.String()
}

func (s *CreateImageBuildResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageBuildResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageBuildResponse) GetBody() *CreateImageBuildResponseBody {
	return s.Body
}

func (s *CreateImageBuildResponse) SetHeaders(v map[string]*string) *CreateImageBuildResponse {
	s.Headers = v
	return s
}

func (s *CreateImageBuildResponse) SetStatusCode(v int32) *CreateImageBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageBuildResponse) SetBody(v *CreateImageBuildResponseBody) *CreateImageBuildResponse {
	s.Body = v
	return s
}

func (s *CreateImageBuildResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
