// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourcesResponseBody) *UpdateResourcesResponse
	GetBody() *UpdateResourcesResponseBody
}

type UpdateResourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourcesResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourcesResponse) GetBody() *UpdateResourcesResponseBody {
	return s.Body
}

func (s *UpdateResourcesResponse) SetHeaders(v map[string]*string) *UpdateResourcesResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourcesResponse) SetStatusCode(v int32) *UpdateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourcesResponse) SetBody(v *UpdateResourcesResponseBody) *UpdateResourcesResponse {
	s.Body = v
	return s
}

func (s *UpdateResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
