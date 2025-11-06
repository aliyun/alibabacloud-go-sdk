// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginStartDate(v int64) *ListServerLockRequest
	GetBeginStartDate() *int64
	SetDomainName(v string) *ListServerLockRequest
	GetDomainName() *string
	SetEndExpireDate(v int64) *ListServerLockRequest
	GetEndExpireDate() *int64
	SetEndStartDate(v int64) *ListServerLockRequest
	GetEndStartDate() *int64
	SetLang(v string) *ListServerLockRequest
	GetLang() *string
	SetLockProductId(v string) *ListServerLockRequest
	GetLockProductId() *string
	SetOrderBy(v string) *ListServerLockRequest
	GetOrderBy() *string
	SetOrderByType(v string) *ListServerLockRequest
	GetOrderByType() *string
	SetPageNum(v int32) *ListServerLockRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListServerLockRequest
	GetPageSize() *int32
	SetServerLockStatus(v int32) *ListServerLockRequest
	GetServerLockStatus() *int32
	SetStartExpireDate(v int64) *ListServerLockRequest
	GetStartExpireDate() *int64
	SetUserClientIp(v string) *ListServerLockRequest
	GetUserClientIp() *string
}

type ListServerLockRequest struct {
	// The start of the time range to query.
	//
	// example:
	//
	// 2021-07-10 17:37:36
	BeginStartDate *int64 `json:"BeginStartDate,omitempty" xml:"BeginStartDate,omitempty"`
	// The domain name for which you want to query the enabled registry lock.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the expiration time.
	//
	// example:
	//
	// 2021-07-10 17:37:36
	EndExpireDate *int64 `json:"EndExpireDate,omitempty" xml:"EndExpireDate,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 2020-07-10 17:37:36
	EndStartDate *int64 `json:"EndStartDate,omitempty" xml:"EndStartDate,omitempty"`
	// The language of the error message to return if the request fails. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the product to which the domain name with the registry lock enabled belongs.
	//
	// example:
	//
	// 1807**
	LockProductId *string `json:"LockProductId,omitempty" xml:"LockProductId,omitempty"`
	// The field that you use to sort the query results.
	//
	// Valid values:
	//
	// 	- EXPIRE_DATE
	//
	// example:
	//
	// EXPIRE_DATE
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order of the information based on which you want to sort the domain names, such as the registration date and expiration date. Valid values: ASC and DESC. The value ASC specifies the ascending order. The value DESC specifies the descending order. If this parameter is not configured, the default value DESC is used.
	//
	// example:
	//
	// DESC
	OrderByType *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the registry lock. Valid values:
	//
	// 	- 1: The registry lock is not enabled.
	//
	// 	- 2: The registry lock is enabled.
	//
	// 	- 3: The registry lock is disabled.
	//
	// example:
	//
	// 1
	ServerLockStatus *int32 `json:"ServerLockStatus,omitempty" xml:"ServerLockStatus,omitempty"`
	// The start of the expiration time.
	//
	// example:
	//
	// 2020-07-10 17:37:36
	StartExpireDate *int64 `json:"StartExpireDate,omitempty" xml:"StartExpireDate,omitempty"`
	// The IP address of the client. For example, you can set the value to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ListServerLockRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServerLockRequest) GoString() string {
	return s.String()
}

func (s *ListServerLockRequest) GetBeginStartDate() *int64 {
	return s.BeginStartDate
}

func (s *ListServerLockRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListServerLockRequest) GetEndExpireDate() *int64 {
	return s.EndExpireDate
}

func (s *ListServerLockRequest) GetEndStartDate() *int64 {
	return s.EndStartDate
}

func (s *ListServerLockRequest) GetLang() *string {
	return s.Lang
}

func (s *ListServerLockRequest) GetLockProductId() *string {
	return s.LockProductId
}

func (s *ListServerLockRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListServerLockRequest) GetOrderByType() *string {
	return s.OrderByType
}

func (s *ListServerLockRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListServerLockRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServerLockRequest) GetServerLockStatus() *int32 {
	return s.ServerLockStatus
}

func (s *ListServerLockRequest) GetStartExpireDate() *int64 {
	return s.StartExpireDate
}

func (s *ListServerLockRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ListServerLockRequest) SetBeginStartDate(v int64) *ListServerLockRequest {
	s.BeginStartDate = &v
	return s
}

func (s *ListServerLockRequest) SetDomainName(v string) *ListServerLockRequest {
	s.DomainName = &v
	return s
}

func (s *ListServerLockRequest) SetEndExpireDate(v int64) *ListServerLockRequest {
	s.EndExpireDate = &v
	return s
}

func (s *ListServerLockRequest) SetEndStartDate(v int64) *ListServerLockRequest {
	s.EndStartDate = &v
	return s
}

func (s *ListServerLockRequest) SetLang(v string) *ListServerLockRequest {
	s.Lang = &v
	return s
}

func (s *ListServerLockRequest) SetLockProductId(v string) *ListServerLockRequest {
	s.LockProductId = &v
	return s
}

func (s *ListServerLockRequest) SetOrderBy(v string) *ListServerLockRequest {
	s.OrderBy = &v
	return s
}

func (s *ListServerLockRequest) SetOrderByType(v string) *ListServerLockRequest {
	s.OrderByType = &v
	return s
}

func (s *ListServerLockRequest) SetPageNum(v int32) *ListServerLockRequest {
	s.PageNum = &v
	return s
}

func (s *ListServerLockRequest) SetPageSize(v int32) *ListServerLockRequest {
	s.PageSize = &v
	return s
}

func (s *ListServerLockRequest) SetServerLockStatus(v int32) *ListServerLockRequest {
	s.ServerLockStatus = &v
	return s
}

func (s *ListServerLockRequest) SetStartExpireDate(v int64) *ListServerLockRequest {
	s.StartExpireDate = &v
	return s
}

func (s *ListServerLockRequest) SetUserClientIp(v string) *ListServerLockRequest {
	s.UserClientIp = &v
	return s
}

func (s *ListServerLockRequest) Validate() error {
	return dara.Validate(s)
}
