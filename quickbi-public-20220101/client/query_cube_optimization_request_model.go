// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubeOptimizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *QueryCubeOptimizationRequest
	GetWorkspaceId() *string
}

type QueryCubeOptimizationRequest struct {
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryCubeOptimizationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCubeOptimizationRequest) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryCubeOptimizationRequest) SetWorkspaceId(v string) *QueryCubeOptimizationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCubeOptimizationRequest) Validate() error {
	return dara.Validate(s)
}
