// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployAirflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowId(v string) *RedeployAirflowRequest
	GetAirflowId() *string
	SetWorkspaceId(v string) *RedeployAirflowRequest
	GetWorkspaceId() *string
}

type RedeployAirflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af-b3a7f110a6vmvn7****
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RedeployAirflowRequest) String() string {
	return dara.Prettify(s)
}

func (s RedeployAirflowRequest) GoString() string {
	return s.String()
}

func (s *RedeployAirflowRequest) GetAirflowId() *string {
	return s.AirflowId
}

func (s *RedeployAirflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RedeployAirflowRequest) SetAirflowId(v string) *RedeployAirflowRequest {
	s.AirflowId = &v
	return s
}

func (s *RedeployAirflowRequest) SetWorkspaceId(v string) *RedeployAirflowRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RedeployAirflowRequest) Validate() error {
	return dara.Validate(s)
}
