// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSecretBlacklistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlackNo(v string) *AddSecretBlacklistRequest
	GetBlackNo() *string
	SetBlackType(v string) *AddSecretBlacklistRequest
	GetBlackType() *string
	SetPoolKey(v string) *AddSecretBlacklistRequest
	GetPoolKey() *string
	SetRemark(v string) *AddSecretBlacklistRequest
	GetRemark() *string
	SetWayControl(v string) *AddSecretBlacklistRequest
	GetWayControl() *string
}

type AddSecretBlacklistRequest struct {
	// The blacklisted phone number. The value type depends on the **BlackType*	- parameter:
	//
	// - **Point-to-point blacklist**: A pair of numbers (A and B numbers) separated by a colon (:).
	//
	// - **Number pool blacklist**: A single number.
	//
	// >(\\*) is not a wildcard.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1825638****:1825248****
	BlackNo *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	// The blacklist type. Valid values:
	//
	// - **POINT_TO_POINT_BLACK**: point-to-point blacklist.
	//
	// - **PARTNER_GLOBAL_NUMBER_BLACK**: number pool blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// POINT_TO_POINT_BLACK
	BlackType *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	// The key of the number pool.
	//
	// You can obtain the **number pool key*	- on the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool) page in the privacy protection console.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2235****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The remarks for the blacklist entry.
	//
	// example:
	//
	// Blacklist remarks
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The call direction control. Valid values:
	//
	// - **PHONEA_REJECT**: The first number in BlackNo can call the second number, but the second number cannot call the first number.
	//
	// - **PHONEB_REJECT**: The first number in BlackNo cannot call the second number, but the second number can call the first number.
	//
	// - **Not specified**: The two numbers in BlackNo cannot call each other.
	//
	// >This parameter takes effect only for point-to-point blacklists.
	//
	// example:
	//
	// PHONEA_REJECT
	WayControl *string `json:"WayControl,omitempty" xml:"WayControl,omitempty"`
}

func (s AddSecretBlacklistRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSecretBlacklistRequest) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistRequest) GetBlackNo() *string {
	return s.BlackNo
}

func (s *AddSecretBlacklistRequest) GetBlackType() *string {
	return s.BlackType
}

func (s *AddSecretBlacklistRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *AddSecretBlacklistRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddSecretBlacklistRequest) GetWayControl() *string {
	return s.WayControl
}

func (s *AddSecretBlacklistRequest) SetBlackNo(v string) *AddSecretBlacklistRequest {
	s.BlackNo = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetBlackType(v string) *AddSecretBlacklistRequest {
	s.BlackType = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetPoolKey(v string) *AddSecretBlacklistRequest {
	s.PoolKey = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetRemark(v string) *AddSecretBlacklistRequest {
	s.Remark = &v
	return s
}

func (s *AddSecretBlacklistRequest) SetWayControl(v string) *AddSecretBlacklistRequest {
	s.WayControl = &v
	return s
}

func (s *AddSecretBlacklistRequest) Validate() error {
	return dara.Validate(s)
}
