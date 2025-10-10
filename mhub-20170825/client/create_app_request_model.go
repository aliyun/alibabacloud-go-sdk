// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v string) *CreateAppRequest
	GetBundleId() *string
	SetEncodedIcon(v string) *CreateAppRequest
	GetEncodedIcon() *string
	SetIndustryId(v string) *CreateAppRequest
	GetIndustryId() *string
	SetName(v string) *CreateAppRequest
	GetName() *string
	SetPackageName(v string) *CreateAppRequest
	GetPackageName() *string
	SetProductId(v string) *CreateAppRequest
	GetProductId() *string
	SetType(v int32) *CreateAppRequest
	GetType() *int32
}

type CreateAppRequest struct {
	// example:
	//
	// com.test.ios
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *string `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateAppRequest) GetEncodedIcon() *string {
	return s.EncodedIcon
}

func (s *CreateAppRequest) GetIndustryId() *string {
	return s.IndustryId
}

func (s *CreateAppRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppRequest) GetPackageName() *string {
	return s.PackageName
}

func (s *CreateAppRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateAppRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateAppRequest) SetBundleId(v string) *CreateAppRequest {
	s.BundleId = &v
	return s
}

func (s *CreateAppRequest) SetEncodedIcon(v string) *CreateAppRequest {
	s.EncodedIcon = &v
	return s
}

func (s *CreateAppRequest) SetIndustryId(v string) *CreateAppRequest {
	s.IndustryId = &v
	return s
}

func (s *CreateAppRequest) SetName(v string) *CreateAppRequest {
	s.Name = &v
	return s
}

func (s *CreateAppRequest) SetPackageName(v string) *CreateAppRequest {
	s.PackageName = &v
	return s
}

func (s *CreateAppRequest) SetProductId(v string) *CreateAppRequest {
	s.ProductId = &v
	return s
}

func (s *CreateAppRequest) SetType(v int32) *CreateAppRequest {
	s.Type = &v
	return s
}

func (s *CreateAppRequest) Validate() error {
	return dara.Validate(s)
}
