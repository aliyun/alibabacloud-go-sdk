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
	// This parameter is required.
	//
	// example:
	//
	// 147776727911^134522727911
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
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
