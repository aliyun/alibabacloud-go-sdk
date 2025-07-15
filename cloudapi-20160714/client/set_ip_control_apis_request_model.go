// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIpControlApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *SetIpControlApisRequest
	GetApiIds() *string
	SetGroupId(v string) *SetIpControlApisRequest
	GetGroupId() *string
	SetIpControlId(v string) *SetIpControlApisRequest
	GetIpControlId() *string
	SetSecurityToken(v string) *SetIpControlApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *SetIpControlApisRequest
	GetStageName() *string
}

type SetIpControlApisRequest struct {
	// The API IDs. Separate multiple API IDs with commas (,). A maximum of 100 API IDs can be entered.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6,46fbb52840d146f186e38e8e70fc8c12
	ApiIds *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **PRE**
	//
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetIpControlApisRequest) String() string {
	return dara.Prettify(s)
}

func (s SetIpControlApisRequest) GoString() string {
	return s.String()
}

func (s *SetIpControlApisRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *SetIpControlApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SetIpControlApisRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *SetIpControlApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetIpControlApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *SetIpControlApisRequest) SetApiIds(v string) *SetIpControlApisRequest {
	s.ApiIds = &v
	return s
}

func (s *SetIpControlApisRequest) SetGroupId(v string) *SetIpControlApisRequest {
	s.GroupId = &v
	return s
}

func (s *SetIpControlApisRequest) SetIpControlId(v string) *SetIpControlApisRequest {
	s.IpControlId = &v
	return s
}

func (s *SetIpControlApisRequest) SetSecurityToken(v string) *SetIpControlApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetIpControlApisRequest) SetStageName(v string) *SetIpControlApisRequest {
	s.StageName = &v
	return s
}

func (s *SetIpControlApisRequest) Validate() error {
	return dara.Validate(s)
}
