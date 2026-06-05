// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSupabaseSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteAppSupabaseSecretsRequest
	GetBizId() *string
	SetKeysJson(v string) *DeleteAppSupabaseSecretsRequest
	GetKeysJson() *string
}

type DeleteAppSupabaseSecretsRequest struct {
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// {"key":"key"}
	KeysJson *string `json:"KeysJson,omitempty" xml:"KeysJson,omitempty"`
}

func (s DeleteAppSupabaseSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSupabaseSecretsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppSupabaseSecretsRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteAppSupabaseSecretsRequest) GetKeysJson() *string {
	return s.KeysJson
}

func (s *DeleteAppSupabaseSecretsRequest) SetBizId(v string) *DeleteAppSupabaseSecretsRequest {
	s.BizId = &v
	return s
}

func (s *DeleteAppSupabaseSecretsRequest) SetKeysJson(v string) *DeleteAppSupabaseSecretsRequest {
	s.KeysJson = &v
	return s
}

func (s *DeleteAppSupabaseSecretsRequest) Validate() error {
	return dara.Validate(s)
}
