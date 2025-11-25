// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSdlProtectedAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpList(v []*string) *DisableSdlProtectedAssetRequest
	GetIpList() []*string
	SetLang(v string) *DisableSdlProtectedAssetRequest
	GetLang() *string
}

type DisableSdlProtectedAssetRequest struct {
	IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DisableSdlProtectedAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSdlProtectedAssetRequest) GoString() string {
	return s.String()
}

func (s *DisableSdlProtectedAssetRequest) GetIpList() []*string {
	return s.IpList
}

func (s *DisableSdlProtectedAssetRequest) GetLang() *string {
	return s.Lang
}

func (s *DisableSdlProtectedAssetRequest) SetIpList(v []*string) *DisableSdlProtectedAssetRequest {
	s.IpList = v
	return s
}

func (s *DisableSdlProtectedAssetRequest) SetLang(v string) *DisableSdlProtectedAssetRequest {
	s.Lang = &v
	return s
}

func (s *DisableSdlProtectedAssetRequest) Validate() error {
	return dara.Validate(s)
}
