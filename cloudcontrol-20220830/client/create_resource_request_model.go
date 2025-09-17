// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *CreateResourceRequest
	GetBody() map[string]interface{}
	SetClientToken(v string) *CreateResourceRequest
	GetClientToken() *string
	SetRegionId(v string) *CreateResourceRequest
	GetRegionId() *string
}

type CreateResourceRequest struct {
	// The request body. The property of the resource, which is specified in JSON format.
	//
	// example:
	//
	// {
	//
	//      "AccountName": "cctest",
	//
	//      "AccountPassword": "Aa1234****"
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

func (s CreateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *CreateResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateResourceRequest) SetBody(v map[string]interface{}) *CreateResourceRequest {
	s.Body = v
	return s
}

func (s *CreateResourceRequest) SetClientToken(v string) *CreateResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateResourceRequest) SetRegionId(v string) *CreateResourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateResourceRequest) Validate() error {
	return dara.Validate(s)
}
