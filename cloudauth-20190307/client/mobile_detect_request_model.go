// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobiles(v string) *MobileDetectRequest
	GetMobiles() *string
	SetParamType(v string) *MobileDetectRequest
	GetParamType() *string
}

type MobileDetectRequest struct {
	// List of phone numbers.
	//
	// example:
	//
	// 19833232569
	Mobiles *string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty"`
	// Encryption method:
	//
	// - normal: plaintext, no encryption
	//
	// - md5: MD5 encryption
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s MobileDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s MobileDetectRequest) GoString() string {
	return s.String()
}

func (s *MobileDetectRequest) GetMobiles() *string {
	return s.Mobiles
}

func (s *MobileDetectRequest) GetParamType() *string {
	return s.ParamType
}

func (s *MobileDetectRequest) SetMobiles(v string) *MobileDetectRequest {
	s.Mobiles = &v
	return s
}

func (s *MobileDetectRequest) SetParamType(v string) *MobileDetectRequest {
	s.ParamType = &v
	return s
}

func (s *MobileDetectRequest) Validate() error {
	return dara.Validate(s)
}
