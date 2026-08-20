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
	// The AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****530646
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// An idempotence parameter that prevents duplicate push notifications caused by retries from the API caller. If you use the same IdempotentToken within 15 minutes, only one push notification is sent, and subsequent calls return the result of the first successful push.
	//
	// >
	//
	// > - The parameter format is a standard 36-character UUID (8-4-4-4-12). Each valid character is a hexadecimal digit in the range 0-9 or a-f, case-insensitive.
	//
	// > - This parameter only prevents duplicate push notifications caused by retries. It cannot prevent duplicate push notifications caused by concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-****-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// The push task.
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
