// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterResourceGroupResponseBody) *ModifyDBClusterResourceGroupResponse
	GetBody() *ModifyDBClusterResourceGroupResponseBody
}

type ModifyDBClusterResourceGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterResourceGroupResponse) GetBody() *ModifyDBClusterResourceGroupResponseBody {
	return s.Body
}

func (s *ModifyDBClusterResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBClusterResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterResourceGroupResponse) SetStatusCode(v int32) *ModifyDBClusterResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterResourceGroupResponse) SetBody(v *ModifyDBClusterResourceGroupResponseBody) *ModifyDBClusterResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
