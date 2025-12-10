// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetAlgorithmTreeRequest
	GetSource() *string
	SetWorkspaceId(v string) *GetAlgorithmTreeRequest
	GetWorkspaceId() *string
}

type GetAlgorithmTreeRequest struct {
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAlgorithmTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmTreeRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmTreeRequest) GetSource() *string {
	return s.Source
}

func (s *GetAlgorithmTreeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAlgorithmTreeRequest) SetSource(v string) *GetAlgorithmTreeRequest {
	s.Source = &v
	return s
}

func (s *GetAlgorithmTreeRequest) SetWorkspaceId(v string) *GetAlgorithmTreeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAlgorithmTreeRequest) Validate() error {
	return dara.Validate(s)
}
