// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGEnterpriseCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSmartAGEnterpriseCodeRequest
	GetClientToken() *string
	SetEnterpriseCode(v string) *UpdateSmartAGEnterpriseCodeRequest
	GetEnterpriseCode() *string
	SetRegionId(v string) *UpdateSmartAGEnterpriseCodeRequest
	GetRegionId() *string
	SetSmartAGId(v string) *UpdateSmartAGEnterpriseCodeRequest
	GetSmartAGId() *string
}

type UpdateSmartAGEnterpriseCodeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enterprise code with you want to associate the SAG APP instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12P**
	EnterpriseCode *string `json:"EnterpriseCode,omitempty" xml:"EnterpriseCode,omitempty"`
	// The ID of the region where the SAG APP instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the SAG APP instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-rz2e23c0e78ema****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s UpdateSmartAGEnterpriseCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGEnterpriseCodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGEnterpriseCodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSmartAGEnterpriseCodeRequest) GetEnterpriseCode() *string {
	return s.EnterpriseCode
}

func (s *UpdateSmartAGEnterpriseCodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAGEnterpriseCodeRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UpdateSmartAGEnterpriseCodeRequest) SetClientToken(v string) *UpdateSmartAGEnterpriseCodeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSmartAGEnterpriseCodeRequest) SetEnterpriseCode(v string) *UpdateSmartAGEnterpriseCodeRequest {
	s.EnterpriseCode = &v
	return s
}

func (s *UpdateSmartAGEnterpriseCodeRequest) SetRegionId(v string) *UpdateSmartAGEnterpriseCodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAGEnterpriseCodeRequest) SetSmartAGId(v string) *UpdateSmartAGEnterpriseCodeRequest {
	s.SmartAGId = &v
	return s
}

func (s *UpdateSmartAGEnterpriseCodeRequest) Validate() error {
	return dara.Validate(s)
}
