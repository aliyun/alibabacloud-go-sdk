// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTrafficControlApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *RemoveTrafficControlApisRequest
	GetApiIds() *string
	SetGroupId(v string) *RemoveTrafficControlApisRequest
	GetGroupId() *string
	SetSecurityToken(v string) *RemoveTrafficControlApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *RemoveTrafficControlApisRequest
	GetStageName() *string
	SetTrafficControlId(v string) *RemoveTrafficControlApisRequest
	GetTrafficControlId() *string
}

type RemoveTrafficControlApisRequest struct {
	// The IDs of the APIs from which you want to unbind a specified throttling policy.
	//
	// 	- If this parameter is not specified, the throttling policy is unbound from all the APIs in the specified environment of the API group.
	//
	// 	- Separate multiple API IDs with commas (,). A maximum of 100 API IDs can be entered.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6,46fbb52840d146f186e38e8e70fc8c12
	ApiIds *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	// The ID of the API group containing the APIs from which you want to unbind a specified throttling policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 4223a10e-eed3-46a6-8b7c-23003f488153
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
	// The ID of the throttling policy that you want to unbind from APIs.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s RemoveTrafficControlApisRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTrafficControlApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveTrafficControlApisRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *RemoveTrafficControlApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveTrafficControlApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveTrafficControlApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *RemoveTrafficControlApisRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *RemoveTrafficControlApisRequest) SetApiIds(v string) *RemoveTrafficControlApisRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) SetGroupId(v string) *RemoveTrafficControlApisRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) SetSecurityToken(v string) *RemoveTrafficControlApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) SetStageName(v string) *RemoveTrafficControlApisRequest {
	s.StageName = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) SetTrafficControlId(v string) *RemoveTrafficControlApisRequest {
	s.TrafficControlId = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) Validate() error {
	return dara.Validate(s)
}
