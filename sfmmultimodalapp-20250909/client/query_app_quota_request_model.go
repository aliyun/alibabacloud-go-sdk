// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryAppQuotaRequest
	GetAppId() *string
	SetWorkspaceId(v string) *QueryAppQuotaRequest
	GetWorkspaceId() *string
}

type QueryAppQuotaRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryAppQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAppQuotaRequest) GoString() string {
	return s.String()
}

func (s *QueryAppQuotaRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryAppQuotaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryAppQuotaRequest) SetAppId(v string) *QueryAppQuotaRequest {
	s.AppId = &v
	return s
}

func (s *QueryAppQuotaRequest) SetWorkspaceId(v string) *QueryAppQuotaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryAppQuotaRequest) Validate() error {
	return dara.Validate(s)
}
