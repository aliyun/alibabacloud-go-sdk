// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupOperateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *HangupOperateRequest
	GetCallId() *string
	SetImmediateHangup(v bool) *HangupOperateRequest
	GetImmediateHangup() *bool
}

type HangupOperateRequest struct {
	// The call ID. Obtain this ID from the response of the [LlmSmartCall](https://help.aliyun.com/document_detail/2862828.html) or [LlmSmartCallEncrypt](https://help.aliyun.com/document_detail/2881065.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14777672****^13452272****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Specifies whether to hang up the call immediately. Valid values:
	//
	// - true (default): Hangs up the call immediately.
	//
	// - false: Hangs up the call after the current playback finishes.
	//
	// example:
	//
	// true
	ImmediateHangup *bool `json:"ImmediateHangup,omitempty" xml:"ImmediateHangup,omitempty"`
}

func (s HangupOperateRequest) String() string {
	return dara.Prettify(s)
}

func (s HangupOperateRequest) GoString() string {
	return s.String()
}

func (s *HangupOperateRequest) GetCallId() *string {
	return s.CallId
}

func (s *HangupOperateRequest) GetImmediateHangup() *bool {
	return s.ImmediateHangup
}

func (s *HangupOperateRequest) SetCallId(v string) *HangupOperateRequest {
	s.CallId = &v
	return s
}

func (s *HangupOperateRequest) SetImmediateHangup(v bool) *HangupOperateRequest {
	s.ImmediateHangup = &v
	return s
}

func (s *HangupOperateRequest) Validate() error {
	return dara.Validate(s)
}
