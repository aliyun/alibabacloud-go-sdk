// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackTemplateByResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStackTemplateByResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStackTemplateByResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStackTemplateByResourcesResponseBody) *UpdateStackTemplateByResourcesResponse
	GetBody() *UpdateStackTemplateByResourcesResponseBody
}

type UpdateStackTemplateByResourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStackTemplateByResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStackTemplateByResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackTemplateByResourcesResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStackTemplateByResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStackTemplateByResourcesResponse) GetBody() *UpdateStackTemplateByResourcesResponseBody {
	return s.Body
}

func (s *UpdateStackTemplateByResourcesResponse) SetHeaders(v map[string]*string) *UpdateStackTemplateByResourcesResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackTemplateByResourcesResponse) SetStatusCode(v int32) *UpdateStackTemplateByResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponse) SetBody(v *UpdateStackTemplateByResourcesResponseBody) *UpdateStackTemplateByResourcesResponse {
	s.Body = v
	return s
}

func (s *UpdateStackTemplateByResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
