// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAppSupabaseSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SaveAppSupabaseSecretsRequest
	GetBizId() *string
	SetSecretsJson(v string) *SaveAppSupabaseSecretsRequest
	GetSecretsJson() *string
}

type SaveAppSupabaseSecretsRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// {"key":"abc"}
	SecretsJson *string `json:"SecretsJson,omitempty" xml:"SecretsJson,omitempty"`
}

func (s SaveAppSupabaseSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAppSupabaseSecretsRequest) GoString() string {
	return s.String()
}

func (s *SaveAppSupabaseSecretsRequest) GetBizId() *string {
	return s.BizId
}

func (s *SaveAppSupabaseSecretsRequest) GetSecretsJson() *string {
	return s.SecretsJson
}

func (s *SaveAppSupabaseSecretsRequest) SetBizId(v string) *SaveAppSupabaseSecretsRequest {
	s.BizId = &v
	return s
}

func (s *SaveAppSupabaseSecretsRequest) SetSecretsJson(v string) *SaveAppSupabaseSecretsRequest {
	s.SecretsJson = &v
	return s
}

func (s *SaveAppSupabaseSecretsRequest) Validate() error {
	return dara.Validate(s)
}
