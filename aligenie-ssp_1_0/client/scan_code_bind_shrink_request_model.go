// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanCodeBindShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindReqShrink(v string) *ScanCodeBindShrinkRequest
	GetBindReqShrink() *string
	SetUserInfoShrink(v string) *ScanCodeBindShrinkRequest
	GetUserInfoShrink() *string
}

type ScanCodeBindShrinkRequest struct {
	// This parameter is required.
	BindReqShrink *string `json:"BindReq,omitempty" xml:"BindReq,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ScanCodeBindShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanCodeBindShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScanCodeBindShrinkRequest) GetBindReqShrink() *string {
	return s.BindReqShrink
}

func (s *ScanCodeBindShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ScanCodeBindShrinkRequest) SetBindReqShrink(v string) *ScanCodeBindShrinkRequest {
	s.BindReqShrink = &v
	return s
}

func (s *ScanCodeBindShrinkRequest) SetUserInfoShrink(v string) *ScanCodeBindShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ScanCodeBindShrinkRequest) Validate() error {
	return dara.Validate(s)
}
