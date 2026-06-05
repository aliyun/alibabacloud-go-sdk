// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSupabaseAuthConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UpdateAppSupabaseAuthConfigRequest
	GetBizId() *string
	SetConfigsJson(v string) *UpdateAppSupabaseAuthConfigRequest
	GetConfigsJson() *string
}

type UpdateAppSupabaseAuthConfigRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// {"key":"value"}
	ConfigsJson *string `json:"ConfigsJson,omitempty" xml:"ConfigsJson,omitempty"`
}

func (s UpdateAppSupabaseAuthConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSupabaseAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppSupabaseAuthConfigRequest) GetBizId() *string {
	return s.BizId
}

func (s *UpdateAppSupabaseAuthConfigRequest) GetConfigsJson() *string {
	return s.ConfigsJson
}

func (s *UpdateAppSupabaseAuthConfigRequest) SetBizId(v string) *UpdateAppSupabaseAuthConfigRequest {
	s.BizId = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigRequest) SetConfigsJson(v string) *UpdateAppSupabaseAuthConfigRequest {
	s.ConfigsJson = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigRequest) Validate() error {
	return dara.Validate(s)
}
