// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVodPackagingAsset interface {
	dara.Model
	String() string
	GoString() string
	SetAssetName(v string) *VodPackagingAsset
	GetAssetName() *string
	SetContentId(v string) *VodPackagingAsset
	GetContentId() *string
	SetCreateTime(v string) *VodPackagingAsset
	GetCreateTime() *string
	SetGroupName(v string) *VodPackagingAsset
	GetGroupName() *string
	SetInput(v *VodPackagingAssetInput) *VodPackagingAsset
	GetInput() *VodPackagingAssetInput
}

type VodPackagingAsset struct {
	AssetName  *string                 `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	ContentId  *string                 `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	CreateTime *string                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GroupName  *string                 `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Input      *VodPackagingAssetInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s VodPackagingAsset) String() string {
	return dara.Prettify(s)
}

func (s VodPackagingAsset) GoString() string {
	return s.String()
}

func (s *VodPackagingAsset) GetAssetName() *string {
	return s.AssetName
}

func (s *VodPackagingAsset) GetContentId() *string {
	return s.ContentId
}

func (s *VodPackagingAsset) GetCreateTime() *string {
	return s.CreateTime
}

func (s *VodPackagingAsset) GetGroupName() *string {
	return s.GroupName
}

func (s *VodPackagingAsset) GetInput() *VodPackagingAssetInput {
	return s.Input
}

func (s *VodPackagingAsset) SetAssetName(v string) *VodPackagingAsset {
	s.AssetName = &v
	return s
}

func (s *VodPackagingAsset) SetContentId(v string) *VodPackagingAsset {
	s.ContentId = &v
	return s
}

func (s *VodPackagingAsset) SetCreateTime(v string) *VodPackagingAsset {
	s.CreateTime = &v
	return s
}

func (s *VodPackagingAsset) SetGroupName(v string) *VodPackagingAsset {
	s.GroupName = &v
	return s
}

func (s *VodPackagingAsset) SetInput(v *VodPackagingAssetInput) *VodPackagingAsset {
	s.Input = v
	return s
}

func (s *VodPackagingAsset) Validate() error {
	return dara.Validate(s)
}

type VodPackagingAssetInput struct {
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s VodPackagingAssetInput) String() string {
	return dara.Prettify(s)
}

func (s VodPackagingAssetInput) GoString() string {
	return s.String()
}

func (s *VodPackagingAssetInput) GetMedia() *string {
	return s.Media
}

func (s *VodPackagingAssetInput) GetType() *string {
	return s.Type
}

func (s *VodPackagingAssetInput) SetMedia(v string) *VodPackagingAssetInput {
	s.Media = &v
	return s
}

func (s *VodPackagingAssetInput) SetType(v string) *VodPackagingAssetInput {
	s.Type = &v
	return s
}

func (s *VodPackagingAssetInput) Validate() error {
	return dara.Validate(s)
}
