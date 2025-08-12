// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlipayUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAlipayUrlRequest
	GetAppId() *string
	SetWorkspaceId(v string) *GetAlipayUrlRequest
	GetWorkspaceId() *string
}

type GetAlipayUrlRequest struct {
	AppId       *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	WorkspaceId *string `json:"workspace_id,omitempty" xml:"workspace_id,omitempty"`
}

func (s GetAlipayUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAlipayUrlRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAlipayUrlRequest) SetAppId(v string) *GetAlipayUrlRequest {
	s.AppId = &v
	return s
}

func (s *GetAlipayUrlRequest) SetWorkspaceId(v string) *GetAlipayUrlRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetAlipayUrlRequest) Validate() error {
	return dara.Validate(s)
}
