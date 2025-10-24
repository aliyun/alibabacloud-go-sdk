// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceMirrorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceMirrorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceMirrorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceMirrorResponseBody) *UpdateServiceMirrorResponse
	GetBody() *UpdateServiceMirrorResponseBody
}

type UpdateServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceMirrorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceMirrorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceMirrorResponse) GetBody() *UpdateServiceMirrorResponseBody {
	return s.Body
}

func (s *UpdateServiceMirrorResponse) SetHeaders(v map[string]*string) *UpdateServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceMirrorResponse) SetStatusCode(v int32) *UpdateServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceMirrorResponse) SetBody(v *UpdateServiceMirrorResponseBody) *UpdateServiceMirrorResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceMirrorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
