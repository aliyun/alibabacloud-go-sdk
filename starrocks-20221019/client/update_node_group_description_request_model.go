// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateNodeGroupDescriptionRequest
	GetDescription() *string
	SetNodeGroupId(v string) *UpdateNodeGroupDescriptionRequest
	GetNodeGroupId() *string
	SetXAcsRamAuthContext(v string) *UpdateNodeGroupDescriptionRequest
	GetXAcsRamAuthContext() *string
}

type UpdateNodeGroupDescriptionRequest struct {
	// example:
	//
	// okcc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId        *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	XAcsRamAuthContext *string `json:"X-Acs-Ram-Auth-Context,omitempty" xml:"X-Acs-Ram-Auth-Context,omitempty"`
}

func (s UpdateNodeGroupDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateNodeGroupDescriptionRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateNodeGroupDescriptionRequest) GetXAcsRamAuthContext() *string {
	return s.XAcsRamAuthContext
}

func (s *UpdateNodeGroupDescriptionRequest) SetDescription(v string) *UpdateNodeGroupDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateNodeGroupDescriptionRequest) SetNodeGroupId(v string) *UpdateNodeGroupDescriptionRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateNodeGroupDescriptionRequest) SetXAcsRamAuthContext(v string) *UpdateNodeGroupDescriptionRequest {
	s.XAcsRamAuthContext = &v
	return s
}

func (s *UpdateNodeGroupDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
