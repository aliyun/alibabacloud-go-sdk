// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultId(v int64) *ModifyDefaultLevelRequest
	GetDefaultId() *int64
	SetLang(v string) *ModifyDefaultLevelRequest
	GetLang() *string
	SetSensitiveIds(v string) *ModifyDefaultLevelRequest
	GetSensitiveIds() *string
}

type ModifyDefaultLevelRequest struct {
	// The default sensitivity level of data that Data Security Center (DSC) cannot classify as sensitive or insensitive. Valid values:
	//
	// 	- **1**: N/A
	//
	// 	- **2**: S1
	//
	// 	- **3**: S2
	//
	// 	- **4**: S3
	//
	// 	- **5**: S4
	//
	// example:
	//
	// 4
	DefaultId *int64 `json:"DefaultId,omitempty" xml:"DefaultId,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sensitivity level ID of data that DSC classifies as sensitive. Separate multiple IDs with commas (,). Valid values:
	//
	// 	- **1**: N/A
	//
	// 	- **2**: S1
	//
	// 	- **3**: S2
	//
	// 	- **4**: S3
	//
	// 	- **5**: S4
	//
	// example:
	//
	// 1,2,3,4
	SensitiveIds *string `json:"SensitiveIds,omitempty" xml:"SensitiveIds,omitempty"`
}

func (s ModifyDefaultLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultLevelRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefaultLevelRequest) GetDefaultId() *int64 {
	return s.DefaultId
}

func (s *ModifyDefaultLevelRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyDefaultLevelRequest) GetSensitiveIds() *string {
	return s.SensitiveIds
}

func (s *ModifyDefaultLevelRequest) SetDefaultId(v int64) *ModifyDefaultLevelRequest {
	s.DefaultId = &v
	return s
}

func (s *ModifyDefaultLevelRequest) SetLang(v string) *ModifyDefaultLevelRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDefaultLevelRequest) SetSensitiveIds(v string) *ModifyDefaultLevelRequest {
	s.SensitiveIds = &v
	return s
}

func (s *ModifyDefaultLevelRequest) Validate() error {
	return dara.Validate(s)
}
