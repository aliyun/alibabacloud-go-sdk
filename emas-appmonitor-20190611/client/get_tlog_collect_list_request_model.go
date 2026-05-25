// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogCollectListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetTlogCollectListRequest
	GetAppKey() *int64
	SetAppVersion(v string) *GetTlogCollectListRequest
	GetAppVersion() *string
	SetAuthor(v string) *GetTlogCollectListRequest
	GetAuthor() *string
	SetBeginDate(v int64) *GetTlogCollectListRequest
	GetBeginDate() *int64
	SetCity(v string) *GetTlogCollectListRequest
	GetCity() *string
	SetCreateBeginDate(v int64) *GetTlogCollectListRequest
	GetCreateBeginDate() *int64
	SetCreateEndDate(v int64) *GetTlogCollectListRequest
	GetCreateEndDate() *int64
	SetDeviceId(v string) *GetTlogCollectListRequest
	GetDeviceId() *string
	SetEndDate(v int64) *GetTlogCollectListRequest
	GetEndDate() *int64
	SetKeyword(v string) *GetTlogCollectListRequest
	GetKeyword() *string
	SetModel(v string) *GetTlogCollectListRequest
	GetModel() *string
	SetOs(v string) *GetTlogCollectListRequest
	GetOs() *string
	SetOsVersion(v string) *GetTlogCollectListRequest
	GetOsVersion() *string
	SetPageIndex(v int32) *GetTlogCollectListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetTlogCollectListRequest
	GetPageSize() *int32
	SetPositiveComment(v string) *GetTlogCollectListRequest
	GetPositiveComment() *string
	SetSourceType(v string) *GetTlogCollectListRequest
	GetSourceType() *string
	SetStatus(v string) *GetTlogCollectListRequest
	GetStatus() *string
	SetUpdateBeginDate(v int64) *GetTlogCollectListRequest
	GetUpdateBeginDate() *int64
	SetUpdateEndDate(v int64) *GetTlogCollectListRequest
	GetUpdateEndDate() *int64
	SetUserNick(v string) *GetTlogCollectListRequest
	GetUserNick() *string
	SetUtdid(v string) *GetTlogCollectListRequest
	GetUtdid() *string
}

type GetTlogCollectListRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 24780725
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// admin
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 1777996800000
	BeginDate *int64 `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 1777996800000
	CreateBeginDate *int64 `json:"CreateBeginDate,omitempty" xml:"CreateBeginDate,omitempty"`
	// example:
	//
	// 1777996800000
	CreateEndDate *int64 `json:"CreateEndDate,omitempty" xml:"CreateEndDate,omitempty"`
	// example:
	//
	// ad-0002nzx5r86f7jzrv0nx-91
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 1777996800000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// user_nick
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// iphone
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1
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
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// positiveComment
	PositiveComment *string `json:"PositiveComment,omitempty" xml:"PositiveComment,omitempty"`
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// START
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1777996800000
	UpdateBeginDate *int64 `json:"UpdateBeginDate,omitempty" xml:"UpdateBeginDate,omitempty"`
	// example:
	//
	// 1777996800000
	UpdateEndDate *int64 `json:"UpdateEndDate,omitempty" xml:"UpdateEndDate,omitempty"`
	// example:
	//
	// user_nick
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	// example:
	//
	// Z70g6V/MXJ8DABtD53eHzn4X
	Utdid *string `json:"Utdid,omitempty" xml:"Utdid,omitempty"`
}

func (s GetTlogCollectListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTlogCollectListRequest) GoString() string {
	return s.String()
}

func (s *GetTlogCollectListRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetTlogCollectListRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetTlogCollectListRequest) GetAuthor() *string {
	return s.Author
}

func (s *GetTlogCollectListRequest) GetBeginDate() *int64 {
	return s.BeginDate
}

func (s *GetTlogCollectListRequest) GetCity() *string {
	return s.City
}

func (s *GetTlogCollectListRequest) GetCreateBeginDate() *int64 {
	return s.CreateBeginDate
}

func (s *GetTlogCollectListRequest) GetCreateEndDate() *int64 {
	return s.CreateEndDate
}

func (s *GetTlogCollectListRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetTlogCollectListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetTlogCollectListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetTlogCollectListRequest) GetModel() *string {
	return s.Model
}

func (s *GetTlogCollectListRequest) GetOs() *string {
	return s.Os
}

func (s *GetTlogCollectListRequest) GetOsVersion() *string {
	return s.OsVersion
}

func (s *GetTlogCollectListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetTlogCollectListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTlogCollectListRequest) GetPositiveComment() *string {
	return s.PositiveComment
}

func (s *GetTlogCollectListRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *GetTlogCollectListRequest) GetStatus() *string {
	return s.Status
}

func (s *GetTlogCollectListRequest) GetUpdateBeginDate() *int64 {
	return s.UpdateBeginDate
}

func (s *GetTlogCollectListRequest) GetUpdateEndDate() *int64 {
	return s.UpdateEndDate
}

func (s *GetTlogCollectListRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *GetTlogCollectListRequest) GetUtdid() *string {
	return s.Utdid
}

func (s *GetTlogCollectListRequest) SetAppKey(v int64) *GetTlogCollectListRequest {
	s.AppKey = &v
	return s
}

func (s *GetTlogCollectListRequest) SetAppVersion(v string) *GetTlogCollectListRequest {
	s.AppVersion = &v
	return s
}

func (s *GetTlogCollectListRequest) SetAuthor(v string) *GetTlogCollectListRequest {
	s.Author = &v
	return s
}

func (s *GetTlogCollectListRequest) SetBeginDate(v int64) *GetTlogCollectListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetTlogCollectListRequest) SetCity(v string) *GetTlogCollectListRequest {
	s.City = &v
	return s
}

func (s *GetTlogCollectListRequest) SetCreateBeginDate(v int64) *GetTlogCollectListRequest {
	s.CreateBeginDate = &v
	return s
}

func (s *GetTlogCollectListRequest) SetCreateEndDate(v int64) *GetTlogCollectListRequest {
	s.CreateEndDate = &v
	return s
}

func (s *GetTlogCollectListRequest) SetDeviceId(v string) *GetTlogCollectListRequest {
	s.DeviceId = &v
	return s
}

func (s *GetTlogCollectListRequest) SetEndDate(v int64) *GetTlogCollectListRequest {
	s.EndDate = &v
	return s
}

func (s *GetTlogCollectListRequest) SetKeyword(v string) *GetTlogCollectListRequest {
	s.Keyword = &v
	return s
}

func (s *GetTlogCollectListRequest) SetModel(v string) *GetTlogCollectListRequest {
	s.Model = &v
	return s
}

func (s *GetTlogCollectListRequest) SetOs(v string) *GetTlogCollectListRequest {
	s.Os = &v
	return s
}

func (s *GetTlogCollectListRequest) SetOsVersion(v string) *GetTlogCollectListRequest {
	s.OsVersion = &v
	return s
}

func (s *GetTlogCollectListRequest) SetPageIndex(v int32) *GetTlogCollectListRequest {
	s.PageIndex = &v
	return s
}

func (s *GetTlogCollectListRequest) SetPageSize(v int32) *GetTlogCollectListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTlogCollectListRequest) SetPositiveComment(v string) *GetTlogCollectListRequest {
	s.PositiveComment = &v
	return s
}

func (s *GetTlogCollectListRequest) SetSourceType(v string) *GetTlogCollectListRequest {
	s.SourceType = &v
	return s
}

func (s *GetTlogCollectListRequest) SetStatus(v string) *GetTlogCollectListRequest {
	s.Status = &v
	return s
}

func (s *GetTlogCollectListRequest) SetUpdateBeginDate(v int64) *GetTlogCollectListRequest {
	s.UpdateBeginDate = &v
	return s
}

func (s *GetTlogCollectListRequest) SetUpdateEndDate(v int64) *GetTlogCollectListRequest {
	s.UpdateEndDate = &v
	return s
}

func (s *GetTlogCollectListRequest) SetUserNick(v string) *GetTlogCollectListRequest {
	s.UserNick = &v
	return s
}

func (s *GetTlogCollectListRequest) SetUtdid(v string) *GetTlogCollectListRequest {
	s.Utdid = &v
	return s
}

func (s *GetTlogCollectListRequest) Validate() error {
	return dara.Validate(s)
}
