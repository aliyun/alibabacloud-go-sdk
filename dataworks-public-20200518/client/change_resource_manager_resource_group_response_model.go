// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceManagerResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeResourceManagerResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeResourceManagerResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeResourceManagerResourceGroupResponseBody) *ChangeResourceManagerResourceGroupResponse
	GetBody() *ChangeResourceManagerResourceGroupResponseBody
}

type ChangeResourceManagerResourceGroupResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceManagerResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceManagerResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceManagerResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceManagerResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeResourceManagerResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeResourceManagerResourceGroupResponse) GetBody() *ChangeResourceManagerResourceGroupResponseBody {
	return s.Body
}

func (s *ChangeResourceManagerResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceManagerResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceManagerResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceManagerResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupResponse) SetBody(v *ChangeResourceManagerResourceGroupResponseBody) *ChangeResourceManagerResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ChangeResourceManagerResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
