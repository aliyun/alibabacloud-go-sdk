// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *PushV2ShrinkRequest
	GetAppKey() *int64
	SetIdempotentToken(v string) *PushV2ShrinkRequest
	GetIdempotentToken() *string
	SetPushTaskShrink(v string) *PushV2ShrinkRequest
	GetPushTaskShrink() *string
}

type PushV2ShrinkRequest struct {
	// AppKey value.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****530646
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// An idempotency token to prevent duplicate pushes caused by client-side retries. If you call this API with the same IdempotentToken within 15 minutes, only one push is sent. Subsequent calls return the result of the first successful push.
	//
	// > - Format the token as a standard 36-character UUID (8-4-4-4-12). Valid characters are hexadecimal digits 0–9 and a–f. Case-insensitive.
	//
	// >
	//
	// > - This parameter prevents duplicates only from retries. It does not prevent duplicates from concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-****-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// Push task definition.
	//
	// This parameter is required.
	PushTaskShrink *string `json:"PushTask,omitempty" xml:"PushTask,omitempty"`
}

func (s PushV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushV2ShrinkRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *PushV2ShrinkRequest) GetIdempotentToken() *string {
	return s.IdempotentToken
}

func (s *PushV2ShrinkRequest) GetPushTaskShrink() *string {
	return s.PushTaskShrink
}

func (s *PushV2ShrinkRequest) SetAppKey(v int64) *PushV2ShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *PushV2ShrinkRequest) SetIdempotentToken(v string) *PushV2ShrinkRequest {
	s.IdempotentToken = &v
	return s
}

func (s *PushV2ShrinkRequest) SetPushTaskShrink(v string) *PushV2ShrinkRequest {
	s.PushTaskShrink = &v
	return s
}

func (s *PushV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
