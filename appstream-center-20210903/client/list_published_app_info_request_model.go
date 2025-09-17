// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListPublishedAppInfoRequest
	GetAppName() *string
	SetBizRegionId(v string) *ListPublishedAppInfoRequest
	GetBizRegionId() *string
	SetCategoryId(v int64) *ListPublishedAppInfoRequest
	GetCategoryId() *int64
	SetCategoryType(v int64) *ListPublishedAppInfoRequest
	GetCategoryType() *int64
	SetClientId(v string) *ListPublishedAppInfoRequest
	GetClientId() *string
	SetClientIp(v string) *ListPublishedAppInfoRequest
	GetClientIp() *string
	SetClientOS(v string) *ListPublishedAppInfoRequest
	GetClientOS() *string
	SetClientVersion(v string) *ListPublishedAppInfoRequest
	GetClientVersion() *string
	SetEndUserId(v string) *ListPublishedAppInfoRequest
	GetEndUserId() *string
	SetLoginRegionId(v string) *ListPublishedAppInfoRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *ListPublishedAppInfoRequest
	GetLoginToken() *string
	SetOrderParam(v string) *ListPublishedAppInfoRequest
	GetOrderParam() *string
	SetProductType(v string) *ListPublishedAppInfoRequest
	GetProductType() *string
	SetSessionId(v string) *ListPublishedAppInfoRequest
	GetSessionId() *string
	SetSortType(v string) *ListPublishedAppInfoRequest
	GetSortType() *string
}

type ListPublishedAppInfoRequest struct {
	// example:
	//
	// Microsoft Word
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// 1
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 1
	CategoryType *int64 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 17b38aaa-761f-44c5-9862-2ad0f5025d15
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 125.80.132.13
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Enterprise\\" 10.0 (Build 14393)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 2.0.1-D-20211008.101607
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-shanghai
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v189fa78c1aff77a0483b16497517322299131027b85bb84bbdc0871988ce8296d8fd891e2fdeaded3bd75f81f639acee8
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OrderParam *string `json:"OrderParam,omitempty" xml:"OrderParam,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// c261a6a1-e242-4f4b-813c-5fe807e49f03
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SortType  *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListPublishedAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInfoRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListPublishedAppInfoRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListPublishedAppInfoRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ListPublishedAppInfoRequest) GetCategoryType() *int64 {
	return s.CategoryType
}

func (s *ListPublishedAppInfoRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ListPublishedAppInfoRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListPublishedAppInfoRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *ListPublishedAppInfoRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ListPublishedAppInfoRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListPublishedAppInfoRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *ListPublishedAppInfoRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ListPublishedAppInfoRequest) GetOrderParam() *string {
	return s.OrderParam
}

func (s *ListPublishedAppInfoRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListPublishedAppInfoRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListPublishedAppInfoRequest) GetSortType() *string {
	return s.SortType
}

func (s *ListPublishedAppInfoRequest) SetAppName(v string) *ListPublishedAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetBizRegionId(v string) *ListPublishedAppInfoRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetCategoryId(v int64) *ListPublishedAppInfoRequest {
	s.CategoryId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetCategoryType(v int64) *ListPublishedAppInfoRequest {
	s.CategoryType = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientId(v string) *ListPublishedAppInfoRequest {
	s.ClientId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientIp(v string) *ListPublishedAppInfoRequest {
	s.ClientIp = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientOS(v string) *ListPublishedAppInfoRequest {
	s.ClientOS = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientVersion(v string) *ListPublishedAppInfoRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetEndUserId(v string) *ListPublishedAppInfoRequest {
	s.EndUserId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetLoginRegionId(v string) *ListPublishedAppInfoRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetLoginToken(v string) *ListPublishedAppInfoRequest {
	s.LoginToken = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetOrderParam(v string) *ListPublishedAppInfoRequest {
	s.OrderParam = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetProductType(v string) *ListPublishedAppInfoRequest {
	s.ProductType = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetSessionId(v string) *ListPublishedAppInfoRequest {
	s.SessionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetSortType(v string) *ListPublishedAppInfoRequest {
	s.SortType = &v
	return s
}

func (s *ListPublishedAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
