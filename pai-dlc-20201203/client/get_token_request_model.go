// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v int64) *GetTokenRequest
	GetExpireTime() *int64
	SetTargetId(v string) *GetTokenRequest
	GetTargetId() *string
	SetTargetType(v string) *GetTokenRequest
	GetTargetType() *string
	SetTokenSettings(v *TokenSettings) *GetTokenRequest
	GetTokenSettings() *TokenSettings
}

type GetTokenRequest struct {
	// The expiration time of the sharing link. Default value: 604800 seconds. Minimum value: 0.
	//
	// example:
	//
	// 60
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the job to be shared.
	//
	// example:
	//
	// dlc*******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the shared job. Valid values: job and tensorboard.
	//
	// example:
	//
	// job
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The permission settings of the sharing token, which controls log download and cross-account access.
	TokenSettings *TokenSettings `json:"TokenSettings,omitempty" xml:"TokenSettings,omitempty"`
}

func (s GetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetTokenRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *GetTokenRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *GetTokenRequest) GetTokenSettings() *TokenSettings {
	return s.TokenSettings
}

func (s *GetTokenRequest) SetExpireTime(v int64) *GetTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenRequest) SetTargetId(v string) *GetTokenRequest {
	s.TargetId = &v
	return s
}

func (s *GetTokenRequest) SetTargetType(v string) *GetTokenRequest {
	s.TargetType = &v
	return s
}

func (s *GetTokenRequest) SetTokenSettings(v *TokenSettings) *GetTokenRequest {
	s.TokenSettings = v
	return s
}

func (s *GetTokenRequest) Validate() error {
	if s.TokenSettings != nil {
		if err := s.TokenSettings.Validate(); err != nil {
			return err
		}
	}
	return nil
}
