// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelTemplateResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelTemplateResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelTemplateResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListModelTemplateResourceGroupResponseBody) *ListModelTemplateResourceGroupResponse
	GetBody() *ListModelTemplateResourceGroupResponseBody
}

type ListModelTemplateResourceGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelTemplateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelTemplateResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListModelTemplateResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelTemplateResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelTemplateResourceGroupResponse) GetBody() *ListModelTemplateResourceGroupResponseBody {
	return s.Body
}

func (s *ListModelTemplateResourceGroupResponse) SetHeaders(v map[string]*string) *ListModelTemplateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListModelTemplateResourceGroupResponse) SetStatusCode(v int32) *ListModelTemplateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelTemplateResourceGroupResponse) SetBody(v *ListModelTemplateResourceGroupResponseBody) *ListModelTemplateResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ListModelTemplateResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
