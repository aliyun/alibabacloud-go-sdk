// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *ModifyAppRequest
	GetAppKey() *string
	SetBundleId(v string) *ModifyAppRequest
	GetBundleId() *string
	SetEncodedIcon(v string) *ModifyAppRequest
	GetEncodedIcon() *string
	SetIndustryId(v int32) *ModifyAppRequest
	GetIndustryId() *int32
	SetName(v string) *ModifyAppRequest
	GetName() *string
	SetPackageName(v string) *ModifyAppRequest
	GetPackageName() *string
}

type ModifyAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// com.test.ios
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32  `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *ModifyAppRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *ModifyAppRequest) GetEncodedIcon() *string {
	return s.EncodedIcon
}

func (s *ModifyAppRequest) GetIndustryId() *int32 {
	return s.IndustryId
}

func (s *ModifyAppRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAppRequest) GetPackageName() *string {
	return s.PackageName
}

func (s *ModifyAppRequest) SetAppKey(v string) *ModifyAppRequest {
	s.AppKey = &v
	return s
}

func (s *ModifyAppRequest) SetBundleId(v string) *ModifyAppRequest {
	s.BundleId = &v
	return s
}

func (s *ModifyAppRequest) SetEncodedIcon(v string) *ModifyAppRequest {
	s.EncodedIcon = &v
	return s
}

func (s *ModifyAppRequest) SetIndustryId(v int32) *ModifyAppRequest {
	s.IndustryId = &v
	return s
}

func (s *ModifyAppRequest) SetName(v string) *ModifyAppRequest {
	s.Name = &v
	return s
}

func (s *ModifyAppRequest) SetPackageName(v string) *ModifyAppRequest {
	s.PackageName = &v
	return s
}

func (s *ModifyAppRequest) Validate() error {
	return dara.Validate(s)
}
