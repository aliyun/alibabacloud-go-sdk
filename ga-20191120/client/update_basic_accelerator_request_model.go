// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateBasicAcceleratorRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *UpdateBasicAcceleratorRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateBasicAcceleratorRequest
	GetDescription() *string
	SetName(v string) *UpdateBasicAcceleratorRequest
	GetName() *string
	SetRegionId(v string) *UpdateBasicAcceleratorRequest
	GetRegionId() *string
}

type UpdateBasicAcceleratorRequest struct {
	// The ID of the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the basic GA instance.
	//
	// The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// BasicAccelerator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the basic GA instance.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// BasicAccelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBasicAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateBasicAcceleratorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateBasicAcceleratorRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateBasicAcceleratorRequest) GetName() *string {
	return s.Name
}

func (s *UpdateBasicAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBasicAcceleratorRequest) SetAcceleratorId(v string) *UpdateBasicAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) SetClientToken(v string) *UpdateBasicAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) SetDescription(v string) *UpdateBasicAcceleratorRequest {
	s.Description = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) SetName(v string) *UpdateBasicAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) SetRegionId(v string) *UpdateBasicAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
