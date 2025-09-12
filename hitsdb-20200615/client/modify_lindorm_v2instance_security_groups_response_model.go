// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2InstanceSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLindormV2InstanceSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLindormV2InstanceSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLindormV2InstanceSecurityGroupsResponseBody) *ModifyLindormV2InstanceSecurityGroupsResponse
	GetBody() *ModifyLindormV2InstanceSecurityGroupsResponseBody
}

type ModifyLindormV2InstanceSecurityGroupsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLindormV2InstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLindormV2InstanceSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2InstanceSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponse) GetBody() *ModifyLindormV2InstanceSecurityGroupsResponseBody {
	return s.Body
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponse) SetHeaders(v map[string]*string) *ModifyLindormV2InstanceSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponse) SetStatusCode(v int32) *ModifyLindormV2InstanceSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponse) SetBody(v *ModifyLindormV2InstanceSecurityGroupsResponseBody) *ModifyLindormV2InstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *ModifyLindormV2InstanceSecurityGroupsResponse) Validate() error {
	return dara.Validate(s)
}
