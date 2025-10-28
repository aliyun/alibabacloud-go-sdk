// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScalingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetScalingRulesRequest
	GetAppId() *string
	SetGroupId(v string) *GetScalingRulesRequest
	GetGroupId() *string
	SetMode(v string) *GetScalingRulesRequest
	GetMode() *string
}

type GetScalingRulesRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33e39be9-3e5f-*********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance group to which the application is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// d8bb9d60-9**************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The type of the scaling rule. You can leave this parameter empty. Valid values:
	//
	// 	- SCALE_IN: scale-in rules
	//
	// 	- SCALE_OUT: scale-out rules
	//
	// example:
	//
	// SCALE_IN
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s GetScalingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *GetScalingRulesRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetScalingRulesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetScalingRulesRequest) GetMode() *string {
	return s.Mode
}

func (s *GetScalingRulesRequest) SetAppId(v string) *GetScalingRulesRequest {
	s.AppId = &v
	return s
}

func (s *GetScalingRulesRequest) SetGroupId(v string) *GetScalingRulesRequest {
	s.GroupId = &v
	return s
}

func (s *GetScalingRulesRequest) SetMode(v string) *GetScalingRulesRequest {
	s.Mode = &v
	return s
}

func (s *GetScalingRulesRequest) Validate() error {
	return dara.Validate(s)
}
