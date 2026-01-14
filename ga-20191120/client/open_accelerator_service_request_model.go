// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAcceleratorServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *OpenAcceleratorServiceRequest
	GetClientToken() *string
	SetRegionId(v string) *OpenAcceleratorServiceRequest
	GetRegionId() *string
}

type OpenAcceleratorServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenAcceleratorServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenAcceleratorServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenAcceleratorServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *OpenAcceleratorServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenAcceleratorServiceRequest) SetClientToken(v string) *OpenAcceleratorServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenAcceleratorServiceRequest) SetRegionId(v string) *OpenAcceleratorServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenAcceleratorServiceRequest) Validate() error {
	return dara.Validate(s)
}
