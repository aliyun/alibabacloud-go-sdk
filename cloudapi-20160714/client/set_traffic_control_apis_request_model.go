// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTrafficControlApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *SetTrafficControlApisRequest
	GetApiIds() *string
	SetGroupId(v string) *SetTrafficControlApisRequest
	GetGroupId() *string
	SetSecurityToken(v string) *SetTrafficControlApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *SetTrafficControlApisRequest
	GetStageName() *string
	SetTrafficControlId(v string) *SetTrafficControlApisRequest
	GetTrafficControlId() *string
}

type SetTrafficControlApisRequest struct {
	// The API ID for the specified operation. Separate multiple API IDs with commas (,). A maximum of 100 API IDs can be entered.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6,46fbb52840d146f186e38e8e70fc8c12
	ApiIds *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	// The ID of the API group containing the APIs to which you want to bind a specified throttling policy.
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
	// 556d15cb-0808-432d-ab07-33e6b961b703
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
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The ID of the throttling policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s SetTrafficControlApisRequest) String() string {
	return dara.Prettify(s)
}

func (s SetTrafficControlApisRequest) GoString() string {
	return s.String()
}

func (s *SetTrafficControlApisRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *SetTrafficControlApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SetTrafficControlApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetTrafficControlApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *SetTrafficControlApisRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *SetTrafficControlApisRequest) SetApiIds(v string) *SetTrafficControlApisRequest {
	s.ApiIds = &v
	return s
}

func (s *SetTrafficControlApisRequest) SetGroupId(v string) *SetTrafficControlApisRequest {
	s.GroupId = &v
	return s
}

func (s *SetTrafficControlApisRequest) SetSecurityToken(v string) *SetTrafficControlApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetTrafficControlApisRequest) SetStageName(v string) *SetTrafficControlApisRequest {
	s.StageName = &v
	return s
}

func (s *SetTrafficControlApisRequest) SetTrafficControlId(v string) *SetTrafficControlApisRequest {
	s.TrafficControlId = &v
	return s
}

func (s *SetTrafficControlApisRequest) Validate() error {
	return dara.Validate(s)
}
