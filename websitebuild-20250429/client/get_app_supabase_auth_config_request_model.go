// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseAuthConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *GetAppSupabaseAuthConfigRequest
	GetAuthType() *string
	SetBizId(v string) *GetAppSupabaseAuthConfigRequest
	GetBizId() *string
}

type GetAppSupabaseAuthConfigRequest struct {
	// example:
	//
	// ALL_AUTH_FLAG
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// WS20250801152639000005
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetAppSupabaseAuthConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseAuthConfigRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *GetAppSupabaseAuthConfigRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppSupabaseAuthConfigRequest) SetAuthType(v string) *GetAppSupabaseAuthConfigRequest {
	s.AuthType = &v
	return s
}

func (s *GetAppSupabaseAuthConfigRequest) SetBizId(v string) *GetAppSupabaseAuthConfigRequest {
	s.BizId = &v
	return s
}

func (s *GetAppSupabaseAuthConfigRequest) Validate() error {
	return dara.Validate(s)
}
