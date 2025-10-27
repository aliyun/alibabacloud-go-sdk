// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMassPushV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *MassPushV2ShrinkRequest
	GetAppKey() *int64
	SetIdempotentToken(v string) *MassPushV2ShrinkRequest
	GetIdempotentToken() *string
	SetPushTasksShrink(v string) *MassPushV2ShrinkRequest
	GetPushTasksShrink() *string
}

type MassPushV2ShrinkRequest struct {
	// This parameter is required.
	AppKey          *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// This parameter is required.
	PushTasksShrink *string `json:"PushTasks,omitempty" xml:"PushTasks,omitempty"`
}

func (s MassPushV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MassPushV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *MassPushV2ShrinkRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *MassPushV2ShrinkRequest) GetIdempotentToken() *string {
	return s.IdempotentToken
}

func (s *MassPushV2ShrinkRequest) GetPushTasksShrink() *string {
	return s.PushTasksShrink
}

func (s *MassPushV2ShrinkRequest) SetAppKey(v int64) *MassPushV2ShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *MassPushV2ShrinkRequest) SetIdempotentToken(v string) *MassPushV2ShrinkRequest {
	s.IdempotentToken = &v
	return s
}

func (s *MassPushV2ShrinkRequest) SetPushTasksShrink(v string) *MassPushV2ShrinkRequest {
	s.PushTasksShrink = &v
	return s
}

func (s *MassPushV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
