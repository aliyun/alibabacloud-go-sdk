// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSdlProtectedAssetResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSdlProtectedAssetResponseBody
  GetRequestId() *string 
}

type EnableSdlProtectedAssetResponseBody struct {
  // example:
  // 
  // 15FCCC52-1E23-57AE-B5EF-3E00A3******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSdlProtectedAssetResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSdlProtectedAssetResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSdlProtectedAssetResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSdlProtectedAssetResponseBody) SetRequestId(v string) *EnableSdlProtectedAssetResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSdlProtectedAssetResponseBody) Validate() error {
  return dara.Validate(s)
}

