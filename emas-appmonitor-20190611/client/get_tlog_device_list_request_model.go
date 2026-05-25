// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogDeviceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetTlogDeviceListRequest
	GetAppKey() *int64
	SetAppVersion(v string) *GetTlogDeviceListRequest
	GetAppVersion() *string
	SetBeginDate(v int64) *GetTlogDeviceListRequest
	GetBeginDate() *int64
	SetBrand(v string) *GetTlogDeviceListRequest
	GetBrand() *string
	SetCity(v string) *GetTlogDeviceListRequest
	GetCity() *string
	SetCreateBeginDate(v int64) *GetTlogDeviceListRequest
	GetCreateBeginDate() *int64
	SetCreateEndDate(v int64) *GetTlogDeviceListRequest
	GetCreateEndDate() *int64
	SetEndDate(v int64) *GetTlogDeviceListRequest
	GetEndDate() *int64
	SetKeyword(v string) *GetTlogDeviceListRequest
	GetKeyword() *string
	SetModel(v string) *GetTlogDeviceListRequest
	GetModel() *string
	SetOs(v string) *GetTlogDeviceListRequest
	GetOs() *string
	SetOsVersion(v string) *GetTlogDeviceListRequest
	GetOsVersion() *string
	SetPageIndex(v int32) *GetTlogDeviceListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetTlogDeviceListRequest
	GetPageSize() *int32
	SetUpdateBeginDate(v int64) *GetTlogDeviceListRequest
	GetUpdateBeginDate() *int64
	SetUpdateEndDate(v int64) *GetTlogDeviceListRequest
	GetUpdateEndDate() *int64
	SetUserNick(v string) *GetTlogDeviceListRequest
	GetUserNick() *string
	SetUtdid(v string) *GetTlogDeviceListRequest
	GetUtdid() *string
}

type GetTlogDeviceListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 24781204
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 1778860800000
	BeginDate *int64 `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// apple
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	City  *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 1633159200000
	CreateBeginDate *int64 `json:"CreateBeginDate,omitempty" xml:"CreateBeginDate,omitempty"`
	// example:
	//
	// 1633159200000
	CreateEndDate *int64 `json:"CreateEndDate,omitempty" xml:"CreateEndDate,omitempty"`
	// example:
	//
	// 1779465599999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// Z70g612312124323Hzn4X
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// iphone16
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1.0
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1633159200000
	UpdateBeginDate *int64 `json:"UpdateBeginDate,omitempty" xml:"UpdateBeginDate,omitempty"`
	// example:
	//
	// 1633159200000
	UpdateEndDate *int64 `json:"UpdateEndDate,omitempty" xml:"UpdateEndDate,omitempty"`
	// example:
	//
	// userNick
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	// example:
	//
	// Z70g612312124323Hzn4X
	Utdid *string `json:"Utdid,omitempty" xml:"Utdid,omitempty"`
}

func (s GetTlogDeviceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTlogDeviceListRequest) GoString() string {
	return s.String()
}

func (s *GetTlogDeviceListRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetTlogDeviceListRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetTlogDeviceListRequest) GetBeginDate() *int64 {
	return s.BeginDate
}

func (s *GetTlogDeviceListRequest) GetBrand() *string {
	return s.Brand
}

func (s *GetTlogDeviceListRequest) GetCity() *string {
	return s.City
}

func (s *GetTlogDeviceListRequest) GetCreateBeginDate() *int64 {
	return s.CreateBeginDate
}

func (s *GetTlogDeviceListRequest) GetCreateEndDate() *int64 {
	return s.CreateEndDate
}

func (s *GetTlogDeviceListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetTlogDeviceListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetTlogDeviceListRequest) GetModel() *string {
	return s.Model
}

func (s *GetTlogDeviceListRequest) GetOs() *string {
	return s.Os
}

func (s *GetTlogDeviceListRequest) GetOsVersion() *string {
	return s.OsVersion
}

func (s *GetTlogDeviceListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetTlogDeviceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTlogDeviceListRequest) GetUpdateBeginDate() *int64 {
	return s.UpdateBeginDate
}

func (s *GetTlogDeviceListRequest) GetUpdateEndDate() *int64 {
	return s.UpdateEndDate
}

func (s *GetTlogDeviceListRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *GetTlogDeviceListRequest) GetUtdid() *string {
	return s.Utdid
}

func (s *GetTlogDeviceListRequest) SetAppKey(v int64) *GetTlogDeviceListRequest {
	s.AppKey = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetAppVersion(v string) *GetTlogDeviceListRequest {
	s.AppVersion = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetBeginDate(v int64) *GetTlogDeviceListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetBrand(v string) *GetTlogDeviceListRequest {
	s.Brand = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetCity(v string) *GetTlogDeviceListRequest {
	s.City = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetCreateBeginDate(v int64) *GetTlogDeviceListRequest {
	s.CreateBeginDate = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetCreateEndDate(v int64) *GetTlogDeviceListRequest {
	s.CreateEndDate = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetEndDate(v int64) *GetTlogDeviceListRequest {
	s.EndDate = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetKeyword(v string) *GetTlogDeviceListRequest {
	s.Keyword = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetModel(v string) *GetTlogDeviceListRequest {
	s.Model = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetOs(v string) *GetTlogDeviceListRequest {
	s.Os = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetOsVersion(v string) *GetTlogDeviceListRequest {
	s.OsVersion = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetPageIndex(v int32) *GetTlogDeviceListRequest {
	s.PageIndex = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetPageSize(v int32) *GetTlogDeviceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetUpdateBeginDate(v int64) *GetTlogDeviceListRequest {
	s.UpdateBeginDate = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetUpdateEndDate(v int64) *GetTlogDeviceListRequest {
	s.UpdateEndDate = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetUserNick(v string) *GetTlogDeviceListRequest {
	s.UserNick = &v
	return s
}

func (s *GetTlogDeviceListRequest) SetUtdid(v string) *GetTlogDeviceListRequest {
	s.Utdid = &v
	return s
}

func (s *GetTlogDeviceListRequest) Validate() error {
	return dara.Validate(s)
}
