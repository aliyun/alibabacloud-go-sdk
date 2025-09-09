// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlrRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *CreateSlrRoleRequest
	GetFeatureType() *int32
	SetLang(v string) *CreateSlrRoleRequest
	GetLang() *string
	SetSourceIp(v string) *CreateSlrRoleRequest
	GetSourceIp() *string
}

type CreateSlrRoleRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateSlrRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *CreateSlrRoleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSlrRoleRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateSlrRoleRequest) SetFeatureType(v int32) *CreateSlrRoleRequest {
	s.FeatureType = &v
	return s
}

func (s *CreateSlrRoleRequest) SetLang(v string) *CreateSlrRoleRequest {
	s.Lang = &v
	return s
}

func (s *CreateSlrRoleRequest) SetSourceIp(v string) *CreateSlrRoleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateSlrRoleRequest) Validate() error {
	return dara.Validate(s)
}
