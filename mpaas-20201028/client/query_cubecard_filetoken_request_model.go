// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubecardFiletokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryCubecardFiletokenRequest
	GetAppId() *string
	SetOnexFlag(v bool) *QueryCubecardFiletokenRequest
	GetOnexFlag() *bool
	SetTenantId(v string) *QueryCubecardFiletokenRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryCubecardFiletokenRequest
	GetWorkspaceId() *string
}

type QueryCubecardFiletokenRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	OnexFlag *bool `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryCubecardFiletokenRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCubecardFiletokenRequest) GoString() string {
	return s.String()
}

func (s *QueryCubecardFiletokenRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryCubecardFiletokenRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *QueryCubecardFiletokenRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryCubecardFiletokenRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryCubecardFiletokenRequest) SetAppId(v string) *QueryCubecardFiletokenRequest {
	s.AppId = &v
	return s
}

func (s *QueryCubecardFiletokenRequest) SetOnexFlag(v bool) *QueryCubecardFiletokenRequest {
	s.OnexFlag = &v
	return s
}

func (s *QueryCubecardFiletokenRequest) SetTenantId(v string) *QueryCubecardFiletokenRequest {
	s.TenantId = &v
	return s
}

func (s *QueryCubecardFiletokenRequest) SetWorkspaceId(v string) *QueryCubecardFiletokenRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCubecardFiletokenRequest) Validate() error {
	return dara.Validate(s)
}
