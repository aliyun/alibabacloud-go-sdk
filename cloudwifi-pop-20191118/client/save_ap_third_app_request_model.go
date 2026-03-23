// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApThirdAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddStyle(v int32) *SaveApThirdAppRequest
	GetAddStyle() *int32
	SetApAssetId(v int64) *SaveApThirdAppRequest
	GetApAssetId() *int64
	SetAppCode(v string) *SaveApThirdAppRequest
	GetAppCode() *string
	SetAppData(v string) *SaveApThirdAppRequest
	GetAppData() *string
	SetAppName(v string) *SaveApThirdAppRequest
	GetAppName() *string
	SetCategory(v int32) *SaveApThirdAppRequest
	GetCategory() *int32
	SetId(v int64) *SaveApThirdAppRequest
	GetId() *int64
	SetMac(v string) *SaveApThirdAppRequest
	GetMac() *string
	SetThirdAppName(v string) *SaveApThirdAppRequest
	GetThirdAppName() *string
	SetVersion(v string) *SaveApThirdAppRequest
	GetVersion() *string
}

type SaveApThirdAppRequest struct {
	AddStyle  *int32 `json:"AddStyle,omitempty" xml:"AddStyle,omitempty"`
	ApAssetId *int64 `json:"ApAssetId,omitempty" xml:"ApAssetId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppData *string `json:"AppData,omitempty" xml:"AppData,omitempty"`
	// This parameter is required.
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Category *int32  `json:"Category,omitempty" xml:"Category,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Mac          *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	ThirdAppName *string `json:"ThirdAppName,omitempty" xml:"ThirdAppName,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SaveApThirdAppRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApThirdAppRequest) GoString() string {
	return s.String()
}

func (s *SaveApThirdAppRequest) GetAddStyle() *int32 {
	return s.AddStyle
}

func (s *SaveApThirdAppRequest) GetApAssetId() *int64 {
	return s.ApAssetId
}

func (s *SaveApThirdAppRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApThirdAppRequest) GetAppData() *string {
	return s.AppData
}

func (s *SaveApThirdAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApThirdAppRequest) GetCategory() *int32 {
	return s.Category
}

func (s *SaveApThirdAppRequest) GetId() *int64 {
	return s.Id
}

func (s *SaveApThirdAppRequest) GetMac() *string {
	return s.Mac
}

func (s *SaveApThirdAppRequest) GetThirdAppName() *string {
	return s.ThirdAppName
}

func (s *SaveApThirdAppRequest) GetVersion() *string {
	return s.Version
}

func (s *SaveApThirdAppRequest) SetAddStyle(v int32) *SaveApThirdAppRequest {
	s.AddStyle = &v
	return s
}

func (s *SaveApThirdAppRequest) SetApAssetId(v int64) *SaveApThirdAppRequest {
	s.ApAssetId = &v
	return s
}

func (s *SaveApThirdAppRequest) SetAppCode(v string) *SaveApThirdAppRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApThirdAppRequest) SetAppData(v string) *SaveApThirdAppRequest {
	s.AppData = &v
	return s
}

func (s *SaveApThirdAppRequest) SetAppName(v string) *SaveApThirdAppRequest {
	s.AppName = &v
	return s
}

func (s *SaveApThirdAppRequest) SetCategory(v int32) *SaveApThirdAppRequest {
	s.Category = &v
	return s
}

func (s *SaveApThirdAppRequest) SetId(v int64) *SaveApThirdAppRequest {
	s.Id = &v
	return s
}

func (s *SaveApThirdAppRequest) SetMac(v string) *SaveApThirdAppRequest {
	s.Mac = &v
	return s
}

func (s *SaveApThirdAppRequest) SetThirdAppName(v string) *SaveApThirdAppRequest {
	s.ThirdAppName = &v
	return s
}

func (s *SaveApThirdAppRequest) SetVersion(v string) *SaveApThirdAppRequest {
	s.Version = &v
	return s
}

func (s *SaveApThirdAppRequest) Validate() error {
	return dara.Validate(s)
}
