// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourceShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateResourceShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateResourceShareResponse
	GetStatusCode() *int32
	SetBody(v *AssociateResourceShareResponseBody) *AssociateResourceShareResponse
	GetBody() *AssociateResourceShareResponseBody
}

type AssociateResourceShareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateResourceShareResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateResourceShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateResourceShareResponse) GetBody() *AssociateResourceShareResponseBody {
	return s.Body
}

func (s *AssociateResourceShareResponse) SetHeaders(v map[string]*string) *AssociateResourceShareResponse {
	s.Headers = v
	return s
}

func (s *AssociateResourceShareResponse) SetStatusCode(v int32) *AssociateResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateResourceShareResponse) SetBody(v *AssociateResourceShareResponseBody) *AssociateResourceShareResponse {
	s.Body = v
	return s
}

func (s *AssociateResourceShareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
