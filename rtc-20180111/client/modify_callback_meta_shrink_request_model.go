// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCallbackMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyCallbackMetaShrinkRequest
	GetAppId() *string
	SetCallbackShrink(v string) *ModifyCallbackMetaShrinkRequest
	GetCallbackShrink() *string
	SetOwnerId(v int64) *ModifyCallbackMetaShrinkRequest
	GetOwnerId() *int64
}

type ModifyCallbackMetaShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	CallbackShrink *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyCallbackMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCallbackMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCallbackMetaShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyCallbackMetaShrinkRequest) GetCallbackShrink() *string {
	return s.CallbackShrink
}

func (s *ModifyCallbackMetaShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCallbackMetaShrinkRequest) SetAppId(v string) *ModifyCallbackMetaShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyCallbackMetaShrinkRequest) SetCallbackShrink(v string) *ModifyCallbackMetaShrinkRequest {
	s.CallbackShrink = &v
	return s
}

func (s *ModifyCallbackMetaShrinkRequest) SetOwnerId(v int64) *ModifyCallbackMetaShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCallbackMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
