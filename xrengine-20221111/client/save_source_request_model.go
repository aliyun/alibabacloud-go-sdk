// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeStatus(v bool) *SaveSourceRequest
	GetChangeStatus() *bool
	SetJwtToken(v string) *SaveSourceRequest
	GetJwtToken() *string
	SetNeedCheck(v bool) *SaveSourceRequest
	GetNeedCheck() *bool
	SetProjectId(v int64) *SaveSourceRequest
	GetProjectId() *int64
}

type SaveSourceRequest struct {
	ChangeStatus *bool   `json:"ChangeStatus,omitempty" xml:"ChangeStatus,omitempty"`
	JwtToken     *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	NeedCheck    *bool   `json:"NeedCheck,omitempty" xml:"NeedCheck,omitempty"`
	ProjectId    *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s SaveSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSourceRequest) GoString() string {
	return s.String()
}

func (s *SaveSourceRequest) GetChangeStatus() *bool {
	return s.ChangeStatus
}

func (s *SaveSourceRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *SaveSourceRequest) GetNeedCheck() *bool {
	return s.NeedCheck
}

func (s *SaveSourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SaveSourceRequest) SetChangeStatus(v bool) *SaveSourceRequest {
	s.ChangeStatus = &v
	return s
}

func (s *SaveSourceRequest) SetJwtToken(v string) *SaveSourceRequest {
	s.JwtToken = &v
	return s
}

func (s *SaveSourceRequest) SetNeedCheck(v bool) *SaveSourceRequest {
	s.NeedCheck = &v
	return s
}

func (s *SaveSourceRequest) SetProjectId(v int64) *SaveSourceRequest {
	s.ProjectId = &v
	return s
}

func (s *SaveSourceRequest) Validate() error {
	return dara.Validate(s)
}
