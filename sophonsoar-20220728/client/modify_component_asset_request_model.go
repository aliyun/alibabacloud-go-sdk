// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComponentAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetConfig(v string) *ModifyComponentAssetRequest
	GetAssetConfig() *string
	SetLang(v string) *ModifyComponentAssetRequest
	GetLang() *string
}

type ModifyComponentAssetRequest struct {
	// The configuration of the asset. The value is a JSON object.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "name": "test asset",
	//
	//     "componentName": "SLS",
	//
	//     "params": [
	//
	//         {
	//
	//             "name": "end_point",
	//
	//             "value": "xxx"
	//
	//         },
	//
	//         {
	//
	//             "name": "sub_id",
	//
	//             "value": "xxxx"
	//
	//         },
	//
	//         {
	//
	//             "name": "access_key",
	//
	//             "value": "xxxx"
	//
	//         }
	//
	//     ]
	//
	// }
	AssetConfig *string `json:"AssetConfig,omitempty" xml:"AssetConfig,omitempty"`
	// The language of the content within the request and response.
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ModifyComponentAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetRequest) GetAssetConfig() *string {
	return s.AssetConfig
}

func (s *ModifyComponentAssetRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyComponentAssetRequest) SetAssetConfig(v string) *ModifyComponentAssetRequest {
	s.AssetConfig = &v
	return s
}

func (s *ModifyComponentAssetRequest) SetLang(v string) *ModifyComponentAssetRequest {
	s.Lang = &v
	return s
}

func (s *ModifyComponentAssetRequest) Validate() error {
	return dara.Validate(s)
}
