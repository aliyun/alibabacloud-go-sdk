// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentsUsersStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetExperimentsUsersStatisticsRequest
	GetSource() *string
	SetWorkspaceId(v string) *GetExperimentsUsersStatisticsRequest
	GetWorkspaceId() *string
}

type GetExperimentsUsersStatisticsRequest struct {
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentsUsersStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentsUsersStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentsUsersStatisticsRequest) GetSource() *string {
	return s.Source
}

func (s *GetExperimentsUsersStatisticsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetExperimentsUsersStatisticsRequest) SetSource(v string) *GetExperimentsUsersStatisticsRequest {
	s.Source = &v
	return s
}

func (s *GetExperimentsUsersStatisticsRequest) SetWorkspaceId(v string) *GetExperimentsUsersStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetExperimentsUsersStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
