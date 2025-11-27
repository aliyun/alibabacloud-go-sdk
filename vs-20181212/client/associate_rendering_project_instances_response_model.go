// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRenderingProjectInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateRenderingProjectInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateRenderingProjectInstancesResponse
	GetStatusCode() *int32
	SetBody(v *AssociateRenderingProjectInstancesResponseBody) *AssociateRenderingProjectInstancesResponse
	GetBody() *AssociateRenderingProjectInstancesResponseBody
}

type AssociateRenderingProjectInstancesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateRenderingProjectInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateRenderingProjectInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateRenderingProjectInstancesResponse) GoString() string {
	return s.String()
}

func (s *AssociateRenderingProjectInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateRenderingProjectInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateRenderingProjectInstancesResponse) GetBody() *AssociateRenderingProjectInstancesResponseBody {
	return s.Body
}

func (s *AssociateRenderingProjectInstancesResponse) SetHeaders(v map[string]*string) *AssociateRenderingProjectInstancesResponse {
	s.Headers = v
	return s
}

func (s *AssociateRenderingProjectInstancesResponse) SetStatusCode(v int32) *AssociateRenderingProjectInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateRenderingProjectInstancesResponse) SetBody(v *AssociateRenderingProjectInstancesResponseBody) *AssociateRenderingProjectInstancesResponse {
	s.Body = v
	return s
}

func (s *AssociateRenderingProjectInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
