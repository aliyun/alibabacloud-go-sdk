// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMappCenterAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMappCenterAppRequest
	GetAppId() *string
	SetWorkspaceId(v string) *QueryMappCenterAppRequest
	GetWorkspaceId() *string
}

type QueryMappCenterAppRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMappCenterAppRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMappCenterAppRequest) GoString() string {
	return s.String()
}

func (s *QueryMappCenterAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMappCenterAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMappCenterAppRequest) SetAppId(v string) *QueryMappCenterAppRequest {
	s.AppId = &v
	return s
}

func (s *QueryMappCenterAppRequest) SetWorkspaceId(v string) *QueryMappCenterAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMappCenterAppRequest) Validate() error {
	return dara.Validate(s)
}
