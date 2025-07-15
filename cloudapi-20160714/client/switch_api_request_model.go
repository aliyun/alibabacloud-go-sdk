// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *SwitchApiRequest
	GetApiId() *string
	SetDescription(v string) *SwitchApiRequest
	GetDescription() *string
	SetGroupId(v string) *SwitchApiRequest
	GetGroupId() *string
	SetHistoryVersion(v string) *SwitchApiRequest
	GetHistoryVersion() *string
	SetSecurityToken(v string) *SwitchApiRequest
	GetSecurityToken() *string
	SetStageName(v string) *SwitchApiRequest
	GetStageName() *string
}

type SwitchApiRequest struct {
	// The API ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d6f679aeb3be4b91b3688e887ca1fe16
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The description. The description can be up to 200 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// for_demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The API group ID.
	//
	// example:
	//
	// 123
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The historical version number of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20160705104552292
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the staging environment
	//
	// 	- **TEST**: the test environment
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SwitchApiRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchApiRequest) GoString() string {
	return s.String()
}

func (s *SwitchApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *SwitchApiRequest) GetDescription() *string {
	return s.Description
}

func (s *SwitchApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SwitchApiRequest) GetHistoryVersion() *string {
	return s.HistoryVersion
}

func (s *SwitchApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SwitchApiRequest) GetStageName() *string {
	return s.StageName
}

func (s *SwitchApiRequest) SetApiId(v string) *SwitchApiRequest {
	s.ApiId = &v
	return s
}

func (s *SwitchApiRequest) SetDescription(v string) *SwitchApiRequest {
	s.Description = &v
	return s
}

func (s *SwitchApiRequest) SetGroupId(v string) *SwitchApiRequest {
	s.GroupId = &v
	return s
}

func (s *SwitchApiRequest) SetHistoryVersion(v string) *SwitchApiRequest {
	s.HistoryVersion = &v
	return s
}

func (s *SwitchApiRequest) SetSecurityToken(v string) *SwitchApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchApiRequest) SetStageName(v string) *SwitchApiRequest {
	s.StageName = &v
	return s
}

func (s *SwitchApiRequest) Validate() error {
	return dara.Validate(s)
}
