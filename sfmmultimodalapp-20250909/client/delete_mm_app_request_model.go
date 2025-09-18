// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMmAppRequest
	GetAppId() *string
	SetWorkspaceId(v string) *DeleteMmAppRequest
	GetWorkspaceId() *string
}

type DeleteMmAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMmAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteMmAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMmAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMmAppRequest) SetAppId(v string) *DeleteMmAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMmAppRequest) SetWorkspaceId(v string) *DeleteMmAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMmAppRequest) Validate() error {
	return dara.Validate(s)
}
