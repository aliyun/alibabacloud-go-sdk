// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceQueuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *ListWorkspaceQueuesRequest
	GetEnvironment() *string
	SetRegionId(v string) *ListWorkspaceQueuesRequest
	GetRegionId() *string
}

type ListWorkspaceQueuesRequest struct {
	// The environment type.
	//
	// Valid values:
	//
	// 	- dev
	//
	// 	- production
	//
	// example:
	//
	// production
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListWorkspaceQueuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ListWorkspaceQueuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkspaceQueuesRequest) SetEnvironment(v string) *ListWorkspaceQueuesRequest {
	s.Environment = &v
	return s
}

func (s *ListWorkspaceQueuesRequest) SetRegionId(v string) *ListWorkspaceQueuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListWorkspaceQueuesRequest) Validate() error {
	return dara.Validate(s)
}
