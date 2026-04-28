// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareLinkTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireSec(v int32) *GetShareLinkTokenRequest
	GetExpireSec() *int32
	SetShareId(v string) *GetShareLinkTokenRequest
	GetShareId() *string
	SetSharePwd(v string) *GetShareLinkTokenRequest
	GetSharePwd() *string
}

type GetShareLinkTokenRequest struct {
	// The validity period of the token. Valid values: (0,7200]. Default value: 7200. Unit: seconds.
	//
	// example:
	//
	// 7200
	ExpireSec *int32 `json:"expire_sec,omitempty" xml:"expire_sec,omitempty"`
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The access code.
	//
	// example:
	//
	// abcF123x
	SharePwd *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
}

func (s GetShareLinkTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetShareLinkTokenRequest) GoString() string {
	return s.String()
}

func (s *GetShareLinkTokenRequest) GetExpireSec() *int32 {
	return s.ExpireSec
}

func (s *GetShareLinkTokenRequest) GetShareId() *string {
	return s.ShareId
}

func (s *GetShareLinkTokenRequest) GetSharePwd() *string {
	return s.SharePwd
}

func (s *GetShareLinkTokenRequest) SetExpireSec(v int32) *GetShareLinkTokenRequest {
	s.ExpireSec = &v
	return s
}

func (s *GetShareLinkTokenRequest) SetShareId(v string) *GetShareLinkTokenRequest {
	s.ShareId = &v
	return s
}

func (s *GetShareLinkTokenRequest) SetSharePwd(v string) *GetShareLinkTokenRequest {
	s.SharePwd = &v
	return s
}

func (s *GetShareLinkTokenRequest) Validate() error {
	return dara.Validate(s)
}
