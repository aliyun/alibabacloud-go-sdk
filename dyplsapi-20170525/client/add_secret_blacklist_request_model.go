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
	// The numbers in the blacklist. A point-to-point blacklist has a pair of numbers separated by a colon (:). A number pool blacklist has only one single number.
	//
	// >  The asterisks (\\*) in the sample value are not wildcards.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1825638****:1825248****
	BlackNo *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	// The blacklist type. Valid values:
	//
	// 	- **POINT_TO_POINT_BLACK**: point-to-point blacklist.
	//
	// 	- **PARTNER_GLOBAL_NUMBER_BLACK**: number pool blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// POINT_TO_POINT_BLACK
	BlackType *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	// The key of the phone number pool.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the Number Pool Management page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2235****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The blacklist remarks.
	//
	// example:
	//
	// Customer service feedback
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The control on the call direction.
	//
	// 	- **PHONEA_REJECT**: The first number in the blacklist can be used to call the second number, but the second number cannot be used to call the first number.
	//
	// 	- **PHONEB_REJECT**: The first number in the blacklist cannot be used to call the second number, but the second number can be used to call the first number.
	//
	// 	- If this parameter is left empty, the two numbers in the blacklist cannot be used to call each other.
	//
	// >  This parameter is available only for a point-to-point blacklist.
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
