// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupAssociateProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceGroupAssociateProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceGroupAssociateProjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceGroupAssociateProjectsResponseBody) *ListResourceGroupAssociateProjectsResponse
	GetBody() *ListResourceGroupAssociateProjectsResponseBody
}

type ListResourceGroupAssociateProjectsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupAssociateProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupAssociateProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupAssociateProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupAssociateProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceGroupAssociateProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceGroupAssociateProjectsResponse) GetBody() *ListResourceGroupAssociateProjectsResponseBody {
	return s.Body
}

func (s *ListResourceGroupAssociateProjectsResponse) SetHeaders(v map[string]*string) *ListResourceGroupAssociateProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupAssociateProjectsResponse) SetStatusCode(v int32) *ListResourceGroupAssociateProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupAssociateProjectsResponse) SetBody(v *ListResourceGroupAssociateProjectsResponseBody) *ListResourceGroupAssociateProjectsResponse {
	s.Body = v
	return s
}

func (s *ListResourceGroupAssociateProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
