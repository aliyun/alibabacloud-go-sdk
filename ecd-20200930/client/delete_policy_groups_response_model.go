// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicyGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicyGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolicyGroupsResponseBody) *DeletePolicyGroupsResponse
	GetBody() *DeletePolicyGroupsResponseBody
}

type DeletePolicyGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicyGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicyGroupsResponse) GetBody() *DeletePolicyGroupsResponseBody {
	return s.Body
}

func (s *DeletePolicyGroupsResponse) SetHeaders(v map[string]*string) *DeletePolicyGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyGroupsResponse) SetStatusCode(v int32) *DeletePolicyGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyGroupsResponse) SetBody(v *DeletePolicyGroupsResponseBody) *DeletePolicyGroupsResponse {
	s.Body = v
	return s
}

func (s *DeletePolicyGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
