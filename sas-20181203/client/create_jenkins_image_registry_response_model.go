// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJenkinsImageRegistryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJenkinsImageRegistryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJenkinsImageRegistryResponse
	GetStatusCode() *int32
	SetBody(v *CreateJenkinsImageRegistryResponseBody) *CreateJenkinsImageRegistryResponse
	GetBody() *CreateJenkinsImageRegistryResponseBody
}

type CreateJenkinsImageRegistryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJenkinsImageRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJenkinsImageRegistryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJenkinsImageRegistryResponse) GoString() string {
	return s.String()
}

func (s *CreateJenkinsImageRegistryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJenkinsImageRegistryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJenkinsImageRegistryResponse) GetBody() *CreateJenkinsImageRegistryResponseBody {
	return s.Body
}

func (s *CreateJenkinsImageRegistryResponse) SetHeaders(v map[string]*string) *CreateJenkinsImageRegistryResponse {
	s.Headers = v
	return s
}

func (s *CreateJenkinsImageRegistryResponse) SetStatusCode(v int32) *CreateJenkinsImageRegistryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJenkinsImageRegistryResponse) SetBody(v *CreateJenkinsImageRegistryResponseBody) *CreateJenkinsImageRegistryResponse {
	s.Body = v
	return s
}

func (s *CreateJenkinsImageRegistryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
