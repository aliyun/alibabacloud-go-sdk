// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMscpRiskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApdidToken(v string) *QueryMscpRiskInfoRequest
	GetApdidToken() *string
	SetAppId(v string) *QueryMscpRiskInfoRequest
	GetAppId() *string
	SetTenantId(v string) *QueryMscpRiskInfoRequest
	GetTenantId() *string
	SetTerminalType(v string) *QueryMscpRiskInfoRequest
	GetTerminalType() *string
	SetWorkspaceId(v string) *QueryMscpRiskInfoRequest
	GetWorkspaceId() *string
}

type QueryMscpRiskInfoRequest struct {
	// ApdidToken
	//
	// example:
	//
	// ApdidToken
	ApdidToken *string `json:"ApdidToken,omitempty" xml:"ApdidToken,omitempty"`
	// AppId
	//
	// example:
	//
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// TenantId
	//
	// example:
	//
	// TenantId
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// TerminalType
	//
	// example:
	//
	// TerminalType
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
	// WorkspaceId
	//
	// example:
	//
	// WorkspaceId
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMscpRiskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMscpRiskInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMscpRiskInfoRequest) GetApdidToken() *string {
	return s.ApdidToken
}

func (s *QueryMscpRiskInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMscpRiskInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMscpRiskInfoRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *QueryMscpRiskInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMscpRiskInfoRequest) SetApdidToken(v string) *QueryMscpRiskInfoRequest {
	s.ApdidToken = &v
	return s
}

func (s *QueryMscpRiskInfoRequest) SetAppId(v string) *QueryMscpRiskInfoRequest {
	s.AppId = &v
	return s
}

func (s *QueryMscpRiskInfoRequest) SetTenantId(v string) *QueryMscpRiskInfoRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMscpRiskInfoRequest) SetTerminalType(v string) *QueryMscpRiskInfoRequest {
	s.TerminalType = &v
	return s
}

func (s *QueryMscpRiskInfoRequest) SetWorkspaceId(v string) *QueryMscpRiskInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMscpRiskInfoRequest) Validate() error {
	return dara.Validate(s)
}
