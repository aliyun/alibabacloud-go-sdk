// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEnterpriseCodeRequest
	GetClientToken() *string
	SetEnterpriseCode(v string) *UpdateEnterpriseCodeRequest
	GetEnterpriseCode() *string
	SetIsDefault(v bool) *UpdateEnterpriseCodeRequest
	GetIsDefault() *bool
	SetRegionId(v string) *UpdateEnterpriseCodeRequest
	GetRegionId() *string
}

type UpdateEnterpriseCodeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12P**
	EnterpriseCode *string `json:"EnterpriseCode,omitempty" xml:"EnterpriseCode,omitempty"`
	// Specifies whether to specify the enterprise code as the default one. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the region to which the enterprise code belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEnterpriseCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseCodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseCodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEnterpriseCodeRequest) GetEnterpriseCode() *string {
	return s.EnterpriseCode
}

func (s *UpdateEnterpriseCodeRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *UpdateEnterpriseCodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnterpriseCodeRequest) SetClientToken(v string) *UpdateEnterpriseCodeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEnterpriseCodeRequest) SetEnterpriseCode(v string) *UpdateEnterpriseCodeRequest {
	s.EnterpriseCode = &v
	return s
}

func (s *UpdateEnterpriseCodeRequest) SetIsDefault(v bool) *UpdateEnterpriseCodeRequest {
	s.IsDefault = &v
	return s
}

func (s *UpdateEnterpriseCodeRequest) SetRegionId(v string) *UpdateEnterpriseCodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnterpriseCodeRequest) Validate() error {
	return dara.Validate(s)
}
