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
	// This parameter is required.
	//
	// example:
	//
	// ****530646
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// c8016d13-6e76-****-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
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
