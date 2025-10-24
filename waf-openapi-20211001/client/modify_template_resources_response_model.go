// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTemplateResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTemplateResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTemplateResourcesResponseBody) *ModifyTemplateResourcesResponse
	GetBody() *ModifyTemplateResourcesResponseBody
}

type ModifyTemplateResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTemplateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTemplateResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateResourcesResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTemplateResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTemplateResourcesResponse) GetBody() *ModifyTemplateResourcesResponseBody {
	return s.Body
}

func (s *ModifyTemplateResourcesResponse) SetHeaders(v map[string]*string) *ModifyTemplateResourcesResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateResourcesResponse) SetStatusCode(v int32) *ModifyTemplateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTemplateResourcesResponse) SetBody(v *ModifyTemplateResourcesResponseBody) *ModifyTemplateResourcesResponse {
	s.Body = v
	return s
}

func (s *ModifyTemplateResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
