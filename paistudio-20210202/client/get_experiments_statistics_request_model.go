// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentsStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetExperimentsStatisticsRequest
	GetSource() *string
	SetWorkspaceIds(v string) *GetExperimentsStatisticsRequest
	GetWorkspaceIds() *string
}

type GetExperimentsStatisticsRequest struct {
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123,345
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s GetExperimentsStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentsStatisticsRequest) GetSource() *string {
	return s.Source
}

func (s *GetExperimentsStatisticsRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *GetExperimentsStatisticsRequest) SetSource(v string) *GetExperimentsStatisticsRequest {
	s.Source = &v
	return s
}

func (s *GetExperimentsStatisticsRequest) SetWorkspaceIds(v string) *GetExperimentsStatisticsRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *GetExperimentsStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
