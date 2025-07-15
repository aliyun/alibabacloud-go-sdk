// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DeployApiRequest
	GetApiId() *string
	SetDescription(v string) *DeployApiRequest
	GetDescription() *string
	SetGroupId(v string) *DeployApiRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DeployApiRequest
	GetSecurityToken() *string
	SetStageName(v string) *DeployApiRequest
	GetStageName() *string
}

type DeployApiRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// d6f679aeb3be4b91b3688e887ca1fe16
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The publishing remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// for_test1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 63be9002440b4778a61122f14c2b2bbb
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **PRE**: the pre-release environment
	//
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DeployApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployApiRequest) GoString() string {
	return s.String()
}

func (s *DeployApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DeployApiRequest) GetDescription() *string {
	return s.Description
}

func (s *DeployApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeployApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeployApiRequest) GetStageName() *string {
	return s.StageName
}

func (s *DeployApiRequest) SetApiId(v string) *DeployApiRequest {
	s.ApiId = &v
	return s
}

func (s *DeployApiRequest) SetDescription(v string) *DeployApiRequest {
	s.Description = &v
	return s
}

func (s *DeployApiRequest) SetGroupId(v string) *DeployApiRequest {
	s.GroupId = &v
	return s
}

func (s *DeployApiRequest) SetSecurityToken(v string) *DeployApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeployApiRequest) SetStageName(v string) *DeployApiRequest {
	s.StageName = &v
	return s
}

func (s *DeployApiRequest) Validate() error {
	return dara.Validate(s)
}
