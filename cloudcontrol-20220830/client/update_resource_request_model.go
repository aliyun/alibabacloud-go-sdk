// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *UpdateResourceRequest
	GetBody() map[string]interface{}
	SetClientToken(v string) *UpdateResourceRequest
	GetClientToken() *string
	SetRegionId(v string) *UpdateResourceRequest
	GetRegionId() *string
}

type UpdateResourceRequest struct {
	// The request body. The property of the resource to be updated is specified in JSON format.
	//
	// example:
	//
	// {
	//
	//      "AccountPassword": "4321****",
	//
	//      "Description": "cctest"
	//
	// }
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. If a cloud service supports idempotence, the parameter takes effect.
	//
	// example:
	//
	// 1e810dfe1468721d0664a49b9d9f74f4
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The region ID. This parameter is required if a cloud service is a regionalized.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *UpdateResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateResourceRequest) SetBody(v map[string]interface{}) *UpdateResourceRequest {
	s.Body = v
	return s
}

func (s *UpdateResourceRequest) SetClientToken(v string) *UpdateResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateResourceRequest) SetRegionId(v string) *UpdateResourceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateResourceRequest) Validate() error {
	return dara.Validate(s)
}
