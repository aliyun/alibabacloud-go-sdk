// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployDifyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RedeployDifyInstanceRequest
	GetClientToken() *string
	SetDataRegion(v string) *RedeployDifyInstanceRequest
	GetDataRegion() *string
	SetDryRun(v string) *RedeployDifyInstanceRequest
	GetDryRun() *string
	SetWorkspaceId(v string) *RedeployDifyInstanceRequest
	GetWorkspaceId() *string
}

type RedeployDifyInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataRegion  *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	DryRun      *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RedeployDifyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RedeployDifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *RedeployDifyInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RedeployDifyInstanceRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *RedeployDifyInstanceRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *RedeployDifyInstanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RedeployDifyInstanceRequest) SetClientToken(v string) *RedeployDifyInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RedeployDifyInstanceRequest) SetDataRegion(v string) *RedeployDifyInstanceRequest {
	s.DataRegion = &v
	return s
}

func (s *RedeployDifyInstanceRequest) SetDryRun(v string) *RedeployDifyInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *RedeployDifyInstanceRequest) SetWorkspaceId(v string) *RedeployDifyInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RedeployDifyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
