// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateResourcesResponse
	GetStatusCode() *int32
	SetBody(v *AssociateResourcesResponseBody) *AssociateResourcesResponse
	GetBody() *AssociateResourcesResponseBody
}

type AssociateResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourcesResponse) GoString() string {
	return s.String()
}

func (s *AssociateResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateResourcesResponse) GetBody() *AssociateResourcesResponseBody {
	return s.Body
}

func (s *AssociateResourcesResponse) SetHeaders(v map[string]*string) *AssociateResourcesResponse {
	s.Headers = v
	return s
}

func (s *AssociateResourcesResponse) SetStatusCode(v int32) *AssociateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateResourcesResponse) SetBody(v *AssociateResourcesResponseBody) *AssociateResourcesResponse {
	s.Body = v
	return s
}

func (s *AssociateResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
