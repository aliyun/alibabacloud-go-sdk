// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAirflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowId(v string) *DeleteAirflowRequest
	GetAirflowId() *string
	SetClientToken(v string) *DeleteAirflowRequest
	GetClientToken() *string
	SetWorkspaceId(v string) *DeleteAirflowRequest
	GetWorkspaceId() *string
}

type DeleteAirflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af-test****
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// example:
	//
	// token-****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86302423828****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteAirflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteAirflowRequest) GetAirflowId() *string {
	return s.AirflowId
}

func (s *DeleteAirflowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAirflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteAirflowRequest) SetAirflowId(v string) *DeleteAirflowRequest {
	s.AirflowId = &v
	return s
}

func (s *DeleteAirflowRequest) SetClientToken(v string) *DeleteAirflowRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAirflowRequest) SetWorkspaceId(v string) *DeleteAirflowRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteAirflowRequest) Validate() error {
	return dara.Validate(s)
}
