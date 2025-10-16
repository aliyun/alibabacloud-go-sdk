// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSdlProtectedAssetRequest interface {
  dara.Model
  String() string
  GoString() string
  SetIpList(v []*string) *EnableSdlProtectedAssetRequest
  GetIpList() []*string 
  SetLang(v string) *EnableSdlProtectedAssetRequest
  GetLang() *string 
}

type EnableSdlProtectedAssetRequest struct {
  IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s EnableSdlProtectedAssetRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSdlProtectedAssetRequest) GoString() string {
  return s.String()
}

func (s *EnableSdlProtectedAssetRequest) GetIpList() []*string  {
  return s.IpList
}

func (s *EnableSdlProtectedAssetRequest) GetLang() *string  {
  return s.Lang
}

func (s *EnableSdlProtectedAssetRequest) SetIpList(v []*string) *EnableSdlProtectedAssetRequest {
  s.IpList = v
  return s
}

func (s *EnableSdlProtectedAssetRequest) SetLang(v string) *EnableSdlProtectedAssetRequest {
  s.Lang = &v
  return s
}

func (s *EnableSdlProtectedAssetRequest) Validate() error {
  return dara.Validate(s)
}

