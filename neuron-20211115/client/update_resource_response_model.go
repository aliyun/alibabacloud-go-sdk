// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceResponse
	GetStatusCode() *int32
	SetBody(v *PdpResource) *UpdateResourceResponse
	GetBody() *PdpResource
}

type UpdateResourceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpResource       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceResponse) GetBody() *PdpResource {
	return s.Body
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetStatusCode(v int32) *UpdateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *PdpResource) *UpdateResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
