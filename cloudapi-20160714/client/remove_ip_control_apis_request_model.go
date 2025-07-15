// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpControlApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *RemoveIpControlApisRequest
	GetApiIds() *string
	SetGroupId(v string) *RemoveIpControlApisRequest
	GetGroupId() *string
	SetIpControlId(v string) *RemoveIpControlApisRequest
	GetIpControlId() *string
	SetSecurityToken(v string) *RemoveIpControlApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *RemoveIpControlApisRequest
	GetStageName() *string
}

type RemoveIpControlApisRequest struct {
	// The IDs of the APIs from which you want to unbind the ACL.
	//
	// 	- If this parameter is not specified, the ACL is unbound from all the APIs in the specified environment of the API group.
	//
	// 	- The IDs of APIs that you want to query. Separate multiple API IDs with commas (,). A maximum of 100 API IDs can be entered.
	//
	// example:
	//
	// 123
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
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveIpControlApisRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpControlApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveIpControlApisRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *RemoveIpControlApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveIpControlApisRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *RemoveIpControlApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveIpControlApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *RemoveIpControlApisRequest) SetApiIds(v string) *RemoveIpControlApisRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveIpControlApisRequest) SetGroupId(v string) *RemoveIpControlApisRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveIpControlApisRequest) SetIpControlId(v string) *RemoveIpControlApisRequest {
	s.IpControlId = &v
	return s
}

func (s *RemoveIpControlApisRequest) SetSecurityToken(v string) *RemoveIpControlApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveIpControlApisRequest) SetStageName(v string) *RemoveIpControlApisRequest {
	s.StageName = &v
	return s
}

func (s *RemoveIpControlApisRequest) Validate() error {
	return dara.Validate(s)
}
