// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProjectToResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateProjectToResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateProjectToResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *AssociateProjectToResourceGroupResponseBody) *AssociateProjectToResourceGroupResponse
	GetBody() *AssociateProjectToResourceGroupResponseBody
}

type AssociateProjectToResourceGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateProjectToResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateProjectToResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateProjectToResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *AssociateProjectToResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateProjectToResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateProjectToResourceGroupResponse) GetBody() *AssociateProjectToResourceGroupResponseBody {
	return s.Body
}

func (s *AssociateProjectToResourceGroupResponse) SetHeaders(v map[string]*string) *AssociateProjectToResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *AssociateProjectToResourceGroupResponse) SetStatusCode(v int32) *AssociateProjectToResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateProjectToResourceGroupResponse) SetBody(v *AssociateProjectToResourceGroupResponseBody) *AssociateProjectToResourceGroupResponse {
	s.Body = v
	return s
}

func (s *AssociateProjectToResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
