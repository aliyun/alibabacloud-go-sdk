// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCorpTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppSecret(v string) *CorpTokenRequest
	GetAppSecret() *string
	SetCorpId(v string) *CorpTokenRequest
	GetCorpId() *string
	SetType(v int32) *CorpTokenRequest
	GetType() *int32
}

type CorpTokenRequest struct {
	// example:
	//
	// Z2FyYmE1YTZjMDAwTixJU1M5LnZlLXtMO3FKbiYqJSM
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open324dfsdafsgcxvxv
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CorpTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CorpTokenRequest) GoString() string {
	return s.String()
}

func (s *CorpTokenRequest) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CorpTokenRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *CorpTokenRequest) GetType() *int32 {
	return s.Type
}

func (s *CorpTokenRequest) SetAppSecret(v string) *CorpTokenRequest {
	s.AppSecret = &v
	return s
}

func (s *CorpTokenRequest) SetCorpId(v string) *CorpTokenRequest {
	s.CorpId = &v
	return s
}

func (s *CorpTokenRequest) SetType(v int32) *CorpTokenRequest {
	s.Type = &v
	return s
}

func (s *CorpTokenRequest) Validate() error {
	return dara.Validate(s)
}
