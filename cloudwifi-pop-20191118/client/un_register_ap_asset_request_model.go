// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnRegisterApAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *UnRegisterApAssetRequest
	GetAppCode() *string
	SetAppName(v string) *UnRegisterApAssetRequest
	GetAppName() *string
	SetAssetApgroupId(v int64) *UnRegisterApAssetRequest
	GetAssetApgroupId() *int64
	SetCategory(v int32) *UnRegisterApAssetRequest
	GetCategory() *int32
	SetId(v int64) *UnRegisterApAssetRequest
	GetId() *int64
	SetMac(v string) *UnRegisterApAssetRequest
	GetMac() *string
	SetSerialNo(v string) *UnRegisterApAssetRequest
	GetSerialNo() *string
	SetUseFor(v int32) *UnRegisterApAssetRequest
	GetUseFor() *int32
}

type UnRegisterApAssetRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AssetApgroupId *int64  `json:"AssetApgroupId,omitempty" xml:"AssetApgroupId,omitempty"`
	Category       *int32  `json:"Category,omitempty" xml:"Category,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Mac      *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	UseFor   *int32  `json:"UseFor,omitempty" xml:"UseFor,omitempty"`
}

func (s UnRegisterApAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s UnRegisterApAssetRequest) GoString() string {
	return s.String()
}

func (s *UnRegisterApAssetRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *UnRegisterApAssetRequest) GetAppName() *string {
	return s.AppName
}

func (s *UnRegisterApAssetRequest) GetAssetApgroupId() *int64 {
	return s.AssetApgroupId
}

func (s *UnRegisterApAssetRequest) GetCategory() *int32 {
	return s.Category
}

func (s *UnRegisterApAssetRequest) GetId() *int64 {
	return s.Id
}

func (s *UnRegisterApAssetRequest) GetMac() *string {
	return s.Mac
}

func (s *UnRegisterApAssetRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *UnRegisterApAssetRequest) GetUseFor() *int32 {
	return s.UseFor
}

func (s *UnRegisterApAssetRequest) SetAppCode(v string) *UnRegisterApAssetRequest {
	s.AppCode = &v
	return s
}

func (s *UnRegisterApAssetRequest) SetAppName(v string) *UnRegisterApAssetRequest {
	s.AppName = &v
	return s
}

func (s *UnRegisterApAssetRequest) SetAssetApgroupId(v int64) *UnRegisterApAssetRequest {
	s.AssetApgroupId = &v
	return s
}

func (s *UnRegisterApAssetRequest) SetCategory(v int32) *UnRegisterApAssetRequest {
	s.Category = &v
	return s
}

func (s *UnRegisterApAssetRequest) SetId(v int64) *UnRegisterApAssetRequest {
	s.Id = &v
	return s
}

func (s *UnRegisterApAssetRequest) SetMac(v string) *UnRegisterApAssetRequest {
	s.Mac = &v
	return s
}

func (s *UnRegisterApAssetRequest) SetSerialNo(v string) *UnRegisterApAssetRequest {
	s.SerialNo = &v
	return s
}

func (s *UnRegisterApAssetRequest) SetUseFor(v int32) *UnRegisterApAssetRequest {
	s.UseFor = &v
	return s
}

func (s *UnRegisterApAssetRequest) Validate() error {
	return dara.Validate(s)
}
