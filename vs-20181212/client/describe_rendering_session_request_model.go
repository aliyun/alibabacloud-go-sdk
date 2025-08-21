// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeRenderingSessionRequest
	GetClientId() *string
	SetProjectId(v string) *DescribeRenderingSessionRequest
	GetProjectId() *string
	SetSessionId(v string) *DescribeRenderingSessionRequest
	GetSessionId() *string
}

type DescribeRenderingSessionRequest struct {
	// example:
	//
	// d27c89d6-4fe3-4855-a89c-ea721c708b0b
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// session-i205217481741918129226
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeRenderingSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingSessionRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenderingSessionRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeRenderingSessionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeRenderingSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeRenderingSessionRequest) SetClientId(v string) *DescribeRenderingSessionRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeRenderingSessionRequest) SetProjectId(v string) *DescribeRenderingSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeRenderingSessionRequest) SetSessionId(v string) *DescribeRenderingSessionRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeRenderingSessionRequest) Validate() error {
	return dara.Validate(s)
}
