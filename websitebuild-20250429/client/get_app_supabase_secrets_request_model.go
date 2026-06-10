// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppSupabaseSecretsRequest
	GetBizId() *string
	SetKeyword(v string) *GetAppSupabaseSecretsRequest
	GetKeyword() *string
}

type GetAppSupabaseSecretsRequest struct {
	// Business ID
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Search keyword
	//
	// example:
	//
	// v2_
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s GetAppSupabaseSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseSecretsRequest) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseSecretsRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppSupabaseSecretsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetAppSupabaseSecretsRequest) SetBizId(v string) *GetAppSupabaseSecretsRequest {
	s.BizId = &v
	return s
}

func (s *GetAppSupabaseSecretsRequest) SetKeyword(v string) *GetAppSupabaseSecretsRequest {
	s.Keyword = &v
	return s
}

func (s *GetAppSupabaseSecretsRequest) Validate() error {
	return dara.Validate(s)
}
