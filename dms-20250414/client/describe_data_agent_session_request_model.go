// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *DescribeDataAgentSessionRequest
	GetDMSUnit() *string
	SetSessionId(v string) *DescribeDataAgentSessionRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *DescribeDataAgentSessionRequest
	GetWorkspaceId() *string
}

type DescribeDataAgentSessionRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// c61n7gmopfvuascr70xmavzrj
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeDataAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *DescribeDataAgentSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeDataAgentSessionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeDataAgentSessionRequest) SetDMSUnit(v string) *DescribeDataAgentSessionRequest {
	s.DMSUnit = &v
	return s
}

func (s *DescribeDataAgentSessionRequest) SetSessionId(v string) *DescribeDataAgentSessionRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeDataAgentSessionRequest) SetWorkspaceId(v string) *DescribeDataAgentSessionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeDataAgentSessionRequest) Validate() error {
	return dara.Validate(s)
}
