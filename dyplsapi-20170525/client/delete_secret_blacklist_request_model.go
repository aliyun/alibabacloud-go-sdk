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
	// The phone numbers in the blacklist. A point-to-point blacklist has a pair of numbers separated by a colon (":"). A number pool blacklist or a platform blacklist has only one single number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18252***383:18252***483
	BlackNo *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	// The blacklist type. Valid values:
	//
	// 	- **POINT_TO_POINT_BLACK**: point-to-point blacklist
	//
	// 	- **PARTNER_GLOBAL_NUMBER_BLACK**: number pool blacklist
	//
	// This parameter is required.
	//
	// example:
	//
	// POINT_TO_POINT_BLACK
	BlackType *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC1232****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The remarks for the blacklist.
	//
	// example:
	//
	// fragile
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The control on the call direction.
	//
	// 	- **PHONEA_REJECT**: The first phone number in the blacklist can be used to call the second phone number, but the second phone number in the blacklist cannot be used to call the first phone number.
	//
	// 	- **PHONEB_REJECT**: The first phone number in the blacklist cannot be used to call the second phone number, but the second phone number in the blacklist can be used to call the first phone number.
	//
	// 	- If this parameter is not specified, the two phone numbers in the blacklist cannot be used to call each other.
	//
	// Valid values:
	//
	// 	- DUPLEX_REJECT
	//
	// 	- PHONEA_REJECT
	//
	// 	- PHONEB_REJECT
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
