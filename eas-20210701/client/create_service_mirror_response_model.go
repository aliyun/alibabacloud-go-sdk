// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceMirrorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceMirrorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceMirrorResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceMirrorResponseBody) *CreateServiceMirrorResponse
	GetBody() *CreateServiceMirrorResponseBody
}

type CreateServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceMirrorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceMirrorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceMirrorResponse) GetBody() *CreateServiceMirrorResponseBody {
	return s.Body
}

func (s *CreateServiceMirrorResponse) SetHeaders(v map[string]*string) *CreateServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceMirrorResponse) SetStatusCode(v int32) *CreateServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceMirrorResponse) SetBody(v *CreateServiceMirrorResponseBody) *CreateServiceMirrorResponse {
	s.Body = v
	return s
}

func (s *CreateServiceMirrorResponse) Validate() error {
	return dara.Validate(s)
}
