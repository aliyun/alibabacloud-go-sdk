// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserByMobileAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobileType(v string) *QueryUserByMobileAccountRequest
	GetMobileType() *string
	SetMobileUserId(v string) *QueryUserByMobileAccountRequest
	GetMobileUserId() *string
}

type QueryUserByMobileAccountRequest struct {
	// The bound mobile type.
	//
	// - DingTalk: ding
	//
	// - WeCom: corp_weixin
	//
	// - Lark: feishu
	//
	// This parameter is required.
	//
	// example:
	//
	// ding
	MobileType *string `json:"MobileType,omitempty" xml:"MobileType,omitempty"`
	// The bound mobile user ID.
	//
	// - DingTalk: The unionId of the DingTalk account.
	//
	// - WeCom: The userId of the WeCom account.
	//
	// - Lark: The userId of the Lark account.
	//
	// 	Notice: The mobileUserId must be obtained by calling the relevant API operations of DingTalk, WeCom, or Lark.
	//
	// This parameter is required.
	//
	// example:
	//
	// sasda
	MobileUserId *string `json:"MobileUserId,omitempty" xml:"MobileUserId,omitempty"`
}

func (s QueryUserByMobileAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserByMobileAccountRequest) GoString() string {
	return s.String()
}

func (s *QueryUserByMobileAccountRequest) GetMobileType() *string {
	return s.MobileType
}

func (s *QueryUserByMobileAccountRequest) GetMobileUserId() *string {
	return s.MobileUserId
}

func (s *QueryUserByMobileAccountRequest) SetMobileType(v string) *QueryUserByMobileAccountRequest {
	s.MobileType = &v
	return s
}

func (s *QueryUserByMobileAccountRequest) SetMobileUserId(v string) *QueryUserByMobileAccountRequest {
	s.MobileUserId = &v
	return s
}

func (s *QueryUserByMobileAccountRequest) Validate() error {
	return dara.Validate(s)
}
