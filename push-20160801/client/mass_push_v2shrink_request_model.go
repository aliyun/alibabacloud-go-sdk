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
	// AppKey value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267***
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// An idempotency token to prevent duplicate pushes caused by API retries. If you call this API with the same IdempotentToken within 15 minutes, only one push is sent. Subsequent calls return the result of the first successful push.
	//
	// > - The token must be a standard 36-character UUID in 8-4-4-4-12 format. Valid characters are hexadecimal digits 0–9 and a–f. Case does not matter.
	//
	// >
	//
	// > - This parameter prevents duplicates only from retries. It does not prevent duplicates from concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-****-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// Batch push tasks.
	//
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
