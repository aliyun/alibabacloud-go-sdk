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
	// The AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267***
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// An idempotent parameter used to prevent duplicate pushes caused by API call retries on the client side. If the same IdempotentToken is used for calls within 15 minutes, only one push is performed, and subsequent calls return the result of the first successful push.
	//
	// >
	//
	// > - The parameter format is a standard 36-character UUID (8-4-4-4-12). Each valid character is a hexadecimal digit in the range 0-9 or a-f, case-insensitive.
	//
	// > - This parameter only prevents duplicate pushes caused by retries. It cannot prevent duplicate pushes caused by concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-****-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// The batch push tasks.
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
