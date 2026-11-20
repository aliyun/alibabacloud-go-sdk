// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretBlacklistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlackNo(v string) *DeleteSecretBlacklistRequest
	GetBlackNo() *string
	SetBlackType(v string) *DeleteSecretBlacklistRequest
	GetBlackType() *string
	SetPoolKey(v string) *DeleteSecretBlacklistRequest
	GetPoolKey() *string
	SetRemark(v string) *DeleteSecretBlacklistRequest
	GetRemark() *string
	SetWayControl(v string) *DeleteSecretBlacklistRequest
	GetWayControl() *string
}

type DeleteSecretBlacklistRequest struct {
	// The blacklisted number. The value type is determined by the **BlackType*	- parameter:
	//
	// - **Point-to-point blacklist**: a pair of numbers separated by a colon (:)
	//
	// - **Number pool blacklist**: a single number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1825****383:1825****483
	BlackNo *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	// The type of the blacklist. Valid values:
	//
	// - **POINT_TO_POINT_BLACK**: point-to-point blacklist
	//
	// - **PARTNER_GLOBAL_NUMBER_BLACK**: number pool blacklist
	//
	// This parameter is required.
	//
	// example:
	//
	// POINT_TO_POINT_BLACK
	BlackType *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	// The key of the number pool.
	//
	// You can obtain the **number pool key*	- on the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool) page in the Phone Number Privacy Protection console.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC1232****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The remark for the blacklist.
	//
	// example:
	//
	// Blacklist remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The call direction control. Valid values:
	//
	// - **PHONEA_REJECT**: The first number in BlackNo can call the second number, but the second number cannot call the first number.
	//
	// - **PHONEB_REJECT**: The first number in BlackNo cannot call the second number, but the second number can call the first number.
	//
	// - **Leave empty**: The two numbers in BlackNo cannot call each other.
	//
	// example:
	//
	// PHONEA_REJECT
	WayControl *string `json:"WayControl,omitempty" xml:"WayControl,omitempty"`
}

func (s DeleteSecretBlacklistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretBlacklistRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistRequest) GetBlackNo() *string {
	return s.BlackNo
}

func (s *DeleteSecretBlacklistRequest) GetBlackType() *string {
	return s.BlackType
}

func (s *DeleteSecretBlacklistRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *DeleteSecretBlacklistRequest) GetRemark() *string {
	return s.Remark
}

func (s *DeleteSecretBlacklistRequest) GetWayControl() *string {
	return s.WayControl
}

func (s *DeleteSecretBlacklistRequest) SetBlackNo(v string) *DeleteSecretBlacklistRequest {
	s.BlackNo = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetBlackType(v string) *DeleteSecretBlacklistRequest {
	s.BlackType = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetPoolKey(v string) *DeleteSecretBlacklistRequest {
	s.PoolKey = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetRemark(v string) *DeleteSecretBlacklistRequest {
	s.Remark = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) SetWayControl(v string) *DeleteSecretBlacklistRequest {
	s.WayControl = &v
	return s
}

func (s *DeleteSecretBlacklistRequest) Validate() error {
	return dara.Validate(s)
}
