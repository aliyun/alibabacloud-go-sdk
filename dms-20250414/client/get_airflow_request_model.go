// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAirflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowId(v string) *GetAirflowRequest
	GetAirflowId() *string
	SetWorkspaceId(v string) *GetAirflowRequest
	GetWorkspaceId() *string
}

type GetAirflowRequest struct {
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
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAirflowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAirflowRequest) GoString() string {
	return s.String()
}

func (s *GetAirflowRequest) GetAirflowId() *string {
	return s.AirflowId
}

func (s *GetAirflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAirflowRequest) SetAirflowId(v string) *GetAirflowRequest {
	s.AirflowId = &v
	return s
}

func (s *GetAirflowRequest) SetWorkspaceId(v string) *GetAirflowRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAirflowRequest) Validate() error {
	return dara.Validate(s)
}
