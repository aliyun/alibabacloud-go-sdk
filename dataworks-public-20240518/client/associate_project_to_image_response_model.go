// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProjectToImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateProjectToImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateProjectToImageResponse
	GetStatusCode() *int32
	SetBody(v *AssociateProjectToImageResponseBody) *AssociateProjectToImageResponse
	GetBody() *AssociateProjectToImageResponseBody
}

type AssociateProjectToImageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateProjectToImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateProjectToImageResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateProjectToImageResponse) GoString() string {
	return s.String()
}

func (s *AssociateProjectToImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateProjectToImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateProjectToImageResponse) GetBody() *AssociateProjectToImageResponseBody {
	return s.Body
}

func (s *AssociateProjectToImageResponse) SetHeaders(v map[string]*string) *AssociateProjectToImageResponse {
	s.Headers = v
	return s
}

func (s *AssociateProjectToImageResponse) SetStatusCode(v int32) *AssociateProjectToImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateProjectToImageResponse) SetBody(v *AssociateProjectToImageResponseBody) *AssociateProjectToImageResponse {
	s.Body = v
	return s
}

func (s *AssociateProjectToImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
