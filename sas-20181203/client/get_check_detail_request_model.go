// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *GetCheckDetailRequest
	GetCheckId() *int64
	SetLang(v string) *GetCheckDetailRequest
	GetLang() *string
	SetRegionId(v string) *GetCheckDetailRequest
	GetRegionId() *string
}

type GetCheckDetailRequest struct {
	// The ID of the check item.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to query the IDs of check items.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [ListCloudAssetInstances](~~ListCloudAssetInstances~~) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCheckDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCheckDetailRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *GetCheckDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCheckDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCheckDetailRequest) SetCheckId(v int64) *GetCheckDetailRequest {
	s.CheckId = &v
	return s
}

func (s *GetCheckDetailRequest) SetLang(v string) *GetCheckDetailRequest {
	s.Lang = &v
	return s
}

func (s *GetCheckDetailRequest) SetRegionId(v string) *GetCheckDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetCheckDetailRequest) Validate() error {
	return dara.Validate(s)
}
