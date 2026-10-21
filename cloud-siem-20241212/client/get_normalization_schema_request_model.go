// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetNormalizationSchemaRequest
	GetLang() *string
	SetNormalizationSchemaId(v string) *GetNormalizationSchemaRequest
	GetNormalizationSchemaId() *string
	SetNormalizationSchemaType(v string) *GetNormalizationSchemaRequest
	GetNormalizationSchemaType() *string
	SetRegionId(v string) *GetNormalizationSchemaRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetNormalizationSchemaRequest
	GetRoleFor() *int64
}

type GetNormalizationSchemaRequest struct {
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required in actual API calls. If you do not specify this parameter, the service returns a 400 error with the message "The specified Lang can not be empty." Explicitly specify the language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the normalized rule category.
	//
	// example:
	//
	// HTTP_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// The type of the normalized schema. Valid values:
	//
	// - log: Log.
	//
	// - entity: Entity. This parameter is required in actual API calls. If you do not specify this parameter, the service returns a 400 error.
	//
	// example:
	//
	// log
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// The region of the data management center for Threat Analysis. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetNormalizationSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaRequest) GetLang() *string {
	return s.Lang
}

func (s *GetNormalizationSchemaRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *GetNormalizationSchemaRequest) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *GetNormalizationSchemaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNormalizationSchemaRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetNormalizationSchemaRequest) SetLang(v string) *GetNormalizationSchemaRequest {
	s.Lang = &v
	return s
}

func (s *GetNormalizationSchemaRequest) SetNormalizationSchemaId(v string) *GetNormalizationSchemaRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *GetNormalizationSchemaRequest) SetNormalizationSchemaType(v string) *GetNormalizationSchemaRequest {
	s.NormalizationSchemaType = &v
	return s
}

func (s *GetNormalizationSchemaRequest) SetRegionId(v string) *GetNormalizationSchemaRequest {
	s.RegionId = &v
	return s
}

func (s *GetNormalizationSchemaRequest) SetRoleFor(v int64) *GetNormalizationSchemaRequest {
	s.RoleFor = &v
	return s
}

func (s *GetNormalizationSchemaRequest) Validate() error {
	return dara.Validate(s)
}
