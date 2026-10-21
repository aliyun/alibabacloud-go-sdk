// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectionStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetDetectionStatisticRequest
	GetLang() *string
	SetRegionId(v string) *GetDetectionStatisticRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetDetectionStatisticRequest
	GetRoleFor() *int64
}

type GetDetectionStatisticRequest struct {
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required in actual API calls. If it is not provided, the service returns a 400 error with the message "The specified Lang can not be empty.". Explicitly specify the language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region of the data management center for threat analysis. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose view the administrator assumes.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetDetectionStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDetectionStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetDetectionStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *GetDetectionStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDetectionStatisticRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetDetectionStatisticRequest) SetLang(v string) *GetDetectionStatisticRequest {
	s.Lang = &v
	return s
}

func (s *GetDetectionStatisticRequest) SetRegionId(v string) *GetDetectionStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *GetDetectionStatisticRequest) SetRoleFor(v int64) *GetDetectionStatisticRequest {
	s.RoleFor = &v
	return s
}

func (s *GetDetectionStatisticRequest) Validate() error {
	return dara.Validate(s)
}
